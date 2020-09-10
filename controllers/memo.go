package controllers

import (
	"encoding/json"
	"goforit/services/MemoService"
	"goforit/utils"
	"strconv"
	"strings"
)

type MemoController struct {
	BaseController
}

type memoReq struct {
	baseRequest
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	ClassifyId    uint64 `json:"classify_id"`
	ClassifyIdStr string `json:"classify_id_str"`
}

// @router / [get]
func (m *MemoController) GetMemo() {
	userId := m.NeedLogin(true)
	var data memoReq
	json.Unmarshal(utils.FormatQuery(m.Input().Encode()), &data)
	pageIndex, _ := strconv.Atoi(data.PageIndex)
	pageSize, _ := strconv.Atoi(data.PageSize)
	title := strings.Trim(data.Title, " ")
	classifyId := data.ClassifyIdStr
	m.Normal(MemoService.GetMemo(userId, title, classifyId, uint64(pageIndex), uint64(pageSize)), "请求成功")
}

// @router / [post]
func (m *MemoController) Create() {
	userId := m.NeedLogin(true)
	var data memoReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	if strings.Trim(data.Title, " ") == "" {
		m.NormalException("标题不能为空")
	}
	if strings.Trim(data.Content, " ") == "" {
		m.NormalException("内容不能为空")
	}
	if data.ClassifyId == 0 {
		m.NormalException("分类ID不能为空")
	}
	if _, count := MemoService.FindMemoClassify(data.ClassifyId, userId); count == 0 {
		m.NormalException("分类不存在")
	}
	if true == MemoService.CreateMemo(data.Title, data.Content, data.ClassifyId, userId) {
		m.Normal(NoneObject{}, "创建成功")
	} else {
		m.NormalException("创建失败")
	}
}

// @router / [put]
func (m *MemoController) Edit() {
	userId := m.NeedLogin(true)
	var data memoReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	if data.ID == 0 {
		m.NormalException("ID不能为空")
	}
	if strings.Trim(data.Title, " ") == "" {
		m.NormalException("标题不能为空")
	}
	if strings.Trim(data.Content, " ") == "" {
		m.NormalException("内容不能为空")
	}
	if data.ClassifyId == 0 {
		m.NormalException("分类ID不能为空")
	}
	if _, count := MemoService.FindMemoClassify(data.ClassifyId, userId); count == 0 {
		m.NormalException("分类不存在")
	}
	if MemoService.EditMemo(data.ID, data.Title, data.Content, data.ClassifyId, userId) > 0 {
		m.Normal(NoneObject{}, "修改成功")
	} else {
		m.NormalException("修改失败")
	}
}

// @router / [delete]
func (m *MemoController) Delete() {
	userId := m.NeedLogin(true)
	var data memoReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	if data.ID == 0 {
		m.NormalException("ID不能为空")
	}
	if MemoService.DeleteMemo(data.ID, userId) > 0 {
		m.Normal(NoneObject{}, "删除成功")
	} else {
		m.NormalException("删除失败")
	}
}

type memoClassifyReq struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

// @router /classify [get]
func (m *MemoController) GetClassify() {
	userId := m.NeedLogin(true)
	m.Normal(MemoService.GetMemoClassify(userId), "请求成功")
}

// @router /classify [post]
func (m *MemoController) CreateClassify() {
	userId := m.NeedLogin(true)
	var data memoClassifyReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	if strings.Trim(data.Name, " ") == "" {
		m.NormalException("分类名称不能为空")
	}
	if true == MemoService.CreateMemoClassify(data.Name, userId) {
		m.Normal(NoneObject{}, "创建成功")
	} else {
		m.NormalException("创建失败")
	}
	m.Normal(NoneObject{}, "请求成功")
}

// @router /classify [put]
func (m *MemoController) EditClassify() {
	userId := m.NeedLogin(true)
	var data memoClassifyReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	if data.ID == 0 {
		m.NormalException("ID不能为空")
	}
	if strings.Trim(data.Name, " ") == "" {
		m.NormalException("分类名称不能为空")
	}
	if result := MemoService.EditMemoClassify(data.Name, data.ID, userId); result == 0 {
		m.NormalException("修改失败")
	} else {
		m.Normal(NoneObject{}, "修改成功")
	}
	m.Normal(NoneObject{}, "请求成功")
}

// @router /classify [delete]
func (m *MemoController) DeleteClassify() {
	userId := m.NeedLogin(true)
	var data memoClassifyReq
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	if data.ID == 0 {
		m.NormalException("ID不能为空")
	}
	if count := MemoService.DeleteMemoClassify(data.ID, userId); count > 0 {
		m.Normal(NoneObject{}, "删除成功")
	} else {
		m.NormalException("删除失败")
	}
	m.Normal(NoneObject{}, "请求成功")
}
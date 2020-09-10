package Page

import "math"

type Lists struct {
	Page pageResult  `json:"page"`
	List interface{} `json:"list"`
}

type pageResult struct {
	ItemTotal uint64 `json:"itemTotal"`
	PageTotal uint64 `json:"pageTotal"`
	PageIndex uint64 `json:"pageIndex"`
	PageSize  uint64 `json:"pageSize"`
}

func Make(count uint64, pageIndex uint64, pageSize uint64) (offset uint64, length uint64, result pageResult) {
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset = (pageIndex - 1) * pageSize
	length = pageSize
	result = pageResult{
		ItemTotal: count,
		PageTotal: uint64(math.Ceil(float64(count) / float64(pageSize))),
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
	return
}

package test

import (
	"fmt"
	_ "goforit/routers"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestS(t *testing.T) {
	str := "123123.0052"
	result, num := IsNum(str)
	fmt.Println(result, num)
}

func IsNum(s string) (bool, float64) {
	result, err := strconv.ParseFloat(s, 64)
	return err == nil, result
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	//r, _ := http.NewRequest("GET", "/v1/object", nil)
	//w := httptest.NewRecorder()
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	//
	//beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	//
	//Convey("Subject: Test Station Endpoint\n", t, func() {
	//	Convey("Status Code Should Be 200", func() {
	//		So(w.Code, ShouldEqual, 200)
	//	})
	//	Convey("The Result Should Not Be Empty", func() {
	//		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	//	})
	//})
}

package service

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"github.com/leo0o/goblog/accountservice/dbclient"
	"github.com/leo0o/goblog/accountservice/model"
	"fmt"
	"encoding/json"
	"net/http"
)

func TestGetAccountWithWrongPath(t *testing.T) {

	convey.Convey("访问一个错的接口地址", t, func() {
		req := httptest.NewRequest("GET", "/xxx/xxxx", nil)
		resp := httptest.NewRecorder()

		convey.Convey("当路由接收到错误的路径时", func() {
			NewRouter().ServeHTTP(resp, req)

			convey.Convey("应该会响应404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)

			})

		})
	})
}


func TestGetAccount(t *testing.T) {
	mock := &dbclient.MockBoltClient{}
	mock.On("QueryAccount", "1111").Return(model.Account{Id:"100001", Name:"leooo."}, nil)
	mock.On("QueryAccount", "2222").Return(model.Account{}, fmt.Errorf("error"))
	DBClient = mock

	convey.Convey("访问 /accounts/1111", t, func() {
		req := httptest.NewRequest("GET", "/accounts/1111", nil)
		rsp := httptest.NewRecorder()

		convey.Convey("当接受到请求时", func() {
			NewRouter().ServeHTTP(rsp, req)

			convey.Convey("应该返回用户 100001", func() {
				account := model.Account{}
				json.Unmarshal(rsp.Body.Bytes(), &account)
				convey.So(account.Id, convey.ShouldEqual, "100001")
				convey.So(account.Name, convey.ShouldEqual, "leooo.")
			})
		})

	})


	convey.Convey("访问 /accounts/2222", t, func() {
		req := httptest.NewRequest("GET", "/accounts/2222", nil)
		rsp := httptest.NewRecorder()

		convey.Convey("当接受到请求时", func() {
			NewRouter().ServeHTTP(rsp, req)

			convey.Convey("应该返回 404", func() {
				convey.So(rsp.Code, convey.ShouldEqual, http.StatusNotFound)
			})
		})

	})
}

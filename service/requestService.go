package service

import (
	"InterfaceTest/common/httpTool"
	"InterfaceTest/model"
	"net/http"
)

type RequestService struct {
}

//表示 excel映射 --> request对象
type RequestExcelObject struct {
	Id     int
	Url    string
	Method string
	param  map[string]string
	group  int
	order  int
}

func (rs *RequestService) ExecRequest(hc *http.Client, url, method string, param map[string]string) (string, error) {
	m := httpTool.MakeHeader()
	tamp := httpTool.Map2Str(param)
	model1 := model.RequestModel{Url: url, Method: method, RequestHead: m, RequestBody: tamp}
	body := model1.SendRequestForResponseBody(hc, nil)
	//fmt.Print(body)
	return body, nil
}

//执行 一组请求 . 返回 一组结果, 如果发生错误则 立即停止该组请求,不会继续往下执行请求
func (rs *RequestService) ExecRequests(hc *http.Client, re *[]RequestExcelObject) ([]string, error) {
	var result []string
	for _, v := range *re {
		request, err := rs.ExecRequest(hc, v.Url, v.Method, v.param)
		if err != nil {
			return nil, err
		}
		result = append(result, request)
	}
	return result, nil
}

func (rs *RequestService) ExecuteQuest() {

}

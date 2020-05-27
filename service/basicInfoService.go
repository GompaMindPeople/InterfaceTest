package service

import (
	"InterfaceTest/common/httpTool"
	"InterfaceTest/model"
	"net/http"
)

type BasicInfoService struct {
}

//nickname: &amp;#039;
//contact_phone: 555-555-5555
//province: 澳门特别行政区
//city: 花地玛堂区
//district:
//user_address: 753 Main Street
//business: 服务器租用防御
//website: idc.zhangqw.cn

//nickname: 1234
//contact_phone:
//province: 北京市
//city: 北京市市辖区
//district: 西城区
//user_address:
//business:
//website:
//,name,phone,province,city,district,user_address,business,website string
func (ls *BasicInfoService) Save(hc *http.Client, param map[string]string) (string, error) {
	m := httpTool.MakeHeader()
	tamp := httpTool.Map2Str(param)

	model1 := model.RequestModel{Url: "http://test.kkclouds.com/user/user.html", Method: "POST", RequestHead: m, RequestBody: tamp}

	body := model1.SendRequestForResponseBody(hc, nil)
	//fmt.Print(body)
	return body, nil
}

func (ls *BasicInfoService) Order(hc *http.Client, param map[string]string) (string, error) {
	m := httpTool.MakeHeader()
	tamp := httpTool.Map2Str(param)
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/control/order.html", Method: "GET", RequestHead: m, RequestBody: tamp}

	body := model1.SendRequestForResponseBody(hc, nil)
	//fmt.Print(body)
	return body, nil
}

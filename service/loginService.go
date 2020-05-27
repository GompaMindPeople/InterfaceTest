package service

import (
	"InterfaceTest/common/httpTool"
	"InterfaceTest/model"
	"net/http"
)

type LoginService struct {
}

func (ls *LoginService) Login(hc *http.Client, username, password string) (string, error) {

	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/passport/cloudlogin.html", Method: "POST", RequestHead: m, RequestBody: "username=" + username + "&password=" + password + "&verify=&phone=&msg_code="}
	body := model1.SendRequestForResponseBody(hc, nil)
	//fmt.Print(body)
	return body, nil
}

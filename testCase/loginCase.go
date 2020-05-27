package testCase

import (
	"InterfaceTest/model"
	"InterfaceTest/service"
)

var ser = service.LoginService{}
var httpClient = model.HttpClient{}
var client = httpClient.NewClient(nil)

//所有的case 都需要登录操作
func Login(username, password string) string {
	result, _ := ser.Login(client.Client, username, password)
	return result
}

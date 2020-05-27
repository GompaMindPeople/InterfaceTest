package testCase

import (
	"InterfaceTest/service"
)

var BasicInfoSer = service.BasicInfoService{}

func Save(m map[string]string) string {

	result, _ := BasicInfoSer.Save(client.Client, m)
	return result
}
func Order(m map[string]string) string {

	result, _ := BasicInfoSer.Order(client.Client, m)
	return result
}

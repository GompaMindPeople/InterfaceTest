package testCase

import (
	"InterfaceTest/common/assert"
	"InterfaceTest/service"
	"testing"
)

func TestLogin(t *testing.T) {

	result := service.ReadExcel("login")
	for _, v := range result {
		loginResult := Login(v["username"], v["password"])
		t.Log(assert.AssertByRegMap(loginResult, v["pattern"], v))
	}

}

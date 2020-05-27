package testCase

import (
	"InterfaceTest/service"
	"testing"
)

//nickname: &amp;#039;
//contact_phone: 555-555-5555
//province: 澳门特别行政区
//city: 花地玛堂区
//district:
//user_address: 753 Main Street
//business: 服务器租用防御
//website: idc.zhangqw.cn

func TestSave(t *testing.T) {

	Login("kkwlcs", "a123456")
	result := service.ReadExcel("basicInfo")
	for _, v := range result {
		t.Log(Save(v))
	}

}

func TestOrder(t *testing.T) {

	Login("kkwlcs", "a123456")
	//result := service.ReadExcel("basicInfo")
	//for _,v  := range result{
	//	t.Log(Save(v))
	//}
	m := map[string]string{}

	m["system"] = ""
	m["type"] = "14"
	m["status"] = ""
	m["starttime"] = ""
	m["endtime"] = ""
	t.Log(Order(m))

}

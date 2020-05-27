package service

import "InterfaceTest/common"

var excelPath = "../case/"

func ReadExcel(name string) []map[string]string {
	return common.ReadExcelForMap(excelPath + name + ".xlsx")
}

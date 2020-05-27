package common

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func ReadExcelForMap(filename string) []map[string]string {

	result := []map[string]string{}
	ReadExcelByRow(filename, func(i int, row *xlsx.Row, m map[int]string) {

		temp := map[string]string{}
		//result= append(result,temp)
		//枚举 每一行的数据
		for k, v := range row.Cells {
			text := v.String()
			temp[m[k]] = text
		}
		result = append(result, temp)
	})

	return result
}

func ReadExcelByRow(filename string, cb func(int, *xlsx.Row, map[int]string)) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("open Excel failed: %s\n", err)
		return
	}
	for _, sheet := range xlFile.Sheets {

		//fmt.Printf("Sheet Name: %s\n", sheet.Name)
		//tmpOra := ora{}
		// 获取标签页(时间)
		//tmpOra.TIME = sheet.Name
		m := map[int]string{}

		if sheet.Rows == nil {
			break
		}
		//先获取第一行, 第一行为 列的key值
		for k, v := range sheet.Rows[1].Cells {
			text := v.String()
			m[k] = text
		}

		rows := len(sheet.Rows)
		//从第二行开始 获取 正式的数据
		for i := 2; i < rows; i++ {
			//result= append(result,temp)
			//枚举 每一行的数据
			cb(i, sheet.Rows[i], m)
		}

	}
}

func ReadExcelForObject(fileName string) interface{} {

	return nil
}

package assert

import "InterfaceTest/common"

//通过正则表达式匹配数据是否已存在
func AssertsByReg(str, pattern string) bool {
	return common.MatchString(pattern, str)
}

func AssertByRegMap(str, pattern string, m map[string]string) string {
	var result = ""
	if AssertsByReg(str, pattern) {
		return "true"
	} else {
		for k, v := range m {
			result = result + k + ":" + v + ","
		}
		result += "\nresult:" + str
	}
	return result
}

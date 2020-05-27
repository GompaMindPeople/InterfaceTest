package common

import (
	"testing"
)

func TestReadExcel(t *testing.T) {
	excel := ReadExcel("../case/login.xlsx")

	for _, v := range excel {
		for k, v2 := range v {
			t.Log("k=", k+",v="+v2+"; ")
		}
		t.Log("\n")
	}
}

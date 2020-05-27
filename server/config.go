package server

import "InterfaceTest/common"

func GetConfigByYaml() map[string]interface{} {
	config := common.ReadConfig("../config/conf.yaml")
	return config
}

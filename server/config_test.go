package server

import (
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	yaml := GetConfigByYaml()

	log.Print(yaml)

}

package model

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

type HttpClient struct {
	Client *http.Client
}

func (hc *HttpClient) NewClient(o *cookiejar.Options) *HttpClient {

	client := http.Client{}
	jar, err := cookiejar.New(o)
	if err != nil {
		log.Print("构造cookiejar时,发生错误->", err)
	}
	client.Jar = jar
	hc.Client = &client
	return hc
}

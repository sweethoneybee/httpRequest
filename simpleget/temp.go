package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values {
		"query" : {"hello world!"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	// 문자열로 "200 OK"
	log.Println("Status:", resp.Status)
	// 수치로 200
	log.Println("StatusCode:", resp.StatusCode)
	// 헤더 출력
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	TIMEOUT     = 60
	METHOD_POST = "POST"
	EXIT_CODE   = -1
	BODY_TYPE   = "application/json"
	ZERO        = 0
)

var client *HttpClient

type HttpClient struct {
	Client *http.Client
}

func NewHttpClient(timeout int) *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			TimeOut: time.Duration(timeout) * time.Second,
		},
	}
}

func InitHttp() error {
	client = NewHttpClient(TIMEOUT)
	return nil
}

type HttpReq struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   string `json:"data"`
}

type Params struct {
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
	Param3 string `json:"param3"`
	Param4 string `json:"param4"`
}

//第一种，http.Post()
func PostHttpRequest(params *Params) error {
	var httpreq HttpReq

	if params == nil {
		return nil
	}

	//接口
	url := fmt.Sprintf("http://xxx.xxx.xxx.xxx/api/v1/post")

	postParams, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("parse post request params failed:%v\n", err)
		return err
	}

	resp, err := http.Post(url, BODY_TYPE, strings.NewReader(string(postParams)))
	if err != nil {
		fmt.Printf("post request failed:%v\n", err)
		return err
	}

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decoder(&httpreq)
	if err != nil {
		fmt.Printf("parse resp body faield:%v\n", err)
		return err
	}

	if httpreq.Errno != ZERO {
		return errors.New("Request failed")
	}

	return nil
}

//第二种，http.NewRequest()
func HttpPostRequest(params *Params) error {
	var httpreq HttpReq

	if params == nil {
		return nil
	}

	url := fmt.Sprintf("http://xxx.xxx.xxx.xxx:xxx/api/v1/post")

	postParams, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("parse post request params failed:%v\n", err)
		return err
	}

	req, err := http.NewRequest(METHOD_POST, URL, strings.NewReader(string(postParams)))
	if err != nil {
		fmt.Printf("new request failed:%v\n", err)
		return err
	}

	req.Header.Add("Content-Type", BODY_TYPE)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("do req failed:%v\n", err)
		return err
	}

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decoder(&httpreq)
	if err != nil {
		fmt.Printf("parse resp body faield:%v\n", err)
		return err
	}

	if httpreq.Errno != ZERO {
		return errors.New("Request failed")
	}

	return nil
}

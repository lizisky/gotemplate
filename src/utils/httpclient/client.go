package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	contentType_json = "application/json;charset=UTF-8"
)

func DoHttpGet(aurl string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(aurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func DoHttpPost(aurl string, bodyData []byte) ([]byte, error) {

	bodyReader := bytes.NewReader(bodyData)
	client := &http.Client{}
	resp, err := client.Post(aurl, contentType_json, bodyReader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

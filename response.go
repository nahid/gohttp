package gohttp

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
)

type Response struct {
	HttpResp *http.Response
}

func (res *Response) GetResp() *http.Response {
	return res.HttpResp
}

func (res *Response) GetStatusCode() int {
	return res.HttpResp.StatusCode
}

func (res *Response) GetBody() io.Reader {
	return res.HttpResp.Body
}


func (res *Response) GetBodyAsByte() ([]byte, error) {
	body, err := ioutil.ReadAll(res.HttpResp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (res *Response) GetBodyAsString() (string, error) {
	body, err := ioutil.ReadAll(res.HttpResp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


func (res Response) GetBodyWithUnmarshal(v interface{}) (error) {
	body, err := res.GetBodyAsByte()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)

	if err != nil {
		return err
	}

	return nil
}



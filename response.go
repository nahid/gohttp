package gohttp

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
)

// Response is a http response struct
type Response struct {
	HttpResp *http.Response
}

// GetResp get net/http original response
func (res *Response) GetResp() *http.Response {
	return res.HttpResp
}

// GetStatusCode returns http status code
func (res *Response) GetStatusCode() int {
	return res.HttpResp.StatusCode
}

// GetBody returns response body
func (res *Response) GetBody() io.Reader {
	if res.HttpResp == nil {
		return nil
	}
	return res.HttpResp.Body
}


// GetBodyAsByte returns response body as byte
func (res *Response) GetBodyAsByte() ([]byte, error) {
	body, err := ioutil.ReadAll(res.HttpResp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetBodyAsString returns resonpose body as string
func (res *Response) GetBodyAsString() (string, error) {
	body, err := ioutil.ReadAll(res.HttpResp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


// GetBodyWithUnmarshal unmarshal response body
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



package gohttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Response is a http response struct
type Response struct {
	resp *http.Response
}

// GetResp get net/http original response
func (res *Response) GetResp() *http.Response {
	return res.resp
}

// GetStatusCode returns http status code
func (res *Response) GetStatusCode() int {
	return res.resp.StatusCode
}

// GetBody returns response body
func (res *Response) GetBody() io.Reader {
	if res.resp == nil {
		return nil
	}
	return res.resp.Body
}

// GetBodyAsByte returns response body as byte
func (res *Response) GetBodyAsByte() ([]byte, error) {
	body, err := ioutil.ReadAll(res.resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetBodyAsString returns resonpose body as string
func (res *Response) GetBodyAsString() (string, error) {
	body, err := ioutil.ReadAll(res.resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetBodyWithUnmarshal unmarshal response body
func (res Response) GetBodyWithUnmarshal(v interface{}) error {
	body, err := res.GetBodyAsByte()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}

	return nil
}

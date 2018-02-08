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

// AsyncResponse is a response struct for asynchronous request
type AsyncResponse struct {
	Resp *Response
	Err  error
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
	defer res.resp.Body.Close()

	return body, nil
}

// GetBodyAsString returns response body as string
func (res *Response) GetBodyAsString() (string, error) {
	body, err := ioutil.ReadAll(res.resp.Body)
	if err != nil {
		return "", err
	}
	defer res.resp.Body.Close()

	return string(body), nil
}

// GetBodyAsJSONRawMessage returns response body as json.RawMessage
func (res *Response) GetBodyAsJSONRawMessage() (json.RawMessage, error) {
	body, err := ioutil.ReadAll(res.resp.Body)
	if err != nil {
		return nil, err
	}
	defer res.resp.Body.Close()

	return json.RawMessage(body), nil
}

// GetBodyWithUnmarshal unmarshal response body
func (res Response) GetBodyWithUnmarshal(v interface{}) error {
	body, err := res.GetBodyAsByte()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

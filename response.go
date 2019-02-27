package gohttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
// if Response is not returned from a Request
// the status code will be 0
func (res *Response) GetStatusCode() int {
	if res.resp == nil {
		return 0
	}
	return res.resp.StatusCode
}

// GetBody returns response body
// It is the caller's responsibility to close Body
func (res *Response) GetBody() io.ReadCloser {
	if res.resp == nil {
		return nil
	}
	return res.resp.Body
}

// GetBodyAsByte returns response body as byte
func (res *Response) GetBodyAsByte() ([]byte, error) {
	body := res.GetBody()
	if body == nil {
		return nil, nil
	}
	defer body.Close()

	byts, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return byts, nil
}

// GetBodyAsString returns response body as string
func (res *Response) GetBodyAsString() (string, error) {
	body, err := res.GetBodyAsByte()
	if err != nil || body == nil {
		return "", err
	}

	return string(body), nil
}

// GetBodyAsJSONRawMessage returns response body as json.RawMessage
func (res *Response) GetBodyAsJSONRawMessage() (json.RawMessage, error) {
	body, err := res.GetBodyAsByte()
	if err != nil || body == nil {
		return nil, err
	}

	return json.RawMessage(body), nil
}

// UnmarshalBody unmarshal response body
func (res *Response) UnmarshalBody(v interface{}) error {
	body, err := res.GetBodyAsByte()
	if err != nil || body == nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

//Protocol returns response proto
func (res *Response) Protocol() string{
	return res.resp.Proto
}

//URL returns response Location
func (res *Response) URL() (*url.URL, error)  {
	return res.resp.Location()
}
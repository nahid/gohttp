package gohttp

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Request is a request type
type Request struct {
	Transport              http.Transport
	Client                 *http.Client
	Cookie                 http.CookieJar
	Timeout                time.Duration
	formVals               *bytes.Buffer
	multipartBuffer        bytes.Buffer
	queryVals              string
	headers                map[string]string
	writer                 *multipart.Writer
	contentType            string
	basicUser, basicPasswd string
}

// NewRequest returns a new request
func NewRequest(opts ...Option) *Request {
	r := &Request{}
	for _, o := range opts {
		o.apply(r)
	}
	return r
}

func (req *Request) createClient() *http.Client {
	tr := &req.Transport
	if req.Client == nil {
		req.Client = &http.Client{
			Transport: tr,
			Timeout:   req.Timeout,
			Jar:       req.Cookie,
		}
	}

	return req.Client
}

// JSON set json data with request
func (req *Request) JSON(jsonBody map[string]interface{}) *Request {

	data, err := json.Marshal(jsonBody)
	if err != nil {
		panic(err)
	}

	req.formVals = bytes.NewBuffer(data)
	req.contentType = "application/json"
	return req
}

// FormData set Post request form parameters
func (req *Request) FormData(formValues map[string]string) *Request {
	vals := url.Values{}
	for key, val := range formValues {
		vals.Add(key, val)
	}

	req.formVals = bytes.NewBuffer([]byte(vals.Encode()))
	req.contentType = "application/x-www-form-urlencoded"

	return req
}

// Body set Post request as body
func (req *Request) Body(formValues []byte) *Request {

	req.formVals = bytes.NewBuffer(formValues)
	req.contentType = "application/octet-stream"

	return req
}

// Text is send text data with post request
func (req *Request) Text(formValues string) *Request {

	req.formVals = bytes.NewBuffer([]byte(formValues))
	req.contentType = "text/plain"

	return req
}

// Query set request query param
func (req *Request) Query(formValues map[string]string) *Request {
	vals := url.Values{}
	for key, val := range formValues {
		vals.Add(key, val)
	}

	req.queryVals = vals.Encode()
	req.contentType = "application/x-www-form-urlencoded"

	return req
}

// Headers set header information
func (req *Request) Headers(headerVals map[string]string) *Request {
	req.headers = headerVals
	return req
}

// BasicAuth make basic authentication
func (req *Request) BasicAuth(username, password string) *Request {
	req.basicUser = username
	req.basicPasswd = password

	return req
}

// Get is a get http request
func (req *Request) Get(url string) (*Response, error) {
	return req.makeRequest("GET", url, req.formVals)
}

// Post is a post http request
func (req *Request) Post(url string) (*Response, error) {
	return req.makeRequest("POST", url, req.formVals)
}

// Put is a put http request
func (req *Request) Put(url string) (*Response, error) {
	return req.makeRequest("PUT", url, req.formVals)
}

// Patch is a patch http request
func (req *Request) Patch(url string) (*Response, error) {
	return req.makeRequest("PATCH", url, req.formVals)
}

// Delete is a delete http request
func (req *Request) Delete(url string) (*Response, error) {
	return req.makeRequest("DELETE", url, req.formVals)
}

// MultipartFormData add form data in multipart request
func (req *Request) MultipartFormData(formData map[string]string) *Request {
	if req.writer == nil {
		req.writer = multipart.NewWriter(&req.multipartBuffer)
	}

	for key, val := range formData {
		req.writer.WriteField(key, val)
	}

	return req
}

// Upload upload a single file
func (req *Request) Upload(name, file string) *Request {
	if req.writer == nil {
		req.writer = multipart.NewWriter(&req.multipartBuffer)
	}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Add file
	fw, err := req.writer.CreateFormFile(name, file)
	if err != nil {
		panic(err)
	}
	if _, err = io.Copy(fw, f); err != nil {
		panic(err)
	}

	req.contentType = req.writer.FormDataContentType()
	req.formVals = &req.multipartBuffer
	return req
}

// Uploads upload multiple files
func (req *Request) Uploads(files map[string]string) *Request {

	for name, file := range files {
		_ = req.Upload(name, file)
	}

	return req
}

// makeRequest makes a http request
func (req *Request) makeRequest(verb, url string, payloads *bytes.Buffer) (*Response, error) {
	response := Response{}
	verb = strings.ToUpper(verb)
	var request *http.Request
	var err error
	client := req.createClient()

	if req.writer != nil {
		req.writer.Close()
	}
	if req.queryVals != "" {
		url += "?" + req.queryVals
	}

	if payloads == nil {
		payloads = bytes.NewBuffer([]byte(``))
	}

	if verb == "GET" {
		request, err = http.NewRequest(verb, url, nil)
	} else {
		request, err = http.NewRequest(verb, url, payloads)
	}

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", req.contentType)

	if req.basicUser != "" && req.basicPasswd != "" {
		request.SetBasicAuth(req.basicUser, req.basicPasswd)
	}

	// set headers from Headers method
	for key, val := range req.headers {
		request.Header.Set(key, val)
	}

	//request.Close = true
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	response.resp = resp
	return &response, nil
}

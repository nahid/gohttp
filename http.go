package gohttp

import (
	"net/url"
	"io"
	"encoding/json"
	"net/http"
	"strings"
	"bytes"
	"mime/multipart"
	"os"
)

// Request is a request type
type Request struct {
	formVals *bytes.Buffer
	multipartBuffer bytes.Buffer
	queryVals string
	headers map[string]string
	writer *multipart.Writer
	contentType string
}

// init set contentType initially
func (req *Request) init()  {
	req.contentType = "application/x-www-form-urlencoded"
}

// Json set json data with request
func (req *Request) Json(formJson map[string]interface{}) *Request  {

	data, err := json.Marshal(formJson)
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

// MultiFormData add form data in multipart request
func (req *Request) MultipartFormData(formData map[string]string) *Request  {
	if req.writer == nil {
		req.writer = multipart.NewWriter(&req.multipartBuffer)
	}

	for key, val := range formData {
		req.writer.WriteField(key, val)
	}

	return req
}

// Upload upload a single file
func (req *Request) Upload(name, file string) (*Request) {
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
func (req *Request) Uploads(files map[string]string) (*Request) {

	for name, file := range files {
		_ = req.Upload(name, file)
	}

	return req
}

// makeRequest makes a http request
func (req *Request) makeRequest(verb, url string, payloads *bytes.Buffer) (*Response, error) {
	client := http.Client{}
	response := Response{}
	verb = strings.ToUpper(verb)
	var data *bytes.Buffer

	if req.writer != nil {
		req.writer.Close()
	}
	if req.queryVals != "" {
		url += "?" + req.queryVals
	}

	if verb == "GET" {
		data = nil
	} else {
		data = payloads
	}

	request, err := http.NewRequest(verb, url, data)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", req.contentType)

	// set headers from Headers method
	for key, val := range req.headers {
		request.Header.Set(key, val)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	response.HttpResp = resp

	return &response, nil
}

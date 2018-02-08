package gohttp

import (
	"bytes"
)

// AsyncGet is a asynchronous get http request
func (req *Request) AsyncGet(url string, ch chan<- *ChannelResponse) {
	go req.makeAsyncRequest("get", url, req.formVals, ch)
}

// AsyncPost is a asynchronous post http request
func (req *Request) AsyncPost(url string, ch chan<- *ChannelResponse) {
	go req.makeAsyncRequest("post", url, req.formVals, ch)
}

// AsyncPut is a asynchronous put http request
func (req *Request) AsyncPut(url string, ch chan<- *ChannelResponse) {
	go req.makeAsyncRequest("put", url, req.formVals, ch)
}

// AsyncDelete is a asynchronous delete http request
func (req *Request) AsyncDelete(url string, ch chan<- *ChannelResponse) {
	go req.makeAsyncRequest("delete", url, req.formVals, ch)
}

// AsyncPatch is a asynchronous patch http request
func (req *Request) AsyncPatch(url string, ch chan<- *ChannelResponse) {
	go req.makeAsyncRequest("patch", url, req.formVals, ch)
}

// makeAsyncRequest generate asynchronous request
func (req *Request) makeAsyncRequest(verb, uri string, payloads *bytes.Buffer, ch chan<- *ChannelResponse) {
	var res *ChannelResponse
	resp, err := req.makeRequest(verb, uri, payloads)

	res = &ChannelResponse{
		Resp: resp,
		Err: err,
	}

	ch <- res
}


package gohttp

import "testing"

// TestAsyncGetRequest tests asynchronous GET request
func TestAsyncGetRequest(t *testing.T) {
	t.Log("Sending GET async request... (expected http code: 200)")

	req := NewRequest()
	ch := make(chan *Response)

	for i:=0; i<=100; i++ {
		req.AsyncGet("http://httpbin.org/get", ch)
	}

	for i:=0; i<=100; i++ {
		resp, err := <- ch

		if err != true {
			t.Error(err)
		}

		if resp.GetStatusCode() != 200 {
			t.Error(
				"For", "GET http://httpbin.org/get",
				"expected", 200,
				"got", resp.GetStatusCode(),
			)
		}
	}


}
package gohttp

import "testing"

// TestAsyncGetRequest tests asynchronous GET request
func TestAsyncGetRequest(t *testing.T) {
	t.Log("Sending GET async request... (expected http code: 200)")

	req := NewRequest()
	ch := make(chan *AsyncResponse)

	for i := 0; i <= 100; i++ {
		req.AsyncGet("http://httpbin.org/get", ch)
	}

	for i := 0; i <= 100; i++ {
		aRes := <-ch

		if aRes.Err != nil {
			t.Error(aRes.Err)
		}

		if aRes.Resp.GetStatusCode() != 200 {
			t.Error(
				"For", "GET http://httpbin.org/get",
				"expected", 200,
				"got", aRes.Resp.GetStatusCode(),
			)
		}
	}

}

// TestAsyncPostRequest tests asynchronous POST request
func TestAsyncPostRequest(t *testing.T) {
	t.Log("Sending POST async request... (expected http code: 200)")

	req := NewRequest()
	ch := make(chan *AsyncResponse)

	for i := 0; i <= 100; i++ {
		req.AsyncPost("http://httpbin.org/post", ch)
	}

	for i := 0; i <= 100; i++ {
		aRes := <-ch

		if aRes.Err != nil {
			t.Error(aRes.Err)
		}

		if aRes.Resp.GetStatusCode() != 200 {
			t.Error(
				"For", "POST http://httpbin.org/post",
				"expected", 200,
				"got", aRes.Resp.GetStatusCode(),
			)
		}
	}

}

// TestAsyncPutRequest tests asynchronous PUT request
func TestAsyncPutRequest(t *testing.T) {
	t.Log("Sending PUT async request... (expected http code: 200)")

	req := NewRequest()
	ch := make(chan *AsyncResponse)

	for i := 0; i <= 100; i++ {
		req.AsyncPut("http://httpbin.org/put", ch)
	}

	for i := 0; i <= 100; i++ {
		aRes := <-ch

		if aRes.Err != nil {
			t.Error(aRes.Err)
		}

		if aRes.Resp.GetStatusCode() != 200 {
			t.Error(
				"For", "PUT http://httpbin.org/put",
				"expected", 200,
				"got", aRes.Resp.GetStatusCode(),
			)
		}
	}

}

// TestAsyncPatchRequest tests asynchronous PATCH request
func TestAsyncPatchRequest(t *testing.T) {
	t.Log("Sending PATCH async request... (expected http code: 200)")

	req := NewRequest()
	ch := make(chan *AsyncResponse)

	for i := 0; i <= 100; i++ {
		req.AsyncPatch("http://httpbin.org/patch", ch)
	}

	for i := 0; i <= 100; i++ {
		aRes := <-ch

		if aRes.Err != nil {
			t.Error(aRes.Err)
		}

		if aRes.Resp.GetStatusCode() != 200 {
			t.Error(
				"For", "PATCH http://httpbin.org/patch",
				"expected", 200,
				"got", aRes.Resp.GetStatusCode(),
			)
		}
	}

}

// TestAsyncDeleteRequest tests asynchronous DELETE request
func TestAsyncDeleteRequest(t *testing.T) {
	t.Log("Sending DELETE async request... (expected http code: 200)")

	req := NewRequest()
	ch := make(chan *AsyncResponse)

	for i := 0; i <= 100; i++ {
		req.AsyncDelete("http://httpbin.org/delete", ch)
	}

	for i := 0; i <= 100; i++ {
		aRes := <-ch

		if aRes.Err != nil {
			t.Error(aRes.Err)
		}

		if aRes.Resp.GetStatusCode() != 200 {
			t.Error(
				"For", "DELETE http://httpbin.org/delete",
				"expected", 200,
				"got", aRes.Resp.GetStatusCode(),
			)
		}
	}

}

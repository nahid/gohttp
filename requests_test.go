package gohttp

import "testing"

// TestGetRequest tests GET request
func TestGetRequest(t *testing.T) {
	t.Log("Sending GET request... (expected http code: 200)")

	req := NewRequest()

	resp, err := req.Get("http://httpbin.org/get")

	if err != nil {
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

// TestPostRequest tests POST request
func TestPostRequest(t *testing.T) {
	t.Log("Sending POST request... (expected http code: 200)")

	req := NewRequest()

	resp, err := req.Post("http://httpbin.org/post")

	if err != nil {
		t.Error(err)
	}

	if resp.GetStatusCode() != 200 {
		t.Error(
			"For", "POST http://httpbin.org/post",
			"expected", 200,
			"got", resp.GetStatusCode(),
		)
	}
}

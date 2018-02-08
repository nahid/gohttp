package gohttp

import "testing"

// TestGetRespResponse tests GetResp response
func TestGetRespResponse(t *testing.T) {
	t.Log("(GetResp expected value)")

	req := NewRequest()

	resp, err := req.Query(map[string]string{"q": "hello"}).Get("http://httpbin.org/get")

	if err != nil {
		t.Error(err)
	}

	if resp.GetResp() == nil {
		t.Error(
			"For", "GetResp",
			"expected", "value",
			"got", "nill | err",
		)
	}
}


// TestGetRespResponse tests GetResp response
func TestGetBodyResponse(t *testing.T) {
	t.Log("(GetBody expected value)")

	req := NewRequest()

	resp, err := req.Query(map[string]string{"q": "hello"}).Get("http://httpbin.org/get")

	if err != nil {
		t.Error(err)
	}

	if resp.GetBody() == nil {
		t.Error(
			"For", "GetBody",
			"expected", "value",
			"got", "nill | err",
		)
	}
}


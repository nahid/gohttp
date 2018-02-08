package gohttp

import "testing"

// TestGetRequest tests GET request
func TestGetRequest(t *testing.T) {
	t.Log("Sending GET request... (expected http code: 200)")

	req := NewRequest()

	resp, err := req.Query(map[string]string{"q": "hello"}).Get("http://httpbin.org/get")

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

	resp, err := req.FormData(map[string]string {
		"name": "Nahid",
	}).Post("http://httpbin.org/post")

	if err != nil {
		t.Error(err)
	}


	if resp.GetStatusCode() != 200 {

		t.Error(
			"For", "POST http://httpbin.org/post",
			"expected", "200",
			"got", resp.GetStatusCode(),
		)
	}
}


// TestPutRequest tests PUT request
func TestPutRequest(t *testing.T) {
	t.Log("Sending PUT request... (expected http code: 200)")

	req := NewRequest()

	resp, err := req.FormData(map[string]string {
		"name": "Nahid",
	}).JSON(map[string]interface{}{
		"website": "www.nahid.im",
	}).Put("http://httpbin.org/put")

	if err != nil {
		t.Error(err)
	}


	if resp.GetStatusCode() != 200 {

		t.Error(
			"For", "PUT http://httpbin.org/put",
			"expected", "200",
			"got", resp.GetStatusCode(),
		)
	}
}

// TestPatchRequest tests POST request
func TestPatchRequest(t *testing.T) {
	t.Log("Sending PATCH request... (expected http code: 200)")

	req := NewRequest()

	resp, err := req.FormData(map[string]string {
		"name": "Nahid",
	}).Headers(map[string]string{
		"Custom-Header": "nothing",
	}).Patch("http://httpbin.org/patch")

	if err != nil {
		t.Error(err)
	}


	if resp.GetStatusCode() != 200 {

		t.Error(
			"For", "PATCH http://httpbin.org/patch",
			"expected", "200",
			"got", resp.GetStatusCode(),
		)
	}
}



// TestPatchRequest tests DELETE request
func TestDeleteRequest(t *testing.T) {
	t.Log("Sending DELETE request... (expected http code: 200)")

	req := NewRequest()

	resp, err := req.Body([]byte{

	}).Headers(map[string]string{
		"Custom-Header": "nothing",
	}).Delete("http://httpbin.org/delete")

	if err != nil {
		t.Error(err)
	}


	if resp.GetStatusCode() != 200 {

		t.Error(
			"For", "DELETE http://httpbin.org/delete",
			"expected", "200",
			"got", resp.GetStatusCode(),
		)
	}
}

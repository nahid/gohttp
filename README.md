# gohttp

[![Go Report Card](https://goreportcard.com/badge/github.com/nahid/gohttp)](https://goreportcard.com/report/github.com/nahid/gohttp)

[![Build Status](https://travis-ci.org/nahid/gohttp.svg?branch=master)](https://travis-ci.org/nahid/gohttp)

HTTP client for Go, its also support asynchronous request

## Installation

```
go get github.com/nahid/gohttp
```

### Example

#### `POST https://httpbin.org/post`

```go
package main

import (
	"github.com/nahid/gohttp"
	"fmt"
)

func main() {
	req := gohttp.NewRequest()

	resp, err := req.
		FormData(map[string]string{"name": "Nahid"}).
		Post("https://httpbin.org/post")

	if err != nil {
		panic(err)
	}

	if resp.GetStatusCode() == 200 {
		var resps map[string]interface{}

		_ = resp.GetBodyWithUnmarshal(&resps)
		fmt.Println(resps["form"])
	}
}
```

#### Async Example

```go
package main

import (
	"github.com/nahid/gohttp"
	"fmt"
)

func main() {
	req := gohttp.NewRequest()
	ch := make(chan *gohttp.Response)

	var users [3]string

	users[0] = "nahid"
	users[1] = "shipu"
	users[2] = "sujan"

	for i:=0; i<len(users); i++ {
		req.
		FormData(map[string]string{"user": users[i]}).
		AsyncPost("http://domain.app/send", ch)
	}


	for i:=0; i<len(users); i++ {
		op := <-ch

		fmt.Println(op.GetBodyAsString())
	}
}
```

### Available Method

- `NewRequest(options ...Option)`

#### Request

- `Get(url string)`
- `Post(url string)`
- `Put(url string)`
- `Patch(url string)`
- `Delete(url string)`

#### Async Request

- `AsyncGet(url string, ch chan)`
- `AsyncPost(url string, ch chan)`
- `AsyncPut(url string, ch chan)`
- `AsyncPatch(url string, ch chan)`
- `AsyncDelete(url string, ch chan)`

#### Data Bindings

- `Headers(data map[string]string)`
- `FormData(data map[string]string)`
- `Json(data map[string]interface{})`
- `Query(data map[string]string{})`
- `Body(body []byte)`
- `Text(text string)`
- `BasicAuth(username, password string)`
- `MultipartFormData(data map[string]string{})`
- `Upload(name, file string)`
- `Uploads(files map[string]string{})`


#### Response

- `GetResp()`
- `GetStatusCode()`
- `GetBody()`
- `GetBodyAsByte()`
- `GetBodyAsString()`
- `GetBodyWithUnmarshal(v interface{})`

See API doc https://godoc.org/github.com/nahid/gohttp
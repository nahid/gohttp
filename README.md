# gohttp
HTTP client for Go

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
	req := gohttp.Request{}

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

package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Println("===> Parsed Body")
		fmt.Println(body)
		err := json.Unmarshal([]byte(body), x);
		fmt.Println(x)
		if  err != nil {
			return
		}
	}
}
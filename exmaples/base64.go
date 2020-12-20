package examples

import (
	"encoding/base64"
	"fmt"
)

func basicAuthenticationHTTP() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}

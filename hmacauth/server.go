package hmacauth

import (
	"io"
	"net/http"
)

// Server runs HMAC authentication demonstration server
func Server() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>HMAC Example</title>
		</head>
		<body>
			<p>foo</p>
		</body>
		</html>
	`
	io.WriteString(w, html)
}

package main

import (
	"fmt"
	"net/http"

	"zhuchong6.com/mini-gin/mgin"
)

func main() {
	r := mgin.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":8080")

}

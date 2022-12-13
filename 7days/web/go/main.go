package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yobol/pl-study/7days/web/go/router"
)

func main() {
	engine := router.New()
	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	log.Fatal(engine.Run(":8080"))
}

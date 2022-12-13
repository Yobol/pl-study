package router

import (
	"fmt"
	"net/http"
)

// New is the constructor of router.Engine.
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// HandlerFunc defines the request handler.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implements the interface 'http.Handler'.
type Engine struct {
	router map[string]HandlerFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s-%s", req.Method, req.URL.Path)
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "NOT FOUND: %s\n", req.URL)
	}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(fn http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		fn = middleware(fn)
	}

	return fn
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	fmt.Println(err, s.port)
	if err != nil {
		return err
	}

	return nil
}

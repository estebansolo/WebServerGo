package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Authenticated() Middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")
			if !flag {
				return
			}

			fn(w, r)
		}
	}
}

func Logging() Middleware {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			fn(w, r)
		}
	}
}

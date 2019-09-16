package main

import (
	"log"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)
type middleware func(http.ResponseWriter, *http.Request) bool

func checkMethodMiddleware(method string) middleware {
	return func(w http.ResponseWriter, r *http.Request) bool {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found\n"))
			return false
		}

		return true
	}
}

func logRequestsMiddleware() middleware {
	return func(w http.ResponseWriter, r *http.Request) bool {
		log.Printf("Endpoint hit: [%v] %v", r.Method, r.URL)
		return true
	}
}

func wrapHandler(middlewares []middleware, handler handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			if middleware(w, r) == false {
				return
			}
		}

		handler(w, r)
	}
}

// Wraps the handling function to provide logging and other features
func route(method string, route string, handler handlerFunc) {
	var middlewares []middleware

	// Add multiple middlewares here
	middlewares = append(middlewares, checkMethodMiddleware(method))
	middlewares = append(middlewares, logRequestsMiddleware())

	http.HandleFunc(route, wrapHandler(middlewares, handler))
}

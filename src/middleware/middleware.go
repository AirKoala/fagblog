package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() {
				log.Printf("%s - %s %s %s took %s\n",
					r.RemoteAddr,
					r.Proto,
					r.Method,
					r.URL,
					time.Since(start).String())
			}()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// // Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
// func Method(m string) Middleware {
//
// 	// Create a new Middleware
// 	return func(f http.HandlerFunc) http.HandlerFunc {
//
// 		// Define the http.HandlerFunc
// 		return func(w http.ResponseWriter, r *http.Request) {
//
// 			// Do middleware things
// 			if r.Method != m {
// 				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 				return
// 			}
//
// 			// Call the next middleware/handler in chain
// 			f(w, r)
// 		}
// 	}
// }

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

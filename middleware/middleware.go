package middleware

import "net/http"


type Middleware interface {
	Handle(handler http.HandlerFunc) http.HandlerFunc
}

func Chain(handler http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return handler
	}
	wrapped := handler
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i].Handle(wrapped)
	}
	return wrapped
}

func ChainMiddleware(pattern string, handler http.HandlerFunc, extras ...Middleware) {
	middlewareSlice := []Middleware{}
	middlewareSlice = append(middlewareSlice, extras...)
	chain := Chain(handler, middlewareSlice...)
	http.HandleFunc(pattern, chain)

}


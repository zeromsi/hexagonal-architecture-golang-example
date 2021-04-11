package middleware

import (
	"log"
	"net/http"
)

type TokenValidator struct {
}

func NewTokenValidator() TokenValidator { return TokenValidator{} }

func (t TokenValidator) Handle(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// you can validate your token here
		log.Println("Executing middleware CheckTokenFromHTTPRequest")
		handler.ServeHTTP(w, r)
		log.Println("Executing middleware CheckTokenFromHTTPRequest again")
	})
}

package api

import (
	"hexagonal-architecture-example/api/utility/middleware"
	"net/http"
)

func Routes() {
	v1ProductHandler := NewProductHandler()
	http.HandleFunc("/api/v1/products/", middleware.Chain(v1ProductHandler.ServeProduct,middleware.TokenValidator{}.ValidateToken))
}

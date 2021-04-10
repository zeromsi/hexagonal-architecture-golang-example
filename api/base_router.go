package api

import (
	"net/http"
)

func Routes() {
	v1ProductHandler := NewProductHandler()
	http.HandleFunc("/api/v1/products/", v1ProductHandler.ServeProduct)
	//middleware.ChainMiddleware("api/v1/products/",v1ProductHandler.ServeProduct,middleware.NewTokenValidator())
}

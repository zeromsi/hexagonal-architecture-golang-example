package utility

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"hexagonal-architecture-example/config"
	"net/http"
)

type RequestError struct {
	Error string `json:"error"`
	Desc  string `json:"error_description"`
}
func ShowError(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	if config.Serializer==string(config.JSON){
	ShowJSONError(w,r,status,data)

	}else if config.Serializer==string(config.XML){
		ShowXMLError(w,r,status,data)

	}else if config.Serializer==string(config.MSGPACK){


	}
}
func ShowJSONError(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "An error occurred while processing your request.")
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	fmt.Fprintf(w, string(body))
}
func ShowXMLError(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	body, err := xml.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "An error occurred while processing your request.")
		return
	}

	w.Header().Set("Content-Type", "application/xml;charset=UTF-8")
	w.WriteHeader(status)
	fmt.Fprintf(w, string(body))
}

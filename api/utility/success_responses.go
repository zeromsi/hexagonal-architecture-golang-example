package utility

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"hexagonal-architecture-example/config"
	"net/http"
)

func ShowSuccess(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	if config.Serializer==string(config.JSON){
		ShowJSONSuccess(w,r,status,data)

	}else if config.Serializer==string(config.XML){
		ShowXMLSuccess(w,r,status,data)

	}else if config.Serializer==string(config.MSGPACK){


	}
}

func ShowJSONSuccess(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(200)
		fmt.Fprintln(w, "Operation Successful")
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	fmt.Fprintf(w, string(body))
}
func ShowXMLSuccess(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	body, err := xml.Marshal(data)
	if err != nil {
		w.WriteHeader(200)
		fmt.Fprintln(w, "Operation Successful")
		return
	}

	w.Header().Set("Content-Type", "application/xml;charset=UTF-8")
	w.WriteHeader(status)
	fmt.Fprintf(w, string(body))
}

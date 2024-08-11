package resp

import (
	"log"
	"net/http"
)

type ErrorResp struct {
	Error string `json:"error"`
}

func WithError(w http.ResponseWriter, statusCode int, errorMsg string) {
	if statusCode > 499 {
		log.Println("Responding with 5XX error:", errorMsg)
	}
	
	WithJSON(w, statusCode, ErrorResp{errorMsg})
}

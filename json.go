package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONRespond is a data encoder for json respond
type JSONRespond struct {
	*json.Encoder
}

// Encode encodes the data to the JSON and write it to ResponseWriter
func (j *JSONRespond) Encode(w http.ResponseWriter, data interface{}) error {
	j.Encoder = json.NewEncoder(w)
	return j.Encoder.Encode(data)
}

// HandleJSON write JSON representation of data to the ResponseWriter
func HandleJSON(statusCode int, data interface{}) http.Handler {
	jsonEncoder := &JSONRespond{}
	errorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"msg":"Internal Server Error"}`)
	})
	return HandleData(jsonEncoder, errorHandler, data)
}

package helpers

import (
	"encoding/json"
	"net/http"
)

type Codec interface {
	Encode(w http.ResponseWriter, v interface{}) error
	Decode(r *http.Request, v interface{}) error
}

type codec struct{}

func (codec) Encode(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

func (codec) Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func NewCodec() Codec {
	return &codec{}
}

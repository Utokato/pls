package resp

import (
	"encoding/json"
	"net/http"
)

const contentTypeJson = "application/json"

func Write(w http.ResponseWriter, r Result) {
	w.Header().Set("content-type", contentTypeJson)
	_ = json.NewEncoder(w).Encode(r)
}

package api

import (
	"encoding/json"
	"net/http"
)

// WriteJSON takes any data, creates a map [key:data], marshals this map into json, and writes it to w
func WriteJSON(w http.ResponseWriter, status int, data any, key string) {
	w.WriteHeader(status)
	m := map[string]any{key: data}
	bs, err := json.Marshal(m)
	if err != nil {
		http.Error(w, "error marshaling the response into json", http.StatusInternalServerError)
		return
	}
	w.Write(bs)
}

// Helper functions for typical use cases
func WriteError(w http.ResponseWriter, status int, errMsg string) {
	WriteJSON(w, status, errMsg, "error")
}

func WriteInternalServerError(w http.ResponseWriter, errMsg string) {
	WriteError(w, http.StatusInternalServerError, errMsg)
}

func WriteBadRequestError(w http.ResponseWriter, errMsg string) {
	WriteError(w, http.StatusBadRequest, errMsg)
}

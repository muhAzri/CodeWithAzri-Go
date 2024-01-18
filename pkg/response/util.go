package response

import (
	"encoding/json"
	"net/http"
)

func Respond(code int, metaData Meta, payload interface{}, w http.ResponseWriter) {
	response := &Response{
		Meta: metaData,
		Data: payload,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle encoding error if needed
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func RespondError(code int, err error, w http.ResponseWriter) {
	response := &Response{
		Error: map[string]string{"error": err.Error()},
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle encoding error if needed
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func RespondErrorMessage(code int, msg string, w http.ResponseWriter) {
	response := &Response{
		Error: map[string]string{"error": msg},
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle encoding error if needed
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

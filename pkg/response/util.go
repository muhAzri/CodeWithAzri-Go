package response

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(code int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func BuildResponse(code int, metaMessage string, metaStatus string, payload interface{}, w http.ResponseWriter) {
	response := Response{
		Meta: Meta{
			Message: metaMessage,
			Code:    code,
			Status:  metaStatus,
		},
		Data: payload,
	}

	respondWithJSON(code, response, w)
}

func Respond(code int, metaData Meta, payload interface{}, w http.ResponseWriter) {

	response := &Response{
		Meta: metaData,
		Data: payload,
	}

	respondWithJSON(code, response, w)
}

func RespondError(code int, err error, w http.ResponseWriter) {
	response := &ResponseError{
		Error: map[string]string{"error": err.Error()},
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
	}

	respondWithJSON(code, response, w)
}

func RespondErrorMessage(code int, msg string, w http.ResponseWriter) {
	response := &ResponseError{
		Error: map[string]string{"error": msg},
		Meta: Meta{
			Message: "Error",
			Code:    code,
			Status:  "error",
		},
	}

	respondWithJSON(code, response, w)
}

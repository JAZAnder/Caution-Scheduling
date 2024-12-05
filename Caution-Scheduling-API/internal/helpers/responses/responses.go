package responses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	logger.Log(4, "HTTP Response", strconv.Itoa(code), "responseMaster", message,  )
	RespondWithJSONNoLog(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	logger.Log(1, "HTTP Response", strconv.Itoa(code), "responseMaster", string(response),  )
}

func RespondWithJSONNoLog(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

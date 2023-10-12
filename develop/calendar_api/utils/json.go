package utils

import (
	"encoding/json"
	"net/http"
)

// Отправка JSON-ответа
func SendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Не удалось закодировать JSON", http.StatusInternalServerError)
	}
}

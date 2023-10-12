package handlers

import (
	"calendar_api/data_store"
	"calendar_api/utils"
	"net/http"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг и валидация параметров запроса
	newEvent, err := utils.ValidateCreateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Добавление события в хранилище
	dataStore.CreateEvent(newEvent)

	// Возвращение JSON-ответа
	response := map[string]string{"result": "Событие успешно создано"}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

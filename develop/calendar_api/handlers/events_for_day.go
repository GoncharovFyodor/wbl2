package handlers

import (
	"calendar_api/data_store"
	"calendar_api/utils"
	"net/http"
)

func EventsForDayHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получить дату из параметров запроса
	date := r.FormValue("date")
	// Валидация параметров
	if date == "" {
		http.Error(w, "Отсутствует параметр даты", http.StatusBadRequest)
		return
	}

	// Получение событий за указанный день
	events := dataStore.GetEventsForDay(date)

	// Возвращение JSON-ответа с найденными событиями
	utils.SendJSONResponse(w, http.StatusOK, events)
}

package handlers

import (
	"calendar_api/data_store"
	"calendar_api/utils"
	"net/http"
)

func EventsForWeekHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получить начальную и конечную дату из параметров запроса
	startDate := r.FormValue("start_date")
	if startDate == "" {
		http.Error(w, "Отсутствует параметр начальной даты", http.StatusBadRequest)
		return
	}
	endDate := r.FormValue("end_date")
	if endDate == "" {
		http.Error(w, "Отсутствует параметр конечной даты", http.StatusBadRequest)
		return
	}

	// Получение событий в указанном диапазоне недель
	events := dataStore.GetEventsForWeek(startDate, endDate)

	// Возвращение JSON-ответа с найденными событиями
	utils.SendJSONResponse(w, http.StatusOK, events)
}

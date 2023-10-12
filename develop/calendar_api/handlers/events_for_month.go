package handlers

import (
	"calendar_api/data_store"
	"calendar_api/utils"
	"net/http"
)

func EventsForMonthHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получить год и месяц из параметров запроса
	year := r.FormValue("year")
	if year == "" {
		http.Error(w, "Отсутствует параметр года", http.StatusBadRequest)
		return
	}
	month := r.FormValue("month")
	if month == "" {
		http.Error(w, "Отсутствует параметр месяца", http.StatusBadRequest)
		return
	}

	// Получение событий в указанном месяце и году
	events := dataStore.GetEventsForMonth(month, year)

	// Возвращение JSON-ответа с найденными событиями
	utils.SendJSONResponse(w, http.StatusOK, events)
}

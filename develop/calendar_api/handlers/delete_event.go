package handlers

import (
	"calendar_api/data_store"
	"calendar_api/utils"
	"net/http"
	"strconv"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, dataStore *data_store.DataStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получить ID обновляемого события
	eventID := r.FormValue("id")
	// Валидация параметров
	if eventID == "" {
		http.Error(w, "Отсутствует идентификатор события", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "Неверный формат идентификатора события", http.StatusBadRequest)
		return
	}

	// Удаление события
	err = dataStore.DeleteEvent(id)
	if err != nil {
		http.Error(w, "Не удалось удалить событие", http.StatusInternalServerError)
		return
	}

	// Возвращение JSON-ответа
	response := map[string]string{"result": "Событие успешно удалено"}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

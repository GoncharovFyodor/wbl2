package handlers

import (
	"calendar_api/data_store"
	"calendar_api/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateEventHandler(t *testing.T) {

	// Фейковый объект dataStore для теста
	dataStore := data_store.NewDataStore()

	// Фейковый HTTP-запрос для теста
	requestBody := "title=Какая-то встреча&date=2023-12-01&details=Никому нет дела"
	req := httptest.NewRequest("POST", "/create_event", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Вызов функции CreateEventHandler с фейковыми параметрами
	CreateEventHandler(w, req, dataStore)

	// Проверка ответа на соответствие ожидаемому
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
	}
}

func TestUpdateEventHandler(t *testing.T) {

	// Фейковый объект dataStore для теста
	dataStore := data_store.NewDataStore()

	// Фейковое событие для теста
	event := domain.Event{
		ID:      1,
		Title:   "Test event",
		Date:    "2023-12-08",
		Details: "Test event details",
	}
	dataStore.CreateEvent(event)

	//Фейковый HTTP-запрос для теста
	requestBody := "id=1&title=Changed event&date=2023-12-01&details=Changed event details"
	req := httptest.NewRequest("POST", "/update_event", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Вызов функции UpdateEventHandler с фейковыми параметрами
	UpdateEventHandler(w, req, dataStore)

	// Проверка ответа на соответствие ожидаемому
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
	}
}

func TestDeleteEventHandler(t *testing.T) {

	// Фейковый объект dataStore для теста
	dataStore := data_store.NewDataStore()

	// Фейковое событие для теста
	event := domain.Event{
		ID:      1,
		Title:   "Test event",
		Date:    "2023-12-08",
		Details: "Test event details",
	}
	dataStore.CreateEvent(event)

	//Фейковый HTTP-запрос для теста
	requestBody := "id=1"
	req := httptest.NewRequest("POST", "/delete_event", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Вызов функции DeleteEventHandler с фейковыми параметрами
	DeleteEventHandler(w, req, dataStore)

	// Проверка ответа на соответствие ожидаемому
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
	}
}

func TestEventsForDayHandler(t *testing.T) {

	// Фейковый объект dataStore для теста
	dataStore := data_store.NewDataStore()

	// Фейковое событие для теста
	event := domain.Event{
		ID:      1,
		Title:   "Test event",
		Date:    "2023-12-08",
		Details: "Test event details",
	}
	dataStore.CreateEvent(event)

	//Фейковый HTTP-запрос для теста
	req := httptest.NewRequest("GET", "/events_for_day?date=2023-12-08", nil)

	// Фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Вызов функции EventsForDayHandler с фейковыми параметрами
	EventsForDayHandler(w, req, dataStore)

	// Проверка ответа на соответствие ожидаемому
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
	}
}

func TestEventsForWeekHandler(t *testing.T) {

	// Фейковый объект dataStore для теста
	dataStore := data_store.NewDataStore()

	// Фейковое событие для теста
	event := domain.Event{
		ID:      1,
		Title:   "Test event",
		Date:    "2023-12-08",
		Details: "Test event details",
	}
	dataStore.CreateEvent(event)

	//Фейковый HTTP-запрос для теста
	req := httptest.NewRequest("GET", "/events_for_day?start_date=2023-12-07&end_date=start_date=2023-12-17", nil)

	// Фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Вызов функции EventsForWeekHandler с фейковыми параметрами
	EventsForWeekHandler(w, req, dataStore)

	// Проверка ответа на соответствие ожидаемому
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
	}
}

func TestEventsForMonthHandler(t *testing.T) {

	// Фейковый объект dataStore для теста
	dataStore := data_store.NewDataStore()

	// Фейковое событие для теста
	event := domain.Event{
		ID:      1,
		Title:   "Test event",
		Date:    "2023-12-08",
		Details: "Test event details",
	}
	dataStore.CreateEvent(event)

	//Фейковый HTTP-запрос для теста
	req := httptest.NewRequest("GET", "/events_for_day?year=2023&month=12", nil)

	// Фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Вызов функции EventsForMonthHandler с фейковыми параметрами
	EventsForMonthHandler(w, req, dataStore)

	// Проверка ответа на соответствие ожидаемому
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус код %d, получен %d", http.StatusOK, w.Code)
	}
}

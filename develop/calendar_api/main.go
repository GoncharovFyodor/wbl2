package main

import (
	"calendar_api/data_store"
	"calendar_api/handlers"
	"log"
	"net/http"
)

func main() {
	dataStore := data_store.NewDataStore()

	// Регистрация хендлеров с передачей им dataStore
	http.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateEventHandler(w, r, dataStore)
	})
	http.HandleFunc("/update_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateEventHandler(w, r, dataStore)
	})
	http.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteEventHandler(w, r, dataStore)
	})
	http.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		handlers.EventsForDayHandler(w, r, dataStore)
	})
	http.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		handlers.EventsForWeekHandler(w, r, dataStore)
	})
	http.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		handlers.EventsForMonthHandler(w, r, dataStore)
	})

	port := ":8080"
	log.Printf("Сервер слушает порт %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}
}

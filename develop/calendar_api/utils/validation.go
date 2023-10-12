package utils

import (
	"calendar_api/domain"
	"fmt"
	"net/http"
	"time"
)

// Валидация параметров создания событий
func ValidateCreateEventParams(r *http.Request) (domain.Event, error) {
	var newEvent domain.Event
	title := r.FormValue("title")
	date := r.FormValue("date")
	details := r.FormValue("details")

	// Валидация параметров
	if title == "" || date == "" || details == "" {
		return newEvent, fmt.Errorf("Отсутствуют обязательные параметры")
	}

	// Проверка формата даты
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return newEvent, fmt.Errorf("Формат даты не соответствует требуемому")
	}

	newEvent = domain.Event{
		Title:   title,
		Date:    date,
		Details: details,
	}

	return newEvent, nil
}

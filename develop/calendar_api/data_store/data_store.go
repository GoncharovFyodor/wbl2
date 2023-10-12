package data_store

import (
	"calendar_api/domain"
	"errors"
	"sync"
)

// Хранилище событий
type DataStore struct {
	events []domain.Event
	mu     sync.Mutex
}

// Создание нового хранилища данных
func NewDataStore() *DataStore {
	return &DataStore{events: make([]domain.Event, 0)}
}

// Создание нового события и его добавление в хранилище
func (ds *DataStore) CreateEvent(event domain.Event) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	// Присвоение уникального ID
	event.ID = len(ds.events) + 1

	// Добавление события в хранилище
	ds.events = append(ds.events, event)
}

// Получение события по его ID
func (ds *DataStore) GetEventByID(id int) (domain.Event, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, event := range ds.events {
		if event.ID == id {
			return event, nil
		}
	}

	return domain.Event{}, errors.New("Событие не найдено")
}

// Обновление события по его ID
func (ds *DataStore) UpdateEvent(id int, updatedEvent domain.Event) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for i, event := range ds.events {
		if event.ID == id {
			// Обновление события
			ds.events[i] = updatedEvent
			return nil
		}
	}

	return errors.New("Событие не найдено")
}

// Удаление события по его ID
func (ds *DataStore) DeleteEvent(id int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for i, event := range ds.events {
		if event.ID == id {
			// Удаление события
			ds.events = append(ds.events[:i], ds.events[i+1:]...)
			return nil
		}
	}

	return errors.New("Событие не найдено")
}

// Получение списка событий на указанный день
func (ds *DataStore) GetEventsForDay(date string) []domain.Event {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	var eventsForDay []domain.Event

	for _, event := range ds.events {
		if event.Date == date {
			// Добавление события к списку событий на указанный день
			eventsForDay = append(eventsForDay, event)
		}
	}

	return eventsForDay
}

// Получение списка событий в указанном диапазоне недель
func (ds *DataStore) GetEventsForWeek(startDate, endDate string) []domain.Event {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	var eventsForWeek []domain.Event

	for _, event := range ds.events {
		if event.Date >= startDate && event.Date <= endDate {
			// Добавление события к списку событий на указанный день
			eventsForWeek = append(eventsForWeek, event)
		}
	}

	return eventsForWeek
}

// Получение списка событий в указанном месяце и году
func (ds *DataStore) GetEventsForMonth(month, year string) []domain.Event {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	var eventsForMonth []domain.Event

	for _, event := range ds.events {
		if event.Date[:7] == year+"-"+month {
			// Добавление события к списку событий на указанный день
			eventsForMonth = append(eventsForMonth, event)
		}
	}

	return eventsForMonth
}

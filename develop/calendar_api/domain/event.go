package domain

// Структура Event для представления ссобытий
type Event struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Details string `json:"details"`
}

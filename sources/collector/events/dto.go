package events

type CreateEventDTO struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Body      string `json:"body"`
	Timestamp string `json:"timestamp"`
}

package todo

import "time"

// Item describes a single todo Item
type Item struct {
	Title string    `json:"title" db:"title"`
	Text  string    `json:"text" db:"text"`
	DueAt time.Time `json:"due_at" db:"due_at"`
}

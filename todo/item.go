package todo

import "time"

// Item describes a single todo Item
type Item struct {
	Title string    `binding:"required" json:"title" db:"title"`
	Text  string    `binding:"required" json:"text" db:"text"`
	DueAt time.Time `binding:"required" json:"due_at" db:"due_at"`
}

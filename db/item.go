package db

import (
	"context"
	"time"

	"github.com/mkock/mytodo/todo"
)

// CreateItem creates the provided item, and returns an error if that failed
func CreateItem(ctx context.Context, i todo.Item) error {
	tx := _db.MustBegin()
	_, err := tx.NamedExecContext(ctx, "INSERT INTO items (title, text, due_at) VALUES (:title, :text, :due_at)", &i)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// AllItems returns all todo items, sorted by due date
func AllItems(ctx context.Context) ([]todo.Item, error) {
	items := []todo.Item{}
	err := _db.SelectContext(ctx, &items, "SELECT title, text, due_at FROM items ORDER BY due_at DESC")
	if err != nil {
		return nil, err
	}
	return items, nil
}

// ItemsByDate returns all todo items that have DueAt within the given day
func ItemsByDate(ctx context.Context, day time.Time) ([]todo.Item, error) {
	items := []todo.Item{}
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := start.Add(24 * time.Hour)
	err := _db.SelectContext(ctx, &items, "SELECT title, text, due_at FROM items WHERE due_at BETWEEN ? AND ?", start, end)
	if err != nil {
		return nil, err
	}
	return items, nil
}

package server

import (
	"net/http"
	"time"

	"github.com/mkock/mytodo/db"
	"github.com/mkock/mytodo/todo"

	"github.com/gin-gonic/gin"
)

func itemText(i todo.Item) string {
	return i.Text + " (" + i.DueAt.Format("2006-01-02") + ")"
}

// handlePing returns a simple response to show that the server is alive
func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// handleTodosToday returns today's todo items, optionally sorted
func handleTodosToday(c *gin.Context) {
	items, err := db.ItemsByDate(c.Request.Context(), time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// handleNewItem creates a new todo item
func handleNewItem(c *gin.Context) {
	var i todo.Item
	err := c.BindJSON(&i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	err = db.CreateItem(c.Request.Context(), i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "todo item created"})
}

// handleGetItems returns all existing todo items, in no particular order
func handleGetItems(c *gin.Context) {
	items, err := db.AllItems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Serve starts a http server and will block until the server is interrupted
func Serve() error {
	r := gin.Default()

	r.GET("/ping", handlePing)
	r.GET("/todos/today", handleTodosToday)
	r.PUT("/todos", handleNewItem)
	r.GET("/todos", handleGetItems)

	err := r.Run()
	return err
}

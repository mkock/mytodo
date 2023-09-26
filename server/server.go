package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// item describes a single todo item
type item struct {
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Date  time.Time `json:"date"`
}

// TODO store should be externalized
var store []item

func init() {
	store = make([]item, 0)
}

func itemText(i item) string {
	return i.Text + " (" + i.Date.Format("2006-01-02") + ")"
}

// handlePing returns a simple response to show that the server is alive
func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// handleTodosToday returns today's todo items, optionally sorted
func handleTodosToday(c *gin.Context) {
	items := make(gin.H)

	currentTime := time.Now()
	startOfToday := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	startOfTomorrow := startOfToday.Add(24 * time.Hour)

	for _, i := range store {
		if i.Date.After(startOfToday) && i.Date.Before(startOfTomorrow) {
			items[i.Title] = itemText(i)
		}
	}
	c.JSON(http.StatusOK, items)
}

func initNewItem(c *gin.Context) {
	var i item
	err := c.ShouldBindWith(&i, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("item", i)
}

func validateNewItem(c *gin.Context) {
	maybeItem, ok := c.Get("item")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "session error"})
		c.Abort()
		return
	}

	i := maybeItem.(item)

	if i.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be empty"})
		c.Abort()
		return
	}
	if i.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text cannot be empty"})
		c.Abort()
		return
	}
	if i.Date.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date cannot be empty"})
		c.Abort()
		return
	}
}

// handleNewItem creates a new todo item
func handleNewItem(c *gin.Context) {
	maybeItem, ok := c.Get("item")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "session error"})
		c.Abort()
		return
	}

	i := maybeItem.(item)

	store = append(store, i)
	c.JSON(http.StatusOK, gin.H{"message": "todo item created"})
}

// handleGetItems returns all existing todo items, in no particular order
func handleGetItems(c *gin.Context) {
	items := make(gin.H)
	for _, i := range store {
		items[i.Title] = itemText(i)
	}
	c.JSON(http.StatusOK, items)
}

// Serve starts a http server and will block until the server is interrupted
func Serve() error {
	r := gin.Default()

	r.GET("/ping", handlePing)
	r.GET("/todos/today", handleTodosToday)
	r.PUT("/todos", initNewItem, validateNewItem, handleNewItem)
	r.GET("/todos", handleGetItems)

	err := r.Run()
	return err
}

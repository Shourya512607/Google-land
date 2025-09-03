package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Message represents the structure of the JSON response.
type Message struct {
	Text string `json:"text"`
}

// GetHello handles the API request to return a greeting.
func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Text: "Hello from the Go backend!"})
}

// PostEcho handles a POST request and echoes the received message.
func PostEcho(c *gin.Context) {
	var receivedMessage Message
	if err := c.ShouldBindJSON(&receivedMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"echo": receivedMessage.Text})
}

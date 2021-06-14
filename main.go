package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var counter int

func main() {
	router := gin.Default()
	router.GET("/api/counter", countHandler)
	router.Run(":4040")
}

func countHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": counter,
	})

	counter++

}

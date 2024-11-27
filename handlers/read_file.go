package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFile(c *gin.Context) {
	var request struct {
		Filename string `json:"filename"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	content, err := readFile(request.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"filename": request.Filename, "content": string(content)})

}

// readFile reads the file from the given filename
func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename) // Read file and return content
}

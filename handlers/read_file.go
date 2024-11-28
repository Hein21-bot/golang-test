package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"fmt"
	"os"
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

func ReadImageFile(c *gin.Context) {

	filename := c.Query("filename")

	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required in the query parameters"})
		return
	}

    content, err := readImageFile(filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file", "details": err.Error()})
        return
    }

    c.Data(http.StatusOK, "image/png", content) // Ensure the MIME type matches the file type
}

func readImageFile(filename string) ([]byte, error) {

    // Construct the full path
    basePath := "./profile/" // Adjust based on where your application runs
    fullPath := filepath.Join(basePath, filename)

    // Check if file exists
    if _, err := os.Stat(fullPath); os.IsNotExist(err) {
        return nil, fmt.Errorf("file does not exist")
    }

    // Read the file
    return ioutil.ReadFile(fullPath)
}


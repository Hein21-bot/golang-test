package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveProfileImage(c *gin.Context) {
	// Parse the uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	// Ensure the folder exists
	profileFolder := "./profile"
	err = os.MkdirAll(profileFolder, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}

	// Validate the file type (only allow JPEG and PNG for this example)
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !contains(allowedExtensions, ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPG and PNG are allowed"})
		return
	}

	// Generate a unique file name to avoid conflicts
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(profileFolder, newFileName)

	// Save the file to the profile folder
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Return the file path in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully",
		"path":    filePath,
	})
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

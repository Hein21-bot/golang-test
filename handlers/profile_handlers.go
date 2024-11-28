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
	// Parse Multipart form to handle multiple file uploads
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files, exists := form.File["images"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images provided"})
		return
	}

	// Ensure the folder exists
	profileFolder := "./profile"
	err = os.MkdirAll(profileFolder, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}

	var uploadedFiles []string
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}

	for _, file := range files {
		// Validate the file type (only allow JPEG and PNG)
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !contains(allowedExtensions, ext) {
			continue // Skip invalid file types
		}

		// Generate a unique file name to avoid conflicts
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(profileFolder, newFileName)

		// Save the file to the profile folder
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			continue // Skip files that fail to save
		}

		// Append the saved file path to the list
		uploadedFiles = append(uploadedFiles, filePath)
	}

	// Return the file paths in the response
	if len(uploadedFiles) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save any images"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Images uploaded successfully",
			"paths":   uploadedFiles,
		})
	}
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

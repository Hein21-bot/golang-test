package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUPC(c *gin.Context) {
	var request struct {
		Country string `json:"name"`
	}
	query := `SELECT upc FROM product group by upc`
	params := map[string]interface{}{
		"country": request.Country,
	}

	// Print the query and parameters for debugging
	res, err := DB.Query(query, params)
	if err != nil || res == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No UPC data found!"})
		return
	}
	// First, assert 'res' to the correct type which seems to be a slice of maps
	resData, ok := res.([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	// Check if there is any data returned
	if len(resData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	// If everything checks out
	c.JSON(http.StatusOK, gin.H{"message": "successful", "Data": res})

}

func GetDataByUPC0(c *gin.Context) {
	var request struct {
		UPC string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	query := `SELECT COUNT() AS count, upc,admin_levels.Country FROM product WHERE upc = $upc  GROUP BY upc,admin_levels.Country;`
	params := map[string]interface{}{
		"upc": request.UPC,
	}

	// Print the query and parameters for debugging
	res, err := DB.Query(query, params)
	if err != nil || res == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No data found!"})
		return
	}
	// First, assert 'res' to the correct type which seems to be a slice of maps
	resData, ok := res.([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	// Check if there is any data returned
	if len(resData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	// If everything checks out
	c.JSON(http.StatusOK, gin.H{"message": "successful", "Data": res})

}

func GetDataByUPC1(c *gin.Context) {
	var request struct {
		UPC     string `json:"name"`
		Country string `json:"country"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	query := `SELECT COUNT() AS count, upc,admin_levels.State, admin_levels.Country FROM product WHERE upc = $upc and admin_levels.Country= $country GROUP BY upc,admin_levels.State, admin_levels.Country;`
	params := map[string]interface{}{
		"upc":     request.UPC,
		"country": request.Country,
	}

	// Print the query and parameters for debugging
	res, err := DB.Query(query, params)
	if err != nil || res == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No data found!"})
		return
	}
	// First, assert 'res' to the correct type which seems to be a slice of maps
	resData, ok := res.([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	// Check if there is any data returned
	if len(resData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	// If everything checks out
	c.JSON(http.StatusOK, gin.H{"message": "successful", "Data": res})

}

func GetDataByUPC2(c *gin.Context) {
	var request struct {
		UPC     string `json:"name"`
		Country string `json:"country"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	query := `SELECT COUNT() AS count, upc,admin_levels.City, admin_levels.Country, admin_levels.State FROM product WHERE upc = $upc  and admin_levels.Country= $country GROUP BY upc,admin_levels.City, admin_levels.Country, admin_levels.State;`
	params := map[string]interface{}{
		"upc":     request.UPC,
		"country": request.Country,
	}

	// Print the query and parameters for debugging
	res, err := DB.Query(query, params)
	if err != nil || res == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No data found!"})
		return
	}
	// First, assert 'res' to the correct type which seems to be a slice of maps
	resData, ok := res.([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	// Check if there is any data returned
	if len(resData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	// If everything checks out
	c.JSON(http.StatusOK, gin.H{"message": "successful", "Data": res})

}

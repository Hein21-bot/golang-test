package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductData(c *gin.Context) {

	query := `select * from product;`
	params := map[string]interface{}{}

	res, err := DB.Query(query, params)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed", "details": err.Error()})
		return
	}

	resData, ok := res.([]interface{})
	if !ok {
		fmt.Println("Unexpected response format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	if len(resData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found for the product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
		"data":    resData,
	})
}



func GetProductDetails(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required in the query parameters"})
		return
	}

	query := `select * from product where id = $id;`
	params := map[string]interface{}{
		"id": id,
	}
	res, err := DB.Query(query, params)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed", "details": err.Error()})
		return
	}

	resData, ok := res.([]interface{})
	if !ok {
		fmt.Println("Unexpected response format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	if len(resData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found for the product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
		"data":    resData,
	})
}

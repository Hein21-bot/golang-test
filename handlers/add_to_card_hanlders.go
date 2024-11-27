package handlers

import (
	"fmt"
	"login_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// func AddToCardCreate(c *gin.Context) {
// 	var newData models.AddToCard
// 	if err := c.BindJSON(&newData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
// 		return
// 	}

// 	query := `CREATE add_to_card SET user_id = $user_id, product_id = $product_id, quantity = $quantity, added_at = $added_at`
// 	params := map[string]interface{}{
// 		"user_id":    newData.User_id,
// 		"product_id": newData.Product_id,
// 		"quantity":   newData.Quantity,
// 		"added_at":   time.Now(),
// 	}

// 	res, err := DB.Query(query, params)
// 	if err != nil {
// 		fmt.Printf("Error executing query: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed", "details": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"message": "Add to card successful", "data": res})
// }

func AddToCartCreate(c *gin.Context) {
	var newData models.AddToCard
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	// Check if user exists
	userCheckQuery := `SELECT * FROM user WHERE id = $user_id`
	userCheckParams := map[string]interface{}{
		"user_id": newData.User_id,
	}
	userRes, err := DB.Query(userCheckQuery, userCheckParams)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if product exists
	productCheckQuery := `SELECT * FROM product WHERE id = $product_id`
	productCheckParams := map[string]interface{}{
		"product_id": newData.Product_id,
	}
	productRes, err := DB.Query(productCheckQuery, productCheckParams)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Insert into add_to_card table
	addToCardQuery := `
		CREATE add_to_card SET 
			user_id = $user_id, 
			product_id = $product_id, 
			quantity = $quantity, 
			added_at = $added_at
	`
	addToCardParams := map[string]interface{}{
		"user_id":    newData.User_id,
		"product_id": newData.Product_id,
		"quantity":   newData.Quantity,
		"added_at":   time.Now(),
	}

	addToCardRes, err := DB.Query(addToCardQuery, addToCardParams)
	if err != nil {
		fmt.Printf("Error creating add_to_card record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to card", "details": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Add to card successful",
		"data": map[string]interface{}{
			"add_to_card": addToCardRes,
			"user":        userRes,
			"product":     productRes,
		},
	})
}

func GetAddToCartData(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required in the query parameters"})
		return
	}

	query := `SELECT *, user_id.*, product_id.* FROM add_to_card WHERE id = $id;`
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
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found for the provided ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
		"data":    resData,
	})
}

func UpdateAddToCart(c *gin.Context) {
	id := c.Query("id")
	var updateData models.AddToCard

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	// Insert into add_to_card table
	addToCardQuery := `update add_to_card set quantity = $quantity where id = $id`
	addToCardParams := map[string]interface{}{
		"id":       id,
		"quantity": updateData.Quantity,
	}

	addToCardRes, err := DB.Query(addToCardQuery, addToCardParams)
	if err != nil {
		fmt.Printf("Error creating add_to_card record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to card", "details": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Add to card successful",
		"data": map[string]interface{}{
			"add_to_card": addToCardRes,
		},
	})
}
func DeleteAddToCart(c *gin.Context) {
	// Get the 'id' from query parameters
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required in query parameters"})
		return
	}

	// Define the delete query
	deleteQuery := `DELETE FROM add_to_card WHERE id = $id`
	deleteParams := map[string]interface{}{
		"id": id,
	}

	// Execute the query
	_, err := DB.Query(deleteQuery, deleteParams)
	if err != nil {
		fmt.Printf("Error deleting add_to_card record: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete the add_to_card record",
			"details": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Record successfully deleted",
		"id":      id,
	})
}

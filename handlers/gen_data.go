package handlers

import (
	"fmt"
	"login_api/models"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
)

// GenerateAndInsertData generates and inserts fake data into the database
func GenerateAndInsertData(c *gin.Context) {
	var requests models.BatchProductRequest
	fmt.Print(requests)

	if err := c.BindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	gofakeit.Seed(0)
	// Generate and insert fake data
	// for i := 0; i < 40; i++ {
	for _, request := range requests {
		fmt.Print(request)
		// product := generateFakeProduct()
		query := `CREATE product SET id = $id, upc = $upc, admin_levels = $admin_levels, time_purchased = $time_purchased, name = $name, price = $price, stock_quantity = $stock_quantity, description = $description ;`
		params := map[string]interface{}{
			"id":             gofakeit.UUID(),
			"upc":            request.UPC,
			"admin_levels":   request.AdminLevels,
			"time_purchased": time.Now(),
			"name":           request.Name,
			"price":          request.Price,
			"stock_quantity": request.StockQuantity,
			"description":    request.Description,
		}

		// Now executing the query with parameters
		_, err := DB.Query(query, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed", "details": err.Error()})
			return
		}

	}
	c.JSON(200, gin.H{"message": "Data generated and inserted successfully"})
}

// generateFakeProduct creates a fake product data structure
func generateFakeProduct() *models.Data {
	// upc := gofakeit.DigitN(12)
	adminLevels := map[string]interface{}{
		"Country": gofakeit.Country(),
		"State":   gofakeit.State(),
		"City":    gofakeit.City(),

		// "Country": "Myanmar",
		// "State":   "Mandalay Region",
		// "City":    "Mandalay",
	}
	timePurchased := []time.Time{time.Now()}

	return &models.Data{
		UPC:           "453982763003",
		AdminLevels:   adminLevels,
		TimePurchased: timePurchased,
	}
}

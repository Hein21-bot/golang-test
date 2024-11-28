package handlers

import (
	"fmt"
	"login_api/models"
	"login_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required in the query parameters"})
		return
	}

	query := `select * from user where id = $id;`
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

func UpdateUserInfo(c *gin.Context) {
	id := c.Query("id")
	var updateData models.User

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	// Hash the password from the newUser object
	hashedPassword, err := utils.HashPassword(updateData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	updateData.Password = hashedPassword

	// Update User Info table
	userInfoQuery := `update user set first_name = $first_name, last_name = $last_name, password = $password, org_name = $org_name, email = $email, phone_number = $phone_number, address = $address, city = $city, state = $state, zip_code = $zip_code, country = $country where id = $id`
	userInfoParams := map[string]interface{}{
		"id":           id,
		"first_name":   updateData.FirstName,
		"last_name":    updateData.LastName,
		"org_name":     updateData.OrgName,
		"email":        updateData.Email,
		"phone_number": updateData.PhoneNo,
		"address":      updateData.Address,
		"city":         updateData.City,
		"state":        updateData.State,
		"zip_code":     updateData.ZipCode,
		"country":      updateData.Country,
		"password":     updateData.Password,
	}

	userInfoRes, err := DB.Query(userInfoQuery, userInfoParams)
	if err != nil {
		fmt.Printf("Error updating user info record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to user info record", "details": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Update user info successful",
		"data": map[string]interface{}{
			"user-info": userInfoRes,
		},
	})
}

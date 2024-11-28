package handlers

import (
	"fmt"
	"login_api/models"
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

	// Update User Info table
	userInfoQuery := `update user set FirstName = $firstname, LastName = $lastname, OrgName = $orgname, email = $email, phoneNo = $phoneNo, address = $address, city = $city, state = $state, zipcode = $zipcode, country = $country where id = $id`
	userInfoParams := map[string]interface{}{
		"id":        id,
		"firstname": updateData.FirstName,
		"lastname":  updateData.LastName,
		"orgname":   updateData.OrgName,
		"email":     updateData.Email,
		"phoneNo":   updateData.PhoneNo,
		"address":   updateData.Address,
		"city":      updateData.City,
		"state":     updateData.State,
		"zipcode":   updateData.ZipCode,
		"country":   updateData.Country,
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

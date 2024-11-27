package handlers

import (
	"login_api/models"
	"login_api/utils"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	// Hash the password from the newUser object
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	query := `CREATE user SET FirstName = $firstname, LastName = $lastname, OrgName = $orgname, Email = $email, PhoneNo = $phoneno, Address = $address, City = $city, State = $state, ZipCode = $zipcode, Country = $country, Password = $password;`
	params := map[string]interface{}{
		"firstname": newUser.FirstName,
		"lastname":  newUser.LastName,
		"orgname":   newUser.OrgName,
		"email":     newUser.Email,
		"phoneno":   newUser.PhoneNo,
		"address":   newUser.Address,
		"city":      newUser.City,
		"state":     newUser.State,
		"zipcode":   newUser.ZipCode,
		"country":   newUser.Country,
		"password":  newUser.Password,
	}

	// Now executing the query with parameters
	res, err := DB.Query(query, params)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed", "details": err.Error()})
		return
	}
	fmt.Printf("Executing query with username: %s and password: %s", newUser.FirstName, newUser.Password)
	c.JSON(http.StatusCreated, gin.H{"message": "User successfully created", "user": res})
}

func Login(c *gin.Context) {
	var loginDetails models.User
	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	query := `SELECT * FROM user WHERE FirstName = $firstname and LastName = $lastname ;`
	params := map[string]interface{}{
		"firstname": loginDetails.FirstName,
		"lastname":  loginDetails.LastName,
	}

	// Execute the query
	res, err := DB.Query(query, params)
	if err != nil || res == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
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

	// Access the first element of the slice, which we expect to be a map
	dataMap, ok := resData[0].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data format error"})
		return
	}

	// Now, access the 'result' key which should contain the user data
	resultData, ok := dataMap["result"].([]interface{})
	if !ok || len(resultData) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Result format error or empty result"})
		return
	}

	// Assume the user data is in the first element of the resultData slice
	userData, ok := resultData[0].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User data format error"})
		return
	}

	hashedPassword := userData["Password"].(string)

	// Example of checking the password (after you've retrieved the actual password to compare with)
	if !utils.CheckPasswordHash(loginDetails.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}
	fmt.Printf("Login success!", userData)

	// If everything checks out
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": userData})

}

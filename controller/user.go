package controller

import (
	"magantas/catalyst/model"
	"time"
	"github.com/gin-gonic/gin"
)

func (server Server) PostUser(c *gin.Context) {

	var user model.Users
	err := c.Bind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bind error , err = " + err.Error()})
		return
	}
	user.CreatedAT = time.Now()

	if user.EmailID != "" && user.Password != "" {

		err = server.Store.CreateUser(&user)
		if err != nil {
			return
		}
		// Display error
		c.JSON(201, gin.H{"success": user})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func (server Server) GetUsers(c *gin.Context) {

	var users []model.Users

	// SELECT * FROM users
	err := server.Store.GetUsers(&users)
	if err != nil {
		return
	}
	// Display JSON result
	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

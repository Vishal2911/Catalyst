package controller

import (
	"magantas/catalyst/model"

	"github.com/gin-gonic/gin"
)

func (server Server) Validateuser(c *gin.Context) (bool,  *model.Users, error ) {

	var user model.Users
	err := c.Bind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bind error , err = " + err.Error()})
	}
	dbuser, err := server.Store.GetUserByEmailAndPassword(&user)
	if err != nil {
		return false,nil,err
	} else if dbuser.EmailID != "" && dbuser.Password != "" {
		return true,dbuser,nil 
	}
	return false, nil , nil
}

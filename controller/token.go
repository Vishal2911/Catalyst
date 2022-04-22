package controller

import (
	
	"magantas/catalyst/model"
	"magantas/catalyst/util"
	"time"

	"github.com/gin-gonic/gin"
)

func (server Server) GenerateToken(c *gin.Context) {

	valid  ,dbuser ,err := server.Validateuser(c)
	if err != nil {
	 	c.JSON(200, "error in validating , err = "+err.Error())
	}
	if valid {
		found := true
		token := model.TokenSchema{
			CreatedBy: dbuser.EmailID,
			CreatedAT: time.Now(),
			Active: true,
			Used: false,
			Token: util.RandStringRunes(),
	
		}
		for found {
			_ , found = server.Store.GetToken(token.Token)
			token.Token = util.RandStringRunes()
		}
	
		server.Store.SaveToken(token)
	
		c.JSON(200, token)
	}else{
		c.JSON(500, "Error while generating token" )
	}

}

func (server Server) RevokeToken(c *gin.Context) {

	valid  ,_ ,err := server.Validateuser(c)
	if err != nil {
	 	c.JSON(200, "error in validating , err = "+err.Error())
	}
	if valid {
		token , _ := c.GetQuery("token")

		ok := server.Store.RevokeToken(token)
		if ok {
			c.JSON(200, "success")
		}
		
	}else {
		c.JSON(500, "Internal Error" )
	}

}


func (server Server) ValidateToken(c *gin.Context) {

		token , _ := c.GetQuery("token")
		daypassed, tokenpresent , tokenused := server.Store.DayPassedSinceTokenCreated(token)
		if tokenpresent && !tokenused && daypassed <= 7 {
			server.Store.SetTokenUsed(token)
			c.JSON(200, "Token Validated Successfully")
			
		}else{
			c.JSON(200, "Try Again")
		}

}




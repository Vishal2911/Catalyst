package controller

import (
	"github.com/gin-gonic/gin"
)

func (server Server) DashBoard(c *gin.Context) {

	dashboard:= server.Store.DashBoard()
	c.JSON(200, dashboard)

}



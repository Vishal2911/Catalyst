package api

import (
	"magantas/catalyst/controller"

	"github.com/gin-gonic/gin"
)

type App struct {
	Controller controller.Server
}

func RunApp() {
	r := gin.Default()
	r.Use(Cors())

	server := controller.NewServer()

	v1 := r.Group("api/v1")
	{
		v1.POST("/createadmin", server.PostUser)
		v1.GET("/admins", server.GetUsers)

		v1.GET("/dashboard", server.DashBoard)

		v1.POST("/generatetoken", server.GenerateToken)
		v1.POST("/revoketoken", server.RevokeToken)
		v1.GET("validatetoken", server.ValidateToken)
	}

	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

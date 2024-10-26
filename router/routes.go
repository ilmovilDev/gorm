package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ilmovilDev/gorm/handlers"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.PATCH("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

}

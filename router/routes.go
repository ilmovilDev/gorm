package router

import (
	"github.com/gin-gonic/gin"
	handlerOpening "github.com/ilmovilDev/gorm/handlers/opening"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handlerOpening.CreateOpeningHandler)
		v1.GET("/openings", handlerOpening.ListOpeningsHandler)
		v1.GET("/opening", handlerOpening.ShowOpeningHandler)
		v1.DELETE("/opening", handlerOpening.DeleteOpeningHandler)
		v1.PATCH("/opening", handlerOpening.UpdateOpeningHandler)
	}

}

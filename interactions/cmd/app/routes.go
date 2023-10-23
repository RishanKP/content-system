package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *Application) Routes() *gin.Engine {

	r := gin.Default()

	router := r.Group("/interactions")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/:id", app.GetInteractionsByStoryID)
	router.POST("/", app.Save)
	router.GET("/analytics/top", app.GetTopContents)

	return r

}

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *Application) Routes() *gin.Engine {

	r := gin.Default()

	router := r.Group("/contents")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", app.All)
	router.GET("/:id", app.GetContentById)
	router.POST("/", app.Create)
	router.PUT("/:id", app.Update)
	router.DELETE("/:id", app.Delete)

	router.GET("/top", app.GetTopContents)
	router.GET("/new", app.Latest)
	router.POST("/uploadByCSV", app.BulkUpload)

	return r

}

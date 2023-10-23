package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *Application) Routes() *gin.Engine {

	r := gin.Default()

	router := r.Group("/users")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", app.All)
	router.GET("/:id", app.GetUserById)
	router.POST("/", app.Create)
	router.PUT("/:id", app.Update)
	router.DELETE("/:id", app.Delete)

	return r

}

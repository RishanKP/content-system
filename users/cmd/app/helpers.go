package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func (app *Application) ServerError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "failed",
		"error":  err.Error(),
	})
}

func (app *Application) ClientError(c *gin.Context, status int) {
	c.JSON(status, gin.H{
		"status":  "failed",
		"message": "bad request",
	})
}

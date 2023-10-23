package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func apiResponse(c *gin.Context, body interface{}, status int) {
	stringBody, _ := json.Marshal(body)

	c.JSON(status, stringBody)
}

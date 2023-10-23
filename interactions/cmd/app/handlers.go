package main

import (
	"content-system/interactions/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInteractions godoc
// @Summary Get Interactions
// @Description endpoint to get interactions for a partiular story by Id
// @Tags endpoints
// @Produce application/json
// @Success 200 {object} models.GetInteractionsByStoryID{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [get]
func (app *Application) GetInteractionsByStoryID(c *gin.Context) {
	interactions, err := app.Interactions.GetInteractionsByStoryID(c.Param("id"))
	if err != nil {
		app.ServerError(c, err)

		return

	}

	app.InfoLog.Println("interactions fetched successfully")

	c.JSON(200, gin.H{
		"status":       "success",
		"message":      "interactions fetched",
		"interactions": interactions,
	})
}

// SaveInteraction godoc
// @Summary Save User Interaction
// @Description endpoint to create a new interaction entry
// @Tags endpoints
// @Produce application/json
// @Param user body models.InteractionsInput true "update user input"
// @Param action query string optional "action type (add/remove); if empty, add by default"
// @Success 200 {object} models.InteractionResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router / [POST]
func (app *Application) Save(c *gin.Context) {

	var u models.InteractionsInput
	c.ShouldBindJSON(&u)

	var user models.User
	url := fmt.Sprintf("http://users:8000/users/%s", u.UserID)
	err := app.getAPIContent(url, &user)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "failed to fetch user data",
			"error":   err,
		})

		return
	}

	if user.User.ID.Hex() == "" || user.User.ID.Hex() == "000000000000000000000000" {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "user not found",
		})

		return
	}

	actionType := c.Query("action")

	_, err = app.Interactions.Save(u, actionType)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("interaction saved successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "action completed",
	})
}

// GetTopContents godoc
// @Summary Get Top Contents
// @Description endpoint to get most viewed/ most liked contents
// @Tags endpoints
// @Produce application/json
// @Param limit query string optional "number of records to be returned"
// @Param param query string true "reads or likes for most viewed or most liked items respectively"
// @Success 200 {object} models.GetTopInteractionsResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /analytics/top [get]
func (app *Application) GetTopContents(c *gin.Context) {
	param := c.Query("param") //reads or likes
	limit := c.Query("limit") //limit

	contents, err := app.Interactions.GetTopContents(param, limit)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	c.JSON(200, gin.H{
		"status":   "success",
		"message":  "data fetched",
		"contents": contents,
	})
}

func (app *Application) getAPIContent(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bodyBytes, data)
	return nil
}

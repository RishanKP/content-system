package main

import (
	"content-system/contents/pkg/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ListContents godoc
// @Summary List all Contents
// @Description endpoint to list all Contents available
// @Tags endpoints
// @Produce application/json
// @Success 200 {object} models.GetAllContentsResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router / [get]
func (app *Application) All(c *gin.Context) {
	contents, err := app.Contents.All()
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("contents fetched successfully")

	c.JSON(200, gin.H{
		"status":   "success",
		"message":  "contents fetched",
		"contents": contents,
	})
}

// ListLatestContents godoc
// @Summary List latest Contents
// @Description endpoint to list contents sorted by published date
// @Tags endpoints
// @Produce application/json
// @Success 200 {object} models.GetAllContentsResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /new [get]
func (app *Application) Latest(c *gin.Context) {
	contents, err := app.Contents.Latest()
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("contents fetched successfully")

	c.JSON(200, gin.H{
		"status":   "success",
		"message":  "contents fetched",
		"contents": contents,
	})
}

// GetContentById godoc
// @Summary Get Content By Id
// @Description endpoint to get particular content by Id
// @Tags endpoints
// @Param id path string true "content id"
// @Produce application/json
// @Success 200 {object} models.GetContentResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [get]
func (app *Application) GetContentById(c *gin.Context) {
	content, err := app.Contents.GetContentById(c.Param("id"))
	if err != nil {
		app.ServerError(c, err)

		return

	}

	app.InfoLog.Println("content fetched successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "content data fetched",
		"content": content,
	})
}

// InsertContent godoc
// @Summary Insert Content
// @Description endpoint to insert new content
// @Tags endpoints
// @Produce application/json
// @Param content body models.InsertContentInput true "create"
// @Success 200 {object} models.InsertContentResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router / [post]
func (app *Application) Create(c *gin.Context) {

	var content models.InsertContentInput
	c.ShouldBindJSON(&content)

	var user models.User
	url := fmt.Sprintf("http://users:8000/users/%s", content.UserID)
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

	insertResult, err := app.Contents.Insert(content)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("content inserted successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "content inserted",
		"id":      insertResult.InsertedID,
	})
}

// UpdateContent godoc
// @Summary Update Content
// @Description endpoint to update a particular content
// @Tags endpoints
// @Produce application/json
// @Param content body models.InsertContentInput true "update"
// @Success 200 {object} models.InsertContentResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [PUT]
func (app *Application) Update(c *gin.Context) {

	var ct models.InsertContentInput
	c.ShouldBindJSON(&ct)

	_, err := app.Contents.Update(c.Param("id"), ct)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("content updated successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "content updated",
		"id":      c.Param("id"),
	})
}

// DeleteContent godoc
// @Summary Delete Content
// @Description endpoint to delete content by Id
// @Tags endpoints
// @Produce application/json
// @Param id body string true "id of content"
// @Produce application/json
// @Success 200 {object} models.InsertContentResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [Delete]
func (app *Application) Delete(c *gin.Context) {

	_, err := app.Contents.Delete(c.Param("id"))
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("content deleted successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "content deleted",
		"id":      c.Param("id"),
	})
}

// BulkUpload godoc
// @Summary Insert Content By CSV
// @Description endpoint to bulk upload contents by csv
// @Tags endpoints
// @Produce application/json
// @Accept mpfd
// @Param file formData file true "csv file"
// @Success 200 {object} models.InsertContentResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /uploadByCSV [post]
func (app *Application) BulkUpload(c *gin.Context) {

	inputFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "success",
			"message": "could not process file",
		})

		return
	}

	file, err := inputFile.Open()
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "success",
			"message": "could not process file",
		})

		return
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "success",
			"message": "could not process file",
		})

		return
	}

	for _, line := range records {
		var fullLine string

		i := 0
		for i = 0; i < len(line)-1; i++ {
			fullLine += line[i] + ","
		}
		fullLine += line[i]

		data := strings.Split(fullLine, "$#") // using $# as delimiter as conventional separators like comma might be a part of Story content
		var content models.InsertContentInput

		content.Title = data[0]
		content.Story = data[1]
		content.UserID = data[2]

		var user models.User
		url := fmt.Sprintf("http://users:8000/users/%s", content.UserID)
		err := app.getAPIContent(url, &user)

		if err != nil {
			//process remaining
			continue
		}

		if user.User.ID.Hex() == "" || user.User.ID.Hex() == "000000000000000000000000" {

			return
		}

		_, err = app.Contents.Insert(content)
		if err != nil {
			//process remaining
			continue
		}

	}

	app.InfoLog.Println("content inserted successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "content inserted",
	})
}

// GetTopContents godoc
// @Summary List top Contents
// @Description endpoint to list top contents (most viewed,most liked)
// @Tags endpoints
// @Produce application/json
// @Success 200 {object} models.GetTopContentsResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /topcontents [get]
func (app *Application) GetTopContents(c *gin.Context) {

	param := c.Query("param")
	if param == "" {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "missing query parameter param. supported values (likes/reads)",
		})

		return
	}

	if param != "reads" && param != "likes" {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "invalid param. supported values (likes/reads)",
		})

		return
	}

	limit := c.Query("limit")

	var tc models.TopContents
	url := fmt.Sprintf("http://interactions:8002/interactions/analytics/top?param=%s&limit=%s", param, limit)
	err := app.getAPIContent(url, &tc)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "failed to fetch analytics data",
			"error":   err,
		})

		return
	}

	topContents, err := app.Contents.GetTopContents(tc)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("contents fetched successfully")

	c.JSON(200, gin.H{
		"status":   "success",
		"message":  "contents fetched",
		"contents": topContents,
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

package main

import (
	"content-system/users/pkg/models"

	"github.com/gin-gonic/gin"
)

// ListUsers godoc
// @Summary List all users
// @Description endpoint to list all users of the system
// @Tags endpoints
// @Produce application/json
// @Success 200 {object} models.GetAllUsersResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router / [get]
func (app *Application) All(c *gin.Context) {
	users, err := app.Users.All()
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("users fetched successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "users fetched",
		"users":   users,
	})
}

// GetUserById godoc
// @Summary Get User
// @Description endpoint to get user details by Id
// @Tags endpoints
// @Param id path string true "user id"
// @Produce application/json
// @Success 200 {object} models.GetUserResponse{}
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [get]
func (app *Application) GetUserById(c *gin.Context) {
	user, err := app.Users.GetUserById(c.Param("id"))
	if err != nil {
		app.ServerError(c, err)

		return

	}

	app.InfoLog.Println("user fetched successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user data fetched",
		"user":    user,
	})
}

// CreateUser godoc
// @Summary Create User
// @Description endpoint to create new user
// @Tags endpoints
// @Produce application/json
// @Param user body models.CreateUserInput true "create"
// @Success 200 {object} models.CreateUserResponse
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router / [post]
func (app *Application) Create(c *gin.Context) {

	var u models.CreateUserInput
	c.ShouldBindJSON(&u)

	insertResult, err := app.Users.Create(u)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("user created  successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user created",
		"id":      insertResult.InsertedID,
	})
}

// UpdateUser godoc
// @Summary Update User
// @Description endpoint to update user details
// @Tags endpoints
// @Produce application/json
// @Param user body models.CreateUserInput true "update user input"
// @Param id body string true "id of user"
// @Success 200 {object} models.CreateUserResponse
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [PUT]
func (app *Application) Update(c *gin.Context) {

	var u models.CreateUserInput
	c.ShouldBindJSON(&u)

	_, err := app.Users.Update(c.Param("id"), u)
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("user updated successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user updated",
		"id":      c.Param("id"),
	})
}

// DeleteUser godoc
// @Summary Delete User
// @Description endpoint to delet user
// @Tags endpoints
// @Produce application/json
// @Param id body string true "id of user"
// @Produce application/json
// @Success 200 {object} models.CreateUserResponse
// @Failure 500 {object} models.InternalServerErrorResponse{}
// @Router /{id} [Delete]
func (app *Application) Delete(c *gin.Context) {

	_, err := app.Users.Delete(c.Param("id"))
	if err != nil {
		app.ServerError(c, err)

		return
	}

	app.InfoLog.Println("user deleted successfully")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user deleted",
		"id":      c.Param("id"),
	})
}

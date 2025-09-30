package controller

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/whatsapp_connect/server/model"
	"github.com/supermario64bit/whatsapp_connect/server/service"
)

type userController struct {
	svc service.UserService
}

type UserController interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
	FindByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

func NewUserController() UserController {
	return &userController{
		svc: service.NewUserService(),
	}
}

func (ctrl *userController) Create(c *gin.Context) {
	var user model.User

	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Body", err))
		return
	}

	new, appErr := ctrl.svc.Create(&user)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	c.JSON(http.StatusAccepted, writeSuccessHttpResponseObj("User created!", "user", new))
}

func (ctrl *userController) Find(c *gin.Context) {
	var filter model.User
	err := c.ShouldBindJSON(&filter)
	if errors.Is(err, io.EOF) {
		filter = model.User{}
	} else {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Body", err))
		return
	}

	set, appErr := ctrl.svc.Find(&filter)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	if set == nil || len(set) == 0 {
		c.JSON(http.StatusOK, writeSuccessHttpResponseObj("No Users Found!", "", nil))
		return
	}

	c.JSON(http.StatusOK, writeSuccessHttpResponseObj("Users Found!", "users", set))
}

func (ctrl *userController) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Params", err))
		return
	}

	user, appErr := ctrl.svc.FindByID(id)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, writeSuccessHttpResponseObj("No users found for the id "+idStr, "", nil))
		return
	}

	c.JSON(http.StatusOK, writeSuccessHttpResponseObj("User Found!", "user", user))
}

func (ctrl *userController) UpdateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Params", err))
		return
	}
	var updates model.User
	err = c.ShouldBindBodyWithJSON(&updates)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Body", err))
		return
	}

	updated, appErr := ctrl.svc.UpdateByID(&updates, id)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	c.JSON(http.StatusAccepted, writeSuccessHttpResponseObj("User updated!", "user", updated))
}

func (ctrl *userController) DeleteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Params", err))
		return
	}

	appErr := ctrl.svc.DeleteByID(id)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	c.JSON(http.StatusOK, writeSuccessHttpResponseObj("User Deleted!", "", nil))
}

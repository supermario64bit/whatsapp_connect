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

type organisationController struct {
	svc service.OrganisationService
}

type OrganisationController interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
	FindByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

func NewOrganisationController() OrganisationController {
	return &organisationController{
		svc: service.NewOrganisationService(),
	}
}

func (ctrl *organisationController) Create(c *gin.Context) {
	var org model.Organisation

	err := c.ShouldBindBodyWithJSON(&org)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Body", err))
		return
	}

	new, appErr := ctrl.svc.Create(&org)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	c.JSON(http.StatusAccepted, writeSuccessHttpResponseObj("Organisation created!", "organisation", new))
}

func (ctrl *organisationController) Find(c *gin.Context) {
	var filter model.Organisation
	err := c.ShouldBindJSON(&filter)
	if errors.Is(err, io.EOF) {
		filter = model.Organisation{}
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
		c.JSON(http.StatusOK, writeSuccessHttpResponseObj("No Organisations Found!", "", nil))
		return
	}

	c.JSON(http.StatusOK, writeSuccessHttpResponseObj("Organisations Found!", "organisations", set))
}

func (ctrl *organisationController) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Params", err))
		return
	}

	org, appErr := ctrl.svc.FindByID(id)
	if appErr != nil {
		c.JSON(appErr.HttpStatus, writeFailedHttpResponseObj(appErr.Message, appErr.Err))
		return
	}

	if org == nil {
		c.JSON(http.StatusNotFound, writeSuccessHttpResponseObj("No organisations found for the id "+idStr, "", nil))
		return
	}

	c.JSON(http.StatusOK, writeSuccessHttpResponseObj("Organisation Found!", "organisation", org))
}

func (ctrl *organisationController) UpdateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, writeFailedHttpResponseObj("Invalid Request Params", err))
		return
	}
	var updates model.Organisation
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

	c.JSON(http.StatusAccepted, writeSuccessHttpResponseObj("Organisation updated!", "organisation", updated))
}

func (ctrl *organisationController) DeleteByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, writeSuccessHttpResponseObj("Organisation Deleted!", "", nil))
}

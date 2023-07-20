package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru2-go/dto"
	"github.com/purawaktra/semeru2-go/functions"
	"github.com/purawaktra/semeru2-go/utils"
	"net/http"
)

type Semeru2RequestHandler struct {
	ctrl Semeru2Controller
}

func CreateSemeru2RequestHandler(ctrl Semeru2Controller) Semeru2RequestHandler {
	return Semeru2RequestHandler{
		ctrl: ctrl,
	}
}

type Semeru2RequestHandlerInterface interface {
	SelectCredentialById(c *gin.Context)
	SelectCredentialByEmail(c *gin.Context)
	InsertCredential(c *gin.Context)
	UpdateCredentialById(c *gin.Context)
	DeleteCredentialById(c *gin.Context)
}

func (rh Semeru2RequestHandler) SelectCredentialById(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectCredentialById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectCredentialById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("SelectCredentialById", requestBody)
	response, err := rh.ctrl.SelectCredentialById(requestBody)
	if err != nil {
		utils.Error(err, "SelectCredentialById", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru2RequestHandler) SelectCredentialByEmail(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectCredentialByEmail", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("SelectCredentialByEmail", requestBody)
	response, err := rh.ctrl.SelectCredentialByEmail(requestBody)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru2RequestHandler) InsertCredential(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "InsertCredential", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("InsertCredential", requestBody)
	response, err := rh.ctrl.InsertCredential(requestBody)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru2RequestHandler) UpdateCredentialById(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "UpdateCredentialById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "UpdateCredentialById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("UpdateCredentialById", requestBody)
	response, err := rh.ctrl.UpdateCredentialById(requestBody)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru2RequestHandler) DeleteCredentialById(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "DeleteCredentialById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "DeleteCredentialById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("DeleteCredentialById", requestBody)
	response, err := rh.ctrl.DeleteCredentialById(requestBody)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

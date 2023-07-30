package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru2-go/dto"
	"github.com/purawaktra/semeru2-go/functions"
	"github.com/purawaktra/semeru2-go/utils"
	"net/http"
	"strconv"
	"time"
)

type Semeru2RequestHandler struct {
	ctrl Semeru2Controller
}

func CreateSemeru2RequestHandler(ctrl Semeru2Controller) Semeru2RequestHandler {
	return Semeru2RequestHandler{
		ctrl: ctrl,
	}
}

func (rh Semeru2RequestHandler) SelectCredentialById(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectCredentialById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectCredentialById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectCredentialById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("SelectCredentialById", header)
	data, err, code := rh.ctrl.SelectCredentialById(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectCredentialById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh Semeru2RequestHandler) SelectCredentialByEmail(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectCredentialByEmail", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("SelectCredentialByEmail", header)
	data, err, code := rh.ctrl.SelectCredentialByEmail(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh Semeru2RequestHandler) InsertCredential(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "InsertCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("InsertCredential", header)
	err, code := rh.ctrl.InsertCredential(bodyBytes)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, nil)
	return
}

func (rh Semeru2RequestHandler) UpdateCredential(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "UpdateCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "UpdateCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "UpdateCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("UpdateCredential", header)
	err, code := rh.ctrl.UpdateCredentialById(bodyBytes)
	if err != nil {
		utils.Error(err, "UpdateCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, nil)
	return
}

func (rh Semeru2RequestHandler) DeleteCredential(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "DeleteCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "DeleteCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "DeleteCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the credential and check err
	utils.Debug("DeleteCredential", header)
	err, code := rh.ctrl.DeleteCredentialById(bodyBytes)
	if err != nil {
		utils.Error(err, "DeleteCredential", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, nil)
	return
}

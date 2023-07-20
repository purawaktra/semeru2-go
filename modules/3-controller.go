package modules

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/purawaktra/semeru2-go/dto"
	"github.com/purawaktra/semeru2-go/utils"
	"strconv"
	"time"
)

type Semeru2Controller struct {
	uc Semeru2Usecase
}

func CreateSemeru2Controller(uc Semeru2Usecase) Semeru2Controller {
	return Semeru2Controller{
		uc: uc,
	}
}

type Semeru2ControllerInterface interface {
	SelectCredentialById(req dto.Request) (any, error)
	SelectCredentialByEmail(req dto.Request) (any, error)
	InsertCredential(req dto.Request) (any, error)
	UpdateCredentialById(req dto.Request) (any, error)
	DeleteCredentialById(req dto.Request) (any, error)
}

func (ctrl Semeru2Controller) SelectCredentialById(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data and check err
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectAccountById", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestCredential
	err = json.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "SelectAccountById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the credential and check err
	credential, err := ctrl.uc.SelectAccountById(requestData.AccountId)
	if err != nil {
		utils.Error(err, "SelectAccountById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectAccountById",
		Data:    credential,
	}

	// create return
	return result, nil
}

func (ctrl Semeru2Controller) SelectCredentialByEmail(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data and check err
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestCredential
	err = json.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the credential and check err
	credential, err := ctrl.uc.SelectCredentialByEmail(requestData.EmailAddress)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectCredentialByEmail",
		Data:    credential,
	}

	// create return
	return result, nil
}

func (ctrl Semeru2Controller) InsertCredential(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data and check err
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "InsertCredential", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestCredential
	err = json.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "InsertCredential", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert request to dto
	query := Credential{
		EmailAddress: requestData.EmailAddress,
		Password:     requestData.Password,
	}

	// call usecase for the credential and check err
	err = ctrl.uc.InsertCredential(query)
	if err != nil {
		utils.Error(err, "InsertCredential", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success InsertCredential",
	}

	// create return
	return result, nil
}

func (ctrl Semeru2Controller) UpdateCredentialById(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data and check err
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestCredential
	err = json.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert request to dto
	query := Credential{
		AccountId:    requestData.AccountId,
		EmailAddress: requestData.EmailAddress,
		Password:     requestData.Password,
	}

	// call usecase for the credential and check err
	err = ctrl.uc.UpdateCredentialById(query)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success UpdateCredentialById",
	}

	// create return
	return result, nil
}

func (ctrl Semeru2Controller) DeleteCredentialById(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data and check err
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestCredential
	err = json.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert request to dto
	query := Credential{
		AccountId: requestData.AccountId,
	}

	// call usecase for the credential and check err
	err = ctrl.uc.DeleteCredentialById(query)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success DeleteCredentialById",
	}

	// create return
	return result, nil
}

package modules

import (
	"encoding/json"
	"github.com/purawaktra/semeru2-go/dto"
	"github.com/purawaktra/semeru2-go/utils"
)

type Semeru2Controller struct {
	uc Semeru2Usecase
}

func CreateSemeru2Controller(uc Semeru2Usecase) Semeru2Controller {
	return Semeru2Controller{
		uc: uc,
	}
}

func (ctrl Semeru2Controller) SelectCredentialById(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCredential
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectCredentialById", requestData)
		return nil, err, "FM"
	}

	// convert request to dto
	query := Credential{
		AccountId: requestData.AccountId,
	}

	// call usecase for the credential and check err
	credential, err, code := ctrl.uc.SelectCredentialById(query)
	if err != nil {
		utils.Error(err, "SelectCredentialById", requestData)
		return nil, err, code
	}

	// create return
	return credential, nil, code
}

func (ctrl Semeru2Controller) SelectCredentialByEmail(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCredential
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", requestData)
		return nil, err, "FM"
	}

	// convert request to dto
	query := Credential{
		EmailAddress: requestData.EmailAddress,
	}

	// call usecase for the credential and check err
	credential, err, code := ctrl.uc.SelectCredentialByEmail(query)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", requestData)
		return nil, err, code
	}

	// create return
	return credential, nil, code
}

func (ctrl Semeru2Controller) InsertCredential(req []byte) (error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCredential
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "InsertCredential", requestData)
		return err, "FM"
	}

	// convert request to dto
	query := Credential{
		EmailAddress: requestData.EmailAddress,
		Password:     requestData.Password,
	}

	// call usecase for the credential and check err
	err, code := ctrl.uc.InsertCredential(query)
	if err != nil {
		utils.Error(err, "InsertCredential", requestData)
		return err, code
	}

	// create return
	return nil, code
}

func (ctrl Semeru2Controller) UpdateCredentialById(req []byte) (error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCredential
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", requestData)
		return err, "FM"
	}

	// convert request to dto
	query := Credential{
		AccountId:    requestData.AccountId,
		EmailAddress: requestData.EmailAddress,
		Password:     requestData.Password,
	}

	// call usecase for the credential and check err
	err, code := ctrl.uc.UpdateCredentialById(query)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", requestData)
		return err, code
	}

	// create return
	return nil, code
}

func (ctrl Semeru2Controller) DeleteCredentialById(req []byte) (error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCredential
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", requestData)
		return err, "FM"
	}

	// convert request to dto
	query := Credential{
		AccountId: requestData.AccountId,
	}

	// call usecase for the credential and check err
	err, code := ctrl.uc.DeleteCredentialById(query)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", requestData)
		return err, code
	}

	// create return
	return nil, code
}

package modules

import (
	"errors"
	"fmt"
	"github.com/purawaktra/semeru2-go/entities"
	"github.com/purawaktra/semeru2-go/functions"
	"github.com/purawaktra/semeru2-go/utils"
)

type Semeru2Usecase struct {
	repo Semeru2Repo
}

func CreateSemeru2Usecase(repo Semeru2Repo) Semeru2Usecase {
	return Semeru2Usecase{
		repo: repo,
	}
}

func (uc Semeru2Usecase) SelectCredentialById(query Credential) (Credential, error, string) {
	// create check input on account id and offset
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "SelectCredentialById", query.AccountId)
		return Credential{}, errors.New("accountId can not be nil, negative or zero"), "FC"
	}

	// convert input to entity
	entity := entities.Credential{AccountId: uint(query.AccountId)}

	// call repo for the account id and check err
	credential, err, code := uc.repo.SelectCredentialById(entity)
	if err != nil {
		utils.Error(err, "SelectCredentialById", query)
		return Credential{}, err, code
	}

	// convert a result to dto and create return
	result := Credential{
		AccountId:    int(credential.AccountId),
		EmailAddress: credential.EmailAddress,
		Password:     credential.Password,
		Salt:         credential.Salt,
	}
	return result, nil, code
}

func (uc Semeru2Usecase) SelectCredentialByEmail(query Credential) (Credential, error, string) {
	// create check input on account email address
	if query.EmailAddress == "" {
		utils.Error(errors.New("emailAddress can not be empty"), "SelectCredentialByEmail", "")
		return Credential{}, errors.New("emailAddress can not be empty"), "FC"
	}

	// convert input to entity
	entity := entities.Credential{EmailAddress: query.EmailAddress}

	// call repo for the account email address and check err
	credential, err, code := uc.repo.SelectCredentialByEmailAddress(entity)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		return Credential{}, err, code
	}

	// convert a result to dto and create return
	result := Credential{
		AccountId:    int(credential.AccountId),
		EmailAddress: credential.EmailAddress,
		Password:     credential.Password,
		Salt:         credential.Salt,
	}
	return result, nil, code
}

func (uc Semeru2Usecase) InsertCredential(query Credential) (error, string) {
	// create check input on credential
	if query.EmailAddress == "" {
		utils.Error(errors.New("emailAddress can not be empty"), "InsertCredential", "")
		return errors.New("emailAddress can not be empty"), "FC"
	}
	if query.Password == "" {
		utils.Error(errors.New("password can not be empty"), "InsertCredential", "")
		return errors.New("password can not be empty"), "FC"
	}

	// generate salt and hash password
	salt := functions.GenerateSalt(utils.CredentialSaltLength)
	hash, err := functions.GenerateSHA1Hash(fmt.Sprintf("%s%s", query.Password, salt))
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		return err, "FH"
	}

	// convert city input to entity
	credential := entities.Credential{
		EmailAddress: query.EmailAddress,
		Password:     hash,
		Salt:         salt,
	}

	// call usecase and check err
	err, code := uc.repo.InsertCredential(credential)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		return err, code
	}

	// create return
	return nil, code
}

func (uc Semeru2Usecase) UpdateCredentialById(query Credential) (error, string) {
	// create check input on credential
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "UpdateCredentialById", query)
		return errors.New("accountId can not be nil, negative or zero"), "FC"
	}
	if query.EmailAddress == "" {
		utils.Error(errors.New("emailAddress can not be empty"), "UpdateCredentialById", "")
		return errors.New("emailAddress can not be empty"), "FC"
	}
	if query.Password == "" {
		utils.Error(errors.New("password can not be empty"), "UpdateCredentialById", "")
		return errors.New("password can not be empty"), "FC"
	}

	// generate salt and hash password
	salt := functions.GenerateSalt(utils.CredentialSaltLength)
	hash, err := functions.GenerateSHA1Hash(fmt.Sprintf("%s%s", query.Password, salt))
	if err != nil {
		utils.Error(err, "UpdateCredentialById", "")
		return err, "FH"
	}

	// convert city input to entity
	credential := entities.Credential{
		AccountId:    uint(query.AccountId),
		EmailAddress: query.EmailAddress,
		Password:     hash,
		Salt:         salt,
	}

	// call usecase and check err
	err, code := uc.repo.UpdateCredentialById(credential)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", "")
		return err, code
	}

	// create return
	return nil, code
}

func (uc Semeru2Usecase) DeleteCredentialById(query Credential) (error, string) {
	// create check input on credential
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "DeleteCredentialById", query)
		return errors.New("accountId can not be nil, negative or zero"), "FH"
	}

	// convert city input to entity
	credential := entities.Credential{
		AccountId: uint(query.AccountId),
	}

	// call usecase and check err
	err, code := uc.repo.DeleteCredentialById(credential)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", "")
		return err, code
	}

	// create return
	return nil, code
}

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

type Semeru2UsecaseInterface interface {
	SelectAccountById(accountId int) (Credential, error)
	SelectCredentialByEmail(emailAddress string) (Credential, error)
	InsertCredential(query Credential) error
	UpdateCredentialById(query Credential) error
	DeleteCredentialById(query Credential) error
}

func (uc Semeru2Usecase) SelectAccountById(accountId int) (Credential, error) {
	// create check input on account id and offset
	if accountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "SelectAccountById", accountId)
		return Credential{}, errors.New("accountId can not be nil, negative or zero")
	}

	// convert input to entity
	query := entities.Credential{AccountId: uint(accountId)}

	// call repo for the account id and check err
	account, err := uc.repo.SelectCredentialById(query)
	if err != nil {
		utils.Error(err, "SelectAccountById", query)
		return Credential{}, err
	}

	// convert a result to dto and create return
	result := Credential{
		AccountId:    int(account.AccountId),
		EmailAddress: account.EmailAddress,
		Password:     account.Password,
		Salt:         account.Salt,
	}
	return result, nil
}

func (uc Semeru2Usecase) SelectCredentialByEmail(emailAddress string) (Credential, error) {
	// create check input on account email address
	if emailAddress == "" {
		utils.Error(errors.New("emailAddress can not be empty"), "SelectCredentialByEmail", "")
		return Credential{}, errors.New("emailAddress can not be empty")
	}

	// convert input to entity
	query := entities.Credential{EmailAddress: emailAddress}

	// call repo for the account email address and check err
	account, err := uc.repo.SelectCredentialByEmailAddress(query)
	if err != nil {
		utils.Error(err, "SelectCredentialByEmail", "")
		return Credential{}, err
	}

	// convert a result to dto and create return
	result := Credential{
		AccountId:    int(account.AccountId),
		EmailAddress: account.EmailAddress,
		Password:     account.Password,
		Salt:         account.Salt,
	}
	return result, nil
}

func (uc Semeru2Usecase) InsertCredential(query Credential) error {
	// create check input on credential
	if query.EmailAddress == "" {
		utils.Error(errors.New("emailAddress can not be empty"), "InsertCredential", "")
		return errors.New("emailAddress can not be empty")
	}
	if query.Password == "" {
		utils.Error(errors.New("password can not be empty"), "InsertCredential", "")
		return errors.New("password can not be empty")
	}

	// generate salt and hash password
	salt := functions.GenerateSalt(utils.CredentialSaltLength)
	hash, err := functions.GenerateSHA1Hash(fmt.Sprintf("%s%s", query.Password, salt))
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		return err
	}

	// convert city input to entity
	credential := entities.Credential{
		EmailAddress: query.EmailAddress,
		Password:     hash,
		Salt:         salt,
	}

	// call usecase and check err
	err = uc.repo.InsertCredential(credential)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		return err
	}

	// create return
	return nil
}

func (uc Semeru2Usecase) UpdateCredentialById(query Credential) error {
	// create check input on credential
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "UpdateCredentialById", query)
		return errors.New("accountId can not be nil, negative or zero")
	}
	if query.EmailAddress == "" {
		utils.Error(errors.New("emailAddress can not be empty"), "UpdateCredentialById", "")
		return errors.New("emailAddress can not be empty")
	}
	if query.Password == "" {
		utils.Error(errors.New("password can not be empty"), "UpdateCredentialById", "")
		return errors.New("password can not be empty")
	}

	// generate salt and hash password
	salt := functions.GenerateSalt(utils.CredentialSaltLength)
	hash, err := functions.GenerateSHA1Hash(fmt.Sprintf("%s%s", query.Password, salt))
	if err != nil {
		utils.Error(err, "UpdateCredentialById", "")
		return err
	}

	// convert city input to entity
	credential := entities.Credential{
		AccountId:    uint(query.AccountId),
		EmailAddress: query.EmailAddress,
		Password:     hash,
		Salt:         salt,
	}

	// call usecase and check err
	err = uc.repo.UpdateCredentialById(credential)
	if err != nil {
		utils.Error(err, "UpdateCredentialById", "")
		return err
	}

	// create return
	return nil
}

func (uc Semeru2Usecase) DeleteCredentialById(query Credential) error {
	// create check input on credential
	if query.AccountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "DeleteCredentialById", query)
		return errors.New("accountId can not be nil, negative or zero")
	}

	// convert city input to entity
	credential := entities.Credential{
		AccountId: uint(query.AccountId),
	}

	// call usecase and check err
	err := uc.repo.DeleteCredentialById(credential)
	if err != nil {
		utils.Error(err, "DeleteCredentialById", "")
		return err
	}

	// create return
	return nil
}

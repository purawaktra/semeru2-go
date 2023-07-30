package modules

import (
	"fmt"
	"github.com/purawaktra/semeru2-go/entities"
	"github.com/purawaktra/semeru2-go/utils"
	"gorm.io/gorm"
)

type Semeru2Repo struct {
	db *gorm.DB
}

func CreateSemeru2Repo(gorm *gorm.DB) Semeru2Repo {
	return Semeru2Repo{
		db: gorm,
	}
}

func (sr Semeru2Repo) SelectCredentialById(query entities.Credential) (entities.Credential, error, string) {
	utils.Debug("SelectCredentialById", query)

	var credential entities.Credential
	tx := sr.db.Raw(
		fmt.Sprint("SELECT account_id, email_address, password, salt FROM credentials WHERE account_id = ", query.AccountId, " LIMIT 1")).Scan(
		&credential)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCredentialById", query)
		return entities.Credential{}, err, "DB"
	}
	return credential, nil, "00"
}

func (sr Semeru2Repo) SelectCredentialByEmailAddress(query entities.Credential) (entities.Credential, error, string) {
	utils.Debug("SelectCredentialByEmailAddress", query)

	var credential entities.Credential
	tx := sr.db.Raw(
		fmt.Sprintf("SELECT account_id, email_address, password, salt FROM credentials WHERE email_address = \"%s\" LIMIT 1", query.EmailAddress)).Scan(
		&credential)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCredentialByEmailAddress", query)
		return entities.Credential{}, err, "DB"
	}
	return credential, nil, "00"
}

func (sr Semeru2Repo) InsertCredential(query entities.Credential) (error, string) {
	utils.Debug("InsertCredential", query)

	tx := sr.db.Exec(
		"INSERT INTO credentials (email_address, password, salt) VALUES (?, ?, ?)",
		query.EmailAddress,
		query.Password,
		query.Salt)
	err := tx.Error
	if err != nil {
		utils.Error(err, "InsertCredential", query)
		return err, "DB"
	}
	return nil, "00"
}

func (sr Semeru2Repo) UpdateCredentialById(query entities.Credential) (error, string) {
	utils.Debug("UpdateCredential", query)

	tx := sr.db.Exec(
		"UPDATE credentials SET email_address = ?, password = ?, salt = ? WHERE account_id = ?",
		query.EmailAddress,
		query.Password,
		query.Salt,
		query.AccountId)
	err := tx.Error
	if err != nil {
		utils.Error(err, "UpdateCredential", query)
		return err, "DB"
	}
	return nil, "00"
}

func (sr Semeru2Repo) DeleteCredentialById(query entities.Credential) (error, string) {
	utils.Debug("DeleteCredentialById", query)

	tx := sr.db.Exec(
		"DELETE FROM credentials WHERE account_id = ?",
		query.AccountId)
	err := tx.Error
	if err != nil {
		utils.Error(err, "DeleteCredentialById", query)
		return err, "DB"
	}
	return nil, "00"
}

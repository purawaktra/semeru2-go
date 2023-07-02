package modules

import (
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

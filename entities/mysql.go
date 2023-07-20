package entities

type Credential struct {
	AccountId    uint   `gorm:"primary_key"`
	EmailAddress string `gorm:"column:email_address"`
	Password     string `gorm:"column:password"`
	Salt         string `gorm:"column:salt"`
}

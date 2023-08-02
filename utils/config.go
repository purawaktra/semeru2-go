package utils

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var (
	AppName        string
	AppHost        string
	AppCode        string
	AppPort        string
	AppEnvironment string
	AppLogLevel    string

	AppAuthUsername string
	AppAuthPassword string

	MysqlHost     string
	MysqlPort     string
	MysqlScheme   string
	MysqlUser     string
	MysqlPassword string

	CredentialSaltLength int
)

func InitConfig() {
	// check app environment on env
	AppEnvironment = os.Getenv("APP_ENV")

	// check for value
	if AppEnvironment == "" {
		// check for config.json
		viper.SetConfigFile(`config.json`)

		// read file
		err := viper.ReadInConfig()
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}

		// check for app env
		AppEnvironment = viper.GetString("app.environment")
	}

	if AppEnvironment == "development" {
		// get variable for app
		AppName = viper.GetString("app.name")
		AppHost = viper.GetString("app.host")
		AppCode = viper.GetString("app.code")
		AppPort = viper.GetString("app.port")
		AppLogLevel = viper.GetString("app.log.level")
		AppAuthUsername = viper.GetString("app.auth.username")
		AppAuthPassword = viper.GetString("app.auth.password")

		// get variable for db
		MysqlHost = viper.GetString("database.mysql.host")
		MysqlPort = viper.GetString("database.mysql.port")
		MysqlScheme = viper.GetString("database.mysql.schema")
		MysqlUser = viper.GetString("database.mysql.user")
		MysqlPassword = viper.GetString("database.mysql.password")

		var err error
		CredentialSaltLength, err = strconv.Atoi(viper.GetString("database.credentials.salt_length"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}

		// create return
		return
	}

	if AppEnvironment == "staging" || AppEnvironment == "production" {
		// get variable for app
		AppName = os.Getenv("APP_NAME")
		AppPort = os.Getenv("APP_PORT")
		AppCode = os.Getenv("APP_CODE")
		AppEnvironment = os.Getenv("APP_ENV")
		AppLogLevel = os.Getenv("APP_LOG_LEVEL")
		AppAuthUsername = os.Getenv("APP_AUTH_USERNAME")
		AppAuthPassword = os.Getenv("APP_AUTH_PASSWORD")

		// get variable for db
		MysqlHost = os.Getenv("MYSQL_HOST")
		MysqlPort = os.Getenv("MYSQL_PORT")
		MysqlScheme = os.Getenv("MYSQL_SCHEME")
		MysqlUser = os.Getenv("MYSQL_USER")
		MysqlPassword = os.Getenv("MYSQL_PASSWORD")

		var err error
		CredentialSaltLength, err = strconv.Atoi(os.Getenv("CREDENTIAL_SALT_LENGTH"))
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}
	}
}

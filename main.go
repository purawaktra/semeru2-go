package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru2-go/modules"
	"github.com/purawaktra/semeru2-go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitLog()

	utils.Log("============= SEMERU2 RUNTIME STARTED =============")
	// create gin engine
	engine := gin.New()

	// create dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		utils.MysqlUser,
		utils.MysqlPassword,
		utils.MysqlHost,
		utils.MysqlPort,
		utils.MysqlScheme)

	// create mysql instance
	gormInstance, err := utils.CreateGorm(dsn)
	if err != nil {
		utils.Error(err, "main", "")
		panic(err)
	}

	// declare architecture
	repository := modules.CreateSemeru2Repo(gormInstance)
	usecase := modules.CreateSemeru2Usecase(repository)
	controller := modules.CreateSemeru2Controller(usecase)
	requestHandler := modules.CreateSemeru2RequestHandler(controller)
	router := modules.CreateSemeru2Router(engine, requestHandler)

	// init routing
	router.Init("/semeru2/api/v1")

	utils.Log("============= SEMERU2 ENGINE STARTING =============")
	utils.Log(fmt.Sprintf("App Address = %s:%s", utils.AppHost, utils.AppPort))

	// init routing to swagger
	swaggerRouter := utils.CreateSwaggerRouter(engine)
	swaggerRouter.InitRouter("/semeru2/api/v1/swagger")

	// start http api engine
	err = engine.Run(fmt.Sprintf("%s:%s", utils.AppHost, utils.AppPort))
	if err != nil {
		utils.Fatal(err, "main", "")
		panic(err)
	}

	utils.Log("============= SEMERU2 ENGINE FAILED =============")
}

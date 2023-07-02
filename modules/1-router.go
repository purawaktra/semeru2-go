package modules

import "github.com/gin-gonic/gin"

type Semeru2Router struct {
	engine *gin.Engine
	rh     Semeru2RequestHandler
}

func CreateSemeru2Router(engine *gin.Engine, rh Semeru2RequestHandler) Semeru2Router {
	return Semeru2Router{
		engine: engine,
		rh:     rh,
	}
}

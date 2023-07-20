package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru2-go/utils"
)

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

func (r Semeru2Router) Init(path string) {
	pathGroup := r.engine.Group(path, utils.CheckBasicAuth)
	pathGroup.POST("/select/credential/id", r.rh.SelectCredentialById)
	pathGroup.POST("/select/credential/email", r.rh.SelectCredentialByEmail)
	pathGroup.POST("/insert/credential", r.rh.InsertCredential)
	pathGroup.POST("/update/credential", r.rh.UpdateCredentialById)
	pathGroup.POST("/delete/credential", r.rh.DeleteCredentialById)
}

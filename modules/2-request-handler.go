package modules

type Semeru2RequestHandler struct {
	ctrl Semeru2Controller
}

func CrateSemeru2RequestHandler(ctrl Semeru2Controller) Semeru2RequestHandler {
	return Semeru2RequestHandler{
		ctrl: ctrl,
	}
}

package modules

type Semeru2Controller struct {
	uc Semeru2Usecase
}

func CreateSemeru2Controller(uc Semeru2Usecase) Semeru2Controller {
	return Semeru2Controller{
		uc: uc,
	}
}

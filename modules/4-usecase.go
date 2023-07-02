package modules

type Semeru2Usecase struct {
	repo Semeru2Repo
}

func CreateSemeru2Usecase(repo Semeru2Repo) Semeru2Usecase {
	return Semeru2Usecase{
		repo: repo,
	}
}

package login

type repoContract interface{}

type service struct {
	repo repoContract
}

func newService(repo repoContract) service {
	return service{repo: repo}
}

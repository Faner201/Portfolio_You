package usecase

type UseCase interface {
	RegisterUser()
	LoginUser()
	CreatePortfolio()
	GetListPorfolio()
	OpenPortfolio()
}

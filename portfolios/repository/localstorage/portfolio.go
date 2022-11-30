package localstorage

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"sync"
)

type portfolioLocalStorage struct {
	portf map[int]*models.Portfolio
	menus map[int]*models.Menu
	mutex *sync.Mutex
}

func NewPortfolioLocalStorage() *portfolioLocalStorage {
	return &portfolioLocalStorage{
		portf: make(map[int]*models.Portfolio),
		menus: make(map[int]*models.Menu),
		mutex: new(sync.Mutex),
	}
}

func (p *portfolioLocalStorage) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User, menu *models.Menu) error {
	portfolio.CreaterName = user.Username
	menu.CreaterName = user.Username

	p.mutex.Lock()
	defer p.mutex.Unlock()
	if portfolio.Name != "" && portfolio.Tags != "" && portfolio.Text != "" && portfolio.URL != "" && menu.ShortText != "" {
		p.menus[menu.ID] = menu
		p.portf[portfolio.ID] = portfolio
		return nil
	}

	return portfolios.ErrCreatePortfolio
}

func (p *portfolioLocalStorage) GetPortfolioByUserName(ctx context.Context, userName string, portfolioID int) (*models.Portfolio, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, portfolio := range p.portf {
		if portfolio.CreaterName == userName && portfolio.ID == portfolioID {
			return portfolio, nil
		}
	}

	return nil, portfolios.ErrGetPortfolioByUserName
}

func (p *portfolioLocalStorage) GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error) {
	menus := make([]*models.Menu, 0)

	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, menu := range p.menus {
		if menu.CreaterName == userName {
			menus = append(menus, menu)
		} else {
			return nil, portfolios.ErrGetListPortfolio
		}
	}

	return menus, nil
}

package localstorage

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"sync"
)

type PortfolioLocalStorage struct {
	portf map[string]*models.Portfolio
	menus map[string]*models.Menu
	mutex *sync.Mutex
}

func NewPortfolioLocalStorage() *PortfolioLocalStorage {
	return &PortfolioLocalStorage{
		portf: make(map[string]*models.Portfolio),
		menus: make(map[string]*models.Menu),
		mutex: new(sync.Mutex),
	}
}

func (p *PortfolioLocalStorage) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User) error {
	portfolio.CreaterUser = user.Username

	p.mutex.Lock()
	defer p.mutex.Unlock()
	if portfolio.Name != "" && portfolio.Struct != nil && portfolio.Photo != nil &&
		portfolio.Text != nil && portfolio.Url != "" && portfolio.CreaterUser != "" {
		p.portf[portfolio.ID] = portfolio
		return nil
	}

	return portfolios.ErrCreatePortfolio
}

func (p *PortfolioLocalStorage) CreateMenuPortfolio(ctx context.Context, user *models.User, menu *models.Menu) error {
	menu.CreaterName = user.Username

	p.mutex.Lock()
	defer p.mutex.Unlock()
	if menu.Name != "" && menu.ShortText != "" {
		p.menus[menu.ID] = menu
		return nil
	}

	return portfolios.ErrCreateMenuPortfolio
}

func (p *PortfolioLocalStorage) GetPortfolioByUserName(ctx context.Context, userName string, portfolioID string) (*models.Portfolio, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, portfolio := range p.portf {
		if portfolio.CreaterUser == userName && portfolio.ID == portfolioID {
			return portfolio, nil
		}
	}

	return nil, portfolios.ErrGetPortfolioByUserName
}

func (p *PortfolioLocalStorage) GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error) {
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

func (p *PortfolioLocalStorage) DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	pf, ex := p.portf[portfolioID]
	if ex && pf.CreaterUser == user.Username {
		delete(p.portf, portfolioID)
		return nil
	}

	return portfolios.ErrDeletePortfolio
}

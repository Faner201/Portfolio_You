package portfolios

import "errors"

var ErrCreatePortfolio = errors.New("Portfolio was not created")
var ErrGetPortfolioByUserName = errors.New("This portfolio does not exist ")
var ErrGetListPortfolio = errors.New("There are no portfolios in your office")

package portfolios

import "errors"

var ErrCreatePortfolio = errors.New("Portfolio was not created")
var ErrGetPortfolioByUserName = errors.New("This portfolio does not exist ")
var ErrGetListPortfolio = errors.New("There are no portfolios in your office")
var ErrDeletePortfolio = errors.New("Failed to delete the portfolio because it does not exist ")
var ErrCreateMenuPortfolio = errors.New("Failed to create a short description, please make sure the data is correct")

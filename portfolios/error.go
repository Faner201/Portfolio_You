package portfolios

import "errors"

var (
	ErrCreatePortfolio        = errors.New("Portfolio was not created")
	ErrGetPortfolioByUserName = errors.New("This portfolio does not exist ")
	ErrGetListPortfolio       = errors.New("There are no portfolios in your office")
	ErrDeletePortfolio        = errors.New("Failed to delete the portfolio because it does not exist ")
	ErrCreateMenuPortfolio    = errors.New("Failed to create a short description, please make sure the data is correct")
	ErrSpecialSymbolName      = errors.New("Please remove the special characters from the name portfolio field")
	ErrNotTextPortfolio       = errors.New("Your portfolio does not have any text fields, please add and fill them in")
	ErrFullnessPortfolio      = errors.New("Your portfolio is empty, please fill it with something")
)

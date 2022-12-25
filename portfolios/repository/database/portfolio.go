package database

import (
	"Portfolio_You/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioRepository struct {
	db *mongo.Collection
}

func NewPortfolioRepository(db *mongo.Database, collection string) *PortfolioRepository {
	return &PortfolioRepository{
		db: db.Collection(collection),
	}
}

func (p PortfolioRepository) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User) error {

	res, err := p.db.InsertOne(ctx, &models.Portfolio{
		ID:          primitive.NewObjectID().Hex(),
		Url:         portfolio.Url,
		CreaterUser: user.Username,
		Name:        portfolio.Name,
		Text:        portfolio.Text,
		Photo:       portfolio.Photo,
		Colors:      portfolio.Colors,
		Struct:      portfolio.Struct,
	})

	if err != nil {
		return err
	}

	portfolio.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (p PortfolioRepository) CreateMenuPortfolio(ctx context.Context, user *models.User, menu *models.Menu) error {
	portfolio := &models.Portfolio{}
	out := p.db.FindOne(ctx, bson.M{
		"name":        menu.Name,
		"createrUser": menu.CreaterName,
	}).Decode(portfolio)

	if out != nil {
		return out
	}

	res, err := p.db.InsertOne(ctx, &models.Menu{
		ID:          portfolio.ID,
		Name:        menu.Name,
		CreaterName: user.Username,
		ShortText:   menu.ShortText,
		Photo:       menu.Photo,
	})

	if err != nil {
		return err
	}

	menu.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (p PortfolioRepository) GetPortfolioByUserName(ctx context.Context, userName, portfolioID string) (*models.Portfolio, error) {

	modelPortfolio := &models.Portfolio{}

	result := p.db.FindOne(ctx, bson.M{
		"id":          portfolioID,
		"createrUser": userName,
	}).Decode(modelPortfolio)

	if result != nil {
		return nil, result
	}

	return &models.Portfolio{
		ID:          modelPortfolio.ID,
		Url:         modelPortfolio.Url,
		CreaterUser: modelPortfolio.CreaterUser,
		Name:        modelPortfolio.Name,
		Text:        modelPortfolio.Text,
		Photo:       modelPortfolio.Photo,
		Colors:      modelPortfolio.Colors,
		Struct:      modelPortfolio.Struct,
	}, nil
}

func (p PortfolioRepository) GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error) {

	cur, err := p.db.Find(ctx, bson.M{
		"createrName": userName,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]*models.Menu, 0)

	for cur.Next(ctx) {
		modelMenu := new(models.Menu)
		err := cur.Decode(modelMenu)
		if err != nil {
			return nil, err
		}

		list = append(list, modelMenu)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (p PortfolioRepository) DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error {

	_, err := p.db.DeleteOne(ctx, bson.M{
		"id":          portfolioID,
		"createrUser": user.Username,
	})

	_, err = p.db.DeleteOne(ctx, bson.M{
		"id":          portfolioID,
		"createrName": user.Username,
	})

	return err
}

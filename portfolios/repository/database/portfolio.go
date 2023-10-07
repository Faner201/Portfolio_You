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

func (p PortfolioRepository) CreatePortfolio(ctx context.Context, user *models.User, portfolio *models.Portfolio) error {

	res, err := p.db.InsertOne(ctx, &models.Portfolio{
		ID:          primitive.NewObjectID().Hex(),
		Url:         portfolio.Url,
		CreaterUser: user.Username,
		Name:        portfolio.Name,
		Texts:       portfolio.Texts,
		Images:      portfolio.Images,
		Colors:      portfolio.Colors,
		Struct:      portfolio.Struct,
	})

	if err != nil {
		return err
	}

	portfolio.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (p PortfolioRepository) CreateMenuPortfolio(ctx context.Context, user *models.User, menuPortfolio *models.Menu) error {
	portfolio := &models.Portfolio{}
	out := p.db.FindOne(ctx, bson.M{
		"name":        menuPortfolio.Name,
		"createrUser": menuPortfolio.CreaterName,
	}).Decode(portfolio)

	if out != nil {
		return out
	}

	res, err := p.db.InsertOne(ctx, &models.Menu{
		ID:          portfolio.ID,
		Name:        menuPortfolio.Name,
		CreaterName: user.Username,
		ShortText:   menuPortfolio.ShortText,
		Image:       menuPortfolio.Image,
	})

	if err != nil {
		return err
	}

	menuPortfolio.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (p PortfolioRepository) GetPortfolioByUserName(ctx context.Context, userName, portfolioID string) (*models.Portfolio, error) {

	modelPortfolio := models.Portfolio{}

	err := p.db.FindOne(ctx, bson.M{
		"id":          portfolioID,
		"createrUser": userName,
	}).Decode(&modelPortfolio)

	if err != nil {
		return nil, err
	}

	return &modelPortfolio, nil
}

func (p PortfolioRepository) GetListPortfolioByUserName(ctx context.Context, userName string) (*[]models.Menu, error) {

	cur, err := p.db.Find(ctx, bson.M{
		"createrName": userName,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	list := []models.Menu{}

	for cur.Next(ctx) {
		modelMenu := new(models.Menu)
		err := cur.Decode(modelMenu)
		if err != nil {
			return nil, err
		}

		list = append(list, *modelMenu)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &list, nil
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

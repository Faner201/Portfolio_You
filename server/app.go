package server

import (
	"Portfolio_You/auth"
	authHttp "Portfolio_You/auth/delivery/http"
	authDatabase "Portfolio_You/auth/repository/database"
	authUseCase "Portfolio_You/auth/usecase"
	"Portfolio_You/portfolios"
	portfHttp "Portfolio_You/portfolios/delivery/http"
	portfDatabase "Portfolio_You/portfolios/repository/database"
	portfUseCase "Portfolio_You/portfolios/usecase"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	httpServer  *http.Server
	portfolioUC portfolios.UseCase
	authUC      auth.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := authDatabase.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	portfolioRepo := portfDatabase.NewPortfolioRepository(db, viper.GetString("mongo.portfolio_collection"))

	return &App{
		portfolioUC: portfUseCase.NewPortfolioUseCase(portfolioRepo),
		authUC: authUseCase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Default(),
	)
	router.Static("/portfolio/open", "/Users/fanfurick/Documents/Profile_You/src")

	authHttp.RegisterHttpEndpoints(router, a.authUC)

	authMiddleware := authHttp.NewAuthMiddleware(a.authUC)
	api := router.Group("", authMiddleware)

	portfHttp.RegisterHttpEndpoints(api, a.portfolioUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}

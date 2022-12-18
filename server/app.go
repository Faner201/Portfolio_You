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
)

type App struct {
	httpServer  *http.Server
	portfolioUC portfolios.UseCase
	authUC      auth.UseCase
}

func NewApp() *App {
	userRepo := authDatabase.NewUserRepository()
	portfolioRepo := portfDatabase.NewPortfolioRepository()

	return &App{
		portfolioUC: portfUseCase.NewPortfolioUseCase(portfolioRepo),
		authUC: authUseCase.NewAuthUseCase(
			userRepo,
			"hash_salt",
			[]byte("signing_key"),
			86400,
			// viper.GetString("auth.hash_salt"),
			// []byte(viper.GetString("auth.signing_key")),
			// viper.GetDuration("auth.token_ttl"),
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

	authHttp.RegisterHttpEndpoints(router, a.authUC)

	// authMiddleware := authHttp.NewAuthMiddleware(a.authUC)
	api := router.Group("")

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

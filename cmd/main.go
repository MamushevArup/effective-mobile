package main

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/handler"
	"github.com/MamushevArup/test-effective-mob/internal/repo/cars"
	serviceCar "github.com/MamushevArup/test-effective-mob/internal/usecase/cars"
	"github.com/MamushevArup/test-effective-mob/pkg/logger"
	"github.com/MamushevArup/test-effective-mob/pkg/psql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title Car API
// @version 3.0
// @description Api related to the cars
// @host localhost:3301
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	lg := logger.NewLogger()

	db := storageInit(ctx, lg)
	defer db.Close()

	lg.Info("connected to the database successfully")

	carsRepo := cars.New(db, lg)

	svc := serviceCar.New(carsRepo)

	hdl := handler.New(svc)

	server := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: hdl.InitRoutes(),
	}

	lg.Info("server started")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			lg.Fatalf("error running server %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	lg.Print("Application Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func storageInit(ctx context.Context, lg *logger.Logger) *pgxpool.Pool {
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	db := os.Getenv("POSTGRES_DATABASE")

	connector, err := psql.NewDBConnector(ctx, lg, user, pwd, host, port, db)
	if err != nil {
		lg.Fatalf(err.Error())
	}
	return connector
}

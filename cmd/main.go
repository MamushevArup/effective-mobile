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
	"log"
	"os"
)

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
	lg.Info("server started")
	go func() {
		if err := hdl.InitRoutes().Run(":" + os.Getenv("SERVER_PORT")); err != nil {
			lg.Fatalf("error running server %v", err)
		}
	}()

	select {}
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

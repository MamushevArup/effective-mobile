package psql

import (
	"context"
	"fmt"
	"github.com/MamushevArup/test-effective-mob/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConnector(ctx context.Context, lg *logger.Logger, user, pwd, host, port, database string) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pwd, host, port, database)
	conn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		lg.Errorf("unable to connect to the database %v\n", err)
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		lg.Errorf("unable to ping database %v\n", err)
		return nil, err
	}
	return conn, nil
}

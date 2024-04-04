package cars

import (
	"github.com/MamushevArup/test-effective-mob/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
	lg *logger.Logger
}

func New(db *pgxpool.Pool, lg *logger.Logger) *Repo {
	return &Repo{
		db: db,
		lg: lg,
	}
}

package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
)

type carManager interface {
	inserter
	deleter
	updater
}

type inserter interface {
	Insert(ctx context.Context, car *models.Car, owner models.Owner) error
}

type deleter interface {
	Delete(ctx context.Context, regNum string) error
}

type updater interface {
	Update(ctx context.Context, car *models.Car, owner *models.Owner) error
}

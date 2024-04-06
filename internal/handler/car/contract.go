package car

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
)

type carCreator interface {
	AddCar(ctx context.Context, car *models.Car) error
}

type carRemover interface {
	RemoveCar(ctx context.Context, regNum string) error
}

type carUpdater interface {
	UpdateCar(ctx context.Context, regNum string, car *models.Car) error
}

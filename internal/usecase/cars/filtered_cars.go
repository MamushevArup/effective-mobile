package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
)

var lim = 20

func (s *UseCase) GetCars(ctx context.Context, car models.Car, limit, offset int) ([]models.Car, error) {
	if limit > lim {
		limit = lim
	}
	if limit <= 0 {
		limit = 1
	}
	if offset < 0 {
		offset = 0
	}
	cars, err := s.carManager.Cars(ctx, car, limit, offset)
	return cars, err
}

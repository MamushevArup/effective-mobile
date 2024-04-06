package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
)

func (s *UseCase) UpdateCar(ctx context.Context, regNum string, car *models.Car) error {
	owner := car.Owner
	car.RegNum = regNum
	err := s.carManager.Update(ctx, car, &owner)
	return err
}

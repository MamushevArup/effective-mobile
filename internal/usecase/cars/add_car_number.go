package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/google/uuid"
)

func (s *UseCase) AddCar(ctx context.Context, reg *models.Car) error {
	var owner models.Owner
	var err error
	reg.Owner.Id = uuid.New()
	owner = reg.Owner
	err = s.carManager.Insert(ctx, reg, owner)
	if err != nil {
		return err
	}
	return nil
}

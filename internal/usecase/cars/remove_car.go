package cars

import (
	"context"
)

func (s *UseCase) RemoveCar(ctx context.Context, regNum string) error {
	err := s.carManager.Delete(ctx, regNum)
	if err != nil {
		return err
	}
	return nil
}

package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
)

type inserter interface {
	Insert(ctx context.Context, car models.Car) error
}

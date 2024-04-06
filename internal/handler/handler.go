package handler

import (
	"github.com/MamushevArup/test-effective-mob/internal/handler/car"
	"github.com/MamushevArup/test-effective-mob/internal/middleware/numbers"
	"github.com/MamushevArup/test-effective-mob/internal/usecase/cars"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	carsUseCase *cars.UseCase
}

func New(carsUseCase *cars.UseCase) *Handler {
	return &Handler{
		carsUseCase: carsUseCase,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/add-car", numbers.ValidateRegistration(), car.JoinCars(h.carsUseCase))
	router.DELETE("/delete-car/:number", numbers.ValidateRegNum(), car.DeleteCar(h.carsUseCase))
	router.PATCH("/update-car/:number", numbers.ValidateRegNum(), car.UpdateCar(h.carsUseCase))

	return router
}

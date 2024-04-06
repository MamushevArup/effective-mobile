package handler

import (
	_ "github.com/MamushevArup/test-effective-mob/docs"
	"github.com/MamushevArup/test-effective-mob/internal/handler/car"
	"github.com/MamushevArup/test-effective-mob/internal/middleware/numbers"
	"github.com/MamushevArup/test-effective-mob/internal/usecase/cars"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/add-car", numbers.ValidateRegistration(), car.JoinCars(h.carsUseCase))
	router.DELETE("/delete-car/:number", numbers.ValidateRegNum(), car.DeleteCar(h.carsUseCase))
	router.PATCH("/update-car/:number", numbers.ValidateRegNum(), car.UpdateCar(h.carsUseCase))
	router.GET("/cars", car.Cars(h.carsUseCase))

	return router
}

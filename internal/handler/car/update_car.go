package car

import (
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateCar
// @Summary UpdateCar
// @Tags car
// @Description update car by any filter except regNum
// @Accept json
// @Produce json
// @Param number path string true "number"
// @Success 200 {string} string "Car updated successfully"
// @Failure 500 {string} string "error"
// @Router /update-car/{number} [patch]
func UpdateCar(cu carUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {

		regNum := c.MustGet("validatedRegNum").(string)

		var car models.Car
		if err := c.BindJSON(&car); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := cu.UpdateCar(c, regNum, &car)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
	}
}

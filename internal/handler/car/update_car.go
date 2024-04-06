package car

import (
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

package car

import (
	"github.com/gin-gonic/gin"
)

// DeleteCar
// @Summary DeleteCar
// @Tags car
// @Description delete car by regNum
// @Accept json
// @Produce json
// @Param number path string true "number"
// @Success 200 {string} string "Car removed successfully"
// @Failure 500 {string} string "error"
// @Router /delete-car/{number} [delete]
func DeleteCar(cr carRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		regNum := c.MustGet("validatedRegNum").(string)

		err := cr.RemoveCar(c, regNum)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Car removed successfully"})
	}
}

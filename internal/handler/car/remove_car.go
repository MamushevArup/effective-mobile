package car

import (
	"github.com/gin-gonic/gin"
)

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

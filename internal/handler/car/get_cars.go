package car

import (
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cars
// @Summary Cars
// @Tags car
// @Description get cars by filter with pagination and offset
// @Accept json
// @Produce json
// @Param model query string false "model"
// @Param mark query string false "mark"
// @Param year query int false "year"
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {array} models.Car "Cars"
// @Failure 404 {string} string "no content by this filter"
// @Router /cars [get]
func Cars(cg carGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var car models.Car
		var pagination models.Pagination
		if err := c.ShouldBindQuery(&car); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindQuery(&pagination); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cars, err := cg.GetCars(c, car, pagination.Limit, pagination.Offset)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no content by this filter"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"cars": cars})
	}
}

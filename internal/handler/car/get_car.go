package car

import (
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

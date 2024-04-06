package numbers

import (
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

const numberPattern = `^[A-Za-z]\d{3}[A-Za-z]{2}\d{3}$`

func ValidateRegistration() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqHandler models.RegNumsReq
		if err := c.BindJSON(&reqHandler); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		re, _ := regexp.Compile(numberPattern)

		for _, regNum := range reqHandler.RegNums {
			if !re.MatchString(regNum) {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid regNum: " + regNum})
				return
			}
		}

		c.Set("validatedRegNums", reqHandler)
		c.Next()
	}
}

func ValidateRegNum() gin.HandlerFunc {
	return func(c *gin.Context) {

		regNum := c.Param("number")

		re, _ := regexp.Compile(numberPattern)

		if !re.MatchString(regNum) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid regNum: " + regNum})
			return
		}

		c.Set("validatedRegNum", regNum)
		c.Next()
	}
}

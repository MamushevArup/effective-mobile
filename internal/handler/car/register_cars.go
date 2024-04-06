package car

import (
	"bytes"
	"encoding/json"
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sync"
)

type Req struct {
	RegNum string `json:"regNum"`
}

func JoinCars(cc carCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		validatedRegNums := c.MustGet("validatedRegNums").(models.RegNumsReq)

		wg := &sync.WaitGroup{}

		err := apiCall(c, wg, validatedRegNums, cc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		wg.Wait()
		c.JSON(http.StatusOK, gin.H{"message": "Cars added successfully"})
	}
}

func apiCall(c *gin.Context, wg *sync.WaitGroup, validatedRegNums models.RegNumsReq, cc carCreator) error {
	var err error

	for _, num := range validatedRegNums.RegNums {
		var carInfo models.Car

		wg.Add(1)
		go func(num string) {
			defer wg.Done()

			reqBody, _ := json.Marshal(Req{RegNum: num})
			resp, err := http.Post(os.Getenv("MOCK_SERVER_ADDRESS"), "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				return
			}
			defer func() {
				err = resp.Body.Close()
				if err != nil {
					return
				}
			}()

			body, _ := io.ReadAll(resp.Body)

			err = json.Unmarshal(body, &carInfo)
			if err != nil {
				return
			}
			err = cc.AddCar(c, &carInfo)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

		}(num)
	}
	return err
}

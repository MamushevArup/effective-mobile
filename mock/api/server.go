package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

type Req struct {
	RegNum string `json:"regNum"`
}

type CarInfo struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
	} `json:"owner"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		return
	}

	r := gin.Default()

	r.POST("/info", func(c *gin.Context) {
		var req Req
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		response := CarInfo{
			RegNum: req.RegNum,
			Mark:   "Lada",
			Model:  "Vesta",
			Year:   2002,
			Owner: struct {
				Name       string `json:"name"`
				Surname    string `json:"surname"`
				Patronymic string `json:"patronymic"`
			}{
				Name:       "John",
				Surname:    "Doe",
				Patronymic: "Smith",
			},
		}

		c.JSON(200, response)
	})
	port := os.Getenv("MOCK_SERVER")

	if port == "" {
		port = "8080"
	}

	go func() {
		err := r.Run(":" + port)
		if err != nil {
			return
		}
	}()

	select {}

}

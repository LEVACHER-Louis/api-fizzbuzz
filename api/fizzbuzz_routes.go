package api

import (
	"api-fizzbuzz/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupFizzbuzzRoutes(groupe *gin.RouterGroup){
	groupe.GET("/", func(c *gin.Context){
		nbMaxStr := c.DefaultQuery ("nombreMax", "15")
		nbMax, conversionOk := strconv.Atoi(nbMaxStr)
		if conversionOk != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erreur":"parametre invalide"})
			return
		}
		res := internal.FizzBuzz(uint(nbMax))
		c.JSON(http.StatusOK, gin.H{"reponse":res})
		return
	})
}

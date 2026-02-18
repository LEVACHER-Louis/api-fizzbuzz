package routes

import (
	"api-fizzbuzz/internal"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// setupFizzbuzzRoutes met en place les routes suivant /fizzbuzz 
// 
// paramètres :
// - groupe : le groupe de routes où les routes seront ajoutées 
func setupFizzbuzzRoutes(groupe *gin.RouterGroup){
	groupe.GET("", func(c *gin.Context){
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

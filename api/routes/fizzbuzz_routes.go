package routes

import (
	"api-fizzbuzz/internal"
	"net/http"
	"github.com/gin-gonic/gin"
)

// setupFizzbuzzRoutes met en place les routes suivant /fizzbuzz 
// 
// paramètres :
// - groupe : le groupe de routes où les routes seront ajoutées 
func setupFizzbuzzRoutes(groupe *gin.RouterGroup){
	groupe.GET("", func(c *gin.Context){
		nbMax, err := defaultQueryUInt(c, "nombreMax", uint(15))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erreur":"parametre invalide"})
			return
		}
		res := internal.FizzBuzz(nbMax)
		c.JSON(http.StatusOK, gin.H{"reponse":res})
		return
	})
}

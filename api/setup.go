package api

import (
	"api-fizzbuzz/api/routes"
	"github.com/gin-gonic/gin"
)

func SetupAPI(cles []string) *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r, cles)
	return r
}

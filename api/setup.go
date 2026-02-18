// Package api contient la logique d'initialisation de l'api
package api

import (
	"api-fizzbuzz/api/routes"
	"github.com/gin-gonic/gin"
)

// SetupAPI initialise l'API.
// 
// paramètres :
// - cles : les cles d'api autorisées à l'accès.
//
// retourne : 
// L'API lançable via gin.Engine.Run 
func SetupAPI(cles map[string]struct{}) *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r, cles)
	return r
}

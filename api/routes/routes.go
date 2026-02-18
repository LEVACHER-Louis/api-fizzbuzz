// Package Routes met à disposition le nécessaire pour initiliser toutes les routes de l'API 
package routes

import (
	"api-fizzbuzz/api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes met en place toutes les routes de l'API ainsi que les middlewares nécessaires au bon fonctionnement
// 
// paramètres :
// - r : Le moteur gin sur lequel ajouter les routes
// - cles : Les clés d'API autorisées à accéder à l'API
func SetupRoutes(r *gin.Engine, cles map[string]struct{}) {
	fizzbuzz := r.Group("/fizzbuzz")
	fizzbuzz.Use(middleware.AuthMiddleware(cles))
	{
		setupFizzbuzzRoutes(fizzbuzz.Group(""))
	}
}

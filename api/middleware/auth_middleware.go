// Package middleware contenant tous les middlewares de l'API 
package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware initialise le middleware vérifiant la validité des clés d'API
//
// paramètres : 
// - cles : l'ensemble des clés autorisés
//
// retourne :
// Le middleware bloquant les requetes non autorisés
func AuthMiddleware(cles map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		cle := c.GetHeader("X-API-Key")
		if _, existe := cles[cle]; !existe {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"clé invalide"})
			c.Abort()
			return 
		}
		c.Next()
	}
}

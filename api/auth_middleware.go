package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cles []string) gin.HandlerFunc {
	ensembleCles := make(map[string]struct{})
	for _, cle := range cles {
		ensembleCles[cle] = struct{}{}
	}
	return func(c *gin.Context) {
		cle := c.GetHeader("X-API-Key")
		if _, existe := ensembleCles[cle]; !existe {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"clef invalide"})
			c.Abort()
			return 
		}
		c.Next()
	}
}

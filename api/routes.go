package api

import "github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine, cles []string) {
	fizzbuzz := r.Group("/fizzbuzz")
	fizzbuzz.Use(AuthMiddleware(cles))
	{
		SetupFizzbuzzRoutes(fizzbuzz.Group("/"))
	}
}

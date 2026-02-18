package routes

import (
	"api-fizzbuzz/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cles []string) {
	fizzbuzz := r.Group("/fizzbuzz")
	fizzbuzz.Use(middleware.AuthMiddleware(cles))
	{
		setupFizzbuzzRoutes(fizzbuzz.Group("/"))
	}
}

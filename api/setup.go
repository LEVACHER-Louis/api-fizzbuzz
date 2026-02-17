package api

import (
	"github.com/gin-gonic/gin"
)

func SetupAPI(cles []string) *gin.Engine {
	r := gin.Default()
	setupRoutes(r, cles)
	return r
}

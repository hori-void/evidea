package api

import "github.com/gin-gonic/gin"

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		RegisterUserRoutes(api)
		// RegisterPostRoutes(api)
	}
}

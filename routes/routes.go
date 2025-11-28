package routes

import (
	"jberaldo/go-api/middleware"
	health "jberaldo/go-api/services"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func RegisterRoutes(r *gin.Engine, db *pg.DB) {
	public := r.Group("/api/v1")
	public.GET("/ping", health.HandlePing)

	authorized := r.Group("/api/v1")
	authorized.Use(middleware.JwtAuth())
	authorized.GET("/api/user")
}

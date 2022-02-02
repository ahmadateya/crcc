package router

import (
	"github.com/gin-gonic/gin"
	//"github.com/ahmadateya/crcc/config"
	//"github.com/ahmadateya/crcc/api/router"
)

func Init() *gin.Engine {
	r := gin.New()
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	InitializeConfig(r)
	getRoutes(r)
	return r
}

// InitializeConfig initialize default configurations
func InitializeConfig(r *gin.Engine) {
	r.Use(CORSMiddleware())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept-Language,  Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT , DELETE , PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

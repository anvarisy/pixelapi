package middlewares

import (
	"log"
	"net/http"

	"github.com/anvarisy/pixelapi/auths"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auths.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := auths.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		if err_token := token.Claims.Valid(); err_token != nil && !token.Valid {
			c.JSON(http.StatusUnauthorized, "Token rejected")
			c.Abort()
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		username, _ := claims["auth_username"].(string)
		log.Println(username)
		err_enforcer := enforcer.LoadPolicy()
		if err_enforcer != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to load policy from DB"})
			return
		}
		is_admin_read, _ := enforcer.Enforce(username, "Admin", "read")
		is_admin_write, _ := enforcer.Enforce(username, "Admin", "write")
		if !is_admin_write && !is_admin_read {
			is_auth, err := enforcer.Enforce(username, obj, act)
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"msg": "Error occurred when authorizing user"})
				return
			}

			if !is_auth {
				c.AbortWithStatusJSON(403, gin.H{"msg": "You are not authorized"})
				return
			}
		}
		c.Next()
	}

}

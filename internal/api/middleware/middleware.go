package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"net/http"
	"strings"
	"webapp/internal/models/response_models"
	"webapp/pkg/utils"
)

const (
	REDIS_TOKEN_PREFIX = "logged_out"
)

// limiting the number of requests per second
// 15 requests per second with a burst of 20

func JWTAuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, response_models.Response{
				ResponseCode: http.StatusUnauthorized,
				Message:      "Authorization header is required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenString)

		ctx := context.Background()
		isLoggedOut, err2 := IsJwtTokenLogout(ctx, redisClient, tokenString)

		if isLoggedOut || err2 != nil {
			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusUnauthorized,
				Message:      "Token is logged out",
			})
			c.Abort()
			return
		}

		if err != nil {
			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusUnauthorized,
				Message:      "Invalid token",
			})
			c.Abort()
			return
		}

		// Pass user information to the next handler
		c.Set("email", claims.Email)
		c.Set("Role", claims.Role)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func IsJwtTokenLogout(ctx context.Context, redisClient *redis.Client, token string) (bool, error) {

	result, err := redisClient.Exists(ctx, REDIS_TOKEN_PREFIX+token).Result()

	// Handle errors appropriately
	if errors.Is(err, redis.Nil) {

		return false, nil

	} else if err != nil {
		// Redis connection or query issue
		return false, err
	}

	// Token found
	if result > 0 {
		return true, nil
	}

	// Token not found
	return false, nil

}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {

	return func(c *gin.Context) {
		role := c.GetString("Role")

		if role != requiredRole {
			c.JSON(http.StatusUnauthorized, response_models.Response{
				ResponseCode: http.StatusUnauthorized,
				Message:      "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

package middleware

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strings"
	"webapp/models/response_models"
	"webapp/utils"

	"github.com/gin-gonic/gin"
)

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
		isLoggedOut, err := IsJwtTokenLogout(ctx, redisClient, tokenString)

		if isLoggedOut {
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
		c.Set("exp", claims.ExpiresAt)
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

	index, err := redisClient.LRange(ctx, "jwt_tokens", 0, -1).Result()

	// Handle errors appropriately
	if errors.Is(err, redis.Nil) {
		// Token not found
		return false, nil
	} else if err != nil {
		// Redis connection or query issue
		return false, err
	}

	for _, t := range index {
		if t == token {
			return true, nil
		}
	}

	// Token found
	return false, nil
}

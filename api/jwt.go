package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rsh456/reservation-system/db"
)

// declarative pattern by declarating a function into another function
func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["Authorization"]
		if !ok {
			return ErrUnAuthorized()
		}
		token = token[len("Bearer "):]
		claims, err := validateToken(token)
		if err != nil {
			return err
		}
		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)

		// Check token expiration
		if time.Now().Unix() > expires {
			return NewError(http.StatusUnauthorized, "token expired")
		}

		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return ErrUnAuthorized()
		}
		// Set the current authenticated user to the context value
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, ErrUnAuthorized()
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, ErrUnAuthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnAuthorized()
	}
	return claims, nil
}

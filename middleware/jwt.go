package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	fmt.Println()
	if !ok {
		return fmt.Errorf("unauthorized")
	}
	claims, err := validateToken(token)
	if err != nil {
		return err
	}
	//Check token expiration
	expiresFloat := claims["expires"].(float64)
	expires := int64(expiresFloat)
	fmt.Println(expires)
	fmt.Println(claims)

	// Check token expiration
	if time.Now().Unix() > expires {
		fmt.Println(time.Now().Unix())
		return fmt.Errorf("token expired")
	}
	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, fmt.Errorf("unuathorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}
	return claims, nil
}

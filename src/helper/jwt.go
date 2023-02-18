package helper

import (
	"errors"
	"shop-api/src/config/env"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const JWT_ERROR = "login to process"

var secret = env.GetSecretJWTEnv()

func GenerateJWT(id int, isAdmin bool) string {
	claims := jwt.MapClaims{
		"id":         id,
		"is_admin":   isAdmin,
		"expires_at": time.Now().Add(1 * time.Hour).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := rawToken.SignedString([]byte(secret))
	return token
}

func ValidateJWT(c *fiber.Ctx) (interface{}, error) {
	header := c.GetReqHeaders()["Authorization"]
	bearer := strings.HasPrefix(strings.ToLower(header), "bearer")
	if !bearer {
		return nil, errors.New(JWT_ERROR)
	}

	strToken := strings.Split(header, " ")
	if len(strToken) != 2 {
		return nil, errors.New(JWT_ERROR)
	}

	token, _ := jwt.Parse(strToken[1], func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New(JWT_ERROR)
	}

	var mapClaims jwt.MapClaims
	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errors.New(JWT_ERROR)
	} else {
		mapClaims = v
	}

	if exp, ok := mapClaims["expires_at"].(float64); !ok {
		return nil, errors.New(JWT_ERROR)
	} else {
		if int64(exp)-time.Now().Unix() <= 0 {
			return nil, errors.New(JWT_ERROR)
		}
	}

	if _, ok := mapClaims["id"].(float64); !ok {
		return nil, errors.New(JWT_ERROR)
	}

	if _, ok := mapClaims["is_admin"].(bool); !ok {
		return nil, errors.New(JWT_ERROR)
	}

	return mapClaims, nil
}

package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"time"
	"user-service/model"
)

type JWT struct {
	Id          uuid.UUID `json:"id"`
	PhoneNumber string    `json:"phoneNumber"`
	Name        string    `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWT{
		Id:          user.Id,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

type JWTForShop struct {
	Id          uuid.UUID `json:"id"`
	PhoneNumber string    `json:"phoneNumber"`
	jwt.StandardClaims
}

func GenerateTokenForShop(userShop model.UserShop) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTForShop{
		Id:          userShop.Id,
		PhoneNumber: userShop.PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_FOR_SHOP")))
}

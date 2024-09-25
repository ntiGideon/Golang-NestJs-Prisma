package helper

import (
	"NestJsStyle/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(jwtPayload *model.JWTPayload) (tokenString string, err error) {
	expirationDate := time.Now().Add(1 * time.Hour)
	subject := uuid.New()
	claims := &model.JWTClaim{
		Username:   jwtPayload.Username,
		Email:      jwtPayload.Email,
		Id:         jwtPayload.Id,
		CustomData: map[string]interface{}{"subject": subject.String()},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
			Audience:  os.Getenv("FRONTEND_URL"),
			Issuer:    os.Getenv("BACKEND_URL"),
			Subject:   subject.String(),
			Id:        jwtPayload.Id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (*model.JWTClaim, error) {
	tokenString, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenString.Claims.(*model.JWTClaim)
	if !ok || tokenString.Valid == false {
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

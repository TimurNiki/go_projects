package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"v4/configs"
	"v4/types"
	"v4/utils"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT( userID int, secret []byte) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)
		token, err := ValidateJWT(tokenString)

		if err != nil {
			log.Printf("Failed to validate JWT: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("Token is not valid")
			permissionDenied(w)
			return
		}


		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(int)

		_, err = store.GetUserByID(userID )
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}
		// Token is valid
		handlerFunc(w, r)
	}
}
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}

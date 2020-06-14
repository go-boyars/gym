package exercise

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-boyars/gym/internal/config"
	"github.com/go-boyars/gym/internal/models"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomJWTClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type RegisterUserAPI struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

func (a *App) registerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var userApi RegisterUserAPI
	err := json.NewDecoder(request.Body).Decode(&userApi)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userApi.Password), 10)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	id := uuid.NewV4().String()
	user := models.User{
		ID:    id,
		Login: userApi.Login,
	}

	err = a.storage.CreateUser(request.Context(), user, string(hash))
	if err != nil {
		SetInternalError(response, err)
		return
	}

	_, err = response.Write([]byte(`{"id": "` + id + `"}`))
	if err != nil {
		// and what? log it
		_ = err
	}
}

type LoginAPI struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

func (a *App) loginHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var loginApi LoginAPI
	err := json.NewDecoder(request.Body).Decode(&loginApi)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	// TODO tx problem

	actualHash, err := a.storage.GetPwhash(request.Context(), loginApi.Login)
	if err != nil {
		// TODO handle not found
		SetInternalError(response, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(actualHash), []byte(loginApi.Password))
	if err != nil {
		response.WriteHeader(http.StatusForbidden)
		_, err := response.Write([]byte(`{"message": "invalid password"}`))
		if err != nil {
			// and what? log it
			_ = err
		}
		return
	}

	id, err := a.storage.GetUserID(request.Context(), loginApi.Login)
	if err != nil {
		// TODO handle not found
		SetInternalError(response, err)
		return
	}

	claims := CustomJWTClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
			Issuer:    config.Application,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		SetInternalError(response, err)
		return
	}

	_, err = response.Write([]byte(`{"token": "` + tokenString + `"}`))
	if err != nil {
		SetInternalError(response, err)
		return
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		authorizationHeader := request.Header.Get("authorization")
		if authorizationHeader == "" {
			// TODO content-type in separate middleware
			response.Header().Add("content-type", "application/json")
			response.WriteHeader(http.StatusForbidden)
			return

		}
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			response.Header().Add("content-type", "application/json")
			response.WriteHeader(http.StatusForbidden)
			return
		}
		jwtClaim, err := ValidateJWT(bearerToken[1])
		if err != nil {
			response.Header().Add("content-type", "application/json")
			response.WriteHeader(http.StatusForbidden)
			return
		}

		context.Set(request, "jwt", jwtClaim)
		next(response, request)
	})
}

func ValidateJWT(t string) (interface{}, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tokenData CustomJWTClaim
		err = mapstructure.Decode(claims, &tokenData)
		if err != nil {
			return nil, fmt.Errorf("unable to decode token")
		}
		return tokenData, nil
	}

	return nil, fmt.Errorf("invalid token")
}

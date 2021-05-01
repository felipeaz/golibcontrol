package jwt

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/platform/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// Auth represents JWT Auth
type Auth struct {
	Cache *redis.Cache
}

func NewAuth(cache *redis.Cache) *Auth {
	return &Auth{
		Cache: cache,
	}
}

// CreateToken generates a JWT Token
func (a Auth) CreateToken(userId uint) (model.TokenDetails, *errors.ApiError) {
	userIdString := strconv.FormatUint(uint64(userId), 10)
	td := model.TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		AccessUuid:  uuid.NewV4().String(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.NewV4().String(),
	}

	// Payload
	atClaims := jwt.MapClaims{
		"authorized":   true,
		"user_id":      userIdString,
		"access_uuid":  td.AccessUuid,
		"refresh_uuid": td.RefreshUuid,
		"exp":          td.AtExpires,
	}

	// Secret key used on signature crypto
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// jwt.SigningMethodHS256 is the method used to generate the signature
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	aToken, err := at.SignedString(jwtSecret)
	if err != nil {
		return model.TokenDetails{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.JWTTokenCreationFailMessage,
			Error:   err.Error(),
		}
	}

	// Create the refresh token
	rtClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userIdString,
		"exp":        td.RtExpires,
	}

	jwtRefreshSecret := []byte(os.Getenv("JWT_REFRESH_SECRET"))
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rToken, err := rt.SignedString(jwtRefreshSecret)
	if err != nil {
		return model.TokenDetails{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.JWTTokenCreationFailMessage,
			Error:   err.Error(),
		}
	}

	td.AccessToken = aToken
	td.RefreshToken = rToken

	return td, nil
}

// TokenValid validates the token
func (a Auth) TokenValid(r *http.Request) error {
	token, err := a.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// UserLoggedIn returns if an error if user is not logged in
func (a Auth) UserLoggedIn(r *http.Request) error {
	usr, err := a.GetAuthUser(r)
	if err != nil {
		return err
	}

	err = a.FetchAuth(&usr)
	if err != nil {
		return err
	}

	return nil
}

// ExtractTokenMetadata retrieves metadata from JWT Token
func (a Auth) ExtractTokenMetadata(r *http.Request) (model.AccessDetails, error) {
	token, err := a.VerifyToken(r)
	if err != nil {
		return model.AccessDetails{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return model.AccessDetails{}, err
		}
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return model.AccessDetails{}, err
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			return model.AccessDetails{}, err
		}

		userIdUint64, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			return model.AccessDetails{}, err
		}

		return model.AccessDetails{
			AccessUuid:  accessUuid,
			RefreshUuid: refreshUuid,
			UserId:      uint(userIdUint64),
		}, nil
	}

	return model.AccessDetails{}, err
}

// VerifyToken validates the token signing method
func (a Auth) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := a.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken retrieves token from request
func (a Auth) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// FetchAuth look for the token on redis
func (a Auth) FetchAuth(authD *model.AccessDetails) error {
	userId := strconv.FormatUint(uint64(authD.UserId), 10)
	_, err := a.Cache.Get(userId)
	if err != nil {
		return err
	}

	return nil
}

// GetAuthUser fetch the user on Redis and return the ID
func (a Auth) GetAuthUser(r *http.Request) (model.AccessDetails, error) {
	userAuth, err := a.ExtractTokenMetadata(r)
	if err != nil {
		return model.AccessDetails{}, err
	}

	return userAuth, nil
}

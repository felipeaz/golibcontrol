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
func CreateToken(userId uint) (model.TokenDetails, *errors.ApiError) {
	td := model.TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		AccessUuid:  uuid.NewV4().String(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.NewV4().String(),
	}

	// Payload
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        td.AtExpires,
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
		"user_id":    userId,
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

// ExtractTokenMetadata retrieves metadata from JWT Token
func (a Auth) ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := a.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
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
func (a Auth) FetchAuth(authD *model.AccessDetails) (uint64, error) {
	userid, err := a.Cache.Get(authD.AccessUuid)
	if err != nil {
		return 0, err
	}

	idString := string(userid)
	userID, _ := strconv.ParseUint(idString, 10, 64)
	return userID, nil
}

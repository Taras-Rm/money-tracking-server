package token

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	Secret string
	Ttl    time.Duration
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{
		Secret: secret,
		Ttl:    ttl,
	}
}

func (t *TokenManager) NewToken(userId int64) (string, error) {
	userIdStr := strconv.Itoa(int(userId))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(t.Ttl).Unix(),
		"id":  userIdStr,
	})

	return token.SignedString([]byte(t.Secret))
}

func (t *TokenManager) Parse(token string) (int64, error) {
	paresedToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.Secret), nil
	})
	if err != nil {
		return 0, err
	}

	res, ok := paresedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("can not parse token")
	}

	id, err := strconv.Atoi(res["id"].(string))
	if !ok {
		return 0, errors.New("invalid id")
	}

	return int64(id), nil
}

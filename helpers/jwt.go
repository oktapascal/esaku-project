package helpers

import (
	"context"
	"esaku-project/app/auths/models/domain"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type ConfigJwt interface {
	GenerateJwtToken(ctx context.Context, secret string, login domain.Login) (string, time.Time, error)
	GenerateJwtRefreshToken(ctx context.Context, secret string, login domain.Login) (string, time.Time, error)
}

type ConfigJwtImpl struct {
}

type Claims struct {
	KodeLokasi string `json:"kode_lokasi"`
	Nik        string `json:"nik"`
	jwt.RegisteredClaims
}

func (config *ConfigJwtImpl) GenerateJwtToken(ctx context.Context, secret string, login domain.Login) (string, time.Time, error) {
	expiration := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		KodeLokasi: login.KodeLokasi,
		Nik:        login.Nik,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiration},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expiration, nil
}

func (config *ConfigJwtImpl) GenerateJwtRefreshToken(ctx context.Context, secret string, login domain.Login) (string, time.Time, error) {
	expiration := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		KodeLokasi: login.KodeLokasi,
		Nik:        login.Nik,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiration},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expiration, nil
}

func NewJwt() ConfigJwt {
	return &ConfigJwtImpl{}
}

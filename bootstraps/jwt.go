package bootstraps

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/configs"
	"esaku-project/types"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT interface {
	GenerateAccessToken(login web.LoginResponse) (string, time.Time, error)
	GenerateRefreshToken(login web.LoginResponse) (string, time.Time, error)
	generateToken(login web.LoginResponse, expiration time.Time, secret []byte) (string, time.Time, error)
	GetJWTKey() string
	GetRefreshKey() string
}

type JWTImpl struct {
	Config configs.Config
}

func NewJWTImpl(config configs.Config) *JWTImpl {
	return &JWTImpl{Config: config}
}

func (config *JWTImpl) GenerateAccessToken(login web.LoginResponse) (string, time.Time, error) {
	expiration := time.Now().Add(1 * time.Hour)

	return config.generateToken(login, expiration, []byte(config.GetJWTKey()))
}

func (config *JWTImpl) GenerateRefreshToken(login web.LoginResponse) (string, time.Time, error) {
	expiration := time.Now().Add(24 * time.Hour)

	return config.generateToken(login, expiration, []byte(config.GetRefreshKey()))
}

func (config *JWTImpl) generateToken(login web.LoginResponse, expiration time.Time, secret []byte) (string, time.Time, error) {
	claims := &types.Claims{
		KodeLokasi: login.KodeLokasi,
		Nik:        login.Nik,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiration},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expiration, nil
}

func (config *JWTImpl) GetJWTKey() string {
	return config.Config.Get("JWT_KEY_TOKEN")
}

func (config *JWTImpl) GetRefreshKey() string {
	return config.Config.Get("JWT_REFRESH_KEY_TOKEN")
}

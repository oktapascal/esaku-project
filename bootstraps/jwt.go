package bootstraps

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/configs"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	KodeLokasi string `json:"kode_lokasi"`
	Nik        string `json:"nik"`
	jwt.RegisteredClaims
}

type JWT interface {
	GenerateAccessToken(login web.LoginResponse) (string, time.Time, error)
	GenerateRefreshToken(login web.LoginResponse) (string, time.Time, error)
	generateToken(login web.LoginResponse, expiration time.Time, secret []byte) (string, time.Time, error)
	getJWTKey() string
	getRefreshKey() string
}

type JWTImpl struct {
	Config configs.Config
}

func NewJWTImpl(config configs.Config) *JWTImpl {
	return &JWTImpl{Config: config}
}

func (config *JWTImpl) GenerateAccessToken(login web.LoginResponse) (string, time.Time, error) {
	expiration := time.Now().Add(1 * time.Hour)

	return config.generateToken(login, expiration, []byte(config.getJWTKey()))
}

func (config *JWTImpl) GenerateRefreshToken(login web.LoginResponse) (string, time.Time, error) {
	expiration := time.Now().Add(24 * time.Hour)

	return config.generateToken(login, expiration, []byte(config.getRefreshKey()))
}

func (config *JWTImpl) generateToken(login web.LoginResponse, expiration time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
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

func (config *JWTImpl) getJWTKey() string {
	return config.Config.Get("JWT_KEY_TOKEN")
}

func (config *JWTImpl) getRefreshKey() string {
	return config.Config.Get("JWT_REFRESH_KEY_TOKEN")
}

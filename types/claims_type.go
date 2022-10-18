package types

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	KodeLokasi string `json:"kode_lokasi"`
	Nik        string `json:"nik"`
	jwt.RegisteredClaims
}

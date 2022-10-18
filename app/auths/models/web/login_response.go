package web

import (
	"esaku-project/app/auths/models/domain"
)

type LoginResponse struct {
	Nik          string `json:"nik"`
	KodeLokasi   string `json:"kode_lokasi"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func ToLoginResponse(login domain.Login) LoginResponse {
	return LoginResponse{
		Nik:          login.Nik,
		KodeLokasi:   login.KodeLokasi,
		Token:        login.Token,
		RefreshToken: login.RefreshToken,
	}
}

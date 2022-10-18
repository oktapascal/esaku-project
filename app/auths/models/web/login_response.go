package web

import (
	"esaku-project/app/auths/models/domain"
)

type LoginResponse struct {
	Nik        string `json:"nik"`
	KodeLokasi string `json:"kode_lokasi"`
}

func ToLoginResponse(login domain.Login) LoginResponse {
	return LoginResponse{
		Nik:        login.Nik,
		KodeLokasi: login.KodeLokasi,
	}
}

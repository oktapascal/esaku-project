package web

import "esaku-project/app/settings/models/domain"

type UserResponse struct {
	Nik     string `json:"nik"`
	Nama    string `json:"nama"`
	Jabatan string `json:"jabatan"`
	NoTelp  string `json:"no_telp"`
	Email   string `json:"email"`
	Foto    string `json:"foto"`
}

func ToUserResponse(user domain.Karyawan) UserResponse {
	return UserResponse{
		Nik:     user.Nik,
		Nama:    user.Nama,
		Jabatan: user.Jabatan,
		NoTelp:  user.NoTelp,
		Email:   user.Email,
		Foto:    user.Foto,
	}
}

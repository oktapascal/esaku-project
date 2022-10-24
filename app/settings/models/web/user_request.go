package web

import "mime/multipart"

type UserRequest struct {
	Nik     string `json:"nik" validate:"required,min=1,max=10"`
	Nama    string `json:"nama" validate:"required,min=1,max=100"`
	Jabatan string `json:"jabatan" validate:"required,min=1,max=100"`
	Email   string `json:"email" validate:"required,email,min=1,max=50"`
	NoTelp  string `json:"no_telp" validate:"required,min=1,max=50"`
}

type PasswordRequest struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserUploadRequest struct {
	Foto *multipart.FileHeader `validate:"file"`
}

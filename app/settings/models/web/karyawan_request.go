package web

import "mime/multipart"

type KaryawanSaveRequest struct {
	Nik       string `json:"nik" validate:"required,min=1,max=10"`
	Nama      string `json:"nama" validate:"required,min=1,max=100"`
	Alamat    string `json:"alamat" validate:"required,min=1,max=150"`
	Jabatan   string `json:"jabatan" validate:"required,min=1,max=100"`
	NoTelp    string `json:"no_telp" validate:"required,numeric,min=1,max=50"`
	Email     string `json:"email" validate:"required,email,min=1,max=50"`
	KodeUnit  string `json:"kode_unit" validate:"required,min=1,max=10"`
	FlagAktif bool   `json:"flag_aktif" validate:"required,min=1,max=1"`
	NoHp      string `json:"no_hp" validate:"required,numeric,min=1,max=50"`
}

type KaryawanUpdateRequest struct {
	Nik       string `json:"nik" validate:"required,min=1,max=10"`
	Nama      string `json:"nama" validate:"required,min=1,max=100"`
	Alamat    string `json:"alamat" validate:"required,min=1,max=150"`
	Jabatan   string `json:"jabatan" validate:"required,min=1,max=100"`
	NoTelp    string `json:"no_telp" validate:"required,numeric,min=1,max=50"`
	Email     string `json:"email" validate:"required,email,min=1,max=50"`
	KodeUnit  string `json:"kode_unit" validate:"required,min=1,max=10"`
	FlagAktif bool   `json:"flag_aktif" validate:"required,min=1,max=1"`
	NoHp      string `json:"no_hp" validate:"required,numeric,min=1,max=50"`
}

type KaryawanUploadRequest struct {
	Nik  string                `validate:"required,min=1,max=10"`
	Foto *multipart.FileHeader `validate:"file"`
}

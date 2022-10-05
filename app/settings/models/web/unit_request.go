package web

type UnitSaveRequest struct {
	KodeUnit  string `json:"kode_unit" validate:"required,min=1,max=10"`
	Nama      string `json:"nama" validate:"required,min=1,max=150"`
	FlagAktif string `json:"flag_aktif" validate:"required,min=1,max=1"`
}

type UnitUpdateRequest struct {
	KodeUnit  string `json:"kode_unit" validate:"required,min=1,max=10"`
	Nama      string `json:"nama" validate:"required,min=1,max=150"`
	FlagAktif string `json:"flag_aktif" validate:"required,min=1,max=1"`
}

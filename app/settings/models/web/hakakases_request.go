package web

type HakAksesSaveRequest struct {
	Nik           string `json:"nik" validate:"required,min=1,max=255"`
	KelompokMenu  string `json:"kelompok_menu" validate:"required,min=1,max=10"`
	StatusAdmin   string `json:"status_admin" validate:"required,min=1,max=1"`
	KelompokAkses string `json:"kelompok_akses" validate:"required,min=1,max=20"`
	Password      string `json:"password" validate:"required,min=1,max=300"`
}

type HakAksesUpdateRequest struct {
	Nik           string `json:"nik" validate:"required,min=1,max=255"`
	KelompokMenu  string `json:"kelompok_menu" validate:"required,min=1,max=10"`
	StatusAdmin   string `json:"status_admin" validate:"required,min=1,max=1"`
	KelompokAkses string `json:"kelompok_akses" validate:"required,min=1,max=20"`
}

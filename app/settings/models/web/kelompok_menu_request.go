package web

type KelompokMenuSaveRequest struct {
	KodeKlp string `json:"kode_klp" validate:"required,min=1,max=20"`
	Nama    string `json:"nama" validate:"required,min=1,max=200"`
}

type KelompokMenuUpdateRequest struct {
	KodeKlp string `json:"kode_klp" validate:"required,min=1,max=20"`
	Nama    string `json:"nama" validate:"required,min=1,max=200"`
}

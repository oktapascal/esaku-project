package web

type FormSaveRequest struct {
	KodeForm string `json:"kode_form" validate:"required,min=1,max=10"`
	Nama     string `json:"nama" validate:"required,min=1,max=100"`
	Program  string `json:"program" validate:"required,min=1,max=200"`
}

type FormUpdateRequest struct {
	KodeForm string `json:"kode_form" validate:"required,min=1,max=10"`
	Nama     string `json:"nama" validate:"required,min=1,max=100"`
	Program  string `json:"program" validate:"required,min=1,max=200"`
}

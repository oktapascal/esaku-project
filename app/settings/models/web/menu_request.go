package web

type menuRequest struct {
	KodeKlpMenu string `json:"kode_klp_menu" validate:"required,min=1,max=20"`
	KodeMenu    string `json:"kode_menu" validate:"required,min=1,max=20"`
	NamaMenu    string `json:"nama_menu" validate:"required,min=1,max=30"`
	Level       string `json:"level_menu" validate:"required,min=1,max=1"`
	Index       string `json:"row_index" validate:"required"`
	Program     string `json:"program" validate:"required"`
}

type MenuSaveRequest struct {
	Payload []menuRequest `validate:"required,min=1,dive"`
}

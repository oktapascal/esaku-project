package web

type dataMenuSub struct {
	KodeMenu string        `json:"kode_menu" validate:"required,min=1,max=10"`
	NamaMenu string        `json:"nama_menu" validate:"required,min=1,max=100"`
	Program  string        `json:"program" validate:"min=1,max=10"`
	Level    uint8         `json:"level" validate:"number"`
	SubMenu  []dataMenuSub `json:"data_menu"`
}

type dataMenuMain struct {
	KodeMenu string        `json:"kode_menu" validate:"required,min=1,max=10"`
	NamaMenu string        `json:"nama_menu" validate:"required,min=1,max=100"`
	Program  string        `json:"program" validate:"min=1,max=10"`
	Level    uint8         `json:"level" validate:"number"`
	SubMenu  []dataMenuSub `json:"data_menu"`
}

type MenuSaveRequest struct {
	KodeKlpMenu string         `json:"kode_klp_menu" validate:"required,min=1,max=20"`
	ListMenu    []dataMenuMain `json:"data_menu" validate:"required,min=1,dive"`
}

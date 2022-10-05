package web

import "esaku-project/app/settings/models/domain"

type UnitResponse struct {
	KodeUnit  string `json:"kode_unit"`
	Nama      string `json:"nama"`
	FlagAktif string `json:"flag_aktif"`
}

func ToUnitResponse(unit domain.Unit) UnitResponse {
	return UnitResponse{
		KodeUnit:  unit.KodeUnit,
		Nama:      unit.Nama,
		FlagAktif: unit.FlagAktif,
	}
}

func ToUnitResponses(units []domain.Unit) []UnitResponse {
	var unitResponses []UnitResponse

	for _, unit := range units {
		unitResponses = append(unitResponses, ToUnitResponse(unit))
	}

	return unitResponses
}

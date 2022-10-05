package web

import "esaku-project/app/settings/models/domain"

type KelompokMenuResponse struct {
	KodeKlp string `json:"kode_klp"`
	Nama    string `json:"nama"`
}

func ToKelompokMenuResponse(klpMenu domain.KelompokMenu) KelompokMenuResponse {
	return KelompokMenuResponse{
		KodeKlp: klpMenu.KodeKlp,
		Nama:    klpMenu.Nama,
	}
}

func ToKelompokMenuResponses(klpMenus []domain.KelompokMenu) []KelompokMenuResponse {
	var klpMenuResponses []KelompokMenuResponse

	for _, klpMenu := range klpMenus {
		klpMenuResponses = append(klpMenuResponses, ToKelompokMenuResponse(klpMenu))
	}

	return klpMenuResponses
}

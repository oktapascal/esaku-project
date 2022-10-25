package web

import "esaku-project/app/settings/models/domain"

type FilterKelompokMenuResponse struct {
	KodeKlp string `json:"kode_klp"`
	Nama    string `json:"nama"`
}

type KelompokMenuResponse struct {
	KodeKlp string `json:"kode_klp"`
	Nama    string `json:"nama"`
}

func ToFilterKelompokMenuResponses(klpMenus []domain.KelompokMenu) []FilterKelompokMenuResponse {
	var klpMenuResponses []FilterKelompokMenuResponse

	for _, klpMenu := range klpMenus {
		klpMenuResponses = append(klpMenuResponses, ToFilterKelompokMenuResponse(klpMenu))
	}

	return klpMenuResponses
}

func ToFilterKelompokMenuResponse(klpMenu domain.KelompokMenu) FilterKelompokMenuResponse {
	return FilterKelompokMenuResponse{
		KodeKlp: klpMenu.KodeKlp,
		Nama:    klpMenu.Nama,
	}
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

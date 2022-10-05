package web

import "esaku-project/app/settings/models/domain"

type FormResponse struct {
	KodeForm string `json:"kode_form"`
	Nama     string `json:"nama"`
	Program  string `json:"program"`
}

func ToFormResponse(form domain.Form) FormResponse {
	return FormResponse{
		KodeForm: form.KodeForm,
		Nama:     form.Nama,
		Program:  form.Program,
	}
}

func ToFormResponses(forms []domain.Form) []FormResponse {
	var formResponses []FormResponse

	for _, form := range forms {
		formResponses = append(formResponses, ToFormResponse(form))
	}

	return formResponses
}

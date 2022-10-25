package web

import "esaku-project/app/settings/models/domain"

type HakAksesListResponse struct {
	Nik          string `json:"nik"`
	Nama         string `json:"nama"`
	KelompokMenu string `json:"kelompok_menu"`
	StatusAdmin  string `json:"status_admin"`
}

type HakAksesDetailResponse struct {
	Nik              string `json:"nik"`
	NamaKaryawan     string `json:"nama_karyawan"`
	StatusAdmin      string `json:"status_admin"`
	DefaultProgram   string `json:"default_program"`
	KelompokAkses    string `json:"kelompok_akses"`
	KodeKelompokMenu string `json:"kode_kelompok_menu"`
	NamaKelompokMenu string `json:"nama_kelompok_menu"`
}

func ToHakAksesListResponse(akses domain.HakAkses) HakAksesListResponse {
	return HakAksesListResponse{
		Nik:          akses.Nik,
		Nama:         akses.Karyawan.Nama,
		KelompokMenu: akses.KodeKlp,
		StatusAdmin:  akses.StatusAdmin,
	}
}

func ToHakAksesListResponses(aksess []domain.HakAkses) []HakAksesListResponse {
	var aksesResponses []HakAksesListResponse

	for _, akses := range aksess {
		aksesResponses = append(aksesResponses, ToHakAksesListResponse(akses))
	}

	return aksesResponses
}

func ToHakAksesDetailResponse(akses domain.HakAkses) HakAksesDetailResponse {
	return HakAksesDetailResponse{
		Nik:              akses.Nik,
		NamaKaryawan:     akses.Karyawan.Nama,
		StatusAdmin:      akses.StatusAdmin,
		KelompokAkses:    akses.KelompokAkses,
		DefaultProgram:   akses.DefaultProgram,
		KodeKelompokMenu: akses.KodeKlp,
		NamaKelompokMenu: akses.KelompokMenu.Nama,
	}
}

package web

import "esaku-project/app/settings/models/domain"

type KaryawanListResponse struct {
	Nik      string `json:"nik"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Jabatan  string `json:"jabatan"`
	NoTelp   string `json:"no_telp"`
	Email    string `json:"email"`
	KodeUnit string `json:"kode_unit"`
}

type KaryawanDetailResponse struct {
	Nik       string `json:"nik"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Jabatan   string `json:"jabatan"`
	NoTelp    string `json:"no_telp"`
	Email     string `json:"email"`
	KodeUnit  string `json:"kode_unit"`
	NoHp      string `json:"no_hp"`
	FlagAktif bool   `json:"flag_aktif"`
	Foto      string `json:"file"`
	NamaUnit  string `json:"nama_unit"`
}

type KaryawanUploadResponse struct {
	Foto string `json:"file"`
}

func ToKaryawanListResponse(karyawan domain.Karyawan) KaryawanListResponse {
	return KaryawanListResponse{
		Nik:      karyawan.Nik,
		Nama:     karyawan.Nama,
		Alamat:   karyawan.Alamat,
		Jabatan:  karyawan.Jabatan,
		NoTelp:   karyawan.NoTelp,
		Email:    karyawan.Email,
		KodeUnit: karyawan.KodeUnit,
	}
}

func ToKaryawanListResponses(karyawans []domain.Karyawan) []KaryawanListResponse {
	var karyawanResponses []KaryawanListResponse

	for _, karyawan := range karyawans {
		karyawanResponses = append(karyawanResponses, ToKaryawanListResponse(karyawan))
	}

	return karyawanResponses
}

func ToKaryawanDetailResponse(karyawan domain.Karyawan) KaryawanDetailResponse {
	return KaryawanDetailResponse{
		Nik:       karyawan.Nik,
		Nama:      karyawan.Nama,
		Alamat:    karyawan.Alamat,
		Jabatan:   karyawan.Jabatan,
		NoTelp:    karyawan.NoTelp,
		Email:     karyawan.Email,
		KodeUnit:  karyawan.KodeUnit,
		NoHp:      karyawan.NoHp,
		FlagAktif: karyawan.FlagAktif,
		Foto:      karyawan.Foto,
		NamaUnit:  karyawan.Unit.Nama,
	}
}

func ToKaryawanUploadResponse(file string) KaryawanUploadResponse {
	return KaryawanUploadResponse{
		Foto: file,
	}
}

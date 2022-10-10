package repository

import (
	"context"
	"database/sql"
	"errors"
	"esaku-project/app/settings/models/domain"
	"esaku-project/helpers"
)

type KaryawanRepositoryImpl struct {
}

func NewKaryawanRepositoryImpl() *KaryawanRepositoryImpl {
	return &KaryawanRepositoryImpl{}
}

func (repository *KaryawanRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan {
	SQL := `insert into karyawan (nik, kode_lokasi, nama, kode_pp, flag_aktif, jabatan, alamat,
    no_telp, no_hp, email) values (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10)`

	_, err := tx.ExecContext(ctx, SQL, karyawan.Nik, karyawan.KodeLokasi, karyawan.Nama, karyawan.KodeUnit,
		karyawan.FlagAktif, karyawan.Jabatan, karyawan.Alamat, karyawan.NoTelp, karyawan.NoHp, karyawan.Email)
	helpers.PanicIfError(err)

	return karyawan
}

func (repository *KaryawanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan {
	SQL := `update karyawan set kode_lokasi = @p1, nama = @p2, kode_pp = @p3, flag_aktif = @p4, 
    jabatan = @p5, alamat = @p6, no_telp = @p7, no_hp = @p8, email = @p9
    where nik = @p10`

	_, err := tx.ExecContext(ctx, SQL, karyawan.KodeLokasi, karyawan.Nama, karyawan.KodeUnit,
		karyawan.FlagAktif, karyawan.Jabatan, karyawan.Alamat, karyawan.NoTelp, karyawan.NoHp, karyawan.Email, karyawan.Nik)
	helpers.PanicIfError(err)

	return karyawan
}

func (repository *KaryawanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) {
	SQL := "delete from karyawan where nik = @p1"

	_, err := tx.ExecContext(ctx, SQL, karyawan.Nik)
	helpers.PanicIfError(err)
}

func (repository *KaryawanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, nik string, kodeLokasi string) (domain.Karyawan, error) {
	SQL := `select a.nik, a.nama, a.alamat, a.jabatan, a.no_telp, a.email, a.kode_pp kode_unit,
	a.flag_aktif, a.no_hp, a.foto, b.nama
	from karyawan a
	inner join pp b on a.kode_pp=b.kode_pp
	where a.nik = @p1 and a.kode_lokasi = @p2`

	rows, err := tx.QueryContext(ctx, SQL, nik, kodeLokasi)
	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	karyawan := domain.Karyawan{}

	if rows.Next() {
		err := rows.Scan(&karyawan.Nik, &karyawan.Nama, &karyawan.Alamat,
			&karyawan.Jabatan, &karyawan.NoTelp, &karyawan.Email,
			&karyawan.KodeUnit, &karyawan.FlagAktif, &karyawan.NoHp,
			&karyawan.Foto, &karyawan.Unit.Nama)

		helpers.PanicIfError(err)

		return karyawan, nil
	} else {
		return karyawan, errors.New("karyawan is not found")
	}
}

func (repository *KaryawanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeLokasi string) []domain.Karyawan {
	SQL := `select nik, nama, alamat, jabatan, no_telp, email, kode_pp
	from karyawan
	where kode_lokasi = @p1`
	rows, err := tx.QueryContext(ctx, SQL, kodeLokasi)
	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	var karyawans []domain.Karyawan

	for rows.Next() {
		karyawan := domain.Karyawan{}

		err := rows.Scan(&karyawan.Nik, &karyawan.Nama, &karyawan.Alamat, &karyawan.Jabatan,
			&karyawan.NoTelp, &karyawan.Email, &karyawan.KodeUnit)

		helpers.PanicIfError(err)

		karyawans = append(karyawans, karyawan)
	}

	return karyawans
}

func (repository *KaryawanRepositoryImpl) UploadImage(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan {
	SQL := "update karyawan set foto = @p1 where nik = @p2"

	_, err := tx.ExecContext(ctx, SQL, karyawan.Foto, karyawan.Nik)
	helpers.PanicIfError(err)

	return karyawan
}

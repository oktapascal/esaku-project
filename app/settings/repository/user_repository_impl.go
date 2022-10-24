package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
	"esaku-project/helpers"
	"github.com/pkg/errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) {
	SQL := `update karyawan set jabatan = @p1, nama = @p2, email = @p3, no_telp = @p4 
    where nik = @p5 and kode_lokasi = @p6`

	_, err := tx.ExecContext(ctx, SQL, karyawan.Jabatan, karyawan.Nama, karyawan.Email, karyawan.NoTelp, karyawan.Nik,
		karyawan.KodeLokasi)
	helpers.PanicIfError(err)
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, hakakses domain.HakAkses) {
	SQL := "update hakakses set password = @p1 where nik = @p2 and kode_lokasi = @p3"

	_, err := tx.ExecContext(ctx, SQL, hakakses.Password, hakakses.Karyawan.Nik, hakakses.Karyawan.KodeLokasi)
	helpers.PanicIfError(err)
}

func (repository *UserRepositoryImpl) UploadImage(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) {
	SQL := "update karyawan set foto = @p1 where nik = @p2 and kode_lokasi = @p3"

	_, err := tx.ExecContext(ctx, SQL, karyawan.Foto, karyawan.Nik, karyawan.KodeLokasi)
	helpers.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, nik string, kodeLokasi string) (domain.Karyawan, error) {
	SQL := `select a.nik, a.nama nama_karyawan, a.jabatan, a.email, a.no_telp
	from karyawan a
	inner join hakakses b on a.nik=b.nik
	where a.nik = @p1 and a.kode_lokasi = @p2`

	rows, err := tx.QueryContext(ctx, SQL, nik, kodeLokasi)
	helpers.PanicIfError(err)

	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	user := domain.Karyawan{}

	if rows.Next() {
		err := rows.Scan(&user.Nik, &user.Nama, &user.Jabatan, &user.Email, &user.NoTelp)
		helpers.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

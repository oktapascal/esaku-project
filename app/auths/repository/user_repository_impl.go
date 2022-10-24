package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
	"esaku-project/helpers"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) {
	SQL := `update karyawan set jabatan = @p1, nama = @p2, kode_pp = @p3, email = @p4, no_telp = @p5 
    where nik = @p6`

	_, err := tx.ExecContext(ctx, SQL, karyawan.Jabatan, karyawan.Nama, karyawan.KodeUnit,
		karyawan.Email, karyawan.NoTelp, karyawan.Nik)
	helpers.PanicIfError(err)
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, hakakses domain.HakAkses) {
	SQL := "update hakakses set password = @p1 where nik = @p2"

	_, err := tx.ExecContext(ctx, SQL, hakakses.Password, hakakses.Karyawan.Nik)
	helpers.PanicIfError(err)
}

func (repository *UserRepositoryImpl) UploadImage(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) {
	SQL := "update karyawan set foto = @p1 where nik = @p2"

	_, err := tx.ExecContext(ctx, SQL, karyawan.Foto, karyawan.Nik)
	helpers.PanicIfError(err)
}

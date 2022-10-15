package repository

import (
	"context"
	"database/sql"
	"errors"
	"esaku-project/app/auths/models/domain"
	"esaku-project/helpers"
)

type LoginRepositoryImpl struct {
}

func NewLoginRepositoryImpl() *LoginRepositoryImpl {
	return &LoginRepositoryImpl{}
}

func (repository *LoginRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, nik string) (domain.Login, error) {
	SQL := `select a.nik, a.password, b.nama nama_user, b.kode_lokasi
	from hakakses a 
	inner join karyawan b on a.nik=b.nik
	where a.nik = @p1`

	rows, err := tx.QueryContext(ctx, SQL, nik)
	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	login := domain.Login{}

	if rows.Next() {
		err := rows.Scan(&login.Nik, &login.Password, &login.NamaUser, &login.KodeLokasi)
		helpers.PanicIfError(err)

		return login, nil
	} else {
		return login, errors.New("username is not found")
	}
}

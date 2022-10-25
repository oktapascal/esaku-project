package repository

import (
	"context"
	"database/sql"
	"errors"
	"esaku-project/app/settings/models/domain"
	"esaku-project/helpers"
)

type HakAksesRepositoryImpl struct {
}

func NewHakAksesRepositoryImpl() *HakAksesRepositoryImpl {
	return &HakAksesRepositoryImpl{}
}

func (repository *HakAksesRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, akses domain.HakAkses) domain.HakAkses {
	SQL := `insert into hakakses (kode_klp_menu, nik, status_admin, klp_akses, password, path_view)
	values (@p1, @p2, @p3, @p4, @p5, @p6)`

	_, err := tx.ExecContext(ctx, SQL, akses.KodeKlp, akses.Nik, akses.StatusAdmin,
		akses.KelompokAkses, akses.Password, akses.DefaultProgram)
	helpers.PanicIfError(err)

	return akses
}

func (repository *HakAksesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, akses domain.HakAkses) domain.HakAkses {
	SQL := `update hakakses set kode_klp_menu = @p1, status_admin = @p2, klp_akses = @p3, path_view = @p4
	where nik = @p5`

	_, err := tx.ExecContext(ctx, SQL, akses.KodeKlp, akses.StatusAdmin, akses.KelompokAkses,
		akses.DefaultProgram, akses.Nik)
	helpers.PanicIfError(err)

	return akses
}

func (repository *HakAksesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, akses domain.HakAkses) {
	SQL := "delete from hakakses where nik = @p1"

	_, err := tx.ExecContext(ctx, SQL, akses.Nik)
	helpers.PanicIfError(err)
}

func (repository *HakAksesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, nik string) (domain.HakAkses, error) {
	SQL := `select a.nik, a.kode_klp_menu, a.status_admin, a.klp_akses, a.path_view, b.nama nama_klp, c.nama nama_karyawan
	from hakakses a
	inner join menu_klp b on a.kode_klp_menu=b.kode_klp
	inner join karyawan c on a.nik=c.nik
	where a.nik = @p1`

	rows, err := tx.QueryContext(ctx, SQL, nik)
	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	akses := domain.HakAkses{}
	if rows.Next() {
		err := rows.Scan(&akses.Nik, &akses.KodeKlp, &akses.StatusAdmin, &akses.KelompokAkses, &akses.DefaultProgram,
			&akses.KelompokMenu.Nama, &akses.Karyawan.Nama)
		helpers.PanicIfError(err)

		return akses, nil
	} else {
		return akses, errors.New("hak akses is not found")
	}
}

func (repository *HakAksesRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.HakAkses {
	SQL := `select a.nik, a.kode_klp_menu, a.status_admin, a.kelompok_akses, b.nama nama_karyawan
	from hakakses a
	inner join karyawan b on a.nik=b.nik`
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	var aksess []domain.HakAkses
	for rows.Next() {
		akses := domain.HakAkses{}

		err := rows.Scan(&akses.Nik, &akses.KodeKlp, &akses.StatusAdmin, &akses.KelompokAkses, &akses.Karyawan.Nama)

		helpers.PanicIfError(err)

		aksess = append(aksess, akses)
	}

	return aksess
}

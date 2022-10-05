package repository

import (
	"context"
	"database/sql"
	"errors"
	"esaku-project/app/settings/models/domain"
	"esaku-project/helpers"
)

type UnitRepositoryImpl struct {
}

func NewUnitRepositoryImpl() *UnitRepositoryImpl {
	return &UnitRepositoryImpl{}
}

func (repository *UnitRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, unit domain.Unit) domain.Unit {
	SQL := "insert into pp (kode_pp, nama, flag_aktif, kode_lokasi) values (@p1, @p2, @p3, @p4)"

	_, err := tx.ExecContext(ctx, SQL, unit.KodeUnit, unit.Nama, unit.FlagAktif, unit.KodeLokasi)

	helpers.PanicIfError(err)

	return unit
}

func (repository *UnitRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, unit domain.Unit) domain.Unit {
	SQL := `update pp set nama = @p1, flag_aktif = @p2
	where kode_pp = @p3 and kode_lokasi = @p4`

	_, err := tx.ExecContext(ctx, SQL, unit.Nama, unit.FlagAktif, unit.KodeUnit, unit.KodeLokasi)

	helpers.PanicIfError(err)

	return unit
}

func (repository *UnitRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, unit domain.Unit) {
	SQL := "delete from pp where kode_pp = @p1 and kode_lokasi = @p2"

	_, err := tx.ExecContext(ctx, SQL, unit.KodeUnit, unit.KodeLokasi)

	helpers.PanicIfError(err)
}

func (repository *UnitRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kodeUnit string, kodeLokasi string) (domain.Unit, error) {
	SQL := "select kode_pp, nama, flag_aktif from pp where kode_pp = @p1 and kode_lokasi = @p2"
	rows, err := tx.QueryContext(ctx, SQL, kodeUnit, kodeLokasi)

	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	unit := domain.Unit{}
	if rows.Next() {
		err := rows.Scan(&unit.KodeUnit, &unit.Nama, &unit.FlagAktif)
		helpers.PanicIfError(err)

		return unit, nil
	} else {
		return unit, errors.New("unit is not found")
	}
}

func (repository *UnitRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeLokasi string) []domain.Unit {
	SQL := "select kode_pp, nama, flag_aktif from pp where kode_lokasi = @p1"
	rows, err := tx.QueryContext(ctx, SQL, kodeLokasi)

	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	var units []domain.Unit

	for rows.Next() {
		unit := domain.Unit{}

		err := rows.Scan(&unit.KodeUnit, &unit.Nama, &unit.FlagAktif)

		helpers.PanicIfError(err)

		units = append(units, unit)
	}

	return units
}

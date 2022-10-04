package repository

import (
	"context"
	"database/sql"
	"errors"
	"esaku-project/app/settings/domain"
	"esaku-project/helpers"
)

type KelompokMenuRepositoryImpl struct {
}

func NewKelompokMenuRepositoryImpl() *KelompokMenuRepositoryImpl {
	return &KelompokMenuRepositoryImpl{}
}

func (repository *KelompokMenuRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, klpMenu domain.KelompokMenu) domain.KelompokMenu {
	SQL := "insert into menu_klp (kode_klp, nama) values (@p1, @p2)"

	_, err := tx.ExecContext(ctx, SQL, klpMenu.KodeKlp, klpMenu.Nama)

	helpers.PanicIfError(err)

	return klpMenu
}

func (repository *KelompokMenuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, klpMenu domain.KelompokMenu) domain.KelompokMenu {
	SQL := "update menu_klp set nama = @p1 where kode_klp = @p2"

	_, err := tx.ExecContext(ctx, SQL, klpMenu.Nama, klpMenu.KodeKlp)

	helpers.PanicIfError(err)

	return klpMenu
}

func (repository *KelompokMenuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, klpMenu domain.KelompokMenu) {
	SQL := "delete from menu_klp where kode_klp = @p1"

	_, err := tx.ExecContext(ctx, SQL, klpMenu.KodeKlp)

	helpers.PanicIfError(err)
}

func (repository *KelompokMenuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kodeKlp string) (domain.KelompokMenu, error) {
	SQL := "select kode_klp, nama from menu_klp where kode_klp = @p1"
	rows, err := tx.QueryContext(ctx, SQL, kodeKlp)

	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	klpMenu := domain.KelompokMenu{}
	if rows.Next() {
		err := rows.Scan(&klpMenu.KodeKlp, &klpMenu.Nama)
		helpers.PanicIfError(err)

		return klpMenu, nil
	} else {
		return klpMenu, errors.New("kelompok menu is not found")
	}
}

func (repository *KelompokMenuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.KelompokMenu {
	SQL := "select kode_klp, nama from menu_klp"
	rows, err := tx.QueryContext(ctx, SQL)

	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	var klpMenus []domain.KelompokMenu

	for rows.Next() {
		klpMenu := domain.KelompokMenu{}

		err := rows.Scan(&klpMenu.KodeKlp, &klpMenu.Nama)

		helpers.PanicIfError(err)

		klpMenus = append(klpMenus, klpMenu)
	}

	return klpMenus
}

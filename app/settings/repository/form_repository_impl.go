package repository

import (
	"context"
	"database/sql"
	"errors"
	"esaku-project/app/settings/domain"
	"esaku-project/helpers"
)

type FormRepositoryImpl struct {
}

func NewFormRepositoryImpl() *FormRepositoryImpl {
	return &FormRepositoryImpl{}
}

func (repository *FormRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, form domain.Form) domain.Form {
	SQL := "insert into m_form (kode_form, nama_form, form) values (@p1, @p2, @p3)"

	_, err := tx.ExecContext(ctx, SQL, form.KodeForm, form.Nama, form.Program)

	helpers.PanicIfError(err)

	return form
}

func (repository *FormRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, form domain.Form) domain.Form {
	SQL := "update m_form set nama_form = @p1, form = @p2 where kode_form = @p3"

	_, err := tx.ExecContext(ctx, SQL, form.Nama, form.Program, form.KodeForm)

	helpers.PanicIfError(err)

	return form
}

func (repository *FormRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, form domain.Form) {
	SQL := "delete from m_form where kode_form = @p1"

	_, err := tx.ExecContext(ctx, SQL, form.KodeForm)

	helpers.PanicIfError(err)
}

func (repository *FormRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kodeForm string) (domain.Form, error) {
	SQL := "select kode_form, nama_form, form from m_form where kode_form = @p1"
	rows, err := tx.QueryContext(ctx, SQL, kodeForm)

	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	form := domain.Form{}
	if rows.Next() {
		err := rows.Scan(&form.KodeForm, &form.Nama, &form.Program)
		helpers.PanicIfError(err)

		return form, nil
	} else {
		return form, errors.New("form is not found")
	}
}

func (repository *FormRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Form {
	SQL := "select kode_form, nama_form, form from m_form"
	rows, err := tx.QueryContext(ctx, SQL)

	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	var forms []domain.Form

	for rows.Next() {
		form := domain.Form{}

		err := rows.Scan(&form.KodeForm, &form.Nama, &form.Program)

		helpers.PanicIfError(err)

		forms = append(forms, form)
	}

	return forms
}

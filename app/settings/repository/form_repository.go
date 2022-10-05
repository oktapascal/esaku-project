package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type FormRepository interface {
	Save(ctx context.Context, tx *sql.Tx, form domain.Form) domain.Form
	Update(ctx context.Context, tx *sql.Tx, form domain.Form) domain.Form
	Delete(ctx context.Context, tx *sql.Tx, form domain.Form)
	FindById(ctx context.Context, tx *sql.Tx, kodeForm string) (domain.Form, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Form
}

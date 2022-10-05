package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type UnitRepository interface {
	Save(ctx context.Context, tx *sql.Tx, unit domain.Unit) domain.Unit
	Update(ctx context.Context, tx *sql.Tx, unit domain.Unit) domain.Unit
	Delete(ctx context.Context, tx *sql.Tx, unit domain.Unit)
	FindById(ctx context.Context, tx *sql.Tx, kodeUnit string, kodeLokasi string) (domain.Unit, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeLokasi string) []domain.Unit
}

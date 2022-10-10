package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type HakAksesRepository interface {
	Save(ctx context.Context, tx *sql.Tx, akses domain.HakAkses) domain.HakAkses
	Update(ctx context.Context, tx *sql.Tx, akses domain.HakAkses) domain.HakAkses
	Delete(ctx context.Context, tx *sql.Tx, akses domain.HakAkses)
	FindById(ctx context.Context, tx *sql.Tx, nik string) (domain.HakAkses, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.HakAkses
}

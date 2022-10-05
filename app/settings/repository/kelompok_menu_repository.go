package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type KelompokMenuRepository interface {
	Save(ctx context.Context, tx *sql.Tx, klpMenu domain.KelompokMenu) domain.KelompokMenu
	Update(ctx context.Context, tx *sql.Tx, klpMenu domain.KelompokMenu) domain.KelompokMenu
	Delete(ctx context.Context, tx *sql.Tx, klpMenu domain.KelompokMenu)
	FindById(ctx context.Context, tx *sql.Tx, kodeKlp string) (domain.KelompokMenu, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.KelompokMenu
}

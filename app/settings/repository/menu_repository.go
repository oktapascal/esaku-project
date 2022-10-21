package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type MenuRepository interface {
	Save(ctx context.Context, tx *sql.Tx, menu []domain.Menu)
	Delete(ctx context.Context, tx *sql.Tx, kodeKlp string)
	FindById(ctx context.Context, tx *sql.Tx, kodeKlp string) []domain.Menu
}

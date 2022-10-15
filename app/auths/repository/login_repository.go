package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/auths/models/domain"
)

type LoginRepository interface {
	Login(ctx context.Context, tx *sql.Tx, nik string) (domain.Login, error)
}

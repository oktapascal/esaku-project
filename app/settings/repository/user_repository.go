package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type UserRepository interface {
	Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan)
	UpdatePassword(ctx context.Context, tx *sql.Tx, hakakses domain.HakAkses)
	FindById(ctx context.Context, tx *sql.Tx, nik string, kodeLokasi string) (domain.Karyawan, error)
	UploadImage(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan)
}

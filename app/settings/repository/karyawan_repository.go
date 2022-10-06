package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
)

type KaryawanRepository interface {
	Save(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan
	Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan
	Delete(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan)
	FindById(ctx context.Context, tx *sql.Tx, nik string, kodeLokasi string) (domain.Karyawan, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeLokasi string) []domain.Karyawan
	UploadImage(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan
}

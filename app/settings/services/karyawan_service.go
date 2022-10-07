package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type KaryawanService interface {
	Save(ctx context.Context, request web.KaryawanSaveRequest) web.KaryawanListResponse
	Update(ctx context.Context, request web.KaryawanUpdateRequest) web.KaryawanListResponse
	Delete(ctx context.Context, nik string)
	FindById(ctx context.Context, nik string) web.KaryawanDetailResponse
	FindAll(ctx context.Context) []web.KaryawanListResponse
	UploadImage(ctx context.Context, request web.KaryawanUploadRequest) web.KaryawanUploadResponse
}

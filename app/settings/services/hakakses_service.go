package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type HakAksesService interface {
	Save(ctx context.Context, request web.HakAksesSaveRequest) web.HakAksesListResponse
	Update(ctx context.Context, request web.HakAksesUpdateRequest) web.HakAksesListResponse
	Delete(ctx context.Context, nik string)
	FindById(ctx context.Context, nik string) web.HakAksesDetailResponse
	FindAll(ctx context.Context) []web.HakAksesListResponse
}

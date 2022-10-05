package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type KelompokMenuService interface {
	Save(ctx context.Context, request web.KelompokMenuSaveRequest) web.KelompokMenuResponse
	Update(ctx context.Context, request web.KelompokMenuUpdateRequest) web.KelompokMenuResponse
	Delete(ctx context.Context, kodeKlp string)
	FindById(ctx context.Context, kodeKlp string) web.KelompokMenuResponse
	FindAll(ctx context.Context) []web.KelompokMenuResponse
}

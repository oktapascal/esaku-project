package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type UnitService interface {
	Save(ctx context.Context, request web.UnitSaveRequest) web.UnitResponse
	Update(ctx context.Context, request web.UnitUpdateRequest) web.UnitResponse
	Delete(ctx context.Context, kodeUnit string)
	FindById(ctx context.Context, kodeUnit string) web.UnitResponse
	FindAll(ctx context.Context) []web.UnitResponse
}

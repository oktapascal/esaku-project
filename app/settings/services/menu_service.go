package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type MenuService interface {
	Save(ctx context.Context, request web.MenuSaveRequest)
	FindById(ctx context.Context, kodeKlp string) web.MenuResponse
}

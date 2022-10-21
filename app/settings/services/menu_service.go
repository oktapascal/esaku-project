package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type MenuService interface {
	Save(ctx context.Context, request web.MenuSaveRequest)
	Delete(ctx context.Context, klpMenu string)
	FindById(ctx context.Context, klpMenu string)
}

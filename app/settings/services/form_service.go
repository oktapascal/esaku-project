package services

import (
	"context"
	"esaku-project/app/settings/models/web"
)

type FormService interface {
	Save(ctx context.Context, request web.FormSaveRequest) web.FormResponse
	Update(ctx context.Context, request web.FormUpdateRequest) web.FormResponse
	Delete(ctx context.Context, kodeForm string)
	FindById(ctx context.Context, kodeForm string) web.FormResponse
	FindAll(ctx context.Context) []web.FormResponse
}

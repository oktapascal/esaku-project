package services

import (
	"context"
	web2 "esaku-project/app/settings/models/web"
)

type UserService interface {
	Update(ctx context.Context, request web2.UserRequest)
	UpdatePassword(ctx context.Context, request web2.PasswordRequest)
	FindById(ctx context.Context) web2.UserResponse
	UploadImage(ctx context.Context, request web2.UserUploadRequest)
}

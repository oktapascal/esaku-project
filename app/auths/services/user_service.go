package services

import (
	"context"
	"esaku-project/app/auths/models/web"
)

type UserService interface {
	Update(ctx context.Context, request web.UserRequest)
	UpdatePassword(ctx context.Context, request web.PasswordRequest)
	FindById(ctx context.Context) web.UserResponse
	UploadImage(ctx context.Context, request web.UserUploadRequest)
}

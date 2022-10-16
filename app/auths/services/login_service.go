package services

import (
	"context"
	"esaku-project/app/auths/models/web"
)

type LoginService interface {
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	Logout(ctx context.Context) (string, string)
}

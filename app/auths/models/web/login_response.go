package web

import (
	"esaku-project/app/auths/models/domain"
	"time"
)

type LoginResponse struct {
	Token             string `json:"token"`
	RefreshToken      string `json:"refresh_token"`
	ExpirationAccess  time.Time
	ExpirationRefresh time.Time
}

func ToLoginResponse(login domain.Login) LoginResponse {
	return LoginResponse{
		Token:             login.Token,
		RefreshToken:      login.RefreshToken,
		ExpirationAccess:  login.ExpirationAccess,
		ExpirationRefresh: login.ExpirationRefresh,
	}
}

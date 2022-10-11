package web

import "esaku-project/app/auths/models/domain"

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func ToLoginResponse(login domain.Login) LoginResponse {
	return LoginResponse{
		Token:        login.Token,
		RefreshToken: login.RefreshToken,
	}
}

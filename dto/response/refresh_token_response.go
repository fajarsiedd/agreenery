package response

import "go-agreenery/entities"

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (tokenResponse RefreshTokenResponse) FromEntity(user entities.User) RefreshTokenResponse {
	return RefreshTokenResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}

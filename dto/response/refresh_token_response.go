package response

import "go-agreenery/entities"

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r RefreshTokenResponse) FromEntity(user entities.User) RefreshTokenResponse {
	return RefreshTokenResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}

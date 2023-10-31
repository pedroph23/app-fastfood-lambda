package apresentacao

type AuthDTO struct {
	AcessToken string `json:"access_token"`
}

func NewAuthDTO(accessToken string) *AuthDTO {
	return &AuthDTO{
		AcessToken: accessToken,
	}
}

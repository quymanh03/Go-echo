package handler

type userLoginResponse struct {
	Token string `json:"token"`
}

func newUserLoginResponse(token string) *userLoginResponse {
	return &userLoginResponse{
		Token: token,
	}
}

package handler

import "beginner/db/models"

type userLoginResponse struct {
	Token string `json:"token"`
}

type getMeResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func newUserLoginResponse(token string) *userLoginResponse {
	return &userLoginResponse{
		Token: token,
	}
}

func newGetMeResponse(user *models.User) *getMeResponse {
	return &getMeResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

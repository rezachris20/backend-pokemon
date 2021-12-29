package response

import "pokemon-list/model"

type UserResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"name"`
	Username string `json:"username"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Nama:     user.Nama,
		Username: user.Username,
	}
}

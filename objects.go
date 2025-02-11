package kenopsiauser

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Verified bool   `json:"verified"`
	AvatarId uint8  `json:"avatarId"`
}

func (user UserResponse) GetId() string {
	return user.Id
}

func (user UserResponse) GetEmail() string {
	return user.Email
}

func (user UserResponse) GetUsername() string {
	return user.Username
}

func (user UserResponse) IsVerified() bool {
	return user.Verified
}

func (user UserResponse) GetAvatarId() uint8 {
	return user.AvatarId
}

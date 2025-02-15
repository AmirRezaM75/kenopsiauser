package kenopsiauser

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Verified bool   `json:"verified"`
	AvatarId uint8  `json:"avatarId"`
}

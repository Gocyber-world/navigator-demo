package response

type UserAccountResponse struct {
	UserID                string `json:"userId"`
	Name                  string `json:"name"`
	BuiltopiaClientUserId string `json:"builtopiaClientUserId"`
}

type UserInfoResponse struct {
	UserID                string `json:"userId"`
	NickName              string `json:"nickName"`
	Email                 string `json:"email"`
	ProfilePicUrl         string `json:"profilePicUrl"`
	AvatarModelUrl        string `json:"avatarModelUrl"`
	BuiltopiaClientUserId string `json:"builtopiaClientUserId"`
}

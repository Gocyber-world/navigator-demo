package request

type OpenLogin struct {
	ClientID     string `json:"clientId"`
	ClientUserID string `json:"clientUserId"`
	Timestamp    uint   `json:"timestamp"`
	Sign         string `json:"sign"`
}

type RegisterUser struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=32"` // 注意这里的长度是对于英文字符来说 中文字符在utf-8下为 3字节长度
	NickName string `json:"nickName" binding:"required,max=64"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type UserEmailConfirmation struct {
	UserId            string `json:"userId"`
	ConfirmationToken string `json:"confirmationToken"`
}

type SysUserRegisterBuiltopiaCustomer struct {
	Email                 string `json:"email" binding:"required,email,max=64"`
	Password              string `json:"password" binding:"required,min=8,max=32"`
	DisplayName           string `json:"displayName" binding:"required,max=64"`
	ProfilePicUrl         string `json:"profilePicUrl" binding:"required,max=64"`
	AvatarModelUrl        string `json:"avatarModelUrl" binding:"required,max=64"`
	BuiltopiaClientUserId string `json:"builtopiaClientUserId" binding:"required,max=64"`
}

type UpdateUserProfile struct {
	NickName       string `json:"nickName" binding:"required,max=64"`
	ProfilePicUrl  string `json:"profilePicUrl" binding:"required,max=64"`
	AvatarModelUrl string `json:"avatarModelUrl" binding:"required,max=64"`
}

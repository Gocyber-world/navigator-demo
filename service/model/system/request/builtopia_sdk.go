package request

type BuiltopiaRegisterCustomer struct {
	Email          string `json:"email" binding:"required,email,max=64"`
	Password       string `json:"password" binding:"required,min=8,max=32"`
	DisplayName    string `json:"displayName" binding:"required,max=64"`
	ProfilePicUrl  string `json:"profilePicUrl" binding:"required,max=64"`
	AvatarModelUrl string `json:"avatarModelUrl" binding:"required,max=64"`
	ClientUserId   string `json:"clientUserId" binding:"required,max=64"`
}

type BuiltopiaUpdateCustomerProfile struct {
	DisplayName    string `json:"displayName" binding:"required,max=64"`
	ProfilePicUrl  string `json:"profilePicUrl" binding:"required,max=64"`
	AvatarModelUrl string `json:"avatarModelUrl" binding:"required,max=64"`
	ClientUserId   string `json:"clientUserId" binding:"required,max=64"`
}

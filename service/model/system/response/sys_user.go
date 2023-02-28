package response

type UserAccountResponse struct {
	UserID                string `json:"userId"`
	Name                  string `json:"name"`
	BuiltopiaClientUserId string `json:"builtopiaClientUserId"`
}

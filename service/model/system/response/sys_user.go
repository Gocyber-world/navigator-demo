package response

type UserAccountResponse struct {
	AccountID string `json:"accountId"`
	Type      string `json:"type"`
	UserID    string `json:"userId"`
	OrgID     string `json:"orgId"`
	Name      string `json:"name"`
}

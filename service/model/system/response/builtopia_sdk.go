package response

type BuiltopiaGetAuthorizedAssetsResponse struct {
	List  []string `json:"list"`
	Total int64    `json:"total"`
}

type BuiltopiaAuthorizeAssetsResponse struct {
	AuthorizedClientAssetIds []string `json:"authorizedClientAssetIds"`
	FailedClientAssetIds     []string `json:"failedClientAssetIds"`
	Failed                   int      `json:"failed"`
}

type BuiltopiaCancelAssetAuthorizationResponse struct {
	CanceledClientAssetIds []string `json:"canceledClientAssetIds"`
	FailedClientAssetIds   []string `json:"failedClientAssetIds"`
	Failed                 int      `json:"failed"`
}

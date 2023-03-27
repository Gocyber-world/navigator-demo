package system

import (
	"encoding/json"
	"errors"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	commResp "github.com/Gocyber-world/navigator-demo/model/common/response"
	request "github.com/Gocyber-world/navigator-demo/model/system/request"
	"github.com/Gocyber-world/navigator-demo/model/system/response"
	"github.com/Gocyber-world/navigator-demo/utils"
)

type BuiltopiaOpenApiService struct{}

var BuiltopiaOpenApiServiceApp = new(BuiltopiaOpenApiService)

func (bs *BuiltopiaOpenApiService) BuiltopiaSdkHelper(url string, method string, data interface{}, query interface{}) (int, commResp.Response, error) {
	respData := commResp.Response{}
	httpStatusCode, respBody, err := utils.SendHttpRequest(global.BUILTOPIA_ENDPOINT+url, method, data, query, global.BUILTOPIA_CLIENT_TOKEN)
	if err != nil {
		logger.Error(err.Error())
		return -1, respData, err
	}
	if err := json.Unmarshal([]byte(respBody), &respData); err != nil {
		logger.Error(err.Error())
		return -1, respData, err
	}
	if respData.Code != 0 || httpStatusCode != 200 {
		logger.Error(respData.Msg)
		return -1, respData, err
	}
	return 0, respData, nil
}

func (bs *BuiltopiaOpenApiService) RegisterCustomer(email string, password string, clientUserId string, displayName string, profilePicUrl string, avatarModelUrl string) error {
	var reqData = request.BuiltopiaRegisterCustomer{
		Email:          email,
		Password:       password,
		ClientUserId:   clientUserId,
		DisplayName:    displayName,
		ProfilePicUrl:  profilePicUrl,
		AvatarModelUrl: avatarModelUrl,
	}
	_, _, err := bs.BuiltopiaSdkHelper("/v2/openapi/customer", "POST", reqData, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (bs *BuiltopiaOpenApiService) UpdateCustomerProfile(clientUserId string, displayName string, profilePicUrl string, avatarModelUrl string) error {
	var reqData = request.BuiltopiaUpdateCustomerProfile{
		ClientUserId:   clientUserId,
		DisplayName:    displayName,
		ProfilePicUrl:  profilePicUrl,
		AvatarModelUrl: avatarModelUrl,
	}
	_, _, err := bs.BuiltopiaSdkHelper("/v2/openapi/account", "PATCH", reqData, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (bs *BuiltopiaOpenApiService) GetAuthorizedAssets(clientUserId string) ([]string, error) {
	query := map[string]interface{}{
		"clientUserId": clientUserId,
	}
	_, respData, err := bs.BuiltopiaSdkHelper("/v2/openapi/asset/authorization", "GET", nil, query)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if resp, ok := respData.Data.(response.BuiltopiaGetAuthorizedAssetsResponse); ok {
		return resp.List, nil
	} else {
		return nil, errors.New("response data type error")
	}
}

func (bs *BuiltopiaOpenApiService) AuthorizeAssetWithUsageLimit(clientUserId string, clientAssetIds map[string]int) (response.BuiltopiaAuthorizeAssetsResponse, error) {
	var reqData = request.BuiltopiaAuthorizeAssets{
		ClientUserId:   clientUserId,
		ClientAssetIds: clientAssetIds,
	}
	_, respData, err := bs.BuiltopiaSdkHelper("/v2/openapi/asset/authorization", "POST", reqData, nil)
	if err != nil {
		logger.Error(err.Error())
		return response.BuiltopiaAuthorizeAssetsResponse{}, err
	}
	if resp, ok := respData.Data.(response.BuiltopiaAuthorizeAssetsResponse); ok {
		return resp, nil
	} else {
		return response.BuiltopiaAuthorizeAssetsResponse{}, errors.New("response data type error")
	}
}

func (bs *BuiltopiaOpenApiService) CancelAssetAuthorization(clientUserId string, clientAssetIds []string) (response.BuiltopiaCancelAssetAuthorizationResponse, error) {
	if len(clientAssetIds) == 0 {
		return response.BuiltopiaCancelAssetAuthorizationResponse{}, errors.New("clientAssetIds is empty")
	}
	if len(clientAssetIds) > 10 {
		return response.BuiltopiaCancelAssetAuthorizationResponse{}, errors.New("clientAssetIds is too long")
	}
	clientAssetIdsParam := clientAssetIds[0]
	for i := 1; i < len(clientAssetIds); i++ {
		clientAssetIdsParam += ("," + clientAssetIds[i])
	}
	query := map[string]interface{}{
		"clientUserId":  clientUserId,
		"clientAssetId": clientAssetIdsParam,
	}
	_, respData, err := bs.BuiltopiaSdkHelper("/v2/openapi/asset/authorization", "DELETE", nil, query)
	if err != nil {
		logger.Error(err.Error())
		return response.BuiltopiaCancelAssetAuthorizationResponse{}, err
	}
	if resp, ok := respData.Data.(response.BuiltopiaCancelAssetAuthorizationResponse); ok {
		return resp, nil
	} else {
		return response.BuiltopiaCancelAssetAuthorizationResponse{}, errors.New("response data type error")
	}
}

package system

import (
	"encoding/json"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	commResp "github.com/Gocyber-world/navigator-demo/model/common/response"
	"github.com/Gocyber-world/navigator-demo/utils"
)

type BuiltopiaOpenApiService struct{}

var BuiltopiaOpenApiServiceApp = new(BuiltopiaOpenApiService)

func (bs *BuiltopiaOpenApiService) BuiltopiaSdkHelper(url string, method string, data interface{}) (int, commResp.Response, error) {
	respData := commResp.Response{}
	httpStatusCode, respBody, err := utils.SendHttpRequest(url, method, data, global.BUILTOPIA_CLIENT_TOKEN)
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
	return nil
}

package dm

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) SetMaxQuotaLevel(request *SetMaxQuotaLevelRequest) (response *SetMaxQuotaLevelResponse, err error) {
	response = CreateSetMaxQuotaLevelResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetMaxQuotaLevelWithChan(request *SetMaxQuotaLevelRequest) (<-chan *SetMaxQuotaLevelResponse, <-chan error) {
	responseChan := make(chan *SetMaxQuotaLevelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetMaxQuotaLevel(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) SetMaxQuotaLevelWithCallback(request *SetMaxQuotaLevelRequest, callback func(response *SetMaxQuotaLevelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetMaxQuotaLevelResponse
		var err error
		defer close(result)
		response, err = client.SetMaxQuotaLevel(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type SetMaxQuotaLevelRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MaxQuotaLevel        string           `position:"Query" name:"MaxQuotaLevel"`
}

type SetMaxQuotaLevelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetMaxQuotaLevelRequest() (request *SetMaxQuotaLevelRequest) {
	request = &SetMaxQuotaLevelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "SetMaxQuotaLevel", "", "")
	return
}

func CreateSetMaxQuotaLevelResponse() (response *SetMaxQuotaLevelResponse) {
	response = &SetMaxQuotaLevelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

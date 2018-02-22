package aegis

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

func (client *Client) ModifyStrategy(request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
	response = CreateModifyStrategyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyStrategyWithChan(request *ModifyStrategyRequest) (<-chan *ModifyStrategyResponse, <-chan error) {
	responseChan := make(chan *ModifyStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyStrategy(request)
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

func (client *Client) ModifyStrategyWithCallback(request *ModifyStrategyRequest, callback func(response *ModifyStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyStrategyResponse
		var err error
		defer close(result)
		response, err = client.ModifyStrategy(request)
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

type ModifyStrategyRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CycleDays       string           `position:"Query" name:"CycleDays"`
	Name            string           `position:"Query" name:"Name"`
	CycleStartTime  string           `position:"Query" name:"CycleStartTime"`
	RiskSubTypeName string           `position:"Query" name:"RiskSubTypeName"`
	Id              string           `position:"Query" name:"Id"`
}

type ModifyStrategyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	TotalCount     int    `json:"TotalCount" xml:"TotalCount"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Result         Result `json:"Result" xml:"Result"`
}

func CreateModifyStrategyRequest() (request *ModifyStrategyRequest) {
	request = &ModifyStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyStrategy", "vipaegis", "openAPI")
	return
}

func CreateModifyStrategyResponse() (response *ModifyStrategyResponse) {
	response = &ModifyStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

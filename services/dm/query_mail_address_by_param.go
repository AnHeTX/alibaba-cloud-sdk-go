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

func (client *Client) QueryMailAddressByParam(request *QueryMailAddressByParamRequest) (response *QueryMailAddressByParamResponse, err error) {
	response = CreateQueryMailAddressByParamResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryMailAddressByParamWithChan(request *QueryMailAddressByParamRequest) (<-chan *QueryMailAddressByParamResponse, <-chan error) {
	responseChan := make(chan *QueryMailAddressByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMailAddressByParam(request)
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

func (client *Client) QueryMailAddressByParamWithCallback(request *QueryMailAddressByParamRequest, callback func(response *QueryMailAddressByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMailAddressByParamResponse
		var err error
		defer close(result)
		response, err = client.QueryMailAddressByParam(request)
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

type QueryMailAddressByParamRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	Sendtype             string           `position:"Query" name:"Sendtype"`
}

type QueryMailAddressByParamResponse struct {
	*responses.BaseResponse
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	TotalCount int                           `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                           `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                           `json:"PageSize" xml:"PageSize"`
	Data       DataInQueryMailAddressByParam `json:"data" xml:"data"`
}

func CreateQueryMailAddressByParamRequest() (request *QueryMailAddressByParamRequest) {
	request = &QueryMailAddressByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryMailAddressByParam", "", "")
	return
}

func CreateQueryMailAddressByParamResponse() (response *QueryMailAddressByParamResponse) {
	response = &QueryMailAddressByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

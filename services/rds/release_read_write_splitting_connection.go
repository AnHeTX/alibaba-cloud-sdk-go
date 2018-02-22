package rds

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

func (client *Client) ReleaseReadWriteSplittingConnection(request *ReleaseReadWriteSplittingConnectionRequest) (response *ReleaseReadWriteSplittingConnectionResponse, err error) {
	response = CreateReleaseReadWriteSplittingConnectionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ReleaseReadWriteSplittingConnectionWithChan(request *ReleaseReadWriteSplittingConnectionRequest) (<-chan *ReleaseReadWriteSplittingConnectionResponse, <-chan error) {
	responseChan := make(chan *ReleaseReadWriteSplittingConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseReadWriteSplittingConnection(request)
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

func (client *Client) ReleaseReadWriteSplittingConnectionWithCallback(request *ReleaseReadWriteSplittingConnectionRequest, callback func(response *ReleaseReadWriteSplittingConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseReadWriteSplittingConnectionResponse
		var err error
		defer close(result)
		response, err = client.ReleaseReadWriteSplittingConnection(request)
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

type ReleaseReadWriteSplittingConnectionRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

type ReleaseReadWriteSplittingConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateReleaseReadWriteSplittingConnectionRequest() (request *ReleaseReadWriteSplittingConnectionRequest) {
	request = &ReleaseReadWriteSplittingConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ReleaseReadWriteSplittingConnection", "", "")
	return
}

func CreateReleaseReadWriteSplittingConnectionResponse() (response *ReleaseReadWriteSplittingConnectionResponse) {
	response = &ReleaseReadWriteSplittingConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package ecs

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

func (client *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (response *ModifyImageSharePermissionResponse, err error) {
	response = CreateModifyImageSharePermissionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyImageSharePermissionWithChan(request *ModifyImageSharePermissionRequest) (<-chan *ModifyImageSharePermissionResponse, <-chan error) {
	responseChan := make(chan *ModifyImageSharePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyImageSharePermission(request)
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

func (client *Client) ModifyImageSharePermissionWithCallback(request *ModifyImageSharePermissionRequest, callback func(response *ModifyImageSharePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyImageSharePermissionResponse
		var err error
		defer close(result)
		response, err = client.ModifyImageSharePermission(request)
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

type ModifyImageSharePermissionRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImageId              string           `position:"Query" name:"ImageId"`
	AddAccount1          string           `position:"Query" name:"AddAccount.1"`
	AddAccount2          string           `position:"Query" name:"AddAccount.2"`
	AddAccount3          string           `position:"Query" name:"AddAccount.3"`
	AddAccount4          string           `position:"Query" name:"AddAccount.4"`
	AddAccount5          string           `position:"Query" name:"AddAccount.5"`
	AddAccount6          string           `position:"Query" name:"AddAccount.6"`
	AddAccount7          string           `position:"Query" name:"AddAccount.7"`
	AddAccount8          string           `position:"Query" name:"AddAccount.8"`
	AddAccount9          string           `position:"Query" name:"AddAccount.9"`
	AddAccount10         string           `position:"Query" name:"AddAccount.10"`
	RemoveAccount1       string           `position:"Query" name:"RemoveAccount.1"`
	RemoveAccount2       string           `position:"Query" name:"RemoveAccount.2"`
	RemoveAccount3       string           `position:"Query" name:"RemoveAccount.3"`
	RemoveAccount4       string           `position:"Query" name:"RemoveAccount.4"`
	RemoveAccount5       string           `position:"Query" name:"RemoveAccount.5"`
	RemoveAccount6       string           `position:"Query" name:"RemoveAccount.6"`
	RemoveAccount7       string           `position:"Query" name:"RemoveAccount.7"`
	RemoveAccount8       string           `position:"Query" name:"RemoveAccount.8"`
	RemoveAccount9       string           `position:"Query" name:"RemoveAccount.9"`
	RemoveAccount10      string           `position:"Query" name:"RemoveAccount.10"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ModifyImageSharePermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
	request = &ModifyImageSharePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageSharePermission", "", "")
	return
}

func CreateModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
	response = &ModifyImageSharePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

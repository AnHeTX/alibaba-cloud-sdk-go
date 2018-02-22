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

func (client *Client) ModifyInstanceAutoRenewalAttribute(request *ModifyInstanceAutoRenewalAttributeRequest) (response *ModifyInstanceAutoRenewalAttributeResponse, err error) {
	response = CreateModifyInstanceAutoRenewalAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyInstanceAutoRenewalAttributeWithChan(request *ModifyInstanceAutoRenewalAttributeRequest) (<-chan *ModifyInstanceAutoRenewalAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceAutoRenewalAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceAutoRenewalAttribute(request)
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

func (client *Client) ModifyInstanceAutoRenewalAttributeWithCallback(request *ModifyInstanceAutoRenewalAttributeRequest, callback func(response *ModifyInstanceAutoRenewalAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceAutoRenewalAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceAutoRenewalAttribute(request)
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

type ModifyInstanceAutoRenewalAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Duration             string           `position:"Query" name:"Duration"`
	AutoRenew            string           `position:"Query" name:"AutoRenew"`
}

type ModifyInstanceAutoRenewalAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyInstanceAutoRenewalAttributeRequest() (request *ModifyInstanceAutoRenewalAttributeRequest) {
	request = &ModifyInstanceAutoRenewalAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyInstanceAutoRenewalAttribute", "", "")
	return
}

func CreateModifyInstanceAutoRenewalAttributeResponse() (response *ModifyInstanceAutoRenewalAttributeResponse) {
	response = &ModifyInstanceAutoRenewalAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

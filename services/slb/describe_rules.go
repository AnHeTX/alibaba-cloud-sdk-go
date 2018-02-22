package slb

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

func (client *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
	response = CreateDescribeRulesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRulesWithChan(request *DescribeRulesRequest) (<-chan *DescribeRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRules(request)
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

func (client *Client) DescribeRulesWithCallback(request *DescribeRulesRequest, callback func(response *DescribeRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRules(request)
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

type DescribeRulesRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
}

type DescribeRulesResponse struct {
	*responses.BaseResponse
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Rules     RulesInDescribeRules `json:"Rules" xml:"Rules"`
}

func CreateDescribeRulesRequest() (request *DescribeRulesRequest) {
	request = &DescribeRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeRules", "", "")
	return
}

func CreateDescribeRulesResponse() (response *DescribeRulesResponse) {
	response = &DescribeRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

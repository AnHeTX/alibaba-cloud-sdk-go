package vpc

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

func (client *Client) AddBandwidthPackageIps(request *AddBandwidthPackageIpsRequest) (response *AddBandwidthPackageIpsResponse, err error) {
	response = CreateAddBandwidthPackageIpsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddBandwidthPackageIpsWithChan(request *AddBandwidthPackageIpsRequest) (<-chan *AddBandwidthPackageIpsResponse, <-chan error) {
	responseChan := make(chan *AddBandwidthPackageIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBandwidthPackageIps(request)
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

func (client *Client) AddBandwidthPackageIpsWithCallback(request *AddBandwidthPackageIpsRequest, callback func(response *AddBandwidthPackageIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBandwidthPackageIpsResponse
		var err error
		defer close(result)
		response, err = client.AddBandwidthPackageIps(request)
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

type AddBandwidthPackageIpsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
	IpCount              string           `position:"Query" name:"IpCount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
}

type AddBandwidthPackageIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddBandwidthPackageIpsRequest() (request *AddBandwidthPackageIpsRequest) {
	request = &AddBandwidthPackageIpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AddBandwidthPackageIps", "", "")
	return
}

func CreateAddBandwidthPackageIpsResponse() (response *AddBandwidthPackageIpsResponse) {
	response = &AddBandwidthPackageIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

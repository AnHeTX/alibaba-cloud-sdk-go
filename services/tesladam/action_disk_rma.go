package tesladam

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

func (client *Client) ActionDiskRma(request *ActionDiskRmaRequest) (response *ActionDiskRmaResponse, err error) {
	response = CreateActionDiskRmaResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ActionDiskRmaWithChan(request *ActionDiskRmaRequest) (<-chan *ActionDiskRmaResponse, <-chan error) {
	responseChan := make(chan *ActionDiskRmaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActionDiskRma(request)
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

func (client *Client) ActionDiskRmaWithCallback(request *ActionDiskRmaRequest, callback func(response *ActionDiskRmaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActionDiskRmaResponse
		var err error
		defer close(result)
		response, err = client.ActionDiskRma(request)
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

type ActionDiskRmaRequest struct {
	*requests.RpcRequest
	Hostname    string `position:"Query" name:"Hostname"`
	DiskMount   string `position:"Query" name:"DiskMount"`
	ExecutionId string `position:"Query" name:"ExecutionId"`
	DiskSlot    string `position:"Query" name:"DiskSlot"`
	DiskName    string `position:"Query" name:"DiskName"`
	DiskSn      string `position:"Query" name:"DiskSn"`
	DiskReason  string `position:"Query" name:"DiskReason"`
}

type ActionDiskRmaResponse struct {
	*responses.BaseResponse
	Status  bool   `json:"Status" xml:"Status"`
	Message string `json:"Message" xml:"Message"`
	Result  string `json:"Result" xml:"Result"`
}

func CreateActionDiskRmaRequest() (request *ActionDiskRmaRequest) {
	request = &ActionDiskRmaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaDam", "2018-01-18", "ActionDiskRma", "", "")
	return
}

func CreateActionDiskRmaResponse() (response *ActionDiskRmaResponse) {
	response = &ActionDiskRmaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

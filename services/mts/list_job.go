package mts

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

func (client *Client) ListJob(request *ListJobRequest) (response *ListJobResponse, err error) {
	response = CreateListJobResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListJobWithChan(request *ListJobRequest) (<-chan *ListJobResponse, <-chan error) {
	responseChan := make(chan *ListJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJob(request)
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

func (client *Client) ListJobWithCallback(request *ListJobRequest, callback func(response *ListJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobResponse
		var err error
		defer close(result)
		response, err = client.ListJob(request)
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

type ListJobRequest struct {
	*requests.RpcRequest
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextPageToken              string           `position:"Query" name:"NextPageToken"`
	MaximumPageSize            requests.Integer `position:"Query" name:"MaximumPageSize"`
	State                      string           `position:"Query" name:"State"`
	StartOfJobCreatedTimeRange string           `position:"Query" name:"StartOfJobCreatedTimeRange"`
	EndOfJobCreatedTimeRange   string           `position:"Query" name:"EndOfJobCreatedTimeRange"`
	PipelineId                 string           `position:"Query" name:"PipelineId"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
}

type ListJobResponse struct {
	*responses.BaseResponse
	RequestId     string           `json:"RequestId" xml:"RequestId"`
	NextPageToken string           `json:"NextPageToken" xml:"NextPageToken"`
	JobList       JobListInListJob `json:"JobList" xml:"JobList"`
}

func CreateListJobRequest() (request *ListJobRequest) {
	request = &ListJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListJob", "", "")
	return
}

func CreateListJobResponse() (response *ListJobResponse) {
	response = &ListJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

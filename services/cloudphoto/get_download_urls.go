package cloudphoto

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

func (client *Client) GetDownloadUrls(request *GetDownloadUrlsRequest) (response *GetDownloadUrlsResponse, err error) {
	response = CreateGetDownloadUrlsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetDownloadUrlsWithChan(request *GetDownloadUrlsRequest) (<-chan *GetDownloadUrlsResponse, <-chan error) {
	responseChan := make(chan *GetDownloadUrlsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDownloadUrls(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) GetDownloadUrlsWithCallback(request *GetDownloadUrlsRequest, callback func(response *GetDownloadUrlsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDownloadUrlsResponse
		var err error
		defer close(result)
		response, err = client.GetDownloadUrls(request)
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

type GetDownloadUrlsRequest struct {
	*requests.RpcRequest
	LibraryId string    `position:"Query" name:"LibraryId"`
	StoreName string    `position:"Query" name:"StoreName"`
	PhotoId   *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
}

type GetDownloadUrlsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Results   struct {
		Result []struct {
			Code        string           `json:"Code" xml:"Code"`
			Message     string           `json:"Message" xml:"Message"`
			PhotoId     requests.Integer `json:"PhotoId" xml:"PhotoId"`
			DownloadUrl string           `json:"DownloadUrl" xml:"DownloadUrl"`
		} `json:"Result" xml:"Result"`
	} `json:"Results" xml:"Results"`
}

func CreateGetDownloadUrlsRequest() (request *GetDownloadUrlsRequest) {
	request = &GetDownloadUrlsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetDownloadUrls", "", "")
	return
}

func CreateGetDownloadUrlsResponse() (response *GetDownloadUrlsResponse) {
	response = &GetDownloadUrlsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
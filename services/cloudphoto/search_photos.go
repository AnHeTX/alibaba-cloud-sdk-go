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

func (client *Client) SearchPhotos(request *SearchPhotosRequest) (response *SearchPhotosResponse, err error) {
	response = CreateSearchPhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SearchPhotosWithChan(request *SearchPhotosRequest) (<-chan *SearchPhotosResponse, <-chan error) {
	responseChan := make(chan *SearchPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchPhotos(request)
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

func (client *Client) SearchPhotosWithCallback(request *SearchPhotosRequest, callback func(response *SearchPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchPhotosResponse
		var err error
		defer close(result)
		response, err = client.SearchPhotos(request)
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

type SearchPhotosRequest struct {
	*requests.RpcRequest
	Size      requests.Integer `position:"Query" name:"Size"`
	LibraryId string           `position:"Query" name:"LibraryId"`
	StoreName string           `position:"Query" name:"StoreName"`
	Page      requests.Integer `position:"Query" name:"Page"`
	Keyword   string           `position:"Query" name:"Keyword"`
}

type SearchPhotosResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Action     string `json:"Action" xml:"Action"`
	Photos     []struct {
		Id              int    `json:"Id" xml:"Id"`
		Title           string `json:"Title" xml:"Title"`
		FileId          string `json:"FileId" xml:"FileId"`
		Location        string `json:"Location" xml:"Location"`
		State           string `json:"State" xml:"State"`
		Md5             string `json:"Md5" xml:"Md5"`
		IsVideo         bool   `json:"IsVideo" xml:"IsVideo"`
		Size            int    `json:"Size" xml:"Size"`
		Width           int    `json:"Width" xml:"Width"`
		Height          int    `json:"Height" xml:"Height"`
		Ctime           int    `json:"Ctime" xml:"Ctime"`
		Mtime           int    `json:"Mtime" xml:"Mtime"`
		TakenAt         int    `json:"TakenAt" xml:"TakenAt"`
		ShareExpireTime int    `json:"ShareExpireTime" xml:"ShareExpireTime"`
	} `json:"Photos" xml:"Photos"`
}

func CreateSearchPhotosRequest() (request *SearchPhotosRequest) {
	request = &SearchPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "SearchPhotos", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateSearchPhotosResponse() (response *SearchPhotosResponse) {
	response = &SearchPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

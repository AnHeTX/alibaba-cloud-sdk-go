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

func (client *Client) RenameAlbum(request *RenameAlbumRequest) (response *RenameAlbumResponse, err error) {
	response = CreateRenameAlbumResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RenameAlbumWithChan(request *RenameAlbumRequest) (<-chan *RenameAlbumResponse, <-chan error) {
	responseChan := make(chan *RenameAlbumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenameAlbum(request)
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

func (client *Client) RenameAlbumWithCallback(request *RenameAlbumRequest, callback func(response *RenameAlbumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenameAlbumResponse
		var err error
		defer close(result)
		response, err = client.RenameAlbum(request)
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

type RenameAlbumRequest struct {
	*requests.RpcRequest
	AlbumName string           `position:"Query" name:"AlbumName"`
	LibraryId string           `position:"Query" name:"LibraryId"`
	AlbumId   requests.Integer `position:"Query" name:"AlbumId"`
	StoreName string           `position:"Query" name:"StoreName"`
}

type RenameAlbumResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

func CreateRenameAlbumRequest() (request *RenameAlbumRequest) {
	request = &RenameAlbumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "RenameAlbum", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateRenameAlbumResponse() (response *RenameAlbumResponse) {
	response = &RenameAlbumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

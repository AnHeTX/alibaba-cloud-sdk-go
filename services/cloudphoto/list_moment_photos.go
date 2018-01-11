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

func (client *Client) ListMomentPhotos(request *ListMomentPhotosRequest) (response *ListMomentPhotosResponse, err error) {
	response = CreateListMomentPhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListMomentPhotosWithChan(request *ListMomentPhotosRequest) (<-chan *ListMomentPhotosResponse, <-chan error) {
	responseChan := make(chan *ListMomentPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMomentPhotos(request)
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

func (client *Client) ListMomentPhotosWithCallback(request *ListMomentPhotosRequest, callback func(response *ListMomentPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMomentPhotosResponse
		var err error
		defer close(result)
		response, err = client.ListMomentPhotos(request)
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

type ListMomentPhotosRequest struct {
	*requests.RpcRequest
	Cursor    string           `position:"Query" name:"Cursor"`
	Size      requests.Integer `position:"Query" name:"Size"`
	LibraryId string           `position:"Query" name:"LibraryId"`
	StoreName string           `position:"Query" name:"StoreName"`
	State     string           `position:"Query" name:"State"`
	MomentId  requests.Integer `position:"Query" name:"MomentId"`
	Direction string           `position:"Query" name:"Direction"`
}

type ListMomentPhotosResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	NextCursor string `json:"NextCursor" xml:"NextCursor"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Action     string `json:"Action" xml:"Action"`
	Results    []struct {
		PhotoId int    `json:"PhotoId" xml:"PhotoId"`
		State   string `json:"State" xml:"State"`
	} `json:"Results" xml:"Results"`
}

func CreateListMomentPhotosRequest() (request *ListMomentPhotosRequest) {
	request = &ListMomentPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListMomentPhotos", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateListMomentPhotosResponse() (response *ListMomentPhotosResponse) {
	response = &ListMomentPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

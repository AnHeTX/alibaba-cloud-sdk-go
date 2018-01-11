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

func (client *Client) ListTimeLines(request *ListTimeLinesRequest) (response *ListTimeLinesResponse, err error) {
	response = CreateListTimeLinesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListTimeLinesWithChan(request *ListTimeLinesRequest) (<-chan *ListTimeLinesResponse, <-chan error) {
	responseChan := make(chan *ListTimeLinesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTimeLines(request)
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

func (client *Client) ListTimeLinesWithCallback(request *ListTimeLinesRequest, callback func(response *ListTimeLinesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTimeLinesResponse
		var err error
		defer close(result)
		response, err = client.ListTimeLines(request)
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

type ListTimeLinesRequest struct {
	*requests.RpcRequest
	Cursor        requests.Integer `position:"Query" name:"Cursor"`
	PhotoSize     requests.Integer `position:"Query" name:"PhotoSize"`
	TimeLineCount requests.Integer `position:"Query" name:"TimeLineCount"`
	LibraryId     string           `position:"Query" name:"LibraryId"`
	StoreName     string           `position:"Query" name:"StoreName"`
	TimeLineUnit  string           `position:"Query" name:"TimeLineUnit"`
	FilterBy      string           `position:"Query" name:"FilterBy"`
	Direction     string           `position:"Query" name:"Direction"`
	Order         string           `position:"Query" name:"Order"`
}

type ListTimeLinesResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	NextCursor int    `json:"NextCursor" xml:"NextCursor"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Action     string `json:"Action" xml:"Action"`
	TimeLines  []struct {
		StartTime   int `json:"StartTime" xml:"StartTime"`
		EndTime     int `json:"EndTime" xml:"EndTime"`
		TotalCount  int `json:"TotalCount" xml:"TotalCount"`
		PhotosCount int `json:"PhotosCount" xml:"PhotosCount"`
		Photos      []struct {
			Id              int    `json:"Id" xml:"Id"`
			Title           string `json:"Title" xml:"Title"`
			Location        string `json:"Location" xml:"Location"`
			FileId          string `json:"FileId" xml:"FileId"`
			State           string `json:"State" xml:"State"`
			Md5             string `json:"Md5" xml:"Md5"`
			IsVideo         bool   `json:"IsVideo" xml:"IsVideo"`
			Remark          string `json:"Remark" xml:"Remark"`
			Size            int    `json:"Size" xml:"Size"`
			Width           int    `json:"Width" xml:"Width"`
			Height          int    `json:"Height" xml:"Height"`
			Ctime           int    `json:"Ctime" xml:"Ctime"`
			Mtime           int    `json:"Mtime" xml:"Mtime"`
			TakenAt         int    `json:"TakenAt" xml:"TakenAt"`
			ShareExpireTime int    `json:"ShareExpireTime" xml:"ShareExpireTime"`
			Like            int    `json:"Like" xml:"Like"`
		} `json:"Photos" xml:"Photos"`
	} `json:"TimeLines" xml:"TimeLines"`
}

func CreateListTimeLinesRequest() (request *ListTimeLinesRequest) {
	request = &ListTimeLinesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTimeLines", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateListTimeLinesResponse() (response *ListTimeLinesResponse) {
	response = &ListTimeLinesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

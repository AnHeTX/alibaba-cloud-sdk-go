package cms

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

// TaskConfigCreate invokes the cms.TaskConfigCreate API synchronously
// api document: https://help.aliyun.com/api/cms/taskconfigcreate.html
func (client *Client) TaskConfigCreate(request *TaskConfigCreateRequest) (response *TaskConfigCreateResponse, err error) {
	response = CreateTaskConfigCreateResponse()
	err = client.DoAction(request, response)
	return
}

// TaskConfigCreateWithChan invokes the cms.TaskConfigCreate API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfigcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigCreateWithChan(request *TaskConfigCreateRequest) (<-chan *TaskConfigCreateResponse, <-chan error) {
	responseChan := make(chan *TaskConfigCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaskConfigCreate(request)
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

// TaskConfigCreateWithCallback invokes the cms.TaskConfigCreate API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfigcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigCreateWithCallback(request *TaskConfigCreateRequest, callback func(response *TaskConfigCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaskConfigCreateResponse
		var err error
		defer close(result)
		response, err = client.TaskConfigCreate(request)
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

// TaskConfigCreateRequest is the request struct for api TaskConfigCreate
type TaskConfigCreateRequest struct {
	*requests.RpcRequest
	InstanceList *[]string        `position:"Query" name:"InstanceList"  type:"Repeated"`
	JsonData     string           `position:"Query" name:"JsonData"`
	TaskType     string           `position:"Query" name:"TaskType"`
	TaskScope    string           `position:"Query" name:"TaskScope"`
	AlertConfig  string           `position:"Query" name:"AlertConfig"`
	GroupId      requests.Integer `position:"Query" name:"GroupId"`
	TaskName     string           `position:"Query" name:"TaskName"`
	GroupName    string           `position:"Query" name:"GroupName"`
}

// TaskConfigCreateResponse is the response struct for api TaskConfigCreate
type TaskConfigCreateResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TaskId       int    `json:"TaskId" xml:"TaskId"`
}

// CreateTaskConfigCreateRequest creates a request to invoke TaskConfigCreate API
func CreateTaskConfigCreateRequest() (request *TaskConfigCreateRequest) {
	request = &TaskConfigCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigCreate", "cms", "openAPI")
	return
}

// CreateTaskConfigCreateResponse creates a response to parse from TaskConfigCreate response
func CreateTaskConfigCreateResponse() (response *TaskConfigCreateResponse) {
	response = &TaskConfigCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package rds

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

func (client *Client) CreateSQLDiagnosis(request *CreateSQLDiagnosisRequest) (response *CreateSQLDiagnosisResponse, err error) {
	response = CreateCreateSQLDiagnosisResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateSQLDiagnosisWithChan(request *CreateSQLDiagnosisRequest) (<-chan *CreateSQLDiagnosisResponse, <-chan error) {
	responseChan := make(chan *CreateSQLDiagnosisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSQLDiagnosis(request)
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

func (client *Client) CreateSQLDiagnosisWithCallback(request *CreateSQLDiagnosisRequest, callback func(response *CreateSQLDiagnosisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSQLDiagnosisResponse
		var err error
		defer close(result)
		response, err = client.CreateSQLDiagnosis(request)
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

type CreateSQLDiagnosisRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
	EndTime      string `position:"Query" name:"EndTime"`
}

type CreateSQLDiagnosisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SQLDiagId string `json:"SQLDiagId" xml:"SQLDiagId"`
}

func CreateCreateSQLDiagnosisRequest() (request *CreateSQLDiagnosisRequest) {
	request = &CreateSQLDiagnosisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateSQLDiagnosis", "", "")
	return
}

func CreateCreateSQLDiagnosisResponse() (response *CreateSQLDiagnosisResponse) {
	response = &CreateSQLDiagnosisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package slb

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

func (client *Client) DescribeLoadBalancerUDPListenerAttribute(request *DescribeLoadBalancerUDPListenerAttributeRequest) (response *DescribeLoadBalancerUDPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerUDPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithChan(request *DescribeLoadBalancerUDPListenerAttributeRequest) (<-chan *DescribeLoadBalancerUDPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerUDPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerUDPListenerAttribute(request)
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

func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithCallback(request *DescribeLoadBalancerUDPListenerAttributeRequest, callback func(response *DescribeLoadBalancerUDPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerUDPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerUDPListenerAttribute(request)
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

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
}

type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	ListenerPort              int    `json:"ListenerPort" xml:"ListenerPort"`
	BackendServerPort         int    `json:"BackendServerPort" xml:"BackendServerPort"`
	Status                    string `json:"Status" xml:"Status"`
	Bandwidth                 int    `json:"Bandwidth" xml:"Bandwidth"`
	Scheduler                 string `json:"Scheduler" xml:"Scheduler"`
	PersistenceTimeout        int    `json:"PersistenceTimeout" xml:"PersistenceTimeout"`
	HealthCheck               string `json:"HealthCheck" xml:"HealthCheck"`
	HealthyThreshold          int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold        int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckInterval       int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckReq            string `json:"HealthCheckReq" xml:"HealthCheckReq"`
	HealthCheckExp            string `json:"HealthCheckExp" xml:"HealthCheckExp"`
	MaxConnection             int    `json:"MaxConnection" xml:"MaxConnection"`
	VServerGroupId            string `json:"VServerGroupId" xml:"VServerGroupId"`
	MasterSlaveServerGroupId  string `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
}

func CreateDescribeLoadBalancerUDPListenerAttributeRequest() (request *DescribeLoadBalancerUDPListenerAttributeRequest) {
	request = &DescribeLoadBalancerUDPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerUDPListenerAttribute", "", "")
	return
}

func CreateDescribeLoadBalancerUDPListenerAttributeResponse() (response *DescribeLoadBalancerUDPListenerAttributeResponse) {
	response = &DescribeLoadBalancerUDPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

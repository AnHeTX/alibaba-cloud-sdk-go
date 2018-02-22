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

func (client *Client) SetLoadBalancerHTTPSListenerAttribute(request *SetLoadBalancerHTTPSListenerAttributeRequest) (response *SetLoadBalancerHTTPSListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerHTTPSListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithChan(request *SetLoadBalancerHTTPSListenerAttributeRequest) (<-chan *SetLoadBalancerHTTPSListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerHTTPSListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerHTTPSListenerAttribute(request)
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

func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithCallback(request *SetLoadBalancerHTTPSListenerAttributeRequest, callback func(response *SetLoadBalancerHTTPSListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerHTTPSListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerHTTPSListenerAttribute(request)
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

type SetLoadBalancerHTTPSListenerAttributeRequest struct {
	*requests.RpcRequest
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId         string           `position:"Query" name:"LoadBalancerId"`
	ListenerPort           requests.Integer `position:"Query" name:"ListenerPort"`
	Bandwidth              requests.Integer `position:"Query" name:"Bandwidth"`
	XForwardedFor          string           `position:"Query" name:"XForwardedFor"`
	Scheduler              string           `position:"Query" name:"Scheduler"`
	StickySession          string           `position:"Query" name:"StickySession"`
	StickySessionType      string           `position:"Query" name:"StickySessionType"`
	CookieTimeout          requests.Integer `position:"Query" name:"CookieTimeout"`
	Cookie                 string           `position:"Query" name:"Cookie"`
	HealthCheck            string           `position:"Query" name:"HealthCheck"`
	HealthCheckDomain      string           `position:"Query" name:"HealthCheckDomain"`
	HealthCheckURI         string           `position:"Query" name:"HealthCheckURI"`
	HealthyThreshold       requests.Integer `position:"Query" name:"HealthyThreshold"`
	UnhealthyThreshold     requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthCheckTimeout     requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	HealthCheckInterval    requests.Integer `position:"Query" name:"HealthCheckInterval"`
	HealthCheckConnectPort requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string           `position:"Query" name:"HealthCheckHttpCode"`
	MaxConnection          requests.Integer `position:"Query" name:"MaxConnection"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId            string           `position:"Query" name:"access_key_id"`
	ServerCertificateId    string           `position:"Query" name:"ServerCertificateId"`
	CACertificateId        string           `position:"Query" name:"CACertificateId"`
	VServerGroup           string           `position:"Query" name:"VServerGroup"`
	VServerGroupId         string           `position:"Query" name:"VServerGroupId"`
	Tags                   string           `position:"Query" name:"Tags"`
	XForwardedForSLBIP     string           `position:"Query" name:"XForwardedFor_SLBIP"`
	XForwardedForSLBID     string           `position:"Query" name:"XForwardedFor_SLBID"`
	XForwardedForProto     string           `position:"Query" name:"XForwardedFor_proto"`
	Gzip                   string           `position:"Query" name:"Gzip"`
}

type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLoadBalancerHTTPSListenerAttributeRequest() (request *SetLoadBalancerHTTPSListenerAttributeRequest) {
	request = &SetLoadBalancerHTTPSListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerHTTPSListenerAttribute", "", "")
	return
}

func CreateSetLoadBalancerHTTPSListenerAttributeResponse() (response *SetLoadBalancerHTTPSListenerAttributeResponse) {
	response = &SetLoadBalancerHTTPSListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

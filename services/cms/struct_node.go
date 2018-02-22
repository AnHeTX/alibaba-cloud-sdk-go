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

type Node struct {
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	SerialNumber     string `json:"SerialNumber" xml:"SerialNumber"`
	HostName         string `json:"HostName" xml:"HostName"`
	AliUid           int    `json:"AliUid" xml:"AliUid"`
	OperatingSystem  string `json:"OperatingSystem" xml:"OperatingSystem"`
	IpGroup          string `json:"IpGroup" xml:"IpGroup"`
	Region           string `json:"Region" xml:"Region"`
	TianjimonVersion string `json:"TianjimonVersion" xml:"TianjimonVersion"`
	EipAddress       string `json:"EipAddress" xml:"EipAddress"`
	EipId            string `json:"EipId" xml:"EipId"`
	AliyunHost       bool   `json:"AliyunHost" xml:"AliyunHost"`
	NatIp            string `json:"NatIp" xml:"NatIp"`
	NetworkType      string `json:"NetworkType" xml:"NetworkType"`
}

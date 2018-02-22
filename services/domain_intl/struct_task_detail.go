package domain_intl

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

type TaskDetail struct {
	TaskNo       string `json:"TaskNo" xml:"TaskNo"`
	TaskDetailNo string `json:"TaskDetailNo" xml:"TaskDetailNo"`
	TaskType     string `json:"TaskType" xml:"TaskType"`
	InstanceId   string `json:"InstanceId" xml:"InstanceId"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
	TaskStatus   string `json:"TaskStatus" xml:"TaskStatus"`
	UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
	CreateTime   string `json:"CreateTime" xml:"CreateTime"`
	TryCount     int    `json:"TryCount" xml:"TryCount"`
	ErrorMsg     string `json:"ErrorMsg" xml:"ErrorMsg"`
}

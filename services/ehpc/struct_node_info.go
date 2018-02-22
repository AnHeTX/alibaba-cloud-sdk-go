package ehpc

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

type NodeInfo struct {
	ExpiredTime     string         `json:"ExpiredTime" xml:"ExpiredTime"`
	LockReason      string         `json:"LockReason" xml:"LockReason"`
	CreatedByEhpc   bool           `json:"CreatedByEhpc" xml:"CreatedByEhpc"`
	AddTime         string         `json:"AddTime" xml:"AddTime"`
	ImageOwnerAlias string         `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
	Expired         bool           `json:"Expired" xml:"Expired"`
	SpotStrategy    string         `json:"SpotStrategy" xml:"SpotStrategy"`
	ImageId         string         `json:"ImageId" xml:"ImageId"`
	Id              string         `json:"Id" xml:"Id"`
	RegionId        string         `json:"RegionId" xml:"RegionId"`
	Role            string         `json:"Role" xml:"Role"`
	Status          string         `json:"Status" xml:"Status"`
	UsedResources   UsedResources  `json:"UsedResources" xml:"UsedResources"`
	TotalResources  TotalResources `json:"TotalResources" xml:"TotalResources"`
}

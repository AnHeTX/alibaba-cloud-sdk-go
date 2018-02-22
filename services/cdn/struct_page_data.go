package cdn

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

type PageData struct {
	DomainName      string                       `json:"DomainName" xml:"DomainName"`
	Cname           string                       `json:"Cname" xml:"Cname"`
	CdnType         string                       `json:"CdnType" xml:"CdnType"`
	DomainStatus    string                       `json:"DomainStatus" xml:"DomainStatus"`
	GmtCreated      string                       `json:"GmtCreated" xml:"GmtCreated"`
	GmtModified     string                       `json:"GmtModified" xml:"GmtModified"`
	Description     string                       `json:"Description" xml:"Description"`
	SourceType      string                       `json:"SourceType" xml:"SourceType"`
	SslProtocol     string                       `json:"SslProtocol" xml:"SslProtocol"`
	ResourceGroupId string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Sandbox         string                       `json:"Sandbox" xml:"Sandbox"`
	Sources         SourcesInDescribeUserDomains `json:"Sources" xml:"Sources"`
}

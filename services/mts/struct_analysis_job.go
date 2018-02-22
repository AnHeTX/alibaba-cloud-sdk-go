package mts

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

type AnalysisJob struct {
	CreationTime     string                             `json:"CreationTime" xml:"CreationTime"`
	Message          string                             `json:"Message" xml:"Message"`
	PipelineId       string                             `json:"PipelineId" xml:"PipelineId"`
	State            string                             `json:"State" xml:"State"`
	Code             string                             `json:"Code" xml:"Code"`
	UserData         string                             `json:"UserData" xml:"UserData"`
	Priority         string                             `json:"Priority" xml:"Priority"`
	Id               string                             `json:"Id" xml:"Id"`
	Percent          int                                `json:"Percent" xml:"Percent"`
	MNSMessageResult MNSMessageResult                   `json:"MNSMessageResult" xml:"MNSMessageResult"`
	InputFile        InputFile                          `json:"InputFile" xml:"InputFile"`
	AnalysisConfig   AnalysisConfig                     `json:"AnalysisConfig" xml:"AnalysisConfig"`
	TemplateList     TemplateListInQueryAnalysisJobList `json:"TemplateList" xml:"TemplateList"`
}

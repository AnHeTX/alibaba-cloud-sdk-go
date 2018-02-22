package ccc

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

type CallDetailRecord struct {
	ContactId          string                            `json:"ContactId" xml:"ContactId"`
	StartTime          int                               `json:"StartTime" xml:"StartTime"`
	Duration           int                               `json:"Duration" xml:"Duration"`
	Satisfaction       int                               `json:"Satisfaction" xml:"Satisfaction"`
	ContactType        string                            `json:"ContactType" xml:"ContactType"`
	ContactDisposition string                            `json:"ContactDisposition" xml:"ContactDisposition"`
	CallingNumber      string                            `json:"CallingNumber" xml:"CallingNumber"`
	CalledNumber       string                            `json:"CalledNumber" xml:"CalledNumber"`
	AgentNames         string                            `json:"AgentNames" xml:"AgentNames"`
	SkillGroupNames    string                            `json:"SkillGroupNames" xml:"SkillGroupNames"`
	InstanceId         string                            `json:"InstanceId" xml:"InstanceId"`
	ExtraAttr          string                            `json:"ExtraAttr" xml:"ExtraAttr"`
	Agents             Agents                            `json:"Agents" xml:"Agents"`
	Recordings         RecordingsInListCallDetailRecords `json:"Recordings" xml:"Recordings"`
}

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) QueryMediaWorkflowExecutionList(request *QueryMediaWorkflowExecutionListRequest) (response *QueryMediaWorkflowExecutionListResponse, err error) {
	response = CreateQueryMediaWorkflowExecutionListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryMediaWorkflowExecutionListWithChan(request *QueryMediaWorkflowExecutionListRequest) (<-chan *QueryMediaWorkflowExecutionListResponse, <-chan error) {
	responseChan := make(chan *QueryMediaWorkflowExecutionListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaWorkflowExecutionList(request)
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

func (client *Client) QueryMediaWorkflowExecutionListWithCallback(request *QueryMediaWorkflowExecutionListRequest, callback func(response *QueryMediaWorkflowExecutionListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaWorkflowExecutionListResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaWorkflowExecutionList(request)
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

type QueryMediaWorkflowExecutionListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RunIds               string           `position:"Query" name:"RunIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryMediaWorkflowExecutionListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	NonExistRunIds struct {
		RunId []string `json:"RunId" xml:"RunId"`
	} `json:"NonExistRunIds" xml:"NonExistRunIds"`
	MediaWorkflowExecutionList struct {
		MediaWorkflowExecution []struct {
			RunId           string `json:"RunId" xml:"RunId"`
			MediaWorkflowId string `json:"MediaWorkflowId" xml:"MediaWorkflowId"`
			Name            string `json:"Name" xml:"Name"`
			State           string `json:"State" xml:"State"`
			MediaId         string `json:"MediaId" xml:"MediaId"`
			CreationTime    string `json:"CreationTime" xml:"CreationTime"`
			Input           struct {
				UserData  string `json:"UserData" xml:"UserData"`
				InputFile struct {
					Bucket   string `json:"Bucket" xml:"Bucket"`
					Location string `json:"Location" xml:"Location"`
					Object   string `json:"Object" xml:"Object"`
				} `json:"InputFile" xml:"InputFile"`
			} `json:"Input" xml:"Input"`
			ActivityList struct {
				Activity []struct {
					Name             string `json:"Name" xml:"Name"`
					Type             string `json:"Type" xml:"Type"`
					JobId            string `json:"JobId" xml:"JobId"`
					State            string `json:"State" xml:"State"`
					Code             string `json:"Code" xml:"Code"`
					Message          string `json:"Message" xml:"Message"`
					StartTime        string `json:"StartTime" xml:"StartTime"`
					EndTime          string `json:"EndTime" xml:"EndTime"`
					MNSMessageResult struct {
						MessageId    string `json:"MessageId" xml:"MessageId"`
						ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
						ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
					} `json:"MNSMessageResult" xml:"MNSMessageResult"`
				} `json:"Activity" xml:"Activity"`
			} `json:"ActivityList" xml:"ActivityList"`
		} `json:"MediaWorkflowExecution" xml:"MediaWorkflowExecution"`
	} `json:"MediaWorkflowExecutionList" xml:"MediaWorkflowExecutionList"`
}

func CreateQueryMediaWorkflowExecutionListRequest() (request *QueryMediaWorkflowExecutionListRequest) {
	request = &QueryMediaWorkflowExecutionListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaWorkflowExecutionList", "mts", "openAPI")
	return
}

func CreateQueryMediaWorkflowExecutionListResponse() (response *QueryMediaWorkflowExecutionListResponse) {
	response = &QueryMediaWorkflowExecutionListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package ecs

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

func (client *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	response = CreateDescribeTasksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeTasksWithChan(request *DescribeTasksRequest) (<-chan *DescribeTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTasks(request)
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

func (client *Client) DescribeTasksWithCallback(request *DescribeTasksRequest, callback func(response *DescribeTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeTasks(request)
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

type DescribeTasksRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TaskAction           string           `position:"Query" name:"TaskAction"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TaskIds              string           `position:"Query" name:"TaskIds"`
	TaskStatus           string           `position:"Query" name:"TaskStatus"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeTasksResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	RegionId   string `json:"RegionId" xml:"RegionId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	TaskSet    struct {
		Task []struct {
			TaskId        string `json:"TaskId" xml:"TaskId"`
			TaskAction    string `json:"TaskAction" xml:"TaskAction"`
			TaskStatus    string `json:"TaskStatus" xml:"TaskStatus"`
			SupportCancel string `json:"SupportCancel" xml:"SupportCancel"`
			CreationTime  string `json:"CreationTime" xml:"CreationTime"`
			FinishedTime  string `json:"FinishedTime" xml:"FinishedTime"`
		} `json:"Task" xml:"Task"`
	} `json:"TaskSet" xml:"TaskSet"`
}

func CreateDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTasks", "ecs", "openAPI")
	return
}

func CreateDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

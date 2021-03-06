package rds

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

func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (response *ModifyDBInstanceMaintainTimeResponse, err error) {
	response = CreateModifyDBInstanceMaintainTimeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyDBInstanceMaintainTimeWithChan(request *ModifyDBInstanceMaintainTimeRequest) (<-chan *ModifyDBInstanceMaintainTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceMaintainTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceMaintainTime(request)
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

func (client *Client) ModifyDBInstanceMaintainTimeWithCallback(request *ModifyDBInstanceMaintainTimeRequest, callback func(response *ModifyDBInstanceMaintainTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceMaintainTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceMaintainTime(request)
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

type ModifyDBInstanceMaintainTimeRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	MaintainTime         string           `position:"Query" name:"MaintainTime"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type ModifyDBInstanceMaintainTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyDBInstanceMaintainTimeRequest() (request *ModifyDBInstanceMaintainTimeRequest) {
	request = &ModifyDBInstanceMaintainTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceMaintainTime", "rds", "openAPI")
	return
}

func CreateModifyDBInstanceMaintainTimeResponse() (response *ModifyDBInstanceMaintainTimeResponse) {
	response = &ModifyDBInstanceMaintainTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

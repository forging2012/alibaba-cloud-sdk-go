package afs

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

func (client *Client) DescribeCaptchaRisk(request *DescribeCaptchaRiskRequest) (response *DescribeCaptchaRiskResponse, err error) {
	response = CreateDescribeCaptchaRiskResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCaptchaRiskWithChan(request *DescribeCaptchaRiskRequest) (<-chan *DescribeCaptchaRiskResponse, <-chan error) {
	responseChan := make(chan *DescribeCaptchaRiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCaptchaRisk(request)
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

func (client *Client) DescribeCaptchaRiskWithCallback(request *DescribeCaptchaRiskRequest, callback func(response *DescribeCaptchaRiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCaptchaRiskResponse
		var err error
		defer close(result)
		response, err = client.DescribeCaptchaRisk(request)
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

type DescribeCaptchaRiskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ConfigName      string           `position:"Query" name:"ConfigName"`
	Time            string           `position:"Query" name:"Time"`
}

type DescribeCaptchaRiskResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BizCode        string `json:"BizCode" xml:"BizCode"`
	NumOfThisMonth int    `json:"NumOfThisMonth" xml:"NumOfThisMonth"`
	NumOfLastMonth int    `json:"NumOfLastMonth" xml:"NumOfLastMonth"`
	RiskLevel      string `json:"RiskLevel" xml:"RiskLevel"`
}

func CreateDescribeCaptchaRiskRequest() (request *DescribeCaptchaRiskRequest) {
	request = &DescribeCaptchaRiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaRisk", "", "")
	return
}

func CreateDescribeCaptchaRiskResponse() (response *DescribeCaptchaRiskResponse) {
	response = &DescribeCaptchaRiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

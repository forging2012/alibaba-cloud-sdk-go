package vpc

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

func (client *Client) DescribeVpcAttribute(request *DescribeVpcAttributeRequest) (response *DescribeVpcAttributeResponse, err error) {
	response = CreateDescribeVpcAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVpcAttributeWithChan(request *DescribeVpcAttributeRequest) (<-chan *DescribeVpcAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcAttribute(request)
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

func (client *Client) DescribeVpcAttributeWithCallback(request *DescribeVpcAttributeRequest, callback func(response *DescribeVpcAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcAttribute(request)
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

type DescribeVpcAttributeRequest struct {
	*requests.RpcRequest
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeVpcAttributeResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	VpcId              string `json:"VpcId" xml:"VpcId"`
	RegionId           string `json:"RegionId" xml:"RegionId"`
	Status             string `json:"Status" xml:"Status"`
	VpcName            string `json:"VpcName" xml:"VpcName"`
	CreationTime       string `json:"CreationTime" xml:"CreationTime"`
	CidrBlock          string `json:"CidrBlock" xml:"CidrBlock"`
	VRouterId          string `json:"VRouterId" xml:"VRouterId"`
	Description        string `json:"Description" xml:"Description"`
	IsDefault          bool   `json:"IsDefault" xml:"IsDefault"`
	ClassicLinkEnabled bool   `json:"ClassicLinkEnabled" xml:"ClassicLinkEnabled"`
	ResourceGroupId    string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	VSwitchIds         struct {
		VSwitchId []string `json:"VSwitchId" xml:"VSwitchId"`
	} `json:"VSwitchIds" xml:"VSwitchIds"`
	UserCidrs struct {
		UserCidr []string `json:"UserCidr" xml:"UserCidr"`
	} `json:"UserCidrs" xml:"UserCidrs"`
	AssociatedCens struct {
		AssociatedCbn []struct {
			CenStatus   string `json:"CenStatus" xml:"CenStatus"`
			CenId       string `json:"CenId" xml:"CenId"`
			CenOwnerUid int    `json:"CenOwnerUid" xml:"CenOwnerUid"`
		} `json:"AssociatedCbn" xml:"AssociatedCbn"`
	} `json:"AssociatedCens" xml:"AssociatedCens"`
	CloudResources struct {
		CloudResourceSetType []struct {
			ResourceType  string `json:"ResourceType" xml:"ResourceType"`
			ResourceCount int    `json:"ResourceCount" xml:"ResourceCount"`
		} `json:"CloudResourceSetType" xml:"CloudResourceSetType"`
	} `json:"CloudResources" xml:"CloudResources"`
}

func CreateDescribeVpcAttributeRequest() (request *DescribeVpcAttributeRequest) {
	request = &DescribeVpcAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpcAttribute", "vpc", "openAPI")
	return
}

func CreateDescribeVpcAttributeResponse() (response *DescribeVpcAttributeResponse) {
	response = &DescribeVpcAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

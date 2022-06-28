package ddoscoo

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

// EnableLayer7CCRule invokes the ddoscoo.EnableLayer7CCRule API synchronously
func (client *Client) EnableLayer7CCRule(request *EnableLayer7CCRuleRequest) (response *EnableLayer7CCRuleResponse, err error) {
	response = CreateEnableLayer7CCRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableLayer7CCRuleWithChan invokes the ddoscoo.EnableLayer7CCRule API asynchronously
func (client *Client) EnableLayer7CCRuleWithChan(request *EnableLayer7CCRuleRequest) (<-chan *EnableLayer7CCRuleResponse, <-chan error) {
	responseChan := make(chan *EnableLayer7CCRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableLayer7CCRule(request)
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

// EnableLayer7CCRuleWithCallback invokes the ddoscoo.EnableLayer7CCRule API asynchronously
func (client *Client) EnableLayer7CCRuleWithCallback(request *EnableLayer7CCRuleRequest, callback func(response *EnableLayer7CCRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableLayer7CCRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableLayer7CCRule(request)
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

// EnableLayer7CCRuleRequest is the request struct for api EnableLayer7CCRule
type EnableLayer7CCRuleRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// EnableLayer7CCRuleResponse is the response struct for api EnableLayer7CCRule
type EnableLayer7CCRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableLayer7CCRuleRequest creates a request to invoke EnableLayer7CCRule API
func CreateEnableLayer7CCRuleRequest() (request *EnableLayer7CCRuleRequest) {
	request = &EnableLayer7CCRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "EnableLayer7CCRule", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableLayer7CCRuleResponse creates a response to parse from EnableLayer7CCRule response
func CreateEnableLayer7CCRuleResponse() (response *EnableLayer7CCRuleResponse) {
	response = &EnableLayer7CCRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

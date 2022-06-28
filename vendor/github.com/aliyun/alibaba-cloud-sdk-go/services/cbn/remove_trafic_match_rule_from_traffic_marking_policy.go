package cbn

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

// RemoveTraficMatchRuleFromTrafficMarkingPolicy invokes the cbn.RemoveTraficMatchRuleFromTrafficMarkingPolicy API synchronously
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicy(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) (response *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, err error) {
	response = CreateRemoveTraficMatchRuleFromTrafficMarkingPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveTraficMatchRuleFromTrafficMarkingPolicyWithChan invokes the cbn.RemoveTraficMatchRuleFromTrafficMarkingPolicy API asynchronously
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicyWithChan(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) (<-chan *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, <-chan error) {
	responseChan := make(chan *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveTraficMatchRuleFromTrafficMarkingPolicy(request)
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

// RemoveTraficMatchRuleFromTrafficMarkingPolicyWithCallback invokes the cbn.RemoveTraficMatchRuleFromTrafficMarkingPolicy API asynchronously
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicyWithCallback(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest, callback func(response *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
		var err error
		defer close(result)
		response, err = client.RemoveTraficMatchRuleFromTrafficMarkingPolicy(request)
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

// RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest is the request struct for api RemoveTraficMatchRuleFromTrafficMarkingPolicy
type RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	TrafficMarkingPolicyId string           `position:"Query" name:"TrafficMarkingPolicyId"`
	DryRun                 requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	TrafficMarkRuleIds     *[]string        `position:"Query" name:"TrafficMarkRuleIds"  type:"Repeated"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse is the response struct for api RemoveTraficMatchRuleFromTrafficMarkingPolicy
type RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveTraficMatchRuleFromTrafficMarkingPolicyRequest creates a request to invoke RemoveTraficMatchRuleFromTrafficMarkingPolicy API
func CreateRemoveTraficMatchRuleFromTrafficMarkingPolicyRequest() (request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) {
	request = &RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "RemoveTraficMatchRuleFromTrafficMarkingPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveTraficMatchRuleFromTrafficMarkingPolicyResponse creates a response to parse from RemoveTraficMatchRuleFromTrafficMarkingPolicy response
func CreateRemoveTraficMatchRuleFromTrafficMarkingPolicyResponse() (response *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) {
	response = &RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

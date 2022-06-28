package dcdn

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

// SetDcdnFullDomainsBlockIP invokes the dcdn.SetDcdnFullDomainsBlockIP API synchronously
func (client *Client) SetDcdnFullDomainsBlockIP(request *SetDcdnFullDomainsBlockIPRequest) (response *SetDcdnFullDomainsBlockIPResponse, err error) {
	response = CreateSetDcdnFullDomainsBlockIPResponse()
	err = client.DoAction(request, response)
	return
}

// SetDcdnFullDomainsBlockIPWithChan invokes the dcdn.SetDcdnFullDomainsBlockIP API asynchronously
func (client *Client) SetDcdnFullDomainsBlockIPWithChan(request *SetDcdnFullDomainsBlockIPRequest) (<-chan *SetDcdnFullDomainsBlockIPResponse, <-chan error) {
	responseChan := make(chan *SetDcdnFullDomainsBlockIPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDcdnFullDomainsBlockIP(request)
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

// SetDcdnFullDomainsBlockIPWithCallback invokes the dcdn.SetDcdnFullDomainsBlockIP API asynchronously
func (client *Client) SetDcdnFullDomainsBlockIPWithCallback(request *SetDcdnFullDomainsBlockIPRequest, callback func(response *SetDcdnFullDomainsBlockIPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDcdnFullDomainsBlockIPResponse
		var err error
		defer close(result)
		response, err = client.SetDcdnFullDomainsBlockIP(request)
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

// SetDcdnFullDomainsBlockIPRequest is the request struct for api SetDcdnFullDomainsBlockIP
type SetDcdnFullDomainsBlockIPRequest struct {
	*requests.RpcRequest
	IPList        string           `position:"Body" name:"IPList"`
	BlockInterval requests.Integer `position:"Body" name:"BlockInterval"`
	OperationType string           `position:"Body" name:"OperationType"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// SetDcdnFullDomainsBlockIPResponse is the response struct for api SetDcdnFullDomainsBlockIP
type SetDcdnFullDomainsBlockIPResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSetDcdnFullDomainsBlockIPRequest creates a request to invoke SetDcdnFullDomainsBlockIP API
func CreateSetDcdnFullDomainsBlockIPRequest() (request *SetDcdnFullDomainsBlockIPRequest) {
	request = &SetDcdnFullDomainsBlockIPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "SetDcdnFullDomainsBlockIP", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDcdnFullDomainsBlockIPResponse creates a response to parse from SetDcdnFullDomainsBlockIP response
func CreateSetDcdnFullDomainsBlockIPResponse() (response *SetDcdnFullDomainsBlockIPResponse) {
	response = &SetDcdnFullDomainsBlockIPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

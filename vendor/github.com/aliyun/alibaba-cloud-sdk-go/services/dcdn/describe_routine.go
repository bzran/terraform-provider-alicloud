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

// DescribeRoutine invokes the dcdn.DescribeRoutine API synchronously
func (client *Client) DescribeRoutine(request *DescribeRoutineRequest) (response *DescribeRoutineResponse, err error) {
	response = CreateDescribeRoutineResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoutineWithChan invokes the dcdn.DescribeRoutine API asynchronously
func (client *Client) DescribeRoutineWithChan(request *DescribeRoutineRequest) (<-chan *DescribeRoutineResponse, <-chan error) {
	responseChan := make(chan *DescribeRoutineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoutine(request)
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

// DescribeRoutineWithCallback invokes the dcdn.DescribeRoutine API asynchronously
func (client *Client) DescribeRoutineWithCallback(request *DescribeRoutineRequest, callback func(response *DescribeRoutineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoutineResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoutine(request)
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

// DescribeRoutineRequest is the request struct for api DescribeRoutine
type DescribeRoutineRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	Name    string           `position:"Body" name:"Name"`
}

// DescribeRoutineResponse is the response struct for api DescribeRoutine
type DescribeRoutineResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Content   map[string]interface{} `json:"Content" xml:"Content"`
}

// CreateDescribeRoutineRequest creates a request to invoke DescribeRoutine API
func CreateDescribeRoutineRequest() (request *DescribeRoutineRequest) {
	request = &DescribeRoutineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeRoutine", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRoutineResponse creates a response to parse from DescribeRoutine response
func CreateDescribeRoutineResponse() (response *DescribeRoutineResponse) {
	response = &DescribeRoutineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

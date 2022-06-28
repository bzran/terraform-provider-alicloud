package emr

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

// ResizeClusterV2 invokes the emr.ResizeClusterV2 API synchronously
func (client *Client) ResizeClusterV2(request *ResizeClusterV2Request) (response *ResizeClusterV2Response, err error) {
	response = CreateResizeClusterV2Response()
	err = client.DoAction(request, response)
	return
}

// ResizeClusterV2WithChan invokes the emr.ResizeClusterV2 API asynchronously
func (client *Client) ResizeClusterV2WithChan(request *ResizeClusterV2Request) (<-chan *ResizeClusterV2Response, <-chan error) {
	responseChan := make(chan *ResizeClusterV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResizeClusterV2(request)
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

// ResizeClusterV2WithCallback invokes the emr.ResizeClusterV2 API asynchronously
func (client *Client) ResizeClusterV2WithCallback(request *ResizeClusterV2Request, callback func(response *ResizeClusterV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResizeClusterV2Response
		var err error
		defer close(result)
		response, err = client.ResizeClusterV2(request)
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

// ResizeClusterV2Request is the request struct for api ResizeClusterV2
type ResizeClusterV2Request struct {
	*requests.RpcRequest
	IsOpenPublicIp    requests.Boolean                    `position:"Query" name:"IsOpenPublicIp"`
	AutoPayOrder      requests.Boolean                    `position:"Query" name:"AutoPayOrder"`
	ClusterId         string                              `position:"Query" name:"ClusterId"`
	VswitchId         string                              `position:"Query" name:"VswitchId"`
	HostComponentInfo *[]ResizeClusterV2HostComponentInfo `position:"Query" name:"HostComponentInfo"  type:"Repeated"`
	ClickhouseConf    string                              `position:"Query" name:"ClickhouseConf"`
	HostGroup         *[]ResizeClusterV2HostGroup         `position:"Query" name:"HostGroup"  type:"Repeated"`
	PromotionInfo     *[]ResizeClusterV2PromotionInfo     `position:"Query" name:"PromotionInfo"  type:"Repeated"`
}

// ResizeClusterV2HostComponentInfo is a repeated param struct in ResizeClusterV2Request
type ResizeClusterV2HostComponentInfo struct {
	HostName          string    `name:"HostName"`
	ComponentNameList *[]string `name:"ComponentNameList" type:"Repeated"`
	ServiceName       string    `name:"ServiceName"`
}

// ResizeClusterV2HostGroup is a repeated param struct in ResizeClusterV2Request
type ResizeClusterV2HostGroup struct {
	Period                          string `name:"Period"`
	SysDiskCapacity                 string `name:"SysDiskCapacity"`
	HostKeyPairName                 string `name:"HostKeyPairName"`
	PrivatePoolOptionsId            string `name:"PrivatePoolOptionsId"`
	DiskCapacity                    string `name:"DiskCapacity"`
	SysDiskType                     string `name:"SysDiskType"`
	ClusterId                       string `name:"ClusterId"`
	DiskType                        string `name:"DiskType"`
	HostGroupName                   string `name:"HostGroupName"`
	VswitchId                       string `name:"VswitchId"`
	DiskCount                       string `name:"DiskCount"`
	AutoRenew                       string `name:"AutoRenew"`
	HostGroupId                     string `name:"HostGroupId"`
	NodeCount                       string `name:"NodeCount"`
	InstanceType                    string `name:"InstanceType"`
	Comment                         string `name:"Comment"`
	ChargeType                      string `name:"ChargeType"`
	CreateType                      string `name:"CreateType"`
	HostPassword                    string `name:"HostPassword"`
	HostGroupType                   string `name:"HostGroupType"`
	PrivatePoolOptionsMatchCriteria string `name:"PrivatePoolOptionsMatchCriteria"`
}

// ResizeClusterV2PromotionInfo is a repeated param struct in ResizeClusterV2Request
type ResizeClusterV2PromotionInfo struct {
	PromotionOptionCode string `name:"PromotionOptionCode"`
	ProductCode         string `name:"ProductCode"`
	PromotionOptionNo   string `name:"PromotionOptionNo"`
}

// ResizeClusterV2Response is the response struct for api ResizeClusterV2
type ResizeClusterV2Response struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
}

// CreateResizeClusterV2Request creates a request to invoke ResizeClusterV2 API
func CreateResizeClusterV2Request() (request *ResizeClusterV2Request) {
	request = &ResizeClusterV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ResizeClusterV2", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResizeClusterV2Response creates a response to parse from ResizeClusterV2 response
func CreateResizeClusterV2Response() (response *ResizeClusterV2Response) {
	response = &ResizeClusterV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// DescribeDcdnBgpBpsData invokes the dcdn.DescribeDcdnBgpBpsData API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnbgpbpsdata.html
func (client *Client) DescribeDcdnBgpBpsData(request *DescribeDcdnBgpBpsDataRequest) (response *DescribeDcdnBgpBpsDataResponse, err error) {
	response = CreateDescribeDcdnBgpBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnBgpBpsDataWithChan invokes the dcdn.DescribeDcdnBgpBpsData API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnbgpbpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnBgpBpsDataWithChan(request *DescribeDcdnBgpBpsDataRequest) (<-chan *DescribeDcdnBgpBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnBgpBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnBgpBpsData(request)
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

// DescribeDcdnBgpBpsDataWithCallback invokes the dcdn.DescribeDcdnBgpBpsData API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnbgpbpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnBgpBpsDataWithCallback(request *DescribeDcdnBgpBpsDataRequest, callback func(response *DescribeDcdnBgpBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnBgpBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnBgpBpsData(request)
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

// DescribeDcdnBgpBpsDataRequest is the request struct for api DescribeDcdnBgpBpsData
type DescribeDcdnBgpBpsDataRequest struct {
	*requests.RpcRequest
	Isp       string           `position:"Query" name:"Isp"`
	StartTime string           `position:"Query" name:"StartTime"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	Interval  string           `position:"Query" name:"Interval"`
}

// DescribeDcdnBgpBpsDataResponse is the response struct for api DescribeDcdnBgpBpsData
type DescribeDcdnBgpBpsDataResponse struct {
	*responses.BaseResponse
	RequestId       string    `json:"RequestId" xml:"RequestId"`
	StartTime       string    `json:"StartTime" xml:"StartTime"`
	EndTime         string    `json:"EndTime" xml:"EndTime"`
	BgpDataInterval []BgpData `json:"BgpDataInterval" xml:"BgpDataInterval"`
}

// CreateDescribeDcdnBgpBpsDataRequest creates a request to invoke DescribeDcdnBgpBpsData API
func CreateDescribeDcdnBgpBpsDataRequest() (request *DescribeDcdnBgpBpsDataRequest) {
	request = &DescribeDcdnBgpBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnBgpBpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnBgpBpsDataResponse creates a response to parse from DescribeDcdnBgpBpsData response
func CreateDescribeDcdnBgpBpsDataResponse() (response *DescribeDcdnBgpBpsDataResponse) {
	response = &DescribeDcdnBgpBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
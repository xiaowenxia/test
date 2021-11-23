package kms

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

// ListCertificates invokes the kms.ListCertificates API synchronously
func (client *Client) ListCertificates(request *ListCertificatesRequest) (response *ListCertificatesResponse, err error) {
	response = CreateListCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListCertificatesWithChan invokes the kms.ListCertificates API asynchronously
func (client *Client) ListCertificatesWithChan(request *ListCertificatesRequest) (<-chan *ListCertificatesResponse, <-chan error) {
	responseChan := make(chan *ListCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCertificates(request)
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

// ListCertificatesWithCallback invokes the kms.ListCertificates API asynchronously
func (client *Client) ListCertificatesWithCallback(request *ListCertificatesRequest, callback func(response *ListCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCertificatesResponse
		var err error
		defer close(result)
		response, err = client.ListCertificates(request)
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

// ListCertificatesRequest is the request struct for api ListCertificates
type ListCertificatesRequest struct {
	*requests.RpcRequest
	Subject    string           `position:"Query" name:"Subject"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Issuer     string           `position:"Query" name:"Issuer"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Status     string           `position:"Query" name:"Status"`
}

// ListCertificatesResponse is the response struct for api ListCertificates
type ListCertificatesResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	TotalSize              int                    `json:"TotalSize" xml:"TotalSize"`
	PageNumber             int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize               int                    `json:"PageSize" xml:"PageSize"`
	CertificateSummaryList CertificateSummaryList `json:"CertificateSummaryList" xml:"CertificateSummaryList"`
}

// CreateListCertificatesRequest creates a request to invoke ListCertificates API
func CreateListCertificatesRequest() (request *ListCertificatesRequest) {
	request = &ListCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListCertificates", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCertificatesResponse creates a response to parse from ListCertificates response
func CreateListCertificatesResponse() (response *ListCertificatesResponse) {
	response = &ListCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

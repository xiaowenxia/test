package sts

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

// AssumeRoleWithServiceIdentity invokes the sts.AssumeRoleWithServiceIdentity API synchronously
// api document: https://help.aliyun.com/api/sts/assumerolewithserviceidentity.html
func (client *Client) AssumeRoleWithServiceIdentity(request *AssumeRoleWithServiceIdentityRequest) (response *AssumeRoleWithServiceIdentityResponse, err error) {
	response = CreateAssumeRoleWithServiceIdentityResponse()
	err = client.DoAction(request, response)
	return
}

// AssumeRoleWithServiceIdentityWithChan invokes the sts.AssumeRoleWithServiceIdentity API asynchronously
// api document: https://help.aliyun.com/api/sts/assumerolewithserviceidentity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssumeRoleWithServiceIdentityWithChan(request *AssumeRoleWithServiceIdentityRequest) (<-chan *AssumeRoleWithServiceIdentityResponse, <-chan error) {
	responseChan := make(chan *AssumeRoleWithServiceIdentityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssumeRoleWithServiceIdentity(request)
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

// AssumeRoleWithServiceIdentityWithCallback invokes the sts.AssumeRoleWithServiceIdentity API asynchronously
// api document: https://help.aliyun.com/api/sts/assumerolewithserviceidentity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssumeRoleWithServiceIdentityWithCallback(request *AssumeRoleWithServiceIdentityRequest, callback func(response *AssumeRoleWithServiceIdentityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssumeRoleWithServiceIdentityResponse
		var err error
		defer close(result)
		response, err = client.AssumeRoleWithServiceIdentity(request)
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

// AssumeRoleWithServiceIdentityRequest is the request struct for api AssumeRoleWithServiceIdentity
type AssumeRoleWithServiceIdentityRequest struct {
	*requests.RpcRequest
	RoleSessionName string           `position:"Query" name:"RoleSessionName"`
	Policy          string           `position:"Query" name:"Policy"`
	AssumeRoleFor   string           `position:"Query" name:"AssumeRoleFor"`
	RoleArn         string           `position:"Query" name:"RoleArn"`
	DurationSeconds requests.Integer `position:"Query" name:"DurationSeconds"`
}

// AssumeRoleWithServiceIdentityResponse is the response struct for api AssumeRoleWithServiceIdentity
type AssumeRoleWithServiceIdentityResponse struct {
	*responses.BaseResponse
	RequestId       string                                         `json:"RequestId" xml:"RequestId"`
	Credentials     CredentialsInAssumeRoleWithServiceIdentity     `json:"Credentials" xml:"Credentials"`
	AssumedRoleUser AssumedRoleUserInAssumeRoleWithServiceIdentity `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

// CreateAssumeRoleWithServiceIdentityRequest creates a request to invoke AssumeRoleWithServiceIdentity API
func CreateAssumeRoleWithServiceIdentityRequest() (request *AssumeRoleWithServiceIdentityRequest) {
	request = &AssumeRoleWithServiceIdentityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sts", "2015-04-01", "AssumeRoleWithServiceIdentity", "sts", "openAPI")
	return
}

// CreateAssumeRoleWithServiceIdentityResponse creates a response to parse from AssumeRoleWithServiceIdentity response
func CreateAssumeRoleWithServiceIdentityResponse() (response *AssumeRoleWithServiceIdentityResponse) {
	response = &AssumeRoleWithServiceIdentityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

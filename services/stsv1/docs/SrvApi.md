# \SrvApi

All URIs are relative to *https://sts.eqix.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SrvGetServiceIdTokenPost**](SrvApi.md#SrvGetServiceIdTokenPost) | **Post** /srv/getServiceIdToken | /srv/getServiceIdToken
[**SrvListClientRegistrationsPost**](SrvApi.md#SrvListClientRegistrationsPost) | **Post** /srv/listClientRegistrations | /srv/listClientRegistrations
[**SrvRegisterClientPost**](SrvApi.md#SrvRegisterClientPost) | **Post** /srv/registerClient | /srv/registerClient



## SrvGetServiceIdTokenPost

> SrvGetServiceIdTokenPost200Response SrvGetServiceIdTokenPost(ctx).SrvGetServiceIdTokenPostRequest(srvGetServiceIdTokenPostRequest).Execute()

/srv/getServiceIdToken



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	srvGetServiceIdTokenPostRequest := *openapiclient.NewSrvGetServiceIdTokenPostRequest("ProjectId_example", []string{"Audiences_example"}) // SrvGetServiceIdTokenPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SrvApi.SrvGetServiceIdTokenPost(context.Background()).SrvGetServiceIdTokenPostRequest(srvGetServiceIdTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SrvApi.SrvGetServiceIdTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SrvGetServiceIdTokenPost`: SrvGetServiceIdTokenPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SrvApi.SrvGetServiceIdTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSrvGetServiceIdTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srvGetServiceIdTokenPostRequest** | [**SrvGetServiceIdTokenPostRequest**](SrvGetServiceIdTokenPostRequest.md) |  | 

### Return type

[**SrvGetServiceIdTokenPost200Response**](SrvGetServiceIdTokenPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrvListClientRegistrationsPost

> SrvListClientRegistrationsPost200Response SrvListClientRegistrationsPost(ctx).SrvListClientRegistrationsPostRequest(srvListClientRegistrationsPostRequest).Execute()

/srv/listClientRegistrations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	srvListClientRegistrationsPostRequest := *openapiclient.NewSrvListClientRegistrationsPostRequest(int32(123)) // SrvListClientRegistrationsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SrvApi.SrvListClientRegistrationsPost(context.Background()).SrvListClientRegistrationsPostRequest(srvListClientRegistrationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SrvApi.SrvListClientRegistrationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SrvListClientRegistrationsPost`: SrvListClientRegistrationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SrvApi.SrvListClientRegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSrvListClientRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srvListClientRegistrationsPostRequest** | [**SrvListClientRegistrationsPostRequest**](SrvListClientRegistrationsPostRequest.md) |  | 

### Return type

[**SrvListClientRegistrationsPost200Response**](SrvListClientRegistrationsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SrvRegisterClientPost

> SrvRegisterClientPost200Response SrvRegisterClientPost(ctx).SrvRegisterClientPostRequest(srvRegisterClientPostRequest).Execute()

/srv/registerClient



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	srvRegisterClientPostRequest := *openapiclient.NewSrvRegisterClientPostRequest(interface{}(123), []interface{}{nil}, []interface{}{nil}, "ServiceAccessPolicyId_example") // SrvRegisterClientPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SrvApi.SrvRegisterClientPost(context.Background()).SrvRegisterClientPostRequest(srvRegisterClientPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SrvApi.SrvRegisterClientPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SrvRegisterClientPost`: SrvRegisterClientPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SrvApi.SrvRegisterClientPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSrvRegisterClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srvRegisterClientPostRequest** | [**SrvRegisterClientPostRequest**](SrvRegisterClientPostRequest.md) |  | 

### Return type

[**SrvRegisterClientPost200Response**](SrvRegisterClientPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ApplicationProcessorsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppProcessorByUuid**](ApplicationProcessorsApi.md#GetAppProcessorByUuid) | **Get** /fabric/v4/appProcessors/{appProcessorId} | Get Processor
[**SearchAppProcessors**](ApplicationProcessorsApi.md#SearchAppProcessors) | **Post** /fabric/v4/appProcessors/search | Search Processors



## GetAppProcessorByUuid

> AppProcessor GetAppProcessorByUuid(ctx, appProcessorId).Execute()

Get Processor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appProcessorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Processor UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationProcessorsApi.GetAppProcessorByUuid(context.Background(), appProcessorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationProcessorsApi.GetAppProcessorByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppProcessorByUuid`: AppProcessor
	fmt.Fprintf(os.Stdout, "Response from `ApplicationProcessorsApi.GetAppProcessorByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appProcessorId** | **string** | App Processor UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppProcessorByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppProcessor**](AppProcessor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppProcessors

> AppProcessorSearchResponse SearchAppProcessors(ctx).AppProcessorSearchRequest(appProcessorSearchRequest).Execute()

Search Processors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appProcessorSearchRequest := *openapiclient.NewAppProcessorSearchRequest() // AppProcessorSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationProcessorsApi.SearchAppProcessors(context.Background()).AppProcessorSearchRequest(appProcessorSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationProcessorsApi.SearchAppProcessors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAppProcessors`: AppProcessorSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationProcessorsApi.SearchAppProcessors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appProcessorSearchRequest** | [**AppProcessorSearchRequest**](AppProcessorSearchRequest.md) |  | 

### Return type

[**AppProcessorSearchResponse**](AppProcessorSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


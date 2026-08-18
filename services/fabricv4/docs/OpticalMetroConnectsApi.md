# \OpticalMetroConnectsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkOpticalConnect**](OpticalMetroConnectsApi.md#CreateBulkOpticalConnect) | **Post** /fabric/v4/opticalConnects/bulk | Create Dual Diverse Optical Metro Connect Service
[**CreateOpticalConnect**](OpticalMetroConnectsApi.md#CreateOpticalConnect) | **Post** /fabric/v4/opticalConnects | Create Optical Metro Connect Service
[**GetOpticalConnectByUuid**](OpticalMetroConnectsApi.md#GetOpticalConnectByUuid) | **Get** /fabric/v4/opticalConnects/{opticalConnectId} | Get Optical Metro Connect Service
[**SearchOpticalConnect**](OpticalMetroConnectsApi.md#SearchOpticalConnect) | **Post** /fabric/v4/opticalConnects/search | Search Optical Metro Connect Services



## CreateBulkOpticalConnect

> OpticalConnectBulk CreateBulkOpticalConnect(ctx).BulkOpticalConnectRequest(bulkOpticalConnectRequest).Execute()

Create Dual Diverse Optical Metro Connect Service



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
	bulkOpticalConnectRequest := *openapiclient.NewBulkOpticalConnectRequest() // BulkOpticalConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpticalMetroConnectsApi.CreateBulkOpticalConnect(context.Background()).BulkOpticalConnectRequest(bulkOpticalConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpticalMetroConnectsApi.CreateBulkOpticalConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkOpticalConnect`: OpticalConnectBulk
	fmt.Fprintf(os.Stdout, "Response from `OpticalMetroConnectsApi.CreateBulkOpticalConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkOpticalConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkOpticalConnectRequest** | [**BulkOpticalConnectRequest**](BulkOpticalConnectRequest.md) |  | 

### Return type

[**OpticalConnectBulk**](OpticalConnectBulk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOpticalConnect

> OpticalConnectResponse CreateOpticalConnect(ctx).OpticalConnectPostRequest(opticalConnectPostRequest).Execute()

Create Optical Metro Connect Service



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
	opticalConnectPostRequest := *openapiclient.NewOpticalConnectPostRequest(openapiclient.OpticalConnectPostRequest_type("OC"), int32(100000), openapiclient.OpticalConnectPostRequest_connectionDestinationType("COLO"), openapiclient.OpticalConnectPostRequest_pathType("DUAL_DIVERSE"), *openapiclient.NewOpticalConnectASideRequest(), *openapiclient.NewOpticalConnectZSideRequest()) // OpticalConnectPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpticalMetroConnectsApi.CreateOpticalConnect(context.Background()).OpticalConnectPostRequest(opticalConnectPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpticalMetroConnectsApi.CreateOpticalConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOpticalConnect`: OpticalConnectResponse
	fmt.Fprintf(os.Stdout, "Response from `OpticalMetroConnectsApi.CreateOpticalConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpticalConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opticalConnectPostRequest** | [**OpticalConnectPostRequest**](OpticalConnectPostRequest.md) |  | 

### Return type

[**OpticalConnectResponse**](OpticalConnectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpticalConnectByUuid

> OpticalConnectResponse GetOpticalConnectByUuid(ctx, opticalConnectId).Execute()

Get Optical Metro Connect Service



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
	opticalConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of an Optical Connect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpticalMetroConnectsApi.GetOpticalConnectByUuid(context.Background(), opticalConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpticalMetroConnectsApi.GetOpticalConnectByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpticalConnectByUuid`: OpticalConnectResponse
	fmt.Fprintf(os.Stdout, "Response from `OpticalMetroConnectsApi.GetOpticalConnectByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opticalConnectId** | **string** | Unique identifier of an Optical Connect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpticalConnectByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpticalConnectResponse**](OpticalConnectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOpticalConnect

> OpticalConnectServiceSearchResponse SearchOpticalConnect(ctx).OpticalConnectSearchRequest(opticalConnectSearchRequest).Execute()

Search Optical Metro Connect Services



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
	opticalConnectSearchRequest := *openapiclient.NewOpticalConnectSearchRequest() // OpticalConnectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpticalMetroConnectsApi.SearchOpticalConnect(context.Background()).OpticalConnectSearchRequest(opticalConnectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpticalMetroConnectsApi.SearchOpticalConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOpticalConnect`: OpticalConnectServiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `OpticalMetroConnectsApi.SearchOpticalConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOpticalConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opticalConnectSearchRequest** | [**OpticalConnectSearchRequest**](OpticalConnectSearchRequest.md) |  | 

### Return type

[**OpticalConnectServiceSearchResponse**](OpticalConnectServiceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


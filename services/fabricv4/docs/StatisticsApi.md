# \StatisticsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectionStatsByPortUuid**](StatisticsApi.md#GetConnectionStatsByPortUuid) | **Get** /fabric/v4/connections/{connectionId}/stats | Get Stats by uuid
[**GetPortStatsByPortUuid**](StatisticsApi.md#GetPortStatsByPortUuid) | **Get** /fabric/v4/ports/{portId}/stats | Get Stats by uuid



## GetConnectionStatsByPortUuid

> Statistics GetConnectionStatsByPortUuid(ctx, connectionId).StartDateTime(startDateTime).EndDateTime(endDateTime).ViewPoint(viewPoint).Execute()

Get Stats by uuid



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionId := "connectionId_example" // string | Connection UUID
	startDateTime := time.Now() // time.Time | startDateTime
	endDateTime := time.Now() // time.Time | endDateTime
	viewPoint := openapiclient.viewPoint("aSide") // ViewPoint | viewPoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsApi.GetConnectionStatsByPortUuid(context.Background(), connectionId).StartDateTime(startDateTime).EndDateTime(endDateTime).ViewPoint(viewPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetConnectionStatsByPortUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionStatsByPortUuid`: Statistics
	fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetConnectionStatsByPortUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionStatsByPortUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDateTime** | **time.Time** | startDateTime | 
 **endDateTime** | **time.Time** | endDateTime | 
 **viewPoint** | [**ViewPoint**](ViewPoint.md) | viewPoint | 

### Return type

[**Statistics**](Statistics.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortStatsByPortUuid

> Statistics GetPortStatsByPortUuid(ctx, portId).StartDateTime(startDateTime).EndDateTime(endDateTime).Execute()

Get Stats by uuid



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
	startDateTime := time.Now() // time.Time | startDateTime
	endDateTime := time.Now() // time.Time | endDateTime

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsApi.GetPortStatsByPortUuid(context.Background(), portId).StartDateTime(startDateTime).EndDateTime(endDateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetPortStatsByPortUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortStatsByPortUuid`: Statistics
	fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetPortStatsByPortUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortStatsByPortUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDateTime** | **time.Time** | startDateTime | 
 **endDateTime** | **time.Time** | endDateTime | 

### Return type

[**Statistics**](Statistics.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


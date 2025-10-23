# \CloudEventsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudEvent**](CloudEventsApi.md#GetCloudEvent) | **Get** /fabric/v4/cloudevents/{cloudEventId} | Get Cloud Event
[**GetCloudEventByAssetId**](CloudEventsApi.md#GetCloudEventByAssetId) | **Get** /fabric/v4/{asset}/{assetId}/cloudevents | Get Cloud Events by Asset Id
[**SearchCloudEvents**](CloudEventsApi.md#SearchCloudEvents) | **Post** /fabric/v4/cloudevents/search | Search Cloud Events



## GetCloudEvent

> CloudEvent GetCloudEvent(ctx, cloudEventId).Execute()

Get Cloud Event



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
	cloudEventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cloud Event UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudEventsApi.GetCloudEvent(context.Background(), cloudEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudEventsApi.GetCloudEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudEvent`: CloudEvent
	fmt.Fprintf(os.Stdout, "Response from `CloudEventsApi.GetCloudEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudEventId** | **string** | Cloud Event UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudEvent**](CloudEvent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudEventByAssetId

> GetCloudEventsByAssetResponse GetCloudEventByAssetId(ctx, asset, assetId).FromDateTime(fromDateTime).ToDateTime(toDateTime).Offset(offset).Limit(limit).Execute()

Get Cloud Events by Asset Id



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
	asset := openapiclient.CloudEventAssetType("ports") // CloudEventAssetType | asset
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | asset UUID
	fromDateTime := time.Now() // time.Time | Start date and time (optional)
	toDateTime := time.Now() // time.Time | End date and time (optional)
	offset := int32(56) // int32 | offset (optional) (default to 0)
	limit := int32(56) // int32 | limit (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudEventsApi.GetCloudEventByAssetId(context.Background(), asset, assetId).FromDateTime(fromDateTime).ToDateTime(toDateTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudEventsApi.GetCloudEventByAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudEventByAssetId`: GetCloudEventsByAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudEventsApi.GetCloudEventByAssetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | [**CloudEventAssetType**](.md) | asset | 
**assetId** | **string** | asset UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudEventByAssetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDateTime** | **time.Time** | Start date and time | 
 **toDateTime** | **time.Time** | End date and time | 
 **offset** | **int32** | offset | [default to 0]
 **limit** | **int32** | limit | [default to 100]

### Return type

[**GetCloudEventsByAssetResponse**](GetCloudEventsByAssetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudEvents

> GetCloudEventsByAssetResponse SearchCloudEvents(ctx).CloudEventSearchRequest(cloudEventSearchRequest).Execute()

Search Cloud Events



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
	cloudEventSearchRequest := *openapiclient.NewCloudEventSearchRequest(*openapiclient.NewCloudEventFilters()) // CloudEventSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudEventsApi.SearchCloudEvents(context.Background()).CloudEventSearchRequest(cloudEventSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudEventsApi.SearchCloudEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCloudEvents`: GetCloudEventsByAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudEventsApi.SearchCloudEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEventSearchRequest** | [**CloudEventSearchRequest**](CloudEventSearchRequest.md) |  | 

### Return type

[**GetCloudEventsByAssetResponse**](GetCloudEventsByAssetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


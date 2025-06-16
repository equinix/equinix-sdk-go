# \MetricsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricByAssetId**](MetricsApi.md#GetMetricByAssetId) | **Get** /fabric/v4/{asset}/{assetId}/metrics | Get Metrics by Asset Id



## GetMetricByAssetId

> GetMetricsByAssetResponse GetMetricByAssetId(ctx, asset, assetId).FromDateTime(fromDateTime).ToDateTime(toDateTime).Offset(offset).Limit(limit).Execute()

Get Metrics by Asset Id



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
	asset := openapiclient.MetricAssetType("ports") // MetricAssetType | asset
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | asset UUID
	fromDateTime := time.Now() // time.Time | Start date and time (optional)
	toDateTime := time.Now() // time.Time | End date and time (optional)
	offset := int32(56) // int32 | offset (optional) (default to 0)
	limit := int32(56) // int32 | limit (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.GetMetricByAssetId(context.Background(), asset, assetId).FromDateTime(fromDateTime).ToDateTime(toDateTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetricByAssetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricByAssetId`: GetMetricsByAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetricByAssetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asset** | [**MetricAssetType**](.md) | asset | 
**assetId** | **string** | asset UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricByAssetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDateTime** | **time.Time** | Start date and time | 
 **toDateTime** | **time.Time** | End date and time | 
 **offset** | **int32** | offset | [default to 0]
 **limit** | **int32** | limit | [default to 20]

### Return type

[**GetMetricsByAssetResponse**](GetMetricsByAssetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \RouteAggregationRulesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteAggregationRule**](RouteAggregationRulesApi.md#CreateRouteAggregationRule) | **Post** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules | Create RARule
[**CreateRouteAggregationRulesInBulk**](RouteAggregationRulesApi.md#CreateRouteAggregationRulesInBulk) | **Post** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/bulk | Bulk RARules
[**DeleteRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#DeleteRouteAggregationRuleByUuid) | **Delete** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | DeleteRARule
[**GetRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#GetRouteAggregationRuleByUuid) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | GetRARule By UUID
[**GetRouteAggregationRuleChangeByUuid**](RouteAggregationRulesApi.md#GetRouteAggregationRuleChangeByUuid) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId}/changes/{changeId} | Get Change By ID
[**GetRouteAggregationRuleChanges**](RouteAggregationRulesApi.md#GetRouteAggregationRuleChanges) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId}/changes | Get All Changes
[**GetRouteAggregationRules**](RouteAggregationRulesApi.md#GetRouteAggregationRules) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules | GetRARules
[**PatchRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#PatchRouteAggregationRuleByUuid) | **Patch** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | PatchRARule
[**ReplaceRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#ReplaceRouteAggregationRuleByUuid) | **Put** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | ReplaceRARule



## CreateRouteAggregationRule

> RouteAggregationRulesData CreateRouteAggregationRule(ctx, routeAggregationId).RouteAggregationRulesBase(routeAggregationRulesBase).Execute()

Create RARule



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRulesBase := *openapiclient.NewRouteAggregationRulesBase("192.168.0.0/24") // RouteAggregationRulesBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.CreateRouteAggregationRule(context.Background(), routeAggregationId).RouteAggregationRulesBase(routeAggregationRulesBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.CreateRouteAggregationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteAggregationRule`: RouteAggregationRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.CreateRouteAggregationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteAggregationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeAggregationRulesBase** | [**RouteAggregationRulesBase**](RouteAggregationRulesBase.md) |  | 

### Return type

[**RouteAggregationRulesData**](RouteAggregationRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouteAggregationRulesInBulk

> GetRouteAggregationRulesResponse CreateRouteAggregationRulesInBulk(ctx, routeAggregationId).RouteAggregationRulesPostRequest(routeAggregationRulesPostRequest).Execute()

Bulk RARules



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRulesPostRequest := *openapiclient.NewRouteAggregationRulesPostRequest() // RouteAggregationRulesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.CreateRouteAggregationRulesInBulk(context.Background(), routeAggregationId).RouteAggregationRulesPostRequest(routeAggregationRulesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.CreateRouteAggregationRulesInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteAggregationRulesInBulk`: GetRouteAggregationRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.CreateRouteAggregationRulesInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteAggregationRulesInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeAggregationRulesPostRequest** | [**RouteAggregationRulesPostRequest**](RouteAggregationRulesPostRequest.md) |  | 

### Return type

[**GetRouteAggregationRulesResponse**](GetRouteAggregationRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteAggregationRuleByUuid

> RouteAggregationRulesData DeleteRouteAggregationRuleByUuid(ctx, routeAggregationId, routeAggregationRuleId).Execute()

DeleteRARule



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRuleId := "routeAggregationRuleId_example" // string | Route Aggregation Rules Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.DeleteRouteAggregationRuleByUuid(context.Background(), routeAggregationId, routeAggregationRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.DeleteRouteAggregationRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRouteAggregationRuleByUuid`: RouteAggregationRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.DeleteRouteAggregationRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**routeAggregationRuleId** | **string** | Route Aggregation Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteAggregationRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteAggregationRulesData**](RouteAggregationRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationRuleByUuid

> RouteAggregationRulesData GetRouteAggregationRuleByUuid(ctx, routeAggregationId, routeAggregationRuleId).Execute()

GetRARule By UUID



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRuleId := "routeAggregationRuleId_example" // string | Route Aggregation Rules Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.GetRouteAggregationRuleByUuid(context.Background(), routeAggregationId, routeAggregationRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.GetRouteAggregationRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationRuleByUuid`: RouteAggregationRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.GetRouteAggregationRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**routeAggregationRuleId** | **string** | Route Aggregation Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteAggregationRulesData**](RouteAggregationRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationRuleChangeByUuid

> RouteAggregationRulesChangeData GetRouteAggregationRuleChangeByUuid(ctx, routeAggregationId, routeAggregationRuleId, changeId).Execute()

Get Change By ID



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRuleId := "routeAggregationRuleId_example" // string | Route Aggregation Rules Id
	changeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Route Aggregation Rule Change UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.GetRouteAggregationRuleChangeByUuid(context.Background(), routeAggregationId, routeAggregationRuleId, changeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.GetRouteAggregationRuleChangeByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationRuleChangeByUuid`: RouteAggregationRulesChangeData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.GetRouteAggregationRuleChangeByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**routeAggregationRuleId** | **string** | Route Aggregation Rules Id | 
**changeId** | **string** | Route Aggregation Rule Change UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationRuleChangeByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RouteAggregationRulesChangeData**](RouteAggregationRulesChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationRuleChanges

> RouteAggregationRulesChangeDataResponse GetRouteAggregationRuleChanges(ctx, routeAggregationId, routeAggregationRuleId).Offset(offset).Limit(limit).Execute()

Get All Changes



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRuleId := "routeAggregationRuleId_example" // string | Route Aggregation Rules Id
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.GetRouteAggregationRuleChanges(context.Background(), routeAggregationId, routeAggregationRuleId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.GetRouteAggregationRuleChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationRuleChanges`: RouteAggregationRulesChangeDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.GetRouteAggregationRuleChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**routeAggregationRuleId** | **string** | Route Aggregation Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationRuleChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**RouteAggregationRulesChangeDataResponse**](RouteAggregationRulesChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationRules

> GetRouteAggregationRulesResponse GetRouteAggregationRules(ctx, routeAggregationId).Offset(offset).Limit(limit).Execute()

GetRARules



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.GetRouteAggregationRules(context.Background(), routeAggregationId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.GetRouteAggregationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationRules`: GetRouteAggregationRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.GetRouteAggregationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetRouteAggregationRulesResponse**](GetRouteAggregationRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRouteAggregationRuleByUuid

> RouteAggregationRulesData PatchRouteAggregationRuleByUuid(ctx, routeAggregationId, routeAggregationRuleId).RouteAggregationRulesPatchRequestItem(routeAggregationRulesPatchRequestItem).Execute()

PatchRARule



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRuleId := "routeAggregationRuleId_example" // string | Route Aggregation Rules Id
	routeAggregationRulesPatchRequestItem := []openapiclient.RouteAggregationRulesPatchRequestItem{*openapiclient.NewRouteAggregationRulesPatchRequestItem("replace", "/prefix", interface{}(123))} // []RouteAggregationRulesPatchRequestItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.PatchRouteAggregationRuleByUuid(context.Background(), routeAggregationId, routeAggregationRuleId).RouteAggregationRulesPatchRequestItem(routeAggregationRulesPatchRequestItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.PatchRouteAggregationRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRouteAggregationRuleByUuid`: RouteAggregationRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.PatchRouteAggregationRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**routeAggregationRuleId** | **string** | Route Aggregation Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRouteAggregationRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routeAggregationRulesPatchRequestItem** | [**[]RouteAggregationRulesPatchRequestItem**](RouteAggregationRulesPatchRequestItem.md) |  | 

### Return type

[**RouteAggregationRulesData**](RouteAggregationRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRouteAggregationRuleByUuid

> RouteAggregationRulesData ReplaceRouteAggregationRuleByUuid(ctx, routeAggregationId, routeAggregationRuleId).RouteAggregationRulesBase(routeAggregationRulesBase).Execute()

ReplaceRARule



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
	routeAggregationId := "routeAggregationId_example" // string | Route Aggregations Id
	routeAggregationRuleId := "routeAggregationRuleId_example" // string | Route Aggregation Rules Id
	routeAggregationRulesBase := *openapiclient.NewRouteAggregationRulesBase("192.168.0.0/24") // RouteAggregationRulesBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.ReplaceRouteAggregationRuleByUuid(context.Background(), routeAggregationId, routeAggregationRuleId).RouteAggregationRulesBase(routeAggregationRulesBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.ReplaceRouteAggregationRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRouteAggregationRuleByUuid`: RouteAggregationRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.ReplaceRouteAggregationRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**routeAggregationRuleId** | **string** | Route Aggregation Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRouteAggregationRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routeAggregationRulesBase** | [**RouteAggregationRulesBase**](RouteAggregationRulesBase.md) |  | 

### Return type

[**RouteAggregationRulesData**](RouteAggregationRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


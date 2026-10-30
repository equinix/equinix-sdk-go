# \RouteAggregationRulesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteAggregationRule**](RouteAggregationRulesApi.md#CreateRouteAggregationRule) | **Post** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules | Create Route Aggregation Rule
[**CreateRouteAggregationRulesInBulk**](RouteAggregationRulesApi.md#CreateRouteAggregationRulesInBulk) | **Post** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/bulk | Bulk Create Route Aggregation Rules
[**DeleteRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#DeleteRouteAggregationRuleByUuid) | **Delete** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | Delete Route Aggregation Rule
[**GetRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#GetRouteAggregationRuleByUuid) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | Get Route Aggregation Rule by ID
[**GetRouteAggregationRuleChangeByUuid**](RouteAggregationRulesApi.md#GetRouteAggregationRuleChangeByUuid) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId}/changes/{changeId} | Get Change By ID
[**GetRouteAggregationRuleChanges**](RouteAggregationRulesApi.md#GetRouteAggregationRuleChanges) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId}/changes | Get Route Aggregation Rule Changes
[**GetRouteAggregationRules**](RouteAggregationRulesApi.md#GetRouteAggregationRules) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules | Get Route Aggregation Rules
[**PatchRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#PatchRouteAggregationRuleByUuid) | **Patch** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | Update Route Aggregation Rule
[**ReplaceRouteAggregationRuleByUuid**](RouteAggregationRulesApi.md#ReplaceRouteAggregationRuleByUuid) | **Put** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/{routeAggregationRuleId} | Replace Route Aggregation Rule
[**SearchRouteAggregationRules**](RouteAggregationRulesApi.md#SearchRouteAggregationRules) | **Post** /fabric/v4/routeAggregations/{routeAggregationId}/routeAggregationRules/search | Search Route Aggregation Rules



## CreateRouteAggregationRule

> RouteAggregationRulesData CreateRouteAggregationRule(ctx, routeAggregationId).RouteAggregationRulesBase(routeAggregationRulesBase).Execute()

Create Route Aggregation Rule



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
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

Bulk Create Route Aggregation Rules



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
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

Delete Route Aggregation Rule



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRuleId := "65b025ef-022b-4180-85cf-82cfc1ab655b" // string | Route Aggregation Rules Id

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

Get Route Aggregation Rule by ID



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRuleId := "65b025ef-022b-4180-85cf-82cfc1ab655b" // string | Route Aggregation Rules Id

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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRuleId := "65b025ef-022b-4180-85cf-82cfc1ab655b" // string | Route Aggregation Rules Id
	changeId := "6d500177-9404-41f2-99e8-2bf1a84d8db5" // string | Route Aggregation Rule Change UUID

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

Get Route Aggregation Rule Changes



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRuleId := "65b025ef-022b-4180-85cf-82cfc1ab655b" // string | Route Aggregation Rules Id
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

Get Route Aggregation Rules



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
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

Update Route Aggregation Rule



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRuleId := "65b025ef-022b-4180-85cf-82cfc1ab655b" // string | Route Aggregation Rules Id
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

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRouteAggregationRuleByUuid

> RouteAggregationRulesData ReplaceRouteAggregationRuleByUuid(ctx, routeAggregationId, routeAggregationRuleId).RouteAggregationRulesBase(routeAggregationRulesBase).Execute()

Replace Route Aggregation Rule



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRuleId := "65b025ef-022b-4180-85cf-82cfc1ab655b" // string | Route Aggregation Rules Id
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


## SearchRouteAggregationRules

> RouteAggregationRulesSearchResponse SearchRouteAggregationRules(ctx, routeAggregationId).RouteAggregationRulesSearchRequest(routeAggregationRulesSearchRequest).Execute()

Search Route Aggregation Rules



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
	routeAggregationId := "695a8471-6595-4ac6-a2f4-b3d96ed3a59d" // string | Route Aggregations Id
	routeAggregationRulesSearchRequest := *openapiclient.NewRouteAggregationRulesSearchRequest() // RouteAggregationRulesSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationRulesApi.SearchRouteAggregationRules(context.Background(), routeAggregationId).RouteAggregationRulesSearchRequest(routeAggregationRulesSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationRulesApi.SearchRouteAggregationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRouteAggregationRules`: RouteAggregationRulesSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationRulesApi.SearchRouteAggregationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRouteAggregationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeAggregationRulesSearchRequest** | [**RouteAggregationRulesSearchRequest**](RouteAggregationRulesSearchRequest.md) |  | 

### Return type

[**RouteAggregationRulesSearchResponse**](RouteAggregationRulesSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


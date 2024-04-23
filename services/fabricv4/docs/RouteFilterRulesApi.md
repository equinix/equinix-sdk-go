# \RouteFilterRulesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteFilterRule**](RouteFilterRulesApi.md#CreateRouteFilterRule) | **Post** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules | Create RFRule
[**CreateRouteFilterRulesInBulk**](RouteFilterRulesApi.md#CreateRouteFilterRulesInBulk) | **Post** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/bulk | Bulk RFRules
[**DeleteRouteFilterRuleByUuid**](RouteFilterRulesApi.md#DeleteRouteFilterRuleByUuid) | **Delete** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | DeleteRFRule
[**GetRouteFilterRuleByUuid**](RouteFilterRulesApi.md#GetRouteFilterRuleByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | GetRFRule By UUID
[**GetRouteFilterRuleChangeByUuid**](RouteFilterRulesApi.md#GetRouteFilterRuleChangeByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId}/changes/{changeId} | Get Change By ID
[**GetRouteFilterRuleChanges**](RouteFilterRulesApi.md#GetRouteFilterRuleChanges) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId}/changes | Get All Changes
[**GetRouteFilterRules**](RouteFilterRulesApi.md#GetRouteFilterRules) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules | GetRFRules
[**PatchRouteFilterRuleByUuid**](RouteFilterRulesApi.md#PatchRouteFilterRuleByUuid) | **Patch** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | PatchRFilterRule
[**ReplaceRouteFilterRuleByUuid**](RouteFilterRulesApi.md#ReplaceRouteFilterRuleByUuid) | **Put** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | ReplaceRFRule



## CreateRouteFilterRule

> RouteFilterRulesData CreateRouteFilterRule(ctx, routeFilterId).RouteFilterRulesBase(routeFilterRulesBase).Execute()

Create RFRule



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRulesBase := *openapiclient.NewRouteFilterRulesBase("192.168.0.0/24") // RouteFilterRulesBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.CreateRouteFilterRule(context.Background(), routeFilterId).RouteFilterRulesBase(routeFilterRulesBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.CreateRouteFilterRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteFilterRule`: RouteFilterRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.CreateRouteFilterRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteFilterRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeFilterRulesBase** | [**RouteFilterRulesBase**](RouteFilterRulesBase.md) |  | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouteFilterRulesInBulk

> GetRouteFilterRulesResponse CreateRouteFilterRulesInBulk(ctx, routeFilterId).RouteFilterRulesPostRequest(routeFilterRulesPostRequest).Execute()

Bulk RFRules



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRulesPostRequest := *openapiclient.NewRouteFilterRulesPostRequest() // RouteFilterRulesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.CreateRouteFilterRulesInBulk(context.Background(), routeFilterId).RouteFilterRulesPostRequest(routeFilterRulesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.CreateRouteFilterRulesInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteFilterRulesInBulk`: GetRouteFilterRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.CreateRouteFilterRulesInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteFilterRulesInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeFilterRulesPostRequest** | [**RouteFilterRulesPostRequest**](RouteFilterRulesPostRequest.md) |  | 

### Return type

[**GetRouteFilterRulesResponse**](GetRouteFilterRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteFilterRuleByUuid

> RouteFilterRulesData DeleteRouteFilterRuleByUuid(ctx, routeFilterId, routeFilterRuleId).Execute()

DeleteRFRule



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRuleId := "routeFilterRuleId_example" // string | Route  Filter  Rules Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.DeleteRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.DeleteRouteFilterRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRouteFilterRuleByUuid`: RouteFilterRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.DeleteRouteFilterRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**routeFilterRuleId** | **string** | Route  Filter  Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteFilterRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterRuleByUuid

> RouteFilterRulesData GetRouteFilterRuleByUuid(ctx, routeFilterId, routeFilterRuleId).Execute()

GetRFRule By UUID



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRuleId := "routeFilterRuleId_example" // string | Route  Filter  Rules Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.GetRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.GetRouteFilterRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterRuleByUuid`: RouteFilterRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.GetRouteFilterRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**routeFilterRuleId** | **string** | Route  Filter  Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterRuleChangeByUuid

> RouteFilterRulesChangeData GetRouteFilterRuleChangeByUuid(ctx, routeFilterId, routeFilterRuleId, changeId).Execute()

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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRuleId := "routeFilterRuleId_example" // string | Route  Filter  Rules Id
	changeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Route Filter Rule Change UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.GetRouteFilterRuleChangeByUuid(context.Background(), routeFilterId, routeFilterRuleId, changeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.GetRouteFilterRuleChangeByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterRuleChangeByUuid`: RouteFilterRulesChangeData
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.GetRouteFilterRuleChangeByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**routeFilterRuleId** | **string** | Route  Filter  Rules Id | 
**changeId** | **string** | Route Filter Rule Change UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterRuleChangeByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RouteFilterRulesChangeData**](RouteFilterRulesChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterRuleChanges

> RouteFilterRulesChangeDataResponse GetRouteFilterRuleChanges(ctx, routeFilterId, routeFilterRuleId).Offset(offset).Limit(limit).Execute()

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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRuleId := "routeFilterRuleId_example" // string | Route  Filter  Rules Id
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.GetRouteFilterRuleChanges(context.Background(), routeFilterId, routeFilterRuleId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.GetRouteFilterRuleChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterRuleChanges`: RouteFilterRulesChangeDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.GetRouteFilterRuleChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**routeFilterRuleId** | **string** | Route  Filter  Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterRuleChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**RouteFilterRulesChangeDataResponse**](RouteFilterRulesChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterRules

> GetRouteFilterRulesResponse GetRouteFilterRules(ctx, routeFilterId).Offset(offset).Limit(limit).Execute()

GetRFRules



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.GetRouteFilterRules(context.Background(), routeFilterId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.GetRouteFilterRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterRules`: GetRouteFilterRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.GetRouteFilterRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetRouteFilterRulesResponse**](GetRouteFilterRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRouteFilterRuleByUuid

> RouteFilterRulesData PatchRouteFilterRuleByUuid(ctx, routeFilterId, routeFilterRuleId).RouteFilterRulesPatchRequestItem(routeFilterRulesPatchRequestItem).Execute()

PatchRFilterRule



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRuleId := "routeFilterRuleId_example" // string | Route  Filter  Rules Id
	routeFilterRulesPatchRequestItem := []openapiclient.RouteFilterRulesPatchRequestItem{*openapiclient.NewRouteFilterRulesPatchRequestItem("replace", "/prefixMatch", map[string]interface{}(123))} // []RouteFilterRulesPatchRequestItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.PatchRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).RouteFilterRulesPatchRequestItem(routeFilterRulesPatchRequestItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.PatchRouteFilterRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRouteFilterRuleByUuid`: RouteFilterRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.PatchRouteFilterRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**routeFilterRuleId** | **string** | Route  Filter  Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRouteFilterRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routeFilterRulesPatchRequestItem** | [**[]RouteFilterRulesPatchRequestItem**](RouteFilterRulesPatchRequestItem.md) |  | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRouteFilterRuleByUuid

> RouteFilterRulesData ReplaceRouteFilterRuleByUuid(ctx, routeFilterId, routeFilterRuleId).RouteFilterRulesBase(routeFilterRulesBase).Execute()

ReplaceRFRule



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
	routeFilterId := "routeFilterId_example" // string | Route Filters Id
	routeFilterRuleId := "routeFilterRuleId_example" // string | Route  Filter  Rules Id
	routeFilterRulesBase := *openapiclient.NewRouteFilterRulesBase("192.168.0.0/24") // RouteFilterRulesBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFilterRulesApi.ReplaceRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).RouteFilterRulesBase(routeFilterRulesBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFilterRulesApi.ReplaceRouteFilterRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRouteFilterRuleByUuid`: RouteFilterRulesData
	fmt.Fprintf(os.Stdout, "Response from `RouteFilterRulesApi.ReplaceRouteFilterRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**routeFilterRuleId** | **string** | Route  Filter  Rules Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRouteFilterRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routeFilterRulesBase** | [**RouteFilterRulesBase**](RouteFilterRulesBase.md) |  | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


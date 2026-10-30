# \RouteAggregationsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachConnectionRouteAggregation**](RouteAggregationsApi.md#AttachConnectionRouteAggregation) | **Put** /fabric/v4/connections/{connectionId}/routeAggregations/{routeAggregationId} | Attach Aggregation
[**CreateRouteAggregation**](RouteAggregationsApi.md#CreateRouteAggregation) | **Post** /fabric/v4/routeAggregations | Create Route Aggregation
[**DeleteRouteAggregationByUuid**](RouteAggregationsApi.md#DeleteRouteAggregationByUuid) | **Delete** /fabric/v4/routeAggregations/{routeAggregationId} | Delete Aggregation
[**DetachConnectionRouteAggregation**](RouteAggregationsApi.md#DetachConnectionRouteAggregation) | **Delete** /fabric/v4/connections/{connectionId}/routeAggregations/{routeAggregationId} | Detach Aggregation
[**GetConnectionRouteAggregationByUuid**](RouteAggregationsApi.md#GetConnectionRouteAggregationByUuid) | **Get** /fabric/v4/connections/{connectionId}/routeAggregations/{routeAggregationId} | Get Aggregation
[**GetConnectionRouteAggregations**](RouteAggregationsApi.md#GetConnectionRouteAggregations) | **Get** /fabric/v4/connections/{connectionId}/routeAggregations | Get Connection Route Aggregations
[**GetRouteAggregationByUuid**](RouteAggregationsApi.md#GetRouteAggregationByUuid) | **Get** /fabric/v4/routeAggregations/{routeAggregationId} | Get Aggregation
[**GetRouteAggregationChangeByUuid**](RouteAggregationsApi.md#GetRouteAggregationChangeByUuid) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/changes/{changeId} | Get Change By ID
[**GetRouteAggregationChanges**](RouteAggregationsApi.md#GetRouteAggregationChanges) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/changes | Get Route Aggregation Changes
[**GetRouteAggregationConnections**](RouteAggregationsApi.md#GetRouteAggregationConnections) | **Get** /fabric/v4/routeAggregations/{routeAggregationId}/connections | Get Route Aggregation Connections
[**PatchRouteAggregationByUuid**](RouteAggregationsApi.md#PatchRouteAggregationByUuid) | **Patch** /fabric/v4/routeAggregations/{routeAggregationId} | Update Route Aggregation
[**SearchCloudRouterRouteAggregationAttachments**](RouteAggregationsApi.md#SearchCloudRouterRouteAggregationAttachments) | **Post** /fabric/v4/routers/{routerId}/routeAggregations/search | Search Cloud Router Route Aggregation Attachments
[**SearchRouteAggregations**](RouteAggregationsApi.md#SearchRouteAggregations) | **Post** /fabric/v4/routeAggregations/search | Search Route Aggregations



## AttachConnectionRouteAggregation

> ConnectionRouteAggregationData AttachConnectionRouteAggregation(ctx, routeAggregationId, connectionId).Execute()

Attach Aggregation



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
	connectionId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.AttachConnectionRouteAggregation(context.Background(), routeAggregationId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.AttachConnectionRouteAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachConnectionRouteAggregation`: ConnectionRouteAggregationData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.AttachConnectionRouteAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachConnectionRouteAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionRouteAggregationData**](ConnectionRouteAggregationData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouteAggregation

> RouteAggregationsData CreateRouteAggregation(ctx).RouteAggregationsBase(routeAggregationsBase).Execute()

Create Route Aggregation



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
	routeAggregationsBase := *openapiclient.NewRouteAggregationsBase(openapiclient.RouteAggregationsBase_type("BGP_IPv4_PREFIX_AGGREGATION"), "My-direct-route-1", *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // RouteAggregationsBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.CreateRouteAggregation(context.Background()).RouteAggregationsBase(routeAggregationsBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.CreateRouteAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteAggregation`: RouteAggregationsData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.CreateRouteAggregation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeAggregationsBase** | [**RouteAggregationsBase**](RouteAggregationsBase.md) |  | 

### Return type

[**RouteAggregationsData**](RouteAggregationsData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteAggregationByUuid

> RouteAggregationsData DeleteRouteAggregationByUuid(ctx, routeAggregationId).Execute()

Delete Aggregation



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.DeleteRouteAggregationByUuid(context.Background(), routeAggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.DeleteRouteAggregationByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRouteAggregationByUuid`: RouteAggregationsData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.DeleteRouteAggregationByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteAggregationByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteAggregationsData**](RouteAggregationsData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachConnectionRouteAggregation

> ConnectionRouteAggregationData DetachConnectionRouteAggregation(ctx, routeAggregationId, connectionId).Execute()

Detach Aggregation



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
	connectionId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.DetachConnectionRouteAggregation(context.Background(), routeAggregationId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.DetachConnectionRouteAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachConnectionRouteAggregation`: ConnectionRouteAggregationData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.DetachConnectionRouteAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachConnectionRouteAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionRouteAggregationData**](ConnectionRouteAggregationData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRouteAggregationByUuid

> ConnectionRouteAggregationData GetConnectionRouteAggregationByUuid(ctx, routeAggregationId, connectionId).Execute()

Get Aggregation



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
	connectionId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.GetConnectionRouteAggregationByUuid(context.Background(), routeAggregationId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.GetConnectionRouteAggregationByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRouteAggregationByUuid`: ConnectionRouteAggregationData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.GetConnectionRouteAggregationByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRouteAggregationByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionRouteAggregationData**](ConnectionRouteAggregationData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRouteAggregations

> GetAllConnectionRouteAggregationsResponse GetConnectionRouteAggregations(ctx, connectionId).Execute()

Get Connection Route Aggregations



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
	connectionId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.GetConnectionRouteAggregations(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.GetConnectionRouteAggregations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRouteAggregations`: GetAllConnectionRouteAggregationsResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.GetConnectionRouteAggregations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRouteAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllConnectionRouteAggregationsResponse**](GetAllConnectionRouteAggregationsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationByUuid

> RouteAggregationsData GetRouteAggregationByUuid(ctx, routeAggregationId).Execute()

Get Aggregation



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.GetRouteAggregationByUuid(context.Background(), routeAggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.GetRouteAggregationByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationByUuid`: RouteAggregationsData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.GetRouteAggregationByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteAggregationsData**](RouteAggregationsData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationChangeByUuid

> RouteAggregationChangeData GetRouteAggregationChangeByUuid(ctx, routeAggregationId, changeId).Execute()

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
	changeId := "6d500177-9404-41f2-99e8-2bf1a84d8db5" // string | Routing Protocol Change UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.GetRouteAggregationChangeByUuid(context.Background(), routeAggregationId, changeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.GetRouteAggregationChangeByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationChangeByUuid`: RouteAggregationChangeData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.GetRouteAggregationChangeByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 
**changeId** | **string** | Routing Protocol Change UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationChangeByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteAggregationChangeData**](RouteAggregationChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationChanges

> RouteAggregationChangeDataResponse GetRouteAggregationChanges(ctx, routeAggregationId).Offset(offset).Limit(limit).Execute()

Get Route Aggregation Changes



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
	resp, r, err := apiClient.RouteAggregationsApi.GetRouteAggregationChanges(context.Background(), routeAggregationId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.GetRouteAggregationChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationChanges`: RouteAggregationChangeDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.GetRouteAggregationChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**RouteAggregationChangeDataResponse**](RouteAggregationChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAggregationConnections

> GetRouteAggregationGetConnectionsResponse GetRouteAggregationConnections(ctx, routeAggregationId).Execute()

Get Route Aggregation Connections



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.GetRouteAggregationConnections(context.Background(), routeAggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.GetRouteAggregationConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAggregationConnections`: GetRouteAggregationGetConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.GetRouteAggregationConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAggregationConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRouteAggregationGetConnectionsResponse**](GetRouteAggregationGetConnectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRouteAggregationByUuid

> RouteAggregationsData PatchRouteAggregationByUuid(ctx, routeAggregationId).RouteAggregationsPatchRequestItem(routeAggregationsPatchRequestItem).Execute()

Update Route Aggregation



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
	routeAggregationsPatchRequestItem := []openapiclient.RouteAggregationsPatchRequestItem{*openapiclient.NewRouteAggregationsPatchRequestItem("replace", "/name", interface{}(123))} // []RouteAggregationsPatchRequestItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.PatchRouteAggregationByUuid(context.Background(), routeAggregationId).RouteAggregationsPatchRequestItem(routeAggregationsPatchRequestItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.PatchRouteAggregationByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRouteAggregationByUuid`: RouteAggregationsData
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.PatchRouteAggregationByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeAggregationId** | **string** | Route Aggregations Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRouteAggregationByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeAggregationsPatchRequestItem** | [**[]RouteAggregationsPatchRequestItem**](RouteAggregationsPatchRequestItem.md) |  | 

### Return type

[**RouteAggregationsData**](RouteAggregationsData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudRouterRouteAggregationAttachments

> CloudRouterRouteAggregationsSearchResponse SearchCloudRouterRouteAggregationAttachments(ctx, routerId).CloudRouterRouteAggregationsSearchBase(cloudRouterRouteAggregationsSearchBase).Execute()

Search Cloud Router Route Aggregation Attachments



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
	routerId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Cloud Router UUID
	cloudRouterRouteAggregationsSearchBase := *openapiclient.NewCloudRouterRouteAggregationsSearchBase() // CloudRouterRouteAggregationsSearchBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.SearchCloudRouterRouteAggregationAttachments(context.Background(), routerId).CloudRouterRouteAggregationsSearchBase(cloudRouterRouteAggregationsSearchBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.SearchCloudRouterRouteAggregationAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCloudRouterRouteAggregationAttachments`: CloudRouterRouteAggregationsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.SearchCloudRouterRouteAggregationAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Cloud Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudRouterRouteAggregationAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRouterRouteAggregationsSearchBase** | [**CloudRouterRouteAggregationsSearchBase**](CloudRouterRouteAggregationsSearchBase.md) |  | 

### Return type

[**CloudRouterRouteAggregationsSearchResponse**](CloudRouterRouteAggregationsSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRouteAggregations

> RouteAggregationsSearchResponse SearchRouteAggregations(ctx).RouteAggregationsSearchRequest(routeAggregationsSearchRequest).Execute()

Search Route Aggregations



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
	routeAggregationsSearchRequest := *openapiclient.NewRouteAggregationsSearchRequest() // RouteAggregationsSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAggregationsApi.SearchRouteAggregations(context.Background()).RouteAggregationsSearchRequest(routeAggregationsSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAggregationsApi.SearchRouteAggregations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRouteAggregations`: RouteAggregationsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteAggregationsApi.SearchRouteAggregations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRouteAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeAggregationsSearchRequest** | [**RouteAggregationsSearchRequest**](RouteAggregationsSearchRequest.md) |  | 

### Return type

[**RouteAggregationsSearchResponse**](RouteAggregationsSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \RouteFiltersApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachConnectionRouteFilter**](RouteFiltersApi.md#AttachConnectionRouteFilter) | **Put** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Attach Route Filter
[**CreateRouteFilter**](RouteFiltersApi.md#CreateRouteFilter) | **Post** /fabric/v4/routeFilters | Create Route Filters
[**DeleteRouteFilterByUuid**](RouteFiltersApi.md#DeleteRouteFilterByUuid) | **Delete** /fabric/v4/routeFilters/{routeFilterId} | Delete Route Filter
[**DetachConnectionRouteFilter**](RouteFiltersApi.md#DetachConnectionRouteFilter) | **Delete** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Detach Route Filter
[**GetConnectionRouteFilterByUuid**](RouteFiltersApi.md#GetConnectionRouteFilterByUuid) | **Get** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Get Route Filter
[**GetConnectionRouteFilters**](RouteFiltersApi.md#GetConnectionRouteFilters) | **Get** /fabric/v4/connections/{connectionId}/routeFilters | Get All Route Filters
[**GetRouteFilterByUuid**](RouteFiltersApi.md#GetRouteFilterByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId} | Get Route Filter By UUID
[**GetRouteFilterChangeByUuid**](RouteFiltersApi.md#GetRouteFilterChangeByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/changes/{changeId} | Get Change By ID
[**GetRouteFilterChanges**](RouteFiltersApi.md#GetRouteFilterChanges) | **Get** /fabric/v4/routeFilters/{routeFilterId}/changes | Get All Changes
[**GetRouteFilterConnections**](RouteFiltersApi.md#GetRouteFilterConnections) | **Get** /fabric/v4/routeFilters/{routeFilterId}/connections | Get All Connections on Route Filter
[**PatchRouteFilterByUuid**](RouteFiltersApi.md#PatchRouteFilterByUuid) | **Patch** /fabric/v4/routeFilters/{routeFilterId} | Patch Route Filter
[**SearchRouteFilters**](RouteFiltersApi.md#SearchRouteFilters) | **Post** /fabric/v4/routeFilters/search | Search Route Filters



## AttachConnectionRouteFilter

> ConnectionRouteFilterData AttachConnectionRouteFilter(ctx, routeFilterId, connectionId).ConnectionRouteFiltersBase(connectionRouteFiltersBase).Execute()

Attach Route Filter



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
	connectionId := "connectionId_example" // string | Connection Id
	connectionRouteFiltersBase := *openapiclient.NewConnectionRouteFiltersBase(openapiclient.ConnectionRouteFiltersBase_direction("INBOUND")) // ConnectionRouteFiltersBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.AttachConnectionRouteFilter(context.Background(), routeFilterId, connectionId).ConnectionRouteFiltersBase(connectionRouteFiltersBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.AttachConnectionRouteFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachConnectionRouteFilter`: ConnectionRouteFilterData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.AttachConnectionRouteFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachConnectionRouteFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionRouteFiltersBase** | [**ConnectionRouteFiltersBase**](ConnectionRouteFiltersBase.md) |  | 

### Return type

[**ConnectionRouteFilterData**](ConnectionRouteFilterData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouteFilter

> RouteFiltersData CreateRouteFilter(ctx).RouteFiltersBase(routeFiltersBase).Execute()

Create Route Filters



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
	routeFiltersBase := *openapiclient.NewRouteFiltersBase(openapiclient.RouteFiltersBase_type("BGP_IPv4_PREFIX_FILTER"), "My-direct-route-1", *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // RouteFiltersBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.CreateRouteFilter(context.Background()).RouteFiltersBase(routeFiltersBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.CreateRouteFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteFilter`: RouteFiltersData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.CreateRouteFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeFiltersBase** | [**RouteFiltersBase**](RouteFiltersBase.md) |  | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteFilterByUuid

> RouteFiltersData DeleteRouteFilterByUuid(ctx, routeFilterId).Execute()

Delete Route Filter



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.DeleteRouteFilterByUuid(context.Background(), routeFilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.DeleteRouteFilterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRouteFilterByUuid`: RouteFiltersData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.DeleteRouteFilterByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteFilterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachConnectionRouteFilter

> ConnectionRouteFilterData DetachConnectionRouteFilter(ctx, routeFilterId, connectionId).Execute()

Detach Route Filter



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
	connectionId := "connectionId_example" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.DetachConnectionRouteFilter(context.Background(), routeFilterId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.DetachConnectionRouteFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachConnectionRouteFilter`: ConnectionRouteFilterData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.DetachConnectionRouteFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachConnectionRouteFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionRouteFilterData**](ConnectionRouteFilterData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRouteFilterByUuid

> ConnectionRouteFilterData GetConnectionRouteFilterByUuid(ctx, routeFilterId, connectionId).Execute()

Get Route Filter



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
	connectionId := "connectionId_example" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.GetConnectionRouteFilterByUuid(context.Background(), routeFilterId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.GetConnectionRouteFilterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRouteFilterByUuid`: ConnectionRouteFilterData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.GetConnectionRouteFilterByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRouteFilterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionRouteFilterData**](ConnectionRouteFilterData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRouteFilters

> GetAllConnectionRouteFiltersResponse GetConnectionRouteFilters(ctx, connectionId).Execute()

Get All Route Filters



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
	connectionId := "connectionId_example" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.GetConnectionRouteFilters(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.GetConnectionRouteFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRouteFilters`: GetAllConnectionRouteFiltersResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.GetConnectionRouteFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRouteFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllConnectionRouteFiltersResponse**](GetAllConnectionRouteFiltersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterByUuid

> RouteFiltersData GetRouteFilterByUuid(ctx, routeFilterId).Execute()

Get Route Filter By UUID



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.GetRouteFilterByUuid(context.Background(), routeFilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.GetRouteFilterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterByUuid`: RouteFiltersData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.GetRouteFilterByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterChangeByUuid

> RouteFilterChangeData GetRouteFilterChangeByUuid(ctx, routeFilterId, changeId).Execute()

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
	changeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Change UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.GetRouteFilterChangeByUuid(context.Background(), routeFilterId, changeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.GetRouteFilterChangeByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterChangeByUuid`: RouteFilterChangeData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.GetRouteFilterChangeByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 
**changeId** | **string** | Routing Protocol Change UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterChangeByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouteFilterChangeData**](RouteFilterChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterChanges

> RouteFilterChangeDataResponse GetRouteFilterChanges(ctx, routeFilterId).Offset(offset).Limit(limit).Execute()

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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.GetRouteFilterChanges(context.Background(), routeFilterId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.GetRouteFilterChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterChanges`: RouteFilterChangeDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.GetRouteFilterChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**RouteFilterChangeDataResponse**](RouteFilterChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFilterConnections

> GetRouteFilterGetConnectionsResponse GetRouteFilterConnections(ctx, routeFilterId).Execute()

Get All Connections on Route Filter



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.GetRouteFilterConnections(context.Background(), routeFilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.GetRouteFilterConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFilterConnections`: GetRouteFilterGetConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.GetRouteFilterConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFilterConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRouteFilterGetConnectionsResponse**](GetRouteFilterGetConnectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRouteFilterByUuid

> RouteFiltersData PatchRouteFilterByUuid(ctx, routeFilterId).RouteFiltersPatchRequestItem(routeFiltersPatchRequestItem).Execute()

Patch Route Filter



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
	routeFiltersPatchRequestItem := []openapiclient.RouteFiltersPatchRequestItem{*openapiclient.NewRouteFiltersPatchRequestItem("replace", "/name", interface{}(123))} // []RouteFiltersPatchRequestItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.PatchRouteFilterByUuid(context.Background(), routeFilterId).RouteFiltersPatchRequestItem(routeFiltersPatchRequestItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.PatchRouteFilterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRouteFilterByUuid`: RouteFiltersData
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.PatchRouteFilterByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeFilterId** | **string** | Route Filters Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRouteFilterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeFiltersPatchRequestItem** | [**[]RouteFiltersPatchRequestItem**](RouteFiltersPatchRequestItem.md) |  | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRouteFilters

> RouteFiltersSearchResponse SearchRouteFilters(ctx).RouteFiltersSearchBase(routeFiltersSearchBase).Execute()

Search Route Filters



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
	routeFiltersSearchBase := *openapiclient.NewRouteFiltersSearchBase() // RouteFiltersSearchBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteFiltersApi.SearchRouteFilters(context.Background()).RouteFiltersSearchBase(routeFiltersSearchBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteFiltersApi.SearchRouteFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRouteFilters`: RouteFiltersSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `RouteFiltersApi.SearchRouteFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRouteFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeFiltersSearchBase** | [**RouteFiltersSearchBase**](RouteFiltersSearchBase.md) |  | 

### Return type

[**RouteFiltersSearchResponse**](RouteFiltersSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


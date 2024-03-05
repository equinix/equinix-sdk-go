# \CloudRoutersApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudRouter**](CloudRoutersApi.md#CreateCloudRouter) | **Post** /fabric/v4/routers | Create Routers
[**CreateCloudRouterAction**](CloudRoutersApi.md#CreateCloudRouterAction) | **Post** /fabric/v4/routers/{routerId}/actions | Route table actions
[**DeleteCloudRouterByUuid**](CloudRoutersApi.md#DeleteCloudRouterByUuid) | **Delete** /fabric/v4/routers/{routerId} | Delete Routers
[**GetCloudRouterActions**](CloudRoutersApi.md#GetCloudRouterActions) | **Get** /fabric/v4/routers/{routerId}/actions | Get actions
[**GetCloudRouterByUuid**](CloudRoutersApi.md#GetCloudRouterByUuid) | **Get** /fabric/v4/routers/{routerId} | Get Routers
[**GetCloudRouterPackageByCode**](CloudRoutersApi.md#GetCloudRouterPackageByCode) | **Get** /fabric/v4/routerPackages/{routerPackageCode} | Get Package Details
[**GetCloudRouterPackages**](CloudRoutersApi.md#GetCloudRouterPackages) | **Get** /fabric/v4/routerPackages | List Packages
[**SearchCloudRouterRoutes**](CloudRoutersApi.md#SearchCloudRouterRoutes) | **Post** /fabric/v4/routers/{routerId}/routes/search | Search Route Table
[**SearchCloudRouters**](CloudRoutersApi.md#SearchCloudRouters) | **Post** /fabric/v4/routers/search | Search Routers
[**UpdateCloudRouterByUuid**](CloudRoutersApi.md#UpdateCloudRouterByUuid) | **Patch** /fabric/v4/routers/{routerId} | Update Routers



## CreateCloudRouter

> CloudRouter CreateCloudRouter(ctx).CloudRouterPostRequest(cloudRouterPostRequest).Execute()

Create Routers



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
	cloudRouterPostRequest := *openapiclient.NewCloudRouterPostRequest() // CloudRouterPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.CreateCloudRouter(context.Background()).CloudRouterPostRequest(cloudRouterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.CreateCloudRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudRouter`: CloudRouter
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.CreateCloudRouter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRouterPostRequest** | [**CloudRouterPostRequest**](CloudRouterPostRequest.md) |  | 

### Return type

[**CloudRouter**](CloudRouter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudRouterAction

> CloudRouterActionResponse CreateCloudRouterAction(ctx, routerId).CloudRouterActionRequest(cloudRouterActionRequest).Execute()

Route table actions



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
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Router UUID
	cloudRouterActionRequest := *openapiclient.NewCloudRouterActionRequest(openapiclient.CloudRouterActionType("BGP_SESSION_STATUS_UPDATE")) // CloudRouterActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.CreateCloudRouterAction(context.Background(), routerId).CloudRouterActionRequest(cloudRouterActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.CreateCloudRouterAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudRouterAction`: CloudRouterActionResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.CreateCloudRouterAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudRouterActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRouterActionRequest** | [**CloudRouterActionRequest**](CloudRouterActionRequest.md) |  | 

### Return type

[**CloudRouterActionResponse**](CloudRouterActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudRouterByUuid

> DeleteCloudRouterByUuid(ctx, routerId).Execute()

Delete Routers



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
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cloud Router UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudRoutersApi.DeleteCloudRouterByUuid(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.DeleteCloudRouterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Cloud Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudRouterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRouterActions

> CloudRouterActionResponse GetCloudRouterActions(ctx, routerId).State(state).Execute()

Get actions



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
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Router UUID
	state := openapiclient.CloudRouterActionState("DONE") // CloudRouterActionState | Action state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.GetCloudRouterActions(context.Background(), routerId).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.GetCloudRouterActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRouterActions`: CloudRouterActionResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.GetCloudRouterActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRouterActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**CloudRouterActionState**](CloudRouterActionState.md) | Action state | 

### Return type

[**CloudRouterActionResponse**](CloudRouterActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRouterByUuid

> CloudRouter GetCloudRouterByUuid(ctx, routerId).Execute()

Get Routers



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
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cloud Router UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.GetCloudRouterByUuid(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.GetCloudRouterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRouterByUuid`: CloudRouter
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.GetCloudRouterByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Cloud Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRouterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRouter**](CloudRouter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRouterPackageByCode

> CloudRouterPackage GetCloudRouterPackageByCode(ctx, routerPackageCode).Execute()

Get Package Details



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
	routerPackageCode := openapiclient.RouterPackageCode("LAB") // RouterPackageCode | Equinix-assigned Cloud Router package identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.GetCloudRouterPackageByCode(context.Background(), routerPackageCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.GetCloudRouterPackageByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRouterPackageByCode`: CloudRouterPackage
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.GetCloudRouterPackageByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerPackageCode** | [**RouterPackageCode**](.md) | Equinix-assigned Cloud Router package identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRouterPackageByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRouterPackage**](CloudRouterPackage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRouterPackages

> PackageResponse GetCloudRouterPackages(ctx).Offset(offset).Limit(limit).Execute()

List Packages



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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.GetCloudRouterPackages(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.GetCloudRouterPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRouterPackages`: PackageResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.GetCloudRouterPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRouterPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**PackageResponse**](PackageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudRouterRoutes

> RouteTableEntrySearchResponse SearchCloudRouterRoutes(ctx, routerId).RouteTableEntrySearchRequest(routeTableEntrySearchRequest).Execute()

Search Route Table



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
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Router UUID
	routeTableEntrySearchRequest := *openapiclient.NewRouteTableEntrySearchRequest() // RouteTableEntrySearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.SearchCloudRouterRoutes(context.Background(), routerId).RouteTableEntrySearchRequest(routeTableEntrySearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.SearchCloudRouterRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCloudRouterRoutes`: RouteTableEntrySearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.SearchCloudRouterRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudRouterRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeTableEntrySearchRequest** | [**RouteTableEntrySearchRequest**](RouteTableEntrySearchRequest.md) |  | 

### Return type

[**RouteTableEntrySearchResponse**](RouteTableEntrySearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudRouters

> SearchResponse SearchCloudRouters(ctx).CloudRouterSearchRequest(cloudRouterSearchRequest).Execute()

Search Routers



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
	cloudRouterSearchRequest := *openapiclient.NewCloudRouterSearchRequest() // CloudRouterSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.SearchCloudRouters(context.Background()).CloudRouterSearchRequest(cloudRouterSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.SearchCloudRouters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCloudRouters`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.SearchCloudRouters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudRoutersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRouterSearchRequest** | [**CloudRouterSearchRequest**](CloudRouterSearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudRouterByUuid

> CloudRouter UpdateCloudRouterByUuid(ctx, routerId).CloudRouterChangeOperation(cloudRouterChangeOperation).Execute()

Update Routers



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
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cloud Router UUID
	cloudRouterChangeOperation := []openapiclient.CloudRouterChangeOperation{*openapiclient.NewCloudRouterChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), "Path_example", interface{}(123))} // []CloudRouterChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRoutersApi.UpdateCloudRouterByUuid(context.Background(), routerId).CloudRouterChangeOperation(cloudRouterChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRoutersApi.UpdateCloudRouterByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudRouterByUuid`: CloudRouter
	fmt.Fprintf(os.Stdout, "Response from `CloudRoutersApi.UpdateCloudRouterByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Cloud Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudRouterByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRouterChangeOperation** | [**[]CloudRouterChangeOperation**](CloudRouterChangeOperation.md) |  | 

### Return type

[**CloudRouter**](CloudRouter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


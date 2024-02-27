# \RoutingProtocolsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionRoutingProtocol**](RoutingProtocolsApi.md#CreateConnectionRoutingProtocol) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols | Create Protocol
[**CreateConnectionRoutingProtocolsInBulk**](RoutingProtocolsApi.md#CreateConnectionRoutingProtocolsInBulk) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols/bulk | Bulk Create Protocol
[**DeleteConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#DeleteConnectionRoutingProtocolByUuid) | **Delete** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Delete Protocol
[**GetConnectionRoutingProtocolAllBgpActions**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolAllBgpActions) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions | Get BGP Actions
[**GetConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolByUuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Get Protocol
[**GetConnectionRoutingProtocols**](RoutingProtocolsApi.md#GetConnectionRoutingProtocols) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols | GetRoutingProtocols
[**GetConnectionRoutingProtocolsBgpActionByUuid**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolsBgpActionByUuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions/{actionId} | Get BGP Action
[**GetConnectionRoutingProtocolsChangeByUuid**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolsChangeByUuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/changes/{changeId} | Get Change By ID
[**GetConnectionRoutingProtocolsChanges**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolsChanges) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/changes | Get Changes
[**PatchConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#PatchConnectionRoutingProtocolByUuid) | **Patch** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Patch Protocol
[**PostConnectionRoutingProtocolBgpActionByUuid**](RoutingProtocolsApi.md#PostConnectionRoutingProtocolBgpActionByUuid) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions | Clear/Reset BGP
[**ReplaceConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#ReplaceConnectionRoutingProtocolByUuid) | **Put** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Replace Protocol
[**ValidateRoutingProtocol**](RoutingProtocolsApi.md#ValidateRoutingProtocol) | **Post** /fabric/v4/routers/{routerId}/validate | Validate Subnet



## CreateConnectionRoutingProtocol

> RoutingProtocolData CreateConnectionRoutingProtocol(ctx, connectionId).RoutingProtocolBase(routingProtocolBase).Execute()

Create Protocol



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
	routingProtocolBase := openapiclient.RoutingProtocolBase{RoutingProtocolBGPType: openapiclient.NewRoutingProtocolBGPType(openapiclient.RoutingProtocolBGPType_type("BGP"))} // RoutingProtocolBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.CreateConnectionRoutingProtocol(context.Background(), connectionId).RoutingProtocolBase(routingProtocolBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.CreateConnectionRoutingProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionRoutingProtocol`: RoutingProtocolData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.CreateConnectionRoutingProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRoutingProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routingProtocolBase** | [**RoutingProtocolBase**](RoutingProtocolBase.md) |  | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionRoutingProtocolsInBulk

> GetResponse CreateConnectionRoutingProtocolsInBulk(ctx, connectionId).ConnectionRoutingProtocolPostRequest(connectionRoutingProtocolPostRequest).Execute()

Bulk Create Protocol



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
	connectionRoutingProtocolPostRequest := *openapiclient.NewConnectionRoutingProtocolPostRequest() // ConnectionRoutingProtocolPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.CreateConnectionRoutingProtocolsInBulk(context.Background(), connectionId).ConnectionRoutingProtocolPostRequest(connectionRoutingProtocolPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.CreateConnectionRoutingProtocolsInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionRoutingProtocolsInBulk`: GetResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.CreateConnectionRoutingProtocolsInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRoutingProtocolsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRoutingProtocolPostRequest** | [**ConnectionRoutingProtocolPostRequest**](ConnectionRoutingProtocolPostRequest.md) |  | 

### Return type

[**GetResponse**](GetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionRoutingProtocolByUuid

> RoutingProtocolData DeleteConnectionRoutingProtocolByUuid(ctx, routingProtocolId, connectionId).Execute()

Delete Protocol



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	connectionId := "connectionId_example" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.DeleteConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.DeleteConnectionRoutingProtocolByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnectionRoutingProtocolByUuid`: RoutingProtocolData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.DeleteConnectionRoutingProtocolByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingProtocolId** | **string** | Routing Protocol Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRoutingProtocolByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRoutingProtocolAllBgpActions

> BGPActionsBulkData GetConnectionRoutingProtocolAllBgpActions(ctx, routingProtocolId, connectionId).Offset(offset).Limit(limit).Execute()

Get BGP Actions



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	connectionId := "connectionId_example" // string | Connection Id
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolAllBgpActions(context.Background(), routingProtocolId, connectionId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.GetConnectionRoutingProtocolAllBgpActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRoutingProtocolAllBgpActions`: BGPActionsBulkData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.GetConnectionRoutingProtocolAllBgpActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingProtocolId** | **string** | Routing Protocol Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRoutingProtocolAllBgpActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**BGPActionsBulkData**](BGPActionsBulkData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRoutingProtocolByUuid

> RoutingProtocolData GetConnectionRoutingProtocolByUuid(ctx, routingProtocolId, connectionId).Execute()

Get Protocol



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	connectionId := "connectionId_example" // string | Connection Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.GetConnectionRoutingProtocolByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRoutingProtocolByUuid`: RoutingProtocolData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.GetConnectionRoutingProtocolByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingProtocolId** | **string** | Routing Protocol Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRoutingProtocolByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRoutingProtocols

> GetResponse GetConnectionRoutingProtocols(ctx, connectionId).Offset(offset).Limit(limit).Execute()

GetRoutingProtocols



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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocols(context.Background(), connectionId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.GetConnectionRoutingProtocols``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRoutingProtocols`: GetResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.GetConnectionRoutingProtocols`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRoutingProtocolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetResponse**](GetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRoutingProtocolsBgpActionByUuid

> BGPActionData GetConnectionRoutingProtocolsBgpActionByUuid(ctx, connectionId, routingProtocolId, actionId).Execute()

Get BGP Action



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	actionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | BGP Action UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolsBgpActionByUuid(context.Background(), connectionId, routingProtocolId, actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.GetConnectionRoutingProtocolsBgpActionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRoutingProtocolsBgpActionByUuid`: BGPActionData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.GetConnectionRoutingProtocolsBgpActionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 
**routingProtocolId** | **string** | Routing Protocol Id | 
**actionId** | **string** | BGP Action UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRoutingProtocolsBgpActionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BGPActionData**](BGPActionData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRoutingProtocolsChangeByUuid

> RoutingProtocolChangeData GetConnectionRoutingProtocolsChangeByUuid(ctx, connectionId, routingProtocolId, changeId).Execute()

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
	connectionId := "connectionId_example" // string | Connection Id
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	changeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Change UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolsChangeByUuid(context.Background(), connectionId, routingProtocolId, changeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.GetConnectionRoutingProtocolsChangeByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRoutingProtocolsChangeByUuid`: RoutingProtocolChangeData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.GetConnectionRoutingProtocolsChangeByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 
**routingProtocolId** | **string** | Routing Protocol Id | 
**changeId** | **string** | Routing Protocol Change UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRoutingProtocolsChangeByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoutingProtocolChangeData**](RoutingProtocolChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionRoutingProtocolsChanges

> RoutingProtocolChangeDataResponse GetConnectionRoutingProtocolsChanges(ctx, connectionId, routingProtocolId).Offset(offset).Limit(limit).Execute()

Get Changes



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolsChanges(context.Background(), connectionId, routingProtocolId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.GetConnectionRoutingProtocolsChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionRoutingProtocolsChanges`: RoutingProtocolChangeDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.GetConnectionRoutingProtocolsChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 
**routingProtocolId** | **string** | Routing Protocol Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRoutingProtocolsChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**RoutingProtocolChangeDataResponse**](RoutingProtocolChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConnectionRoutingProtocolByUuid

> RoutingProtocolData PatchConnectionRoutingProtocolByUuid(ctx, routingProtocolId, connectionId).ConnectionChangeOperation(connectionChangeOperation).Execute()

Patch Protocol



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	connectionId := "connectionId_example" // string | Connection Id
	connectionChangeOperation := []openapiclient.ConnectionChangeOperation{*openapiclient.NewConnectionChangeOperation("add", "/ipv6", map[string]interface{}(123))} // []ConnectionChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.PatchConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).ConnectionChangeOperation(connectionChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.PatchConnectionRoutingProtocolByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConnectionRoutingProtocolByUuid`: RoutingProtocolData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.PatchConnectionRoutingProtocolByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingProtocolId** | **string** | Routing Protocol Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConnectionRoutingProtocolByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionChangeOperation** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md) |  | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectionRoutingProtocolBgpActionByUuid

> BGPActionData PostConnectionRoutingProtocolBgpActionByUuid(ctx, routingProtocolId, connectionId).BGPActionRequest(bGPActionRequest).Execute()

Clear/Reset BGP



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	connectionId := "connectionId_example" // string | Connection Id
	bGPActionRequest := *openapiclient.NewBGPActionRequest(openapiclient.BGPActions("CLEAR_BGPIPV4")) // BGPActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.PostConnectionRoutingProtocolBgpActionByUuid(context.Background(), routingProtocolId, connectionId).BGPActionRequest(bGPActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.PostConnectionRoutingProtocolBgpActionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectionRoutingProtocolBgpActionByUuid`: BGPActionData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.PostConnectionRoutingProtocolBgpActionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingProtocolId** | **string** | Routing Protocol Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectionRoutingProtocolBgpActionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bGPActionRequest** | [**BGPActionRequest**](BGPActionRequest.md) |  | 

### Return type

[**BGPActionData**](BGPActionData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceConnectionRoutingProtocolByUuid

> RoutingProtocolData ReplaceConnectionRoutingProtocolByUuid(ctx, routingProtocolId, connectionId).RoutingProtocolBase(routingProtocolBase).Execute()

Replace Protocol



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
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol Id
	connectionId := "connectionId_example" // string | Connection Id
	routingProtocolBase := openapiclient.RoutingProtocolBase{RoutingProtocolBGPType: openapiclient.NewRoutingProtocolBGPType(openapiclient.RoutingProtocolBGPType_type("BGP"))} // RoutingProtocolBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.ReplaceConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).RoutingProtocolBase(routingProtocolBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.ReplaceConnectionRoutingProtocolByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceConnectionRoutingProtocolByUuid`: RoutingProtocolData
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.ReplaceConnectionRoutingProtocolByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingProtocolId** | **string** | Routing Protocol Id | 
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceConnectionRoutingProtocolByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routingProtocolBase** | [**RoutingProtocolBase**](RoutingProtocolBase.md) |  | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRoutingProtocol

> ValidateSubnetResponse ValidateRoutingProtocol(ctx, routerId).ValidateRequest(validateRequest).Execute()

Validate Subnet



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
	validateRequest := *openapiclient.NewValidateRequest() // ValidateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingProtocolsApi.ValidateRoutingProtocol(context.Background(), routerId).ValidateRequest(validateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingProtocolsApi.ValidateRoutingProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateRoutingProtocol`: ValidateSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingProtocolsApi.ValidateRoutingProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | Cloud Router UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRoutingProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validateRequest** | [**ValidateRequest**](ValidateRequest.md) |  | 

### Return type

[**ValidateSubnetResponse**](ValidateSubnetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


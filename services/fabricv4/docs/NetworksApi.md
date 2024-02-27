# \NetworksApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](NetworksApi.md#CreateNetwork) | **Post** /fabric/v4/networks | Create Network
[**DeleteNetworkByUuid**](NetworksApi.md#DeleteNetworkByUuid) | **Delete** /fabric/v4/networks/{networkId} | Delete Network By ID
[**GetConnectionsByNetworkUuid**](NetworksApi.md#GetConnectionsByNetworkUuid) | **Get** /fabric/v4/networks/{networkId}/connections | Get Connections
[**GetNetworkByUuid**](NetworksApi.md#GetNetworkByUuid) | **Get** /fabric/v4/networks/{networkId} | Get Network By ID
[**GetNetworkChangeByUuid**](NetworksApi.md#GetNetworkChangeByUuid) | **Get** /fabric/v4/networks/{networkId}/changes/{changeId} | Get Change By ID
[**GetNetworkChanges**](NetworksApi.md#GetNetworkChanges) | **Get** /fabric/v4/networks/{networkId}/changes | Get Network Changes
[**SearchNetworks**](NetworksApi.md#SearchNetworks) | **Post** /fabric/v4/networks/search | Search Network
[**UpdateNetworkByUuid**](NetworksApi.md#UpdateNetworkByUuid) | **Patch** /fabric/v4/networks/{networkId} | Update Network By ID



## CreateNetwork

> Network CreateNetwork(ctx).NetworkPostRequest(networkPostRequest).Execute()

Create Network



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
	networkPostRequest := *openapiclient.NewNetworkPostRequest(openapiclient.NetworkType("EVPLAN"), "Name_example", openapiclient.NetworkScope("REGIONAL"), []openapiclient.SimplifiedNotification{*openapiclient.NewSimplifiedNotification(openapiclient.SimplifiedNotification_type("NOTIFICATION"), []string{"Emails_example"})}) // NetworkPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.CreateNetwork(context.Background()).NetworkPostRequest(networkPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetwork`: Network
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkPostRequest** | [**NetworkPostRequest**](NetworkPostRequest.md) |  | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkByUuid

> Network DeleteNetworkByUuid(ctx, networkId).Execute()

Delete Network By ID



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
	networkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.DeleteNetworkByUuid(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkByUuid`: Network
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsByNetworkUuid

> NetworkConnections GetConnectionsByNetworkUuid(ctx, networkId).Execute()

Get Connections



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
	networkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.GetConnectionsByNetworkUuid(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetConnectionsByNetworkUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsByNetworkUuid`: NetworkConnections
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetConnectionsByNetworkUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsByNetworkUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkConnections**](NetworkConnections.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkByUuid

> Network GetNetworkByUuid(ctx, networkId).Execute()

Get Network By ID



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
	networkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.GetNetworkByUuid(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkByUuid`: Network
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkChangeByUuid

> NetworkChange GetNetworkChangeByUuid(ctx, networkId, changeId).Execute()

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
	networkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network UUID
	changeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network Change UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.GetNetworkChangeByUuid(context.Background(), networkId, changeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkChangeByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkChangeByUuid`: NetworkChange
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkChangeByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network UUID | 
**changeId** | **string** | Network Change UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkChangeByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkChange**](NetworkChange.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkChanges

> NetworkChangeResponse GetNetworkChanges(ctx, networkId).Execute()

Get Network Changes



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
	networkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.GetNetworkChanges(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkChanges`: NetworkChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkChangeResponse**](NetworkChangeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchNetworks

> NetworkSearchResponse SearchNetworks(ctx).NetworkSearchRequest(networkSearchRequest).Execute()

Search Network



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
	networkSearchRequest := *openapiclient.NewNetworkSearchRequest() // NetworkSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.SearchNetworks(context.Background()).NetworkSearchRequest(networkSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.SearchNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNetworks`: NetworkSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.SearchNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkSearchRequest** | [**NetworkSearchRequest**](NetworkSearchRequest.md) |  | 

### Return type

[**NetworkSearchResponse**](NetworkSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkByUuid

> Network UpdateNetworkByUuid(ctx, networkId).NetworkChangeOperation(networkChangeOperation).Execute()

Update Network By ID



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
	networkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Network UUID
	networkChangeOperation := []openapiclient.NetworkChangeOperation{*openapiclient.NewNetworkChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), "/name", map[string]interface{}(123))} // []NetworkChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksApi.UpdateNetworkByUuid(context.Background(), networkId).NetworkChangeOperation(networkChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkByUuid`: Network
	fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkChangeOperation** | [**[]NetworkChangeOperation**](NetworkChangeOperation.md) |  | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


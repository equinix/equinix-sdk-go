# \GatewaysApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](GatewaysApi.md#CreateGateway) | **Post** /fabric/v4/gateways | Create Gateway
[**DeleteGatewayByUuid**](GatewaysApi.md#DeleteGatewayByUuid) | **Delete** /fabric/v4/gateways/{gatewayId} | Delete Gateway
[**GetGatewayByUuid**](GatewaysApi.md#GetGatewayByUuid) | **Get** /fabric/v4/gateways/{gatewayId} | Get Gateway
[**UpdateGatewayByUuid**](GatewaysApi.md#UpdateGatewayByUuid) | **Patch** /fabric/v4/gateways/{gatewayId} | Update Gateway by ID



## CreateGateway

> Gateway CreateGateway(ctx).GatewayPostRequest(gatewayPostRequest).Execute()

Create Gateway



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
	gatewayPostRequest := *openapiclient.NewGatewayPostRequest(openapiclient.GatewayType("VPN_GW"), "test-gateway", int32(123), int32(123), *openapiclient.NewRouter(), *openapiclient.NewSimplifiedAccount(), *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // GatewayPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysApi.CreateGateway(context.Background()).GatewayPostRequest(gatewayPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.CreateGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGateway`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.CreateGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayPostRequest** | [**GatewayPostRequest**](GatewayPostRequest.md) |  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGatewayByUuid

> Gateway DeleteGatewayByUuid(ctx, gatewayId).Execute()

Delete Gateway



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
	gatewayId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Gateway UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysApi.DeleteGatewayByUuid(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.DeleteGatewayByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGatewayByUuid`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.DeleteGatewayByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Gateway**](Gateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayByUuid

> Gateway GetGatewayByUuid(ctx, gatewayId).Execute()

Get Gateway



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
	gatewayId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Gateway UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysApi.GetGatewayByUuid(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.GetGatewayByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGatewayByUuid`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.GetGatewayByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Gateway**](Gateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGatewayByUuid

> Gateway UpdateGatewayByUuid(ctx, gatewayId).GatewayChangeOperation(gatewayChangeOperation).Execute()

Update Gateway by ID



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
	gatewayId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Gateway UUID
	gatewayChangeOperation := []openapiclient.GatewayChangeOperation{*openapiclient.NewGatewayChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), "Path_example", map[string]interface{}(123))} // []GatewayChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysApi.UpdateGatewayByUuid(context.Background(), gatewayId).GatewayChangeOperation(gatewayChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.UpdateGatewayByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGatewayByUuid`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.UpdateGatewayByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayChangeOperation** | [**[]GatewayChangeOperation**](GatewayChangeOperation.md) |  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


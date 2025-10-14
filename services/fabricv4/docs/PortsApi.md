# \PortsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToLag**](PortsApi.md#AddToLag) | **Post** /fabric/v4/ports/{portId}/physicalPorts/bulk | Add to Lag
[**CreatePort**](PortsApi.md#CreatePort) | **Post** /fabric/v4/ports | Create Port
[**DeletePort**](PortsApi.md#DeletePort) | **Delete** /fabric/v4/ports/{portId} | Delete a single port
[**GetPortByUuid**](PortsApi.md#GetPortByUuid) | **Get** /fabric/v4/ports/{portId} | Get Port by uuid
[**GetPorts**](PortsApi.md#GetPorts) | **Get** /fabric/v4/ports | Get All Ports
[**GetVlans**](PortsApi.md#GetVlans) | **Get** /fabric/v4/ports/{portUuid}/linkProtocols | Get Vlans
[**SearchPorts**](PortsApi.md#SearchPorts) | **Post** /fabric/v4/ports/search | Search ports
[**UpdatePortByUuid**](PortsApi.md#UpdatePortByUuid) | **Patch** /fabric/v4/ports/{portId} | Update by UUID



## AddToLag

> AllPhysicalPortsResponse AddToLag(ctx, portId).BulkPhysicalPort(bulkPhysicalPort).Execute()

Add to Lag



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
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
	bulkPhysicalPort := *openapiclient.NewBulkPhysicalPort() // BulkPhysicalPort | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.AddToLag(context.Background(), portId).BulkPhysicalPort(bulkPhysicalPort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.AddToLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddToLag`: AllPhysicalPortsResponse
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.AddToLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkPhysicalPort** | [**BulkPhysicalPort**](BulkPhysicalPort.md) |  | 

### Return type

[**AllPhysicalPortsResponse**](AllPhysicalPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePort

> Port CreatePort(ctx).PortRequest(portRequest).DryRun(dryRun).Execute()

Create Port



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
	portRequest := *openapiclient.NewPortRequest(openapiclient.PortType("XF_PORT"), int32(123), openapiclient.Port_physicalPortsType("1000BASE_LX"), openapiclient.Port_connectivitySourceType("COLO"), *openapiclient.NewSimplifiedAccount(), *openapiclient.NewSimplifiedLocation(), *openapiclient.NewPortEncapsulation(), *openapiclient.NewPortSettings()) // PortRequest | 
	dryRun := true // bool | option to verify that API calls will succeed (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.CreatePort(context.Background()).PortRequest(portRequest).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.CreatePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePort`: Port
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.CreatePort`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portRequest** | [**PortRequest**](PortRequest.md) |  | 
 **dryRun** | **bool** | option to verify that API calls will succeed | [default to false]

### Return type

[**Port**](Port.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePort

> Port DeletePort(ctx, portId).DryRun(dryRun).Execute()

Delete a single port



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
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
	dryRun := true // bool | option to verify that API calls will succeed (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.DeletePort(context.Background(), portId).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.DeletePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePort`: Port
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.DeletePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** | option to verify that API calls will succeed | [default to false]

### Return type

[**Port**](Port.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortByUuid

> Port GetPortByUuid(ctx, portId).Execute()

Get Port by uuid



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
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.GetPortByUuid(context.Background(), portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.GetPortByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortByUuid`: Port
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.GetPortByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Port**](Port.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPorts

> AllPortsResponse GetPorts(ctx).Name(name).Execute()

Get All Ports



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
	name := "name_example" // string | port name to be provided if specific port(s) to be retrieved (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.GetPorts(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.GetPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPorts`: AllPortsResponse
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.GetPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | port name to be provided if specific port(s) to be retrieved | 

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlans

> LinkProtocolGetResponse GetVlans(ctx, portUuid).Execute()

Get Vlans



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
	portUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.GetVlans(context.Background(), portUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.GetVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVlans`: LinkProtocolGetResponse
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.GetVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portUuid** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkProtocolGetResponse**](LinkProtocolGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPorts

> AllPortsResponse SearchPorts(ctx).PortV4SearchRequest(portV4SearchRequest).Execute()

Search ports



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
	portV4SearchRequest := *openapiclient.NewPortV4SearchRequest(*openapiclient.NewPortExpression()) // PortV4SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.SearchPorts(context.Background()).PortV4SearchRequest(portV4SearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.SearchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPorts`: AllPortsResponse
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.SearchPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portV4SearchRequest** | [**PortV4SearchRequest**](PortV4SearchRequest.md) |  | 

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePortByUuid

> AllPortsResponse UpdatePortByUuid(ctx, portId).PortChangeOperation(portChangeOperation).DryRun(dryRun).Execute()

Update by UUID



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
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID
	portChangeOperation := []openapiclient.PortChangeOperation{*openapiclient.NewPortChangeOperation("replace", "/name", interface{}(123))} // []PortChangeOperation | 
	dryRun := true // bool | option to verify that API calls will succeed (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsApi.UpdatePortByUuid(context.Background(), portId).PortChangeOperation(portChangeOperation).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.UpdatePortByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortByUuid`: AllPortsResponse
	fmt.Fprintf(os.Stdout, "Response from `PortsApi.UpdatePortByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portChangeOperation** | [**[]PortChangeOperation**](PortChangeOperation.md) |  | 
 **dryRun** | **bool** | option to verify that API calls will succeed | [default to false]

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


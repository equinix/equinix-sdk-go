# \FabricOneApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterconnect**](FabricOneApi.md#CreateInterconnect) | **Post** /fabric/v4/interconnects | Create Interconnect
[**DeleteInterconnectByUuid**](FabricOneApi.md#DeleteInterconnectByUuid) | **Delete** /fabric/v4/interconnects/{interconnectId} | Delete Interconnect By ID
[**GetInterconnectByUuid**](FabricOneApi.md#GetInterconnectByUuid) | **Get** /fabric/v4/interconnects/{interconnectId} | Get Interconnect By ID
[**GetInterconnectPackages**](FabricOneApi.md#GetInterconnectPackages) | **Get** /fabric/v4/interconnectPackages | Get All Interconnect Packages
[**SearchInterconnects**](FabricOneApi.md#SearchInterconnects) | **Post** /fabric/v4/interconnects/search | Search Interconnects



## CreateInterconnect

> Interconnect CreateInterconnect(ctx).InterconnectPostRequest(interconnectPostRequest).Execute()

Create Interconnect



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
	interconnectPostRequest := *openapiclient.NewInterconnectPostRequest(openapiclient.InterconnectPostRequest_type("XF_IC"), "test ic sv", *openapiclient.NewInterconnectLocationRequest(), *openapiclient.NewInterconnectPackage("Type_example", "Code_example"), *openapiclient.NewSimplifiedAccount(), *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5"), []openapiclient.InterconnectNotification{*openapiclient.NewInterconnectNotification()}) // InterconnectPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FabricOneApi.CreateInterconnect(context.Background()).InterconnectPostRequest(interconnectPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FabricOneApi.CreateInterconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInterconnect`: Interconnect
	fmt.Fprintf(os.Stdout, "Response from `FabricOneApi.CreateInterconnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interconnectPostRequest** | [**InterconnectPostRequest**](InterconnectPostRequest.md) |  | 

### Return type

[**Interconnect**](Interconnect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterconnectByUuid

> Interconnect DeleteInterconnectByUuid(ctx, interconnectId).Execute()

Delete Interconnect By ID



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
	interconnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Interconnect UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FabricOneApi.DeleteInterconnectByUuid(context.Background(), interconnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FabricOneApi.DeleteInterconnectByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInterconnectByUuid`: Interconnect
	fmt.Fprintf(os.Stdout, "Response from `FabricOneApi.DeleteInterconnectByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interconnectId** | **string** | Interconnect UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterconnectByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Interconnect**](Interconnect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnectByUuid

> Interconnect GetInterconnectByUuid(ctx, interconnectId).Execute()

Get Interconnect By ID



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
	interconnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Interconnect UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FabricOneApi.GetInterconnectByUuid(context.Background(), interconnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FabricOneApi.GetInterconnectByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterconnectByUuid`: Interconnect
	fmt.Fprintf(os.Stdout, "Response from `FabricOneApi.GetInterconnectByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interconnectId** | **string** | Interconnect UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Interconnect**](Interconnect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnectPackages

> AllInterconnectPackagesResponse GetInterconnectPackages(ctx).Offset(offset).Limit(limit).Execute()

Get All Interconnect Packages



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
	resp, r, err := apiClient.FabricOneApi.GetInterconnectPackages(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FabricOneApi.GetInterconnectPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterconnectPackages`: AllInterconnectPackagesResponse
	fmt.Fprintf(os.Stdout, "Response from `FabricOneApi.GetInterconnectPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**AllInterconnectPackagesResponse**](AllInterconnectPackagesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchInterconnects

> InterconnectSearchResponse SearchInterconnects(ctx).InterconnectSearchRequest(interconnectSearchRequest).Execute()

Search Interconnects



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
	interconnectSearchRequest := *openapiclient.NewInterconnectSearchRequest() // InterconnectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FabricOneApi.SearchInterconnects(context.Background()).InterconnectSearchRequest(interconnectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FabricOneApi.SearchInterconnects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchInterconnects`: InterconnectSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `FabricOneApi.SearchInterconnects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchInterconnectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interconnectSearchRequest** | [**InterconnectSearchRequest**](InterconnectSearchRequest.md) |  | 

### Return type

[**InterconnectSearchResponse**](InterconnectSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


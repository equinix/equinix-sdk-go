# \InternetAccessServicesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEiaService**](InternetAccessServicesApi.md#CreateEiaService) | **Post** /fabric/v4/internetAccessServices | Creates Internet Access Service
[**DeleteEiaService**](InternetAccessServicesApi.md#DeleteEiaService) | **Delete** /fabric/v4/internetAccessServices/{uuid} | Delete Internet Access Service by UUID
[**GetEiaService**](InternetAccessServicesApi.md#GetEiaService) | **Get** /fabric/v4/internetAccessServices/{uuid} | Retrieve Internet Access Service by UUID
[**PatchEiaService**](InternetAccessServicesApi.md#PatchEiaService) | **Patch** /fabric/v4/internetAccessServices/{uuid} | Patch Internet Access Service by UUID
[**SearchEiaServices**](InternetAccessServicesApi.md#SearchEiaServices) | **Post** /fabric/v4/internetAccessServices/search | Search for Internet Access Services



## CreateEiaService

> InternetAccessService CreateEiaService(ctx).InternetAccessPostRequest(internetAccessPostRequest).Execute()

Creates Internet Access Service



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
	internetAccessPostRequest := *openapiclient.NewInternetAccessPostRequest(openapiclient.InternetAccessServiceType("SINGLE_IA"), "Name_example", *openapiclient.NewInternetAccessRoutingProtocolRequest(openapiclient.InternetAccessRoutingProtocolType("BGP"), []openapiclient.InternetAccessCustomerRouteRequest{*openapiclient.NewInternetAccessCustomerRouteRequest(*openapiclient.NewInternetAccessIpBlockRequest("Uuid_example"))}), *openapiclient.NewInternetAccessPostRequestBilling(openapiclient.InternetAccessBillingType("FIXED")), *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5"), *openapiclient.NewInternetAccessAccount("AccountNumber_example")) // InternetAccessPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetAccessServicesApi.CreateEiaService(context.Background()).InternetAccessPostRequest(internetAccessPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetAccessServicesApi.CreateEiaService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEiaService`: InternetAccessService
	fmt.Fprintf(os.Stdout, "Response from `InternetAccessServicesApi.CreateEiaService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEiaServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internetAccessPostRequest** | [**InternetAccessPostRequest**](InternetAccessPostRequest.md) |  | 

### Return type

[**InternetAccessService**](InternetAccessService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEiaService

> InternetAccessService DeleteEiaService(ctx, uuid).Execute()

Delete Internet Access Service by UUID



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the EIA Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetAccessServicesApi.DeleteEiaService(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetAccessServicesApi.DeleteEiaService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEiaService`: InternetAccessService
	fmt.Fprintf(os.Stdout, "Response from `InternetAccessServicesApi.DeleteEiaService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the EIA Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEiaServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternetAccessService**](InternetAccessService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEiaService

> InternetAccessService GetEiaService(ctx, uuid).Execute()

Retrieve Internet Access Service by UUID



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the EIA Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetAccessServicesApi.GetEiaService(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetAccessServicesApi.GetEiaService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEiaService`: InternetAccessService
	fmt.Fprintf(os.Stdout, "Response from `InternetAccessServicesApi.GetEiaService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the EIA Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEiaServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternetAccessService**](InternetAccessService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEiaService

> InternetAccessService PatchEiaService(ctx, uuid).InternetAccessPatchOperationUpdate(internetAccessPatchOperationUpdate).Execute()

Patch Internet Access Service by UUID



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the EIA Service
	internetAccessPatchOperationUpdate := []openapiclient.InternetAccessPatchOperationUpdate{*openapiclient.NewInternetAccessPatchOperationUpdate(openapiclient.InternetAccessPatchOperationUpdateAllowedOp("replace"), "/bandwidth")} // []InternetAccessPatchOperationUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetAccessServicesApi.PatchEiaService(context.Background(), uuid).InternetAccessPatchOperationUpdate(internetAccessPatchOperationUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetAccessServicesApi.PatchEiaService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEiaService`: InternetAccessService
	fmt.Fprintf(os.Stdout, "Response from `InternetAccessServicesApi.PatchEiaService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the EIA Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEiaServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **internetAccessPatchOperationUpdate** | [**[]InternetAccessPatchOperationUpdate**](InternetAccessPatchOperationUpdate.md) |  | 

### Return type

[**InternetAccessService**](InternetAccessService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEiaServices

> InternetAccessServices SearchEiaServices(ctx).InternetAccessSearchRequest(internetAccessSearchRequest).Execute()

Search for Internet Access Services



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
	internetAccessSearchRequest := *openapiclient.NewInternetAccessSearchRequest() // InternetAccessSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetAccessServicesApi.SearchEiaServices(context.Background()).InternetAccessSearchRequest(internetAccessSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetAccessServicesApi.SearchEiaServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEiaServices`: InternetAccessServices
	fmt.Fprintf(os.Stdout, "Response from `InternetAccessServicesApi.SearchEiaServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEiaServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internetAccessSearchRequest** | [**InternetAccessSearchRequest**](InternetAccessSearchRequest.md) |  | 

### Return type

[**InternetAccessServices**](InternetAccessServices.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


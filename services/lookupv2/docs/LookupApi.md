# \LookupApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocationsByPermissionCode**](LookupApi.md#GetLocationsByPermissionCode) | **Get** /colocations/v2/locations | Get Locations by permission code
[**RetrieveAllPatchPanels**](LookupApi.md#RetrieveAllPatchPanels) | **Get** /colocations/v2/patchPanels | Retrieve all patch panels
[**RetrieveListOfConnectionServices**](LookupApi.md#RetrieveListOfConnectionServices) | **Get** /colocations/v2/connectionServices | Retrieve list of connection services
[**RetrieveListOfProviders**](LookupApi.md#RetrieveListOfProviders) | **Get** /colocations/v2/providers | Retrieve list of providers
[**RetrievePatchPanelDetails**](LookupApi.md#RetrievePatchPanelDetails) | **Get** /colocations/v2/patchPanels/{patchPanelId} | Retrieve patch panel details



## GetLocationsByPermissionCode

> LocationsDetailsResponse GetLocationsByPermissionCode(ctx).PermissionCode(permissionCode).Ibxs(ibxs).ProviderAccountNumber(providerAccountNumber).ASideIbx(aSideIbx).ConnectionService(connectionService).Details(details).Execute()

Get Locations by permission code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/lookupv2"
)

func main() {
	permissionCode := openapiclient.Get_Locations_by_permission_code_permissionCode_parameter("CROSS_CONNECT") // GetLocationsByPermissionCodePermissionCodeParameter | List of Permission code to be filtered
	ibxs := []string{"Inner_example"} // []string | List of IBXs to be filtered (optional)
	providerAccountNumber := "providerAccountNumber_example" // string | The service provider's account number (Z-side) linked to their cage. Mandatory when used together with `aSideIbx`. This is only applicable when permissionCode is `CROSS_CONNECT` (optional)
	aSideIbx := "aSideIbx_example" // string | A-Side IBX details to fetch the (Z-side). This is only applicable when permissionCode is `CROSS_CONNECT` (optional)
	connectionService := "connectionService_example" // string | Type of connection service to fetch the Z-side details. This is only applicable when permissionCode is `CROSS_CONNECT` and is required when searching for zSide details. (optional)
	details := true // bool | When `true`, API response returns cage, cabinet and account details (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupApi.GetLocationsByPermissionCode(context.Background()).PermissionCode(permissionCode).Ibxs(ibxs).ProviderAccountNumber(providerAccountNumber).ASideIbx(aSideIbx).ConnectionService(connectionService).Details(details).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetLocationsByPermissionCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocationsByPermissionCode`: LocationsDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetLocationsByPermissionCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsByPermissionCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionCode** | [**GetLocationsByPermissionCodePermissionCodeParameter**](GetLocationsByPermissionCodePermissionCodeParameter.md) | List of Permission code to be filtered | 
 **ibxs** | **[]string** | List of IBXs to be filtered | 
 **providerAccountNumber** | **string** | The service provider&#39;s account number (Z-side) linked to their cage. Mandatory when used together with &#x60;aSideIbx&#x60;. This is only applicable when permissionCode is &#x60;CROSS_CONNECT&#x60; | 
 **aSideIbx** | **string** | A-Side IBX details to fetch the (Z-side). This is only applicable when permissionCode is &#x60;CROSS_CONNECT&#x60; | 
 **connectionService** | **string** | Type of connection service to fetch the Z-side details. This is only applicable when permissionCode is &#x60;CROSS_CONNECT&#x60; and is required when searching for zSide details. | 
 **details** | **bool** | When &#x60;true&#x60;, API response returns cage, cabinet and account details | [default to false]

### Return type

[**LocationsDetailsResponse**](LocationsDetailsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllPatchPanels

> []PatchPanelResponse RetrieveAllPatchPanels(ctx).CabinetId(cabinetId).ProviderAccountNumber(providerAccountNumber).ASideIbx(aSideIbx).AccountNumber(accountNumber).Execute()

Retrieve all patch panels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/lookupv2"
)

func main() {
	cabinetId := "cabinetId_example" // string | ID of the cabinet
	providerAccountNumber := "providerAccountNumber_example" // string | The service provider's account number (Z-side) linked to their cage. Mandatory when used together with `aSideIbx` (optional)
	aSideIbx := "aSideIbx_example" // string | The IBX location code for A-Side. When used together with `providerAccountNumber`, this returns Z-side patch panel details. (optional)
	accountNumber := "accountNumber_example" // string | Account number is only required when cabinet is shared. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupApi.RetrieveAllPatchPanels(context.Background()).CabinetId(cabinetId).ProviderAccountNumber(providerAccountNumber).ASideIbx(aSideIbx).AccountNumber(accountNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.RetrieveAllPatchPanels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllPatchPanels`: []PatchPanelResponse
	fmt.Fprintf(os.Stdout, "Response from `LookupApi.RetrieveAllPatchPanels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllPatchPanelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cabinetId** | **string** | ID of the cabinet | 
 **providerAccountNumber** | **string** | The service provider&#39;s account number (Z-side) linked to their cage. Mandatory when used together with &#x60;aSideIbx&#x60; | 
 **aSideIbx** | **string** | The IBX location code for A-Side. When used together with &#x60;providerAccountNumber&#x60;, this returns Z-side patch panel details. | 
 **accountNumber** | **string** | Account number is only required when cabinet is shared. | 

### Return type

[**[]PatchPanelResponse**](PatchPanelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListOfConnectionServices

> []ConnectionServicesDetailsInner RetrieveListOfConnectionServices(ctx).Ibx(ibx).Execute()

Retrieve list of connection services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/lookupv2"
)

func main() {
	ibx := "ibx_example" // string | IBX

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupApi.RetrieveListOfConnectionServices(context.Background()).Ibx(ibx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.RetrieveListOfConnectionServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveListOfConnectionServices`: []ConnectionServicesDetailsInner
	fmt.Fprintf(os.Stdout, "Response from `LookupApi.RetrieveListOfConnectionServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListOfConnectionServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibx** | **string** | IBX | 

### Return type

[**[]ConnectionServicesDetailsInner**](ConnectionServicesDetailsInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListOfProviders

> []Provider RetrieveListOfProviders(ctx).CageId(cageId).AccountNumber(accountNumber).Execute()

Retrieve list of providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/lookupv2"
)

func main() {
	cageId := "cageId_example" // string | Cage ID
	accountNumber := "accountNumber_example" // string | Account number of A-Side cage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupApi.RetrieveListOfProviders(context.Background()).CageId(cageId).AccountNumber(accountNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.RetrieveListOfProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveListOfProviders`: []Provider
	fmt.Fprintf(os.Stdout, "Response from `LookupApi.RetrieveListOfProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListOfProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cageId** | **string** | Cage ID | 
 **accountNumber** | **string** | Account number of A-Side cage | 

### Return type

[**[]Provider**](Provider.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePatchPanelDetails

> PatchPanelDetails RetrievePatchPanelDetails(ctx, patchPanelId).ProviderAccountNumber(providerAccountNumber).ASideIbx(aSideIbx).AccountNumber(accountNumber).Execute()

Retrieve patch panel details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/lookupv2"
)

func main() {
	patchPanelId := "patchPanelId_example" // string | ID of patch panel
	providerAccountNumber := "providerAccountNumber_example" // string | The service provider's account number (Z-side) linked to their cage. Mandatory when used together with `aSideIbx` (optional)
	aSideIbx := "aSideIbx_example" // string | The IBX location code for A-Side. When used together with `providerAccountNumber` and `accountNumber`, this returns Z-side patch panel details. (optional)
	accountNumber := "accountNumber_example" // string | A-Side cage account number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupApi.RetrievePatchPanelDetails(context.Background(), patchPanelId).ProviderAccountNumber(providerAccountNumber).ASideIbx(aSideIbx).AccountNumber(accountNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.RetrievePatchPanelDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievePatchPanelDetails`: PatchPanelDetails
	fmt.Fprintf(os.Stdout, "Response from `LookupApi.RetrievePatchPanelDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchPanelId** | **string** | ID of patch panel | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePatchPanelDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerAccountNumber** | **string** | The service provider&#39;s account number (Z-side) linked to their cage. Mandatory when used together with &#x60;aSideIbx&#x60; | 
 **aSideIbx** | **string** | The IBX location code for A-Side. When used together with &#x60;providerAccountNumber&#x60; and &#x60;accountNumber&#x60;, this returns Z-side patch panel details. | 
 **accountNumber** | **string** | A-Side cage account number | 

### Return type

[**PatchPanelDetails**](PatchPanelDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


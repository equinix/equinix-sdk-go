# \FirmwareSetsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationFirmwareSets**](FirmwareSetsApi.md#GetOrganizationFirmwareSets) | **Get** /organizations/{id}/firmware-sets | Get Organization&#39;s Firmware Sets
[**GetProjectFirmwareSets**](FirmwareSetsApi.md#GetProjectFirmwareSets) | **Get** /projects/{id}/firmware-sets | Get Project&#39;s Firmware Sets



## GetOrganizationFirmwareSets

> FirmwareSetListResponse GetOrganizationFirmwareSets(ctx, id).Page(page).PerPage(perPage).Execute()

Get Organization's Firmware Sets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix-labs/metal-go/metal/v1"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization UUID
    page := int32(56) // int32 | page number to return (optional)
    perPage := int32(56) // int32 | items returned per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareSetsApi.GetOrganizationFirmwareSets(context.Background(), id).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareSetsApi.GetOrganizationFirmwareSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFirmwareSets`: FirmwareSetListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirmwareSetsApi.GetOrganizationFirmwareSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFirmwareSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | page number to return | 
 **perPage** | **int32** | items returned per page. | 

### Return type

[**FirmwareSetListResponse**](FirmwareSetListResponse.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectFirmwareSets

> FirmwareSetListResponse GetProjectFirmwareSets(ctx, id).Page(page).PerPage(perPage).Execute()

Get Project's Firmware Sets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix-labs/metal-go/metal/v1"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    page := int32(56) // int32 | page number to return (optional)
    perPage := int32(56) // int32 | items returned per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareSetsApi.GetProjectFirmwareSets(context.Background(), id).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareSetsApi.GetProjectFirmwareSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectFirmwareSets`: FirmwareSetListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirmwareSetsApi.GetProjectFirmwareSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectFirmwareSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | page number to return | 
 **perPage** | **int32** | items returned per page. | 

### Return type

[**FirmwareSetListResponse**](FirmwareSetListResponse.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


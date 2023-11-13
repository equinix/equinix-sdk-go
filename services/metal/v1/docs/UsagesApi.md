# \UsagesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindDeviceUsages**](UsagesApi.md#FindDeviceUsages) | **Get** /devices/{id}/usages | Retrieve all usages for device
[**FindProjectUsage**](UsagesApi.md#FindProjectUsage) | **Get** /projects/{id}/usages | Retrieve all usages for project



## FindDeviceUsages

> DeviceUsageList FindDeviceUsages(ctx, id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

Retrieve all usages for device



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Device UUID
    createdAfter := "createdAfter_example" // string | Filter usages created after this date (optional)
    createdBefore := "createdBefore_example" // string | Filter usages created before this date (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsagesApi.FindDeviceUsages(context.Background(), id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesApi.FindDeviceUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDeviceUsages`: DeviceUsageList
    fmt.Fprintf(os.Stdout, "Response from `UsagesApi.FindDeviceUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **string** | Filter usages created after this date | 
 **createdBefore** | **string** | Filter usages created before this date | 

### Return type

[**DeviceUsageList**](DeviceUsageList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectUsage

> ProjectUsageList FindProjectUsage(ctx, id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()

Retrieve all usages for project



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
    createdAfter := "createdAfter_example" // string | Filter usages created after this date (optional)
    createdBefore := "createdBefore_example" // string | Filter usages created before this date (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsagesApi.FindProjectUsage(context.Background(), id).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesApi.FindProjectUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectUsage`: ProjectUsageList
    fmt.Fprintf(os.Stdout, "Response from `UsagesApi.FindProjectUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAfter** | **string** | Filter usages created after this date | 
 **createdBefore** | **string** | Filter usages created before this date | 

### Return type

[**ProjectUsageList**](ProjectUsageList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


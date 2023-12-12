# \CapacityApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCapacityForFacility**](CapacityApi.md#CheckCapacityForFacility) | **Post** /capacity | Check capacity
[**CheckCapacityForMetro**](CapacityApi.md#CheckCapacityForMetro) | **Post** /capacity/metros | Check capacity for a metro
[**FindCapacityForFacility**](CapacityApi.md#FindCapacityForFacility) | **Get** /capacity | View capacity
[**FindCapacityForMetro**](CapacityApi.md#FindCapacityForMetro) | **Get** /capacity/metros | View capacity for metros
[**FindOrganizationCapacityPerFacility**](CapacityApi.md#FindOrganizationCapacityPerFacility) | **Get** /organizations/{id}/capacity | View available hardware plans per Facility for given organization
[**FindOrganizationCapacityPerMetro**](CapacityApi.md#FindOrganizationCapacityPerMetro) | **Get** /organizations/{id}/capacity/metros | View available hardware plans per Metro for given organization



## CheckCapacityForFacility

> CapacityCheckPerFacilityList CheckCapacityForFacility(ctx).CapacityInput(capacityInput).Execute()

Check capacity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
)

func main() {
    capacityInput := *openapiclient.NewCapacityInput() // CapacityInput | Facility to check capacity in

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapacityApi.CheckCapacityForFacility(context.Background()).CapacityInput(capacityInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.CheckCapacityForFacility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCapacityForFacility`: CapacityCheckPerFacilityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.CheckCapacityForFacility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCapacityForFacilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capacityInput** | [**CapacityInput**](CapacityInput.md) | Facility to check capacity in | 

### Return type

[**CapacityCheckPerFacilityList**](CapacityCheckPerFacilityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckCapacityForMetro

> CapacityCheckPerMetroList CheckCapacityForMetro(ctx).CapacityInput(capacityInput).Execute()

Check capacity for a metro



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
)

func main() {
    capacityInput := *openapiclient.NewCapacityInput() // CapacityInput | Metro to check capacity in

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapacityApi.CheckCapacityForMetro(context.Background()).CapacityInput(capacityInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.CheckCapacityForMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCapacityForMetro`: CapacityCheckPerMetroList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.CheckCapacityForMetro`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCapacityForMetroRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capacityInput** | [**CapacityInput**](CapacityInput.md) | Metro to check capacity in | 

### Return type

[**CapacityCheckPerMetroList**](CapacityCheckPerMetroList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCapacityForFacility

> CapacityList FindCapacityForFacility(ctx).Execute()

View capacity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapacityApi.FindCapacityForFacility(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.FindCapacityForFacility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCapacityForFacility`: CapacityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.FindCapacityForFacility`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindCapacityForFacilityRequest struct via the builder pattern


### Return type

[**CapacityList**](CapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCapacityForMetro

> CapacityList FindCapacityForMetro(ctx).Execute()

View capacity for metros



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapacityApi.FindCapacityForMetro(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.FindCapacityForMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCapacityForMetro`: CapacityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.FindCapacityForMetro`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindCapacityForMetroRequest struct via the builder pattern


### Return type

[**CapacityList**](CapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOrganizationCapacityPerFacility

> CapacityList FindOrganizationCapacityPerFacility(ctx, id).Execute()

View available hardware plans per Facility for given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapacityApi.FindOrganizationCapacityPerFacility(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.FindOrganizationCapacityPerFacility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOrganizationCapacityPerFacility`: CapacityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.FindOrganizationCapacityPerFacility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOrganizationCapacityPerFacilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapacityList**](CapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOrganizationCapacityPerMetro

> CapacityList FindOrganizationCapacityPerMetro(ctx, id).Execute()

View available hardware plans per Metro for given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapacityApi.FindOrganizationCapacityPerMetro(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityApi.FindOrganizationCapacityPerMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOrganizationCapacityPerMetro`: CapacityList
    fmt.Fprintf(os.Stdout, "Response from `CapacityApi.FindOrganizationCapacityPerMetro`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOrganizationCapacityPerMetroRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapacityList**](CapacityList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


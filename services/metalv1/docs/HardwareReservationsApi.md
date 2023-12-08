# \HardwareReservationsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateHardwareReservation**](HardwareReservationsApi.md#ActivateHardwareReservation) | **Post** /hardware-reservations/{id}/activate | Activate a spare hardware reservation
[**FindHardwareReservationById**](HardwareReservationsApi.md#FindHardwareReservationById) | **Get** /hardware-reservations/{id} | Retrieve a hardware reservation
[**FindProjectHardwareReservations**](HardwareReservationsApi.md#FindProjectHardwareReservations) | **Get** /projects/{id}/hardware-reservations | Retrieve all hardware reservations for a given project
[**MoveHardwareReservation**](HardwareReservationsApi.md#MoveHardwareReservation) | **Post** /hardware-reservations/{id}/move | Move a hardware reservation



## ActivateHardwareReservation

> HardwareReservation ActivateHardwareReservation(ctx, id).Include(include).Exclude(exclude).ActivateHardwareReservationRequest(activateHardwareReservationRequest).Execute()

Activate a spare hardware reservation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Hardware Reservation UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    activateHardwareReservationRequest := *openapiclient.NewActivateHardwareReservationRequest() // ActivateHardwareReservationRequest | Note to attach to the reservation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.ActivateHardwareReservation(context.Background(), id).Include(include).Exclude(exclude).ActivateHardwareReservationRequest(activateHardwareReservationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.ActivateHardwareReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateHardwareReservation`: HardwareReservation
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.ActivateHardwareReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Hardware Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateHardwareReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **activateHardwareReservationRequest** | [**ActivateHardwareReservationRequest**](ActivateHardwareReservationRequest.md) | Note to attach to the reservation | 

### Return type

[**HardwareReservation**](HardwareReservation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindHardwareReservationById

> HardwareReservation FindHardwareReservationById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve a hardware reservation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | HardwareReservation UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.FindHardwareReservationById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.FindHardwareReservationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindHardwareReservationById`: HardwareReservation
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.FindHardwareReservationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | HardwareReservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindHardwareReservationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**HardwareReservation**](HardwareReservation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectHardwareReservations

> HardwareReservationList FindProjectHardwareReservations(ctx, id).Query(query).State(state).Provisionable(provisionable).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all hardware reservations for a given project


FindProjectHardwareReservations is a paginated API operation. You can specify page number and per page parameters with `Execute`, or to fetch and return all pages of results as a single slice you can call `ExecuteWithPagination()`. `ExecuteWithPagination()`, while convenient, can significantly increase the overhead of API calls in terms of time, network utilization, and memory. Excludes can make the call more efficient by removing any unneeded nested fields that are included by default. If any of the requests fail, the list of results will be nil and the error returned will represent the failed request.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    query := "query_example" // string | Search by facility code, plan name, project name, reservation short ID or device hostname (optional)
    state := openapiclient.findProjectHardwareReservations_state_parameter("active") // FindProjectHardwareReservationsStateParameter | Filter by hardware reservation state (optional)
    provisionable := openapiclient.findProjectHardwareReservations_provisionable_parameter("only") // FindProjectHardwareReservationsProvisionableParameter | Filter hardware reservation that is provisionable (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.FindProjectHardwareReservations(context.Background(), id).Query(query).State(state).Provisionable(provisionable).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.FindProjectHardwareReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectHardwareReservations`: HardwareReservationList
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.FindProjectHardwareReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectHardwareReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Search by facility code, plan name, project name, reservation short ID or device hostname | 
 **state** | [**FindProjectHardwareReservationsStateParameter**](FindProjectHardwareReservationsStateParameter.md) | Filter by hardware reservation state | 
 **provisionable** | [**FindProjectHardwareReservationsProvisionableParameter**](FindProjectHardwareReservationsProvisionableParameter.md) | Filter hardware reservation that is provisionable | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**HardwareReservationList**](HardwareReservationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveHardwareReservation

> HardwareReservation MoveHardwareReservation(ctx, id).MoveHardwareReservationRequest(moveHardwareReservationRequest).Include(include).Exclude(exclude).Execute()

Move a hardware reservation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Hardware Reservation UUID
    moveHardwareReservationRequest := *openapiclient.NewMoveHardwareReservationRequest() // MoveHardwareReservationRequest | Destination Project UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HardwareReservationsApi.MoveHardwareReservation(context.Background(), id).MoveHardwareReservationRequest(moveHardwareReservationRequest).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HardwareReservationsApi.MoveHardwareReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveHardwareReservation`: HardwareReservation
    fmt.Fprintf(os.Stdout, "Response from `HardwareReservationsApi.MoveHardwareReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Hardware Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveHardwareReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveHardwareReservationRequest** | [**MoveHardwareReservationRequest**](MoveHardwareReservationRequest.md) | Destination Project UUID | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**HardwareReservation**](HardwareReservation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


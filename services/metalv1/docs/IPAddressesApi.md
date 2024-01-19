# \IPAddressesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIPAddress**](IPAddressesApi.md#DeleteIPAddress) | **Delete** /ips/{id} | Unassign an ip address
[**FindIPAddressById**](IPAddressesApi.md#FindIPAddressById) | **Get** /ips/{id} | Retrieve an ip address
[**FindIPAddressCustomdata**](IPAddressesApi.md#FindIPAddressCustomdata) | **Get** /ips/{id}/customdata | Retrieve the custom metadata of an IP Reservation or IP Assignment
[**FindIPAvailabilities**](IPAddressesApi.md#FindIPAvailabilities) | **Get** /ips/{id}/available | Retrieve all available subnets of a particular reservation
[**FindIPReservations**](IPAddressesApi.md#FindIPReservations) | **Get** /projects/{id}/ips | Retrieve all ip reservations
[**RequestIPReservation**](IPAddressesApi.md#RequestIPReservation) | **Post** /projects/{id}/ips | Requesting IP reservations
[**UpdateIPAddress**](IPAddressesApi.md#UpdateIPAddress) | **Patch** /ips/{id} | Update an ip address



## DeleteIPAddress

> DeleteIPAddress(ctx, id).Execute()

Unassign an ip address



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | IP Address UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IPAddressesApi.DeleteIPAddress(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.DeleteIPAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IP Address UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAddressById

> FindIPAddressById200Response FindIPAddressById(ctx, id).Include(include).Exclude(exclude).Execute()

Retrieve an ip address



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | IP Address UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPAddressesApi.FindIPAddressById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAddressById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAddressById`: FindIPAddressById200Response
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPAddressById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IP Address UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAddressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindIPAddressById200Response**](FindIPAddressById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAddressCustomdata

> FindIPAddressCustomdata(ctx, id).Execute()

Retrieve the custom metadata of an IP Reservation or IP Assignment



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Ip Reservation UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IPAddressesApi.FindIPAddressCustomdata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAddressCustomdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Ip Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAddressCustomdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPAvailabilities

> IPAvailabilitiesList FindIPAvailabilities(ctx, id).Cidr(cidr).Execute()

Retrieve all available subnets of a particular reservation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | IP Reservation UUID
    cidr := openapiclient.findIPAvailabilities_cidr_parameter("20") // FindIPAvailabilitiesCidrParameter | Size of subnets in bits

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPAddressesApi.FindIPAvailabilities(context.Background(), id).Cidr(cidr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPAvailabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPAvailabilities`: IPAvailabilitiesList
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IP Reservation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cidr** | [**FindIPAvailabilitiesCidrParameter**](FindIPAvailabilitiesCidrParameter.md) | Size of subnets in bits | 

### Return type

[**IPAvailabilitiesList**](IPAvailabilitiesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIPReservations

> IPReservationList FindIPReservations(ctx, id).Types(types).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Retrieve all ip reservations


FindIPReservations is a paginated API operation. You can specify page number and per page parameters with `Execute`, or to fetch and return all pages of results as a single slice you can call `ExecuteWithPagination()`. `ExecuteWithPagination()`, while convenient, can significantly increase the overhead of API calls in terms of time, network utilization, and memory. Excludes can make the call more efficient by removing any unneeded nested fields that are included by default. If any of the requests fail, the list of results will be nil and the error returned will represent the failed request.

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
    types := []openapiclient.FindIPReservationsTypesParameterInner{openapiclient.findIPReservations_types_parameter_inner("global_ipv4")} // []FindIPReservationsTypesParameterInner | Filter project IP reservations by reservation type (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPAddressesApi.FindIPReservations(context.Background(), id).Types(types).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindIPReservations`: IPReservationList
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.FindIPReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIPReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **types** | [**[]FindIPReservationsTypesParameterInner**](FindIPReservationsTypesParameterInner.md) | Filter project IP reservations by reservation type | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 250]

### Return type

[**IPReservationList**](IPReservationList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestIPReservation

> RequestIPReservation201Response RequestIPReservation(ctx, id).RequestIPReservationRequest(requestIPReservationRequest).Include(include).Exclude(exclude).Execute()

Requesting IP reservations



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
    requestIPReservationRequest := openapiclient.requestIPReservation_request{IPReservationRequestInput: openapiclient.NewIPReservationRequestInput(int32(123), "Type_example")} // RequestIPReservationRequest | IP Reservation Request to create
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPAddressesApi.RequestIPReservation(context.Background(), id).RequestIPReservationRequest(requestIPReservationRequest).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.RequestIPReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestIPReservation`: RequestIPReservation201Response
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.RequestIPReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestIPReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestIPReservationRequest** | [**RequestIPReservationRequest**](RequestIPReservationRequest.md) | IP Reservation Request to create | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**RequestIPReservation201Response**](RequestIPReservation201Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPAddress

> FindIPAddressById200Response UpdateIPAddress(ctx, id).Include(include).Exclude(exclude).IPAssignmentUpdateInput(iPAssignmentUpdateInput).Execute()

Update an ip address



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | IP Address UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    iPAssignmentUpdateInput := *openapiclient.NewIPAssignmentUpdateInput() // IPAssignmentUpdateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPAddressesApi.UpdateIPAddress(context.Background(), id).Include(include).Exclude(exclude).IPAssignmentUpdateInput(iPAssignmentUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.UpdateIPAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIPAddress`: FindIPAddressById200Response
    fmt.Fprintf(os.Stdout, "Response from `IPAddressesApi.UpdateIPAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | IP Address UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **iPAssignmentUpdateInput** | [**IPAssignmentUpdateInput**](IPAssignmentUpdateInput.md) |  | 

### Return type

[**FindIPAddressById200Response**](FindIPAddressById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


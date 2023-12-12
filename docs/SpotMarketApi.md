# \SpotMarketApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpotMarketRequest**](SpotMarketApi.md#CreateSpotMarketRequest) | **Post** /projects/{id}/spot-market-requests | Create a spot market request
[**DeleteSpotMarketRequest**](SpotMarketApi.md#DeleteSpotMarketRequest) | **Delete** /spot-market-requests/{id} | Delete the spot market request
[**FindMetroSpotMarketPrices**](SpotMarketApi.md#FindMetroSpotMarketPrices) | **Get** /market/spot/prices/metros | Get current spot market prices for metros
[**FindSpotMarketPrices**](SpotMarketApi.md#FindSpotMarketPrices) | **Get** /market/spot/prices | Get current spot market prices
[**FindSpotMarketPricesHistory**](SpotMarketApi.md#FindSpotMarketPricesHistory) | **Get** /market/spot/prices/history | Get spot market prices for a given period of time
[**FindSpotMarketRequestById**](SpotMarketApi.md#FindSpotMarketRequestById) | **Get** /spot-market-requests/{id} | Retrieve a spot market request
[**ListSpotMarketRequests**](SpotMarketApi.md#ListSpotMarketRequests) | **Get** /projects/{id}/spot-market-requests | List spot market requests



## CreateSpotMarketRequest

> SpotMarketRequest CreateSpotMarketRequest(ctx, id).SpotMarketRequestCreateInput(spotMarketRequestCreateInput).Execute()

Create a spot market request



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
    spotMarketRequestCreateInput := *openapiclient.NewSpotMarketRequestCreateInput() // SpotMarketRequestCreateInput | Spot Market Request to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpotMarketApi.CreateSpotMarketRequest(context.Background(), id).SpotMarketRequestCreateInput(spotMarketRequestCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.CreateSpotMarketRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpotMarketRequest`: SpotMarketRequest
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketApi.CreateSpotMarketRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpotMarketRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spotMarketRequestCreateInput** | [**SpotMarketRequestCreateInput**](SpotMarketRequestCreateInput.md) | Spot Market Request to create | 

### Return type

[**SpotMarketRequest**](SpotMarketRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpotMarketRequest

> DeleteSpotMarketRequest(ctx, id).ForceTermination(forceTermination).Execute()

Delete the spot market request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SpotMarketRequest UUID
    forceTermination := true // bool | Terminate associated spot instances (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpotMarketApi.DeleteSpotMarketRequest(context.Background(), id).ForceTermination(forceTermination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.DeleteSpotMarketRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SpotMarketRequest UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpotMarketRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceTermination** | **bool** | Terminate associated spot instances | 

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


## FindMetroSpotMarketPrices

> SpotMarketPricesPerMetroList FindMetroSpotMarketPrices(ctx).Metro(metro).Plan(plan).Execute()

Get current spot market prices for metros



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
    metro := "metro_example" // string | Metro to filter spot market prices (optional)
    plan := "plan_example" // string | Plan to filter spot market prices (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpotMarketApi.FindMetroSpotMarketPrices(context.Background()).Metro(metro).Plan(plan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.FindMetroSpotMarketPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetroSpotMarketPrices`: SpotMarketPricesPerMetroList
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketApi.FindMetroSpotMarketPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindMetroSpotMarketPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metro** | **string** | Metro to filter spot market prices | 
 **plan** | **string** | Plan to filter spot market prices | 

### Return type

[**SpotMarketPricesPerMetroList**](SpotMarketPricesPerMetroList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSpotMarketPrices

> SpotMarketPricesList FindSpotMarketPrices(ctx).Facility(facility).Plan(plan).Execute()

Get current spot market prices



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
    facility := "facility_example" // string | Facility to check spot market prices (optional)
    plan := "plan_example" // string | Plan to check spot market prices (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpotMarketApi.FindSpotMarketPrices(context.Background()).Facility(facility).Plan(plan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.FindSpotMarketPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSpotMarketPrices`: SpotMarketPricesList
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketApi.FindSpotMarketPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSpotMarketPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | **string** | Facility to check spot market prices | 
 **plan** | **string** | Plan to check spot market prices | 

### Return type

[**SpotMarketPricesList**](SpotMarketPricesList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSpotMarketPricesHistory

> SpotPricesHistoryReport FindSpotMarketPricesHistory(ctx).Facility(facility).Plan(plan).From(from).Until(until).Metro(metro).Execute()

Get spot market prices for a given period of time



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
    facility := "facility_example" // string | Facility to check spot market prices
    plan := "plan_example" // string | Plan to check spot market prices
    from := "from_example" // string | Timestamp from range
    until := "until_example" // string | Timestamp to range
    metro := "metro_example" // string | Metro to check spot market price history (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpotMarketApi.FindSpotMarketPricesHistory(context.Background()).Facility(facility).Plan(plan).From(from).Until(until).Metro(metro).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.FindSpotMarketPricesHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSpotMarketPricesHistory`: SpotPricesHistoryReport
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketApi.FindSpotMarketPricesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSpotMarketPricesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | **string** | Facility to check spot market prices | 
 **plan** | **string** | Plan to check spot market prices | 
 **from** | **string** | Timestamp from range | 
 **until** | **string** | Timestamp to range | 
 **metro** | **string** | Metro to check spot market price history | 

### Return type

[**SpotPricesHistoryReport**](SpotPricesHistoryReport.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSpotMarketRequestById

> SpotMarketRequest FindSpotMarketRequestById(ctx, id).Include(include).Execute()

Retrieve a spot market request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SpotMarketRequest UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpotMarketApi.FindSpotMarketRequestById(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.FindSpotMarketRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSpotMarketRequestById`: SpotMarketRequest
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketApi.FindSpotMarketRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SpotMarketRequest UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSpotMarketRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SpotMarketRequest**](SpotMarketRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpotMarketRequests

> SpotMarketRequestList ListSpotMarketRequests(ctx, id).Execute()

List spot market requests



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpotMarketApi.ListSpotMarketRequests(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpotMarketApi.ListSpotMarketRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpotMarketRequests`: SpotMarketRequestList
    fmt.Fprintf(os.Stdout, "Response from `SpotMarketApi.ListSpotMarketRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpotMarketRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpotMarketRequestList**](SpotMarketRequestList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


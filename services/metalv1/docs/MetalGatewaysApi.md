# \MetalGatewaysApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpDynamicNeighbor**](MetalGatewaysApi.md#CreateBgpDynamicNeighbor) | **Post** /metal-gateways/{id}/bgp-dynamic-neighbors | Create a VRF BGP Dynamic Neighbor range
[**CreateMetalGateway**](MetalGatewaysApi.md#CreateMetalGateway) | **Post** /projects/{project_id}/metal-gateways | Create a metal gateway
[**CreateMetalGatewayElasticIp**](MetalGatewaysApi.md#CreateMetalGatewayElasticIp) | **Post** /metal-gateways/{id}/ips | Create a Metal Gateway Elastic IP
[**DeleteMetalGateway**](MetalGatewaysApi.md#DeleteMetalGateway) | **Delete** /metal-gateways/{id} | Deletes the metal gateway
[**FindMetalGatewayById**](MetalGatewaysApi.md#FindMetalGatewayById) | **Get** /metal-gateways/{id} | Returns the metal gateway
[**FindMetalGatewaysByProject**](MetalGatewaysApi.md#FindMetalGatewaysByProject) | **Get** /projects/{project_id}/metal-gateways | Returns all metal gateways for a project
[**GetBgpDynamicNeighbors**](MetalGatewaysApi.md#GetBgpDynamicNeighbors) | **Get** /metal-gateways/{id}/bgp-dynamic-neighbors | List BGP Dynamic Neighbors
[**GetMetalGatewayElasticIps**](MetalGatewaysApi.md#GetMetalGatewayElasticIps) | **Get** /metal-gateways/{id}/ips | List Metal Gateway Elastic IPs



## CreateBgpDynamicNeighbor

> BgpDynamicNeighbor CreateBgpDynamicNeighbor(ctx, id).BgpDynamicNeighborCreateInput(bgpDynamicNeighborCreateInput).Include(include).Exclude(exclude).Execute()

Create a VRF BGP Dynamic Neighbor range



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metal Gateway UUID
    bgpDynamicNeighborCreateInput := *openapiclient.NewBgpDynamicNeighborCreateInput("192.168.1.0/25", int64(12345)) // BgpDynamicNeighborCreateInput | 
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.CreateBgpDynamicNeighbor(context.Background(), id).BgpDynamicNeighborCreateInput(bgpDynamicNeighborCreateInput).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.CreateBgpDynamicNeighbor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBgpDynamicNeighbor`: BgpDynamicNeighbor
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.CreateBgpDynamicNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpDynamicNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bgpDynamicNeighborCreateInput** | [**BgpDynamicNeighborCreateInput**](BgpDynamicNeighborCreateInput.md) |  | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**BgpDynamicNeighbor**](BgpDynamicNeighbor.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMetalGateway

> FindMetalGatewayById200Response CreateMetalGateway(ctx, projectId).CreateMetalGatewayRequest(createMetalGatewayRequest).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Create a metal gateway



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    createMetalGatewayRequest := openapiclient.createMetalGateway_request{MetalGatewayCreateInput: openapiclient.NewMetalGatewayCreateInput("VirtualNetworkId_example")} // CreateMetalGatewayRequest | Metal Gateway to create
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.CreateMetalGateway(context.Background(), projectId).CreateMetalGatewayRequest(createMetalGatewayRequest).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.CreateMetalGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetalGateway`: FindMetalGatewayById200Response
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.CreateMetalGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetalGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMetalGatewayRequest** | [**CreateMetalGatewayRequest**](CreateMetalGatewayRequest.md) | Metal Gateway to create | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**FindMetalGatewayById200Response**](FindMetalGatewayById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMetalGatewayElasticIp

> IPAssignment CreateMetalGatewayElasticIp(ctx, id).MetalGatewayElasticIpCreateInput(metalGatewayElasticIpCreateInput).Include(include).Exclude(exclude).Execute()

Create a Metal Gateway Elastic IP



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metal Gateway UUID
    metalGatewayElasticIpCreateInput := *openapiclient.NewMetalGatewayElasticIpCreateInput("147.75.234.8/31", "192.168.12.13") // MetalGatewayElasticIpCreateInput | 
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.CreateMetalGatewayElasticIp(context.Background(), id).MetalGatewayElasticIpCreateInput(metalGatewayElasticIpCreateInput).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.CreateMetalGatewayElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetalGatewayElasticIp`: IPAssignment
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.CreateMetalGatewayElasticIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetalGatewayElasticIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metalGatewayElasticIpCreateInput** | [**MetalGatewayElasticIpCreateInput**](MetalGatewayElasticIpCreateInput.md) |  | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**IPAssignment**](IPAssignment.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetalGateway

> FindMetalGatewayById200Response DeleteMetalGateway(ctx, id).Include(include).Exclude(exclude).Execute()

Deletes the metal gateway



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metal Gateway UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.DeleteMetalGateway(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.DeleteMetalGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMetalGateway`: FindMetalGatewayById200Response
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.DeleteMetalGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetalGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindMetalGatewayById200Response**](FindMetalGatewayById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindMetalGatewayById

> FindMetalGatewayById200Response FindMetalGatewayById(ctx, id).Include(include).Exclude(exclude).Execute()

Returns the metal gateway



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metal Gateway UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.FindMetalGatewayById(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.FindMetalGatewayById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetalGatewayById`: FindMetalGatewayById200Response
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.FindMetalGatewayById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMetalGatewayByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**FindMetalGatewayById200Response**](FindMetalGatewayById200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindMetalGatewaysByProject

> MetalGatewayList FindMetalGatewaysByProject(ctx, projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()

Returns all metal gateways for a project


FindMetalGatewaysByProject is a paginated API operation. You can specify page number and per page parameters with `Execute`, or to fetch and return all pages of results as a single slice you can call `ExecuteWithPagination()`. `ExecuteWithPagination()`, while convenient, can significantly increase the overhead of API calls in terms of time, network utilization, and memory. Excludes can make the call more efficient by removing any unneeded nested fields that are included by default. If any of the requests fail, the list of results will be nil and the error returned will represent the failed request.

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.FindMetalGatewaysByProject(context.Background(), projectId).Include(include).Exclude(exclude).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.FindMetalGatewaysByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetalGatewaysByProject`: MetalGatewayList
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.FindMetalGatewaysByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMetalGatewaysByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**MetalGatewayList**](MetalGatewayList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpDynamicNeighbors

> BgpDynamicNeighborList GetBgpDynamicNeighbors(ctx, id).Include(include).Exclude(exclude).Execute()

List BGP Dynamic Neighbors



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metal Gateway UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.GetBgpDynamicNeighbors(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.GetBgpDynamicNeighbors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBgpDynamicNeighbors`: BgpDynamicNeighborList
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.GetBgpDynamicNeighbors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpDynamicNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**BgpDynamicNeighborList**](BgpDynamicNeighborList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetalGatewayElasticIps

> IPAssignmentList GetMetalGatewayElasticIps(ctx, id).Include(include).Exclude(exclude).Execute()

List Metal Gateway Elastic IPs



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Metal Gateway UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetalGatewaysApi.GetMetalGatewayElasticIps(context.Background(), id).Include(include).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.GetMetalGatewayElasticIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetalGatewayElasticIps`: IPAssignmentList
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.GetMetalGatewayElasticIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetalGatewayElasticIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 

### Return type

[**IPAssignmentList**](IPAssignmentList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


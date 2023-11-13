# \TransferRequestsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptTransferRequest**](TransferRequestsApi.md#AcceptTransferRequest) | **Put** /transfers/{id} | Accept a transfer request
[**DeclineTransferRequest**](TransferRequestsApi.md#DeclineTransferRequest) | **Delete** /transfers/{id} | Decline a transfer request
[**FindTransferRequestById**](TransferRequestsApi.md#FindTransferRequestById) | **Get** /transfers/{id} | View a transfer request



## AcceptTransferRequest

> AcceptTransferRequest(ctx, id).Execute()

Accept a transfer request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transfer request UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransferRequestsApi.AcceptTransferRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferRequestsApi.AcceptTransferRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transfer request UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptTransferRequestRequest struct via the builder pattern


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


## DeclineTransferRequest

> DeclineTransferRequest(ctx, id).Execute()

Decline a transfer request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transfer request UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransferRequestsApi.DeclineTransferRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferRequestsApi.DeclineTransferRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transfer request UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineTransferRequestRequest struct via the builder pattern


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


## FindTransferRequestById

> TransferRequest FindTransferRequestById(ctx, id).Include(include).Execute()

View a transfer request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transfer request UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransferRequestsApi.FindTransferRequestById(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferRequestsApi.FindTransferRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindTransferRequestById`: TransferRequest
    fmt.Fprintf(os.Stdout, "Response from `TransferRequestsApi.FindTransferRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transfer request UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTransferRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**TransferRequest**](TransferRequest.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


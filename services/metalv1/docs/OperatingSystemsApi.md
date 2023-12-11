# \OperatingSystemsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindOperatingSystemVersion**](OperatingSystemsApi.md#FindOperatingSystemVersion) | **Get** /operating-system-versions | Retrieve all operating system versions
[**FindOperatingSystems**](OperatingSystemsApi.md#FindOperatingSystems) | **Get** /operating-systems | Retrieve all operating systems



## FindOperatingSystemVersion

> OperatingSystemList FindOperatingSystemVersion(ctx).Execute()

Retrieve all operating system versions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatingSystemsApi.FindOperatingSystemVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatingSystemsApi.FindOperatingSystemVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOperatingSystemVersion`: OperatingSystemList
    fmt.Fprintf(os.Stdout, "Response from `OperatingSystemsApi.FindOperatingSystemVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindOperatingSystemVersionRequest struct via the builder pattern


### Return type

[**OperatingSystemList**](OperatingSystemList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOperatingSystems

> OperatingSystemList FindOperatingSystems(ctx).Execute()

Retrieve all operating systems



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatingSystemsApi.FindOperatingSystems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatingSystemsApi.FindOperatingSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOperatingSystems`: OperatingSystemList
    fmt.Fprintf(os.Stdout, "Response from `OperatingSystemsApi.FindOperatingSystems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindOperatingSystemsRequest struct via the builder pattern


### Return type

[**OperatingSystemList**](OperatingSystemList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


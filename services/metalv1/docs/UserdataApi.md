# \UserdataApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateUserdata**](UserdataApi.md#ValidateUserdata) | **Post** /userdata/validate | Validate user data



## ValidateUserdata

> ValidateUserdata(ctx).Userdata(userdata).Execute()

Validate user data



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
    userdata := "userdata_example" // string | Userdata to validate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserdataApi.ValidateUserdata(context.Background()).Userdata(userdata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserdataApi.ValidateUserdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateUserdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userdata** | **string** | Userdata to validate | 

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


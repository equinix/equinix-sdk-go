# \TwoFactorAuthApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableTfaApp**](TwoFactorAuthApi.md#DisableTfaApp) | **Delete** /user/otp/app | Disable two factor authentication
[**DisableTfaSms**](TwoFactorAuthApi.md#DisableTfaSms) | **Delete** /user/otp/sms | Disable two factor authentication
[**EnableTfaApp**](TwoFactorAuthApi.md#EnableTfaApp) | **Post** /user/otp/app | Enable two factor auth using app
[**EnableTfaSms**](TwoFactorAuthApi.md#EnableTfaSms) | **Post** /user/otp/sms | Enable two factor auth using sms



## DisableTfaApp

> DisableTfaApp(ctx).XOtpToken(xOtpToken).Execute()

Disable two factor authentication



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
    xOtpToken := "xOtpToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TwoFactorAuthApi.DisableTfaApp(context.Background()).XOtpToken(xOtpToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthApi.DisableTfaApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableTfaAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOtpToken** | **string** |  | 

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


## DisableTfaSms

> DisableTfaSms(ctx).XOtpToken(xOtpToken).Execute()

Disable two factor authentication



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
    xOtpToken := "xOtpToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TwoFactorAuthApi.DisableTfaSms(context.Background()).XOtpToken(xOtpToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthApi.DisableTfaSms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableTfaSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOtpToken** | **string** |  | 

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


## EnableTfaApp

> EnableTfaApp(ctx).XOtpToken(xOtpToken).Execute()

Enable two factor auth using app



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
    xOtpToken := "xOtpToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TwoFactorAuthApi.EnableTfaApp(context.Background()).XOtpToken(xOtpToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthApi.EnableTfaApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableTfaAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOtpToken** | **string** |  | 

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


## EnableTfaSms

> EnableTfaSms(ctx).XOtpToken(xOtpToken).Execute()

Enable two factor auth using sms



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
    xOtpToken := "xOtpToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TwoFactorAuthApi.EnableTfaSms(context.Background()).XOtpToken(xOtpToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthApi.EnableTfaSms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableTfaSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOtpToken** | **string** |  | 

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


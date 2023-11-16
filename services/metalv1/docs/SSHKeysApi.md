# \SSHKeysApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectSSHKey**](SSHKeysApi.md#CreateProjectSSHKey) | **Post** /projects/{id}/ssh-keys | Create a ssh key for the given project
[**CreateSSHKey**](SSHKeysApi.md#CreateSSHKey) | **Post** /ssh-keys | Create a ssh key for the current user
[**DeleteSSHKey**](SSHKeysApi.md#DeleteSSHKey) | **Delete** /ssh-keys/{id} | Delete the ssh key
[**FindDeviceSSHKeys**](SSHKeysApi.md#FindDeviceSSHKeys) | **Get** /devices/{id}/ssh-keys | Retrieve a device&#39;s ssh keys
[**FindProjectSSHKeys**](SSHKeysApi.md#FindProjectSSHKeys) | **Get** /projects/{id}/ssh-keys | Retrieve a project&#39;s ssh keys
[**FindSSHKeyById**](SSHKeysApi.md#FindSSHKeyById) | **Get** /ssh-keys/{id} | Retrieve a ssh key
[**FindSSHKeys**](SSHKeysApi.md#FindSSHKeys) | **Get** /ssh-keys | Retrieve all ssh keys
[**UpdateSSHKey**](SSHKeysApi.md#UpdateSSHKey) | **Put** /ssh-keys/{id} | Update the ssh key



## CreateProjectSSHKey

> SSHKey CreateProjectSSHKey(ctx, id).SSHKeyCreateInput(sSHKeyCreateInput).Include(include).Execute()

Create a ssh key for the given project



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
    sSHKeyCreateInput := *openapiclient.NewSSHKeyCreateInput() // SSHKeyCreateInput | ssh key to create
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.CreateProjectSSHKey(context.Background(), id).SSHKeyCreateInput(sSHKeyCreateInput).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.CreateProjectSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.CreateProjectSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHKeyCreateInput** | [**SSHKeyCreateInput**](SSHKeyCreateInput.md) | ssh key to create | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHKey

> SSHKey CreateSSHKey(ctx).SSHKeyCreateInput(sSHKeyCreateInput).Include(include).Execute()

Create a ssh key for the current user



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
    sSHKeyCreateInput := *openapiclient.NewSSHKeyCreateInput() // SSHKeyCreateInput | ssh key to create
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.CreateSSHKey(context.Background()).SSHKeyCreateInput(sSHKeyCreateInput).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.CreateSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.CreateSSHKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSHKeyCreateInput** | [**SSHKeyCreateInput**](SSHKeyCreateInput.md) | ssh key to create | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHKey

> DeleteSSHKey(ctx, id).Execute()

Delete the ssh key



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ssh key UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSHKeysApi.DeleteSSHKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.DeleteSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ssh key UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHKeyRequest struct via the builder pattern


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


## FindDeviceSSHKeys

> SSHKeyList FindDeviceSSHKeys(ctx, id).SearchString(searchString).Include(include).Execute()

Retrieve a device's ssh keys



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
    searchString := "searchString_example" // string | Search by key, label, or fingerprint (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.FindDeviceSSHKeys(context.Background(), id).SearchString(searchString).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.FindDeviceSSHKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDeviceSSHKeys`: SSHKeyList
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.FindDeviceSSHKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDeviceSSHKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchString** | **string** | Search by key, label, or fingerprint | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKeyList**](SSHKeyList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindProjectSSHKeys

> SSHKeyList FindProjectSSHKeys(ctx, id).Query(query).Include(include).Execute()

Retrieve a project's ssh keys



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
    query := "query_example" // string | Search by key, label, or fingerprint (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.FindProjectSSHKeys(context.Background(), id).Query(query).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.FindProjectSSHKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindProjectSSHKeys`: SSHKeyList
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.FindProjectSSHKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindProjectSSHKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Search by key, label, or fingerprint | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKeyList**](SSHKeyList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSSHKeyById

> SSHKey FindSSHKeyById(ctx, id).Include(include).Execute()

Retrieve a ssh key



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SSH Key UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.FindSSHKeyById(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.FindSSHKeyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSSHKeyById`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.FindSSHKeyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SSH Key UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSSHKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSSHKeys

> SSHKeyList FindSSHKeys(ctx).Search(search).Include(include).Execute()

Retrieve all ssh keys



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
    search := "search_example" // string | Search by key, label, or fingerprint (optional)
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.FindSSHKeys(context.Background()).Search(search).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.FindSSHKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSSHKeys`: SSHKeyList
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.FindSSHKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSSHKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search by key, label, or fingerprint | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKeyList**](SSHKeyList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHKey

> SSHKey UpdateSSHKey(ctx, id).SSHKeyInput(sSHKeyInput).Include(include).Execute()

Update the ssh key



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SSH Key UUID
    sSHKeyInput := *openapiclient.NewSSHKeyInput() // SSHKeyInput | ssh key to update
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.UpdateSSHKey(context.Background(), id).SSHKeyInput(sSHKeyInput).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.UpdateSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.UpdateSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SSH Key UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sSHKeyInput** | [**SSHKeyInput**](SSHKeyInput.md) | ssh key to update | 
 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


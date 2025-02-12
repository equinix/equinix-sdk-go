# \RolesApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationRole**](RolesApi.md#GetOrganizationRole) | **Get** /organizations/{id}/roles/{role_id} | Get details about a specific role
[**ListOrganizationRoles**](RolesApi.md#ListOrganizationRoles) | **Get** /organizations/{id}/roles | List available roles



## GetOrganizationRole

> Role GetOrganizationRole(ctx, id, roleId).Execute()

Get details about a specific role



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationRole(context.Background(), id, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationRoles

> RoleList ListOrganizationRoles(ctx, id).Execute()

List available roles



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
    resp, r, err := apiClient.RolesApi.ListOrganizationRoles(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListOrganizationRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationRoles`: RoleList
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListOrganizationRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Organization UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleList**](RoleList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


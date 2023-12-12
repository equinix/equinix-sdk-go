# \InvitationsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitation**](InvitationsApi.md#AcceptInvitation) | **Put** /invitations/{id} | Accept an invitation
[**DeclineInvitation**](InvitationsApi.md#DeclineInvitation) | **Delete** /invitations/{id} | Decline an invitation
[**FindInvitationById**](InvitationsApi.md#FindInvitationById) | **Get** /invitations/{id} | View an invitation



## AcceptInvitation

> Membership AcceptInvitation(ctx, id).Include(include).Execute()

Accept an invitation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invitation UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.AcceptInvitation(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.AcceptInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptInvitation`: Membership
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.AcceptInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invitation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**Membership**](Membership.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineInvitation

> DeclineInvitation(ctx, id).Execute()

Decline an invitation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invitation UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InvitationsApi.DeclineInvitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.DeclineInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invitation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineInvitationRequest struct via the builder pattern


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


## FindInvitationById

> Invitation FindInvitationById(ctx, id).Include(include).Execute()

View an invitation



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invitation UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.FindInvitationById(context.Background(), id).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.FindInvitationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindInvitationById`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.FindInvitationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invitation UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInvitationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


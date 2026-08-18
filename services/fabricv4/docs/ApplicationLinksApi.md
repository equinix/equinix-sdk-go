# \ApplicationLinksApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachAppDomainToAppLink**](ApplicationLinksApi.md#AttachAppDomainToAppLink) | **Put** /fabric/v4/appLinks/{appLinkId}/appDomains/{appDomainId} | Attach App Domain to App Link
[**AttachAppServiceToAppLink**](ApplicationLinksApi.md#AttachAppServiceToAppLink) | **Put** /fabric/v4/appLinks/{appLinkId}/appServices/{appServiceId} | Attach App Service to App Link
[**CreateAppLink**](ApplicationLinksApi.md#CreateAppLink) | **Post** /fabric/v4/appLinks | Create App Link
[**DeleteAppLinkByUuid**](ApplicationLinksApi.md#DeleteAppLinkByUuid) | **Delete** /fabric/v4/appLinks/{appLinkId} | Delete App Link
[**DetachAppDomainFromAppLink**](ApplicationLinksApi.md#DetachAppDomainFromAppLink) | **Delete** /fabric/v4/appLinks/{appLinkId}/appDomains/{appDomainId} | Detach App Domain from App Link
[**DetachAppServiceFromAppLink**](ApplicationLinksApi.md#DetachAppServiceFromAppLink) | **Delete** /fabric/v4/appLinks/{appLinkId}/appServices/{appServiceId} | Detach App Service from App Link
[**GetAppLinkByUuid**](ApplicationLinksApi.md#GetAppLinkByUuid) | **Get** /fabric/v4/appLinks/{appLinkId} | Get App Link
[**GetAttachedAppDomainByUuid**](ApplicationLinksApi.md#GetAttachedAppDomainByUuid) | **Get** /fabric/v4/appLinks/{appLinkId}/appDomains/{appDomainId} | Get attached App Domain for App Link
[**GetAttachedAppDomainsByAppLinkId**](ApplicationLinksApi.md#GetAttachedAppDomainsByAppLinkId) | **Get** /fabric/v4/appLinks/{appLinkId}/appDomains | Get attached App Domains for App Link
[**GetAttachedAppServiceByUuid**](ApplicationLinksApi.md#GetAttachedAppServiceByUuid) | **Get** /fabric/v4/appLinks/{appLinkId}/appServices/{appServiceId} | Get attached App Service for App Link
[**GetAttachedAppServicesByAppLinkId**](ApplicationLinksApi.md#GetAttachedAppServicesByAppLinkId) | **Get** /fabric/v4/appLinks/{appLinkId}/appServices | Get attached App Services for App Link
[**SearchAppLinks**](ApplicationLinksApi.md#SearchAppLinks) | **Post** /fabric/v4/appLinks/search | Search App Links
[**SearchAttachedAppDomains**](ApplicationLinksApi.md#SearchAttachedAppDomains) | **Post** /fabric/v4/appLinks/{appLinkId}/appDomains/search | Search attached App Domain to App Link
[**SearchAttachedAppServices**](ApplicationLinksApi.md#SearchAttachedAppServices) | **Post** /fabric/v4/appLinks/{appLinkId}/appServices/search | Search attached App Service to App Link
[**UpdateAppLinkByUuid**](ApplicationLinksApi.md#UpdateAppLinkByUuid) | **Patch** /fabric/v4/appLinks/{appLinkId} | Update App Link
[**UpdateAppServiceAttachmentToAppLink**](ApplicationLinksApi.md#UpdateAppServiceAttachmentToAppLink) | **Patch** /fabric/v4/appLinks/{appLinkId}/appServices/{appServiceId} | Update App Service attachment to App Link



## AttachAppDomainToAppLink

> AppLinkAppDomainAttachment AttachAppDomainToAppLink(ctx, appLinkId, appDomainId).Execute()

Attach App Domain to App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.AttachAppDomainToAppLink(context.Background(), appLinkId, appDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.AttachAppDomainToAppLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachAppDomainToAppLink`: AppLinkAppDomainAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.AttachAppDomainToAppLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachAppDomainToAppLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppLinkAppDomainAttachment**](AppLinkAppDomainAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachAppServiceToAppLink

> AppLinkAppServiceAttachment AttachAppServiceToAppLink(ctx, appLinkId, appServiceId).AppLinkAttachServiceRequest(appLinkAttachServiceRequest).Execute()

Attach App Service to App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID
	appLinkAttachServiceRequest := *openapiclient.NewAppLinkAttachServiceRequest("GLOBAL", "10.10.1.110") // AppLinkAttachServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.AttachAppServiceToAppLink(context.Background(), appLinkId, appServiceId).AppLinkAttachServiceRequest(appLinkAttachServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.AttachAppServiceToAppLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachAppServiceToAppLink`: AppLinkAppServiceAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.AttachAppServiceToAppLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachAppServiceToAppLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appLinkAttachServiceRequest** | [**AppLinkAttachServiceRequest**](AppLinkAttachServiceRequest.md) |  | 

### Return type

[**AppLinkAppServiceAttachment**](AppLinkAppServiceAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppLink

> AppLink CreateAppLink(ctx).AppLinkPostRequest(appLinkPostRequest).Execute()

Create App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkPostRequest := *openapiclient.NewAppLinkPostRequest(openapiclient.AppLinkType("APP_LINK"), "Atlassian App Link", *openapiclient.NewAppLinkPostRequestRouter("36204584-2ae1-44f6-980f-08ecd555d9c6"), int32(1000), *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // AppLinkPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.CreateAppLink(context.Background()).AppLinkPostRequest(appLinkPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.CreateAppLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppLink`: AppLink
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.CreateAppLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appLinkPostRequest** | [**AppLinkPostRequest**](AppLinkPostRequest.md) |  | 

### Return type

[**AppLink**](AppLink.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppLinkByUuid

> AppLink DeleteAppLinkByUuid(ctx, appLinkId).Execute()

Delete App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.DeleteAppLinkByUuid(context.Background(), appLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.DeleteAppLinkByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppLinkByUuid`: AppLink
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.DeleteAppLinkByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppLinkByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppLink**](AppLink.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachAppDomainFromAppLink

> AppLinkAppDomainAttachment DetachAppDomainFromAppLink(ctx, appLinkId, appDomainId).Execute()

Detach App Domain from App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.DetachAppDomainFromAppLink(context.Background(), appLinkId, appDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.DetachAppDomainFromAppLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachAppDomainFromAppLink`: AppLinkAppDomainAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.DetachAppDomainFromAppLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachAppDomainFromAppLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppLinkAppDomainAttachment**](AppLinkAppDomainAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachAppServiceFromAppLink

> AppLinkAppServiceAttachment DetachAppServiceFromAppLink(ctx, appLinkId, appServiceId).Execute()

Detach App Service from App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.DetachAppServiceFromAppLink(context.Background(), appLinkId, appServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.DetachAppServiceFromAppLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachAppServiceFromAppLink`: AppLinkAppServiceAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.DetachAppServiceFromAppLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachAppServiceFromAppLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppLinkAppServiceAttachment**](AppLinkAppServiceAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppLinkByUuid

> AppLink GetAppLinkByUuid(ctx, appLinkId).Execute()

Get App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.GetAppLinkByUuid(context.Background(), appLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.GetAppLinkByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppLinkByUuid`: AppLink
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.GetAppLinkByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppLinkByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppLink**](AppLink.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppDomainByUuid

> AppLinkAppDomainAttachment GetAttachedAppDomainByUuid(ctx, appLinkId, appDomainId).Execute()

Get attached App Domain for App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.GetAttachedAppDomainByUuid(context.Background(), appLinkId, appDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.GetAttachedAppDomainByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppDomainByUuid`: AppLinkAppDomainAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.GetAttachedAppDomainByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppDomainByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppLinkAppDomainAttachment**](AppLinkAppDomainAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppDomainsByAppLinkId

> AppLinkAttachedAppDomains GetAttachedAppDomainsByAppLinkId(ctx, appLinkId).Offset(offset).Limit(limit).AttachmentStatus(attachmentStatus).Order(order).Style(style).Execute()

Get attached App Domains for App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	attachmentStatus := []openapiclient.AppLinkAttachState{openapiclient.AppLinkAttachState("ATTACHING")} // []AppLinkAttachState | Filter attached App Domains by one or more attachment lifecycle states. (optional)
	order := openapiclient.AttachedAppDomainOrder("DESC") // AttachedAppDomainOrder | Sort order for attached App Domains. (optional) (default to "DESC")
	style := openapiclient.Style("MIN") // Style | Detail level of the response. (optional) (default to "MEDIUM")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.GetAttachedAppDomainsByAppLinkId(context.Background(), appLinkId).Offset(offset).Limit(limit).AttachmentStatus(attachmentStatus).Order(order).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.GetAttachedAppDomainsByAppLinkId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppDomainsByAppLinkId`: AppLinkAttachedAppDomains
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.GetAttachedAppDomainsByAppLinkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppDomainsByAppLinkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **attachmentStatus** | [**[]AppLinkAttachState**](AppLinkAttachState.md) | Filter attached App Domains by one or more attachment lifecycle states. | 
 **order** | [**AttachedAppDomainOrder**](AttachedAppDomainOrder.md) | Sort order for attached App Domains. | [default to &quot;DESC&quot;]
 **style** | [**Style**](Style.md) | Detail level of the response. | [default to &quot;MEDIUM&quot;]

### Return type

[**AppLinkAttachedAppDomains**](AppLinkAttachedAppDomains.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppServiceByUuid

> AppLinkAppServiceAttachment GetAttachedAppServiceByUuid(ctx, appLinkId, appServiceId).Execute()

Get attached App Service for App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.GetAttachedAppServiceByUuid(context.Background(), appLinkId, appServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.GetAttachedAppServiceByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppServiceByUuid`: AppLinkAppServiceAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.GetAttachedAppServiceByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppServiceByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppLinkAppServiceAttachment**](AppLinkAppServiceAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppServicesByAppLinkId

> AppLinkAttachedAppServices GetAttachedAppServicesByAppLinkId(ctx, appLinkId).Offset(offset).Limit(limit).AttachmentStatus(attachmentStatus).Order(order).Style(style).Execute()

Get attached App Services for App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	attachmentStatus := []openapiclient.AppLinkAttachState{openapiclient.AppLinkAttachState("ATTACHING")} // []AppLinkAttachState | Filter attached App Services by one or more attachment lifecycle states. (optional)
	order := openapiclient.AttachedAppServiceOrder("DESC") // AttachedAppServiceOrder | Sort order for attached App Services. (optional) (default to "DESC")
	style := openapiclient.Style("MIN") // Style | Detail level of the response. (optional) (default to "MEDIUM")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.GetAttachedAppServicesByAppLinkId(context.Background(), appLinkId).Offset(offset).Limit(limit).AttachmentStatus(attachmentStatus).Order(order).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.GetAttachedAppServicesByAppLinkId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppServicesByAppLinkId`: AppLinkAttachedAppServices
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.GetAttachedAppServicesByAppLinkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppServicesByAppLinkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **attachmentStatus** | [**[]AppLinkAttachState**](AppLinkAttachState.md) | Filter attached App Services by one or more attachment lifecycle states. | 
 **order** | [**AttachedAppServiceOrder**](AttachedAppServiceOrder.md) | Sort order for attached App Services. | [default to &quot;DESC&quot;]
 **style** | [**Style**](Style.md) | Detail level of the response. | [default to &quot;MEDIUM&quot;]

### Return type

[**AppLinkAttachedAppServices**](AppLinkAttachedAppServices.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppLinks

> AppLinkSearchResponse SearchAppLinks(ctx).AppLinkSearchRequest(appLinkSearchRequest).Execute()

Search App Links



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkSearchRequest := *openapiclient.NewAppLinkSearchRequest() // AppLinkSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.SearchAppLinks(context.Background()).AppLinkSearchRequest(appLinkSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.SearchAppLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAppLinks`: AppLinkSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.SearchAppLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appLinkSearchRequest** | [**AppLinkSearchRequest**](AppLinkSearchRequest.md) |  | 

### Return type

[**AppLinkSearchResponse**](AppLinkSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAttachedAppDomains

> AppLinkAttachDomainSearchResponse SearchAttachedAppDomains(ctx, appLinkId).AppLinkAttachDomainSearchRequest(appLinkAttachDomainSearchRequest).Execute()

Search attached App Domain to App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appLinkAttachDomainSearchRequest := *openapiclient.NewAppLinkAttachDomainSearchRequest() // AppLinkAttachDomainSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.SearchAttachedAppDomains(context.Background(), appLinkId).AppLinkAttachDomainSearchRequest(appLinkAttachDomainSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.SearchAttachedAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAttachedAppDomains`: AppLinkAttachDomainSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.SearchAttachedAppDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAttachedAppDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appLinkAttachDomainSearchRequest** | [**AppLinkAttachDomainSearchRequest**](AppLinkAttachDomainSearchRequest.md) |  | 

### Return type

[**AppLinkAttachDomainSearchResponse**](AppLinkAttachDomainSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAttachedAppServices

> AppLinkAttachServiceSearchResponse SearchAttachedAppServices(ctx, appLinkId).AppLinkAttachServiceSearchRequest(appLinkAttachServiceSearchRequest).Execute()

Search attached App Service to App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appLinkAttachServiceSearchRequest := *openapiclient.NewAppLinkAttachServiceSearchRequest() // AppLinkAttachServiceSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.SearchAttachedAppServices(context.Background(), appLinkId).AppLinkAttachServiceSearchRequest(appLinkAttachServiceSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.SearchAttachedAppServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAttachedAppServices`: AppLinkAttachServiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.SearchAttachedAppServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAttachedAppServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appLinkAttachServiceSearchRequest** | [**AppLinkAttachServiceSearchRequest**](AppLinkAttachServiceSearchRequest.md) |  | 

### Return type

[**AppLinkAttachServiceSearchResponse**](AppLinkAttachServiceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppLinkByUuid

> AppLink UpdateAppLinkByUuid(ctx, appLinkId).AppLinkChangeOperation(appLinkChangeOperation).Execute()

Update App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appLinkChangeOperation := []openapiclient.AppLinkChangeOperation{*openapiclient.NewAppLinkChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), openapiclient.AppLinkChangeOperation_path("/name"), map[string]interface{}(123))} // []AppLinkChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.UpdateAppLinkByUuid(context.Background(), appLinkId).AppLinkChangeOperation(appLinkChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.UpdateAppLinkByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppLinkByUuid`: AppLink
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.UpdateAppLinkByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppLinkByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appLinkChangeOperation** | [**[]AppLinkChangeOperation**](AppLinkChangeOperation.md) |  | 

### Return type

[**AppLink**](AppLink.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceAttachmentToAppLink

> AppLinkAppServiceAttachment UpdateAppServiceAttachmentToAppLink(ctx, appLinkId, appServiceId).AppLinkAppServiceAttachmentChangeOperation(appLinkAppServiceAttachmentChangeOperation).Execute()

Update App Service attachment to App Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	appLinkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Link UUID
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID
	appLinkAppServiceAttachmentChangeOperation := []openapiclient.AppLinkAppServiceAttachmentChangeOperation{*openapiclient.NewAppLinkAppServiceAttachmentChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), openapiclient.AppLinkAppServiceAttachmentChangeOperation_path("/destinationIp"), map[string]interface{}(123))} // []AppLinkAppServiceAttachmentChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationLinksApi.UpdateAppServiceAttachmentToAppLink(context.Background(), appLinkId, appServiceId).AppLinkAppServiceAttachmentChangeOperation(appLinkAppServiceAttachmentChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLinksApi.UpdateAppServiceAttachmentToAppLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceAttachmentToAppLink`: AppLinkAppServiceAttachment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationLinksApi.UpdateAppServiceAttachmentToAppLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLinkId** | **string** | App Link UUID | 
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceAttachmentToAppLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appLinkAppServiceAttachmentChangeOperation** | [**[]AppLinkAppServiceAttachmentChangeOperation**](AppLinkAppServiceAttachmentChangeOperation.md) |  | 

### Return type

[**AppLinkAppServiceAttachment**](AppLinkAppServiceAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


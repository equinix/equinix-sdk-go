# \ApplicationDomainsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppDomain**](ApplicationDomainsApi.md#CreateAppDomain) | **Post** /fabric/v4/appDomains | Create App Domain
[**DeleteAppDomainByUuid**](ApplicationDomainsApi.md#DeleteAppDomainByUuid) | **Delete** /fabric/v4/appDomains/{appDomainId} | Delete App Domain
[**GetAppDomainByUuid**](ApplicationDomainsApi.md#GetAppDomainByUuid) | **Get** /fabric/v4/appDomains/{appDomainId} | Get App Domain
[**GetAttachedAppLinksByAppDomainId**](ApplicationDomainsApi.md#GetAttachedAppLinksByAppDomainId) | **Get** /fabric/v4/appDomains/{appDomainId}/appLinks | Get attached App Links for App Domain
[**SearchAppDomains**](ApplicationDomainsApi.md#SearchAppDomains) | **Post** /fabric/v4/appDomains/search | Search App Domains
[**UpdateAppDomainByUuid**](ApplicationDomainsApi.md#UpdateAppDomainByUuid) | **Patch** /fabric/v4/appDomains/{appDomainId} | Update App Domain



## CreateAppDomain

> AppDomain CreateAppDomain(ctx).AppDomainPostRequest(appDomainPostRequest).Execute()

Create App Domain



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
	appDomainPostRequest := *openapiclient.NewAppDomainPostRequest(openapiclient.AppDomainType("APP_DOMAIN"), "atlassian.net", *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // AppDomainPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDomainsApi.CreateAppDomain(context.Background()).AppDomainPostRequest(appDomainPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDomainsApi.CreateAppDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppDomain`: AppDomain
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDomainsApi.CreateAppDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDomainPostRequest** | [**AppDomainPostRequest**](AppDomainPostRequest.md) |  | 

### Return type

[**AppDomain**](AppDomain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppDomainByUuid

> AppDomain DeleteAppDomainByUuid(ctx, appDomainId).Execute()

Delete App Domain



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
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDomainsApi.DeleteAppDomainByUuid(context.Background(), appDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDomainsApi.DeleteAppDomainByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppDomainByUuid`: AppDomain
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDomainsApi.DeleteAppDomainByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppDomainByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDomain**](AppDomain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDomainByUuid

> AppDomain GetAppDomainByUuid(ctx, appDomainId).Execute()

Get App Domain



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
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDomainsApi.GetAppDomainByUuid(context.Background(), appDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDomainsApi.GetAppDomainByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppDomainByUuid`: AppDomain
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDomainsApi.GetAppDomainByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDomainByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppDomain**](AppDomain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppLinksByAppDomainId

> AppDomainAttachedAppLinks GetAttachedAppLinksByAppDomainId(ctx, appDomainId).Offset(offset).Limit(limit).State(state).Order(order).Style(style).Execute()

Get attached App Links for App Domain



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
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	state := []openapiclient.AppLinkState{openapiclient.AppLinkState("PROVISIONING")} // []AppLinkState | Filter attached App Links by one or more lifecycle states. (optional)
	order := openapiclient.AttachedAppLinkOrder("DESC") // AttachedAppLinkOrder | Sort order for attached App Links. (optional) (default to "DESC")
	style := openapiclient.Style("MIN") // Style | Detail level of the response. (optional) (default to "MEDIUM")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDomainsApi.GetAttachedAppLinksByAppDomainId(context.Background(), appDomainId).Offset(offset).Limit(limit).State(state).Order(order).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDomainsApi.GetAttachedAppLinksByAppDomainId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppLinksByAppDomainId`: AppDomainAttachedAppLinks
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDomainsApi.GetAttachedAppLinksByAppDomainId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppLinksByAppDomainIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **state** | [**[]AppLinkState**](AppLinkState.md) | Filter attached App Links by one or more lifecycle states. | 
 **order** | [**AttachedAppLinkOrder**](AttachedAppLinkOrder.md) | Sort order for attached App Links. | [default to &quot;DESC&quot;]
 **style** | [**Style**](Style.md) | Detail level of the response. | [default to &quot;MEDIUM&quot;]

### Return type

[**AppDomainAttachedAppLinks**](AppDomainAttachedAppLinks.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppDomains

> AppDomainSearchResponse SearchAppDomains(ctx).AppDomainSearchRequest(appDomainSearchRequest).Execute()

Search App Domains



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
	appDomainSearchRequest := *openapiclient.NewAppDomainSearchRequest() // AppDomainSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDomainsApi.SearchAppDomains(context.Background()).AppDomainSearchRequest(appDomainSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDomainsApi.SearchAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAppDomains`: AppDomainSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDomainsApi.SearchAppDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDomainSearchRequest** | [**AppDomainSearchRequest**](AppDomainSearchRequest.md) |  | 

### Return type

[**AppDomainSearchResponse**](AppDomainSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppDomainByUuid

> AppDomain UpdateAppDomainByUuid(ctx, appDomainId).AppDomainChangeOperation(appDomainChangeOperation).Execute()

Update App Domain



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
	appDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Domain UUID
	appDomainChangeOperation := []openapiclient.AppDomainChangeOperation{*openapiclient.NewAppDomainChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), openapiclient.AppDomainChangeOperation_path("/description"), map[string]interface{}(123))} // []AppDomainChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationDomainsApi.UpdateAppDomainByUuid(context.Background(), appDomainId).AppDomainChangeOperation(appDomainChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDomainsApi.UpdateAppDomainByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppDomainByUuid`: AppDomain
	fmt.Fprintf(os.Stdout, "Response from `ApplicationDomainsApi.UpdateAppDomainByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appDomainId** | **string** | App Domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppDomainByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appDomainChangeOperation** | [**[]AppDomainChangeOperation**](AppDomainChangeOperation.md) |  | 

### Return type

[**AppDomain**](AppDomain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


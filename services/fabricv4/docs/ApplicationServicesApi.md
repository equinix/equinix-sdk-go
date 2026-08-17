# \ApplicationServicesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppService**](ApplicationServicesApi.md#CreateAppService) | **Post** /fabric/v4/appServices | Create App Service
[**DeleteAppServiceByUuid**](ApplicationServicesApi.md#DeleteAppServiceByUuid) | **Delete** /fabric/v4/appServices/{appServiceId} | Delete App Service
[**GetAppServiceByUuid**](ApplicationServicesApi.md#GetAppServiceByUuid) | **Get** /fabric/v4/appServices/{appServiceId} | Get App Service
[**GetAttachedAppLinksByAppServiceId**](ApplicationServicesApi.md#GetAttachedAppLinksByAppServiceId) | **Get** /fabric/v4/appServices/{appServiceId}/appLinks | Get attached App Links for App Service
[**GetAttachedAppSubscriptionsByAppServiceId**](ApplicationServicesApi.md#GetAttachedAppSubscriptionsByAppServiceId) | **Get** /fabric/v4/appServices/{appServiceId}/appSubscriptions | Get attached App Subscriptions for App Service
[**SearchAppServices**](ApplicationServicesApi.md#SearchAppServices) | **Post** /fabric/v4/appServices/search | Search App Services
[**SearchAttachedAppSubscriptionsByAppServiceId**](ApplicationServicesApi.md#SearchAttachedAppSubscriptionsByAppServiceId) | **Post** /fabric/v4/appServices/{appServiceId}/appSubscriptions/search | Search attached App Subscriptions
[**UpdateAppServiceByUuid**](ApplicationServicesApi.md#UpdateAppServiceByUuid) | **Patch** /fabric/v4/appServices/{appServiceId} | Update App Service



## CreateAppService

> AppService CreateAppService(ctx).AppServicePostRequest(appServicePostRequest).Execute()

Create App Service



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
	appServicePostRequest := *openapiclient.NewAppServicePostRequest(openapiclient.AppServiceType("APP_SERVICE"), "Atlassian App Service", "equinixjira.atlassian.net", *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // AppServicePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.CreateAppService(context.Background()).AppServicePostRequest(appServicePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.CreateAppService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppService`: AppService
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.CreateAppService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appServicePostRequest** | [**AppServicePostRequest**](AppServicePostRequest.md) |  | 

### Return type

[**AppService**](AppService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppServiceByUuid

> AppService DeleteAppServiceByUuid(ctx, appServiceId).Execute()

Delete App Service



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
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.DeleteAppServiceByUuid(context.Background(), appServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.DeleteAppServiceByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppServiceByUuid`: AppService
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.DeleteAppServiceByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppService**](AppService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppServiceByUuid

> AppService GetAppServiceByUuid(ctx, appServiceId).Execute()

Get App Service



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
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.GetAppServiceByUuid(context.Background(), appServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.GetAppServiceByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppServiceByUuid`: AppService
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.GetAppServiceByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppService**](AppService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppLinksByAppServiceId

> AppServiceAttachedAppLinks GetAttachedAppLinksByAppServiceId(ctx, appServiceId).Offset(offset).Limit(limit).State(state).Order(order).Style(style).Execute()

Get attached App Links for App Service



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
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	state := []openapiclient.AppLinkState{openapiclient.AppLinkState("PROVISIONING")} // []AppLinkState | Filter attached App Links by one or more lifecycle states. (optional)
	order := openapiclient.AttachedAppLinkOrder("DESC") // AttachedAppLinkOrder | Sort order for attached App Links. (optional) (default to "DESC")
	style := openapiclient.Style("MIN") // Style | Detail level of the response. (optional) (default to "MEDIUM")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.GetAttachedAppLinksByAppServiceId(context.Background(), appServiceId).Offset(offset).Limit(limit).State(state).Order(order).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.GetAttachedAppLinksByAppServiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppLinksByAppServiceId`: AppServiceAttachedAppLinks
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.GetAttachedAppLinksByAppServiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppLinksByAppServiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **state** | [**[]AppLinkState**](AppLinkState.md) | Filter attached App Links by one or more lifecycle states. | 
 **order** | [**AttachedAppLinkOrder**](AttachedAppLinkOrder.md) | Sort order for attached App Links. | [default to &quot;DESC&quot;]
 **style** | [**Style**](Style.md) | Detail level of the response. | [default to &quot;MEDIUM&quot;]

### Return type

[**AppServiceAttachedAppLinks**](AppServiceAttachedAppLinks.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachedAppSubscriptionsByAppServiceId

> AppServiceAttachedAppSubscriptions GetAttachedAppSubscriptionsByAppServiceId(ctx, appServiceId).Offset(offset).Limit(limit).State(state).Order(order).Style(style).Execute()

Get attached App Subscriptions for App Service



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
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	state := []openapiclient.AppSubscriptionState{openapiclient.AppSubscriptionState("PROVISIONING")} // []AppSubscriptionState | Filter attached App Subscriptions by one or more lifecycle states. (optional)
	order := openapiclient.AttachedAppSubscriptionOrder("DESC") // AttachedAppSubscriptionOrder | Sort order for attached App Subscriptions. (optional) (default to "DESC")
	style := openapiclient.Style("MIN") // Style | Detail level of the response. (optional) (default to "MEDIUM")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.GetAttachedAppSubscriptionsByAppServiceId(context.Background(), appServiceId).Offset(offset).Limit(limit).State(state).Order(order).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.GetAttachedAppSubscriptionsByAppServiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachedAppSubscriptionsByAppServiceId`: AppServiceAttachedAppSubscriptions
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.GetAttachedAppSubscriptionsByAppServiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachedAppSubscriptionsByAppServiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **state** | [**[]AppSubscriptionState**](AppSubscriptionState.md) | Filter attached App Subscriptions by one or more lifecycle states. | 
 **order** | [**AttachedAppSubscriptionOrder**](AttachedAppSubscriptionOrder.md) | Sort order for attached App Subscriptions. | [default to &quot;DESC&quot;]
 **style** | [**Style**](Style.md) | Detail level of the response. | [default to &quot;MEDIUM&quot;]

### Return type

[**AppServiceAttachedAppSubscriptions**](AppServiceAttachedAppSubscriptions.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppServices

> AppServiceSearchResponse SearchAppServices(ctx).AppServiceSearchRequest(appServiceSearchRequest).Execute()

Search App Services



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
	appServiceSearchRequest := *openapiclient.NewAppServiceSearchRequest() // AppServiceSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.SearchAppServices(context.Background()).AppServiceSearchRequest(appServiceSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.SearchAppServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAppServices`: AppServiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.SearchAppServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appServiceSearchRequest** | [**AppServiceSearchRequest**](AppServiceSearchRequest.md) |  | 

### Return type

[**AppServiceSearchResponse**](AppServiceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAttachedAppSubscriptionsByAppServiceId

> AppServiceAttachedAppSubscriptionSearchResponse SearchAttachedAppSubscriptionsByAppServiceId(ctx, appServiceId).AppServiceAttachedAppSubscriptionSearchRequest(appServiceAttachedAppSubscriptionSearchRequest).Execute()

Search attached App Subscriptions



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
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID
	appServiceAttachedAppSubscriptionSearchRequest := *openapiclient.NewAppServiceAttachedAppSubscriptionSearchRequest() // AppServiceAttachedAppSubscriptionSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.SearchAttachedAppSubscriptionsByAppServiceId(context.Background(), appServiceId).AppServiceAttachedAppSubscriptionSearchRequest(appServiceAttachedAppSubscriptionSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.SearchAttachedAppSubscriptionsByAppServiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAttachedAppSubscriptionsByAppServiceId`: AppServiceAttachedAppSubscriptionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.SearchAttachedAppSubscriptionsByAppServiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAttachedAppSubscriptionsByAppServiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appServiceAttachedAppSubscriptionSearchRequest** | [**AppServiceAttachedAppSubscriptionSearchRequest**](AppServiceAttachedAppSubscriptionSearchRequest.md) |  | 

### Return type

[**AppServiceAttachedAppSubscriptionSearchResponse**](AppServiceAttachedAppSubscriptionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceByUuid

> AppService UpdateAppServiceByUuid(ctx, appServiceId).AppServiceChangeOperation(appServiceChangeOperation).Execute()

Update App Service



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
	appServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Service UUID
	appServiceChangeOperation := []openapiclient.AppServiceChangeOperation{*openapiclient.NewAppServiceChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), openapiclient.AppServiceChangeOperation_path("/name"), map[string]interface{}(123))} // []AppServiceChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationServicesApi.UpdateAppServiceByUuid(context.Background(), appServiceId).AppServiceChangeOperation(appServiceChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationServicesApi.UpdateAppServiceByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceByUuid`: AppService
	fmt.Fprintf(os.Stdout, "Response from `ApplicationServicesApi.UpdateAppServiceByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appServiceId** | **string** | App Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appServiceChangeOperation** | [**[]AppServiceChangeOperation**](AppServiceChangeOperation.md) |  | 

### Return type

[**AppService**](AppService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


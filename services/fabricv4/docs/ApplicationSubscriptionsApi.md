# \ApplicationSubscriptionsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppSubscription**](ApplicationSubscriptionsApi.md#CreateAppSubscription) | **Post** /fabric/v4/appSubscriptions | Create App Subscription
[**DeleteAppSubscriptionByUuid**](ApplicationSubscriptionsApi.md#DeleteAppSubscriptionByUuid) | **Delete** /fabric/v4/appSubscriptions/{appSubscriptionId} | Delete App Subscription
[**GetAppSubscriptionByUuid**](ApplicationSubscriptionsApi.md#GetAppSubscriptionByUuid) | **Get** /fabric/v4/appSubscriptions/{appSubscriptionId} | Get App Subscription
[**SearchAppSubscriptions**](ApplicationSubscriptionsApi.md#SearchAppSubscriptions) | **Post** /fabric/v4/appSubscriptions/search | Search App Subscriptions
[**UpdateAppSubscriptionByUuid**](ApplicationSubscriptionsApi.md#UpdateAppSubscriptionByUuid) | **Patch** /fabric/v4/appSubscriptions/{appSubscriptionId} | Update App Subscription



## CreateAppSubscription

> AppSubscription CreateAppSubscription(ctx).AppSubscriptionPostRequest(appSubscriptionPostRequest).Execute()

Create App Subscription



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
	appSubscriptionPostRequest := *openapiclient.NewAppSubscriptionPostRequest(openapiclient.AppSubscriptionType("APP_SUBSCRIPTION"), *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // AppSubscriptionPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSubscriptionsApi.CreateAppSubscription(context.Background()).AppSubscriptionPostRequest(appSubscriptionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSubscriptionsApi.CreateAppSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppSubscription`: AppSubscription
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSubscriptionsApi.CreateAppSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appSubscriptionPostRequest** | [**AppSubscriptionPostRequest**](AppSubscriptionPostRequest.md) |  | 

### Return type

[**AppSubscription**](AppSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppSubscriptionByUuid

> AppSubscription DeleteAppSubscriptionByUuid(ctx, appSubscriptionId).Execute()

Delete App Subscription



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
	appSubscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Subscription UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSubscriptionsApi.DeleteAppSubscriptionByUuid(context.Background(), appSubscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSubscriptionsApi.DeleteAppSubscriptionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppSubscriptionByUuid`: AppSubscription
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSubscriptionsApi.DeleteAppSubscriptionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSubscriptionId** | **string** | App Subscription UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppSubscriptionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppSubscription**](AppSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppSubscriptionByUuid

> AppSubscription GetAppSubscriptionByUuid(ctx, appSubscriptionId).Execute()

Get App Subscription



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
	appSubscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Subscription UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSubscriptionsApi.GetAppSubscriptionByUuid(context.Background(), appSubscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSubscriptionsApi.GetAppSubscriptionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppSubscriptionByUuid`: AppSubscription
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSubscriptionsApi.GetAppSubscriptionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSubscriptionId** | **string** | App Subscription UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppSubscriptionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppSubscription**](AppSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppSubscriptions

> AppSubscriptionSearchResponse SearchAppSubscriptions(ctx).AppSubscriptionSearchRequest(appSubscriptionSearchRequest).Execute()

Search App Subscriptions



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
	appSubscriptionSearchRequest := *openapiclient.NewAppSubscriptionSearchRequest() // AppSubscriptionSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSubscriptionsApi.SearchAppSubscriptions(context.Background()).AppSubscriptionSearchRequest(appSubscriptionSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSubscriptionsApi.SearchAppSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAppSubscriptions`: AppSubscriptionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSubscriptionsApi.SearchAppSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appSubscriptionSearchRequest** | [**AppSubscriptionSearchRequest**](AppSubscriptionSearchRequest.md) |  | 

### Return type

[**AppSubscriptionSearchResponse**](AppSubscriptionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSubscriptionByUuid

> AppSubscription UpdateAppSubscriptionByUuid(ctx, appSubscriptionId).AppSubscriptionChangeOperation(appSubscriptionChangeOperation).Execute()

Update App Subscription



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
	appSubscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App Subscription UUID
	appSubscriptionChangeOperation := []openapiclient.AppSubscriptionChangeOperation{*openapiclient.NewAppSubscriptionChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), "Path_example", map[string]interface{}(123))} // []AppSubscriptionChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSubscriptionsApi.UpdateAppSubscriptionByUuid(context.Background(), appSubscriptionId).AppSubscriptionChangeOperation(appSubscriptionChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSubscriptionsApi.UpdateAppSubscriptionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppSubscriptionByUuid`: AppSubscription
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSubscriptionsApi.UpdateAppSubscriptionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSubscriptionId** | **string** | App Subscription UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSubscriptionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appSubscriptionChangeOperation** | [**[]AppSubscriptionChangeOperation**](AppSubscriptionChangeOperation.md) |  | 

### Return type

[**AppSubscription**](AppSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


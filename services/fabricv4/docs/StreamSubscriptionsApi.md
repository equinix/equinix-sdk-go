# \StreamSubscriptionsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamSubscriptions**](StreamSubscriptionsApi.md#CreateStreamSubscriptions) | **Post** /fabric/v4/streams/{streamId}/subscriptions | Create Subscription
[**DeleteStreamSubscriptionByUuid**](StreamSubscriptionsApi.md#DeleteStreamSubscriptionByUuid) | **Delete** /fabric/v4/streams/{streamId}/subscriptions/{subscriptionId} | Delete Subscription
[**GetStreamSubscriptionByUuid**](StreamSubscriptionsApi.md#GetStreamSubscriptionByUuid) | **Get** /fabric/v4/streams/{streamId}/subscriptions/{subscriptionId} | Get Subscription
[**GetStreamSubscriptions**](StreamSubscriptionsApi.md#GetStreamSubscriptions) | **Get** /fabric/v4/streams/{streamId}/subscriptions | Get Subscriptions
[**UpdateStreamSubscriptionByUuid**](StreamSubscriptionsApi.md#UpdateStreamSubscriptionByUuid) | **Put** /fabric/v4/streams/{streamId}/subscriptions/{subscriptionId} | Update Subscription



## CreateStreamSubscriptions

> StreamSubscription CreateStreamSubscriptions(ctx, streamId).StreamSubscriptionPostRequest(streamSubscriptionPostRequest).Execute()

Create Subscription



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
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID
	streamSubscriptionPostRequest := *openapiclient.NewStreamSubscriptionPostRequest(openapiclient.StreamSubscriptionPostRequest_type("STREAM_SUBSCRIPTION"), "Name_example", *openapiclient.NewStreamSubscriptionSink(openapiclient.StreamSubscriptionSink_type("DATADOG"))) // StreamSubscriptionPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamSubscriptionsApi.CreateStreamSubscriptions(context.Background(), streamId).StreamSubscriptionPostRequest(streamSubscriptionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamSubscriptionsApi.CreateStreamSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStreamSubscriptions`: StreamSubscription
	fmt.Fprintf(os.Stdout, "Response from `StreamSubscriptionsApi.CreateStreamSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamSubscriptionPostRequest** | [**StreamSubscriptionPostRequest**](StreamSubscriptionPostRequest.md) |  | 

### Return type

[**StreamSubscription**](StreamSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamSubscriptionByUuid

> StreamSubscription DeleteStreamSubscriptionByUuid(ctx, streamId, subscriptionId).Execute()

Delete Subscription



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
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream Subscription UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamSubscriptionsApi.DeleteStreamSubscriptionByUuid(context.Background(), streamId, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamSubscriptionsApi.DeleteStreamSubscriptionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStreamSubscriptionByUuid`: StreamSubscription
	fmt.Fprintf(os.Stdout, "Response from `StreamSubscriptionsApi.DeleteStreamSubscriptionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 
**subscriptionId** | **string** | Stream Subscription UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamSubscriptionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamSubscription**](StreamSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamSubscriptionByUuid

> StreamSubscription GetStreamSubscriptionByUuid(ctx, streamId, subscriptionId).Execute()

Get Subscription



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
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream Subscription UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamSubscriptionsApi.GetStreamSubscriptionByUuid(context.Background(), streamId, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamSubscriptionsApi.GetStreamSubscriptionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamSubscriptionByUuid`: StreamSubscription
	fmt.Fprintf(os.Stdout, "Response from `StreamSubscriptionsApi.GetStreamSubscriptionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 
**subscriptionId** | **string** | Stream Subscription UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamSubscriptionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamSubscription**](StreamSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamSubscriptions

> GetAllStreamSubscriptionResponse GetStreamSubscriptions(ctx, streamId).Offset(offset).Limit(limit).Execute()

Get Subscriptions



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
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamSubscriptionsApi.GetStreamSubscriptions(context.Background(), streamId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamSubscriptionsApi.GetStreamSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamSubscriptions`: GetAllStreamSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamSubscriptionsApi.GetStreamSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetAllStreamSubscriptionResponse**](GetAllStreamSubscriptionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamSubscriptionByUuid

> StreamSubscription UpdateStreamSubscriptionByUuid(ctx, streamId, subscriptionId).StreamSubscriptionPutRequest(streamSubscriptionPutRequest).Execute()

Update Subscription



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
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream Subscription UUID
	streamSubscriptionPutRequest := *openapiclient.NewStreamSubscriptionPutRequest(openapiclient.StreamSubscriptionPostRequest_type("STREAM_SUBSCRIPTION"), "Name_example", *openapiclient.NewStreamSubscriptionSink(openapiclient.StreamSubscriptionSink_type("DATADOG"))) // StreamSubscriptionPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamSubscriptionsApi.UpdateStreamSubscriptionByUuid(context.Background(), streamId, subscriptionId).StreamSubscriptionPutRequest(streamSubscriptionPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamSubscriptionsApi.UpdateStreamSubscriptionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStreamSubscriptionByUuid`: StreamSubscription
	fmt.Fprintf(os.Stdout, "Response from `StreamSubscriptionsApi.UpdateStreamSubscriptionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 
**subscriptionId** | **string** | Stream Subscription UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamSubscriptionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **streamSubscriptionPutRequest** | [**StreamSubscriptionPutRequest**](StreamSubscriptionPutRequest.md) |  | 

### Return type

[**StreamSubscription**](StreamSubscription.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


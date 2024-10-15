# \StreamsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreams**](StreamsApi.md#CreateStreams) | **Post** /fabric/v4/streams | Create Stream
[**DeleteStreamAssetByUuid**](StreamsApi.md#DeleteStreamAssetByUuid) | **Delete** /fabric/v4/streams/{streamId}/{asset}/{assetId} | Detach Asset
[**DeleteStreamByUuid**](StreamsApi.md#DeleteStreamByUuid) | **Delete** /fabric/v4/streams/{streamId} | Delete Stream
[**GetStreamAssetByUuid**](StreamsApi.md#GetStreamAssetByUuid) | **Get** /fabric/v4/streams/{streamId}/{asset}/{assetId} | Get Asset
[**GetStreamByUuid**](StreamsApi.md#GetStreamByUuid) | **Get** /fabric/v4/streams/{streamId} | Get Stream
[**GetStreams**](StreamsApi.md#GetStreams) | **Get** /fabric/v4/streams | Get Streams
[**GetStreamsAssets**](StreamsApi.md#GetStreamsAssets) | **Post** /fabric/v4/streamAssets/search | Get Assets
[**GetSubscriptionsInStream**](StreamsApi.md#GetSubscriptionsInStream) | **Get** /fabric/v4/streams/{streamId}/streamSubscriptions | Get Stream&#39;s Subs
[**UpdateStreamAssetByUuid**](StreamsApi.md#UpdateStreamAssetByUuid) | **Put** /fabric/v4/streams/{streamId}/{asset}/{assetId} | Attach Asset
[**UpdateStreamByUuid**](StreamsApi.md#UpdateStreamByUuid) | **Put** /fabric/v4/streams/{streamId} | Update Stream



## CreateStreams

> Stream CreateStreams(ctx).StreamPostRequest(streamPostRequest).Execute()

Create Stream



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
	streamPostRequest := *openapiclient.NewStreamPostRequest() // StreamPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.CreateStreams(context.Background()).StreamPostRequest(streamPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.CreateStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStreams`: Stream
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.CreateStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamPostRequest** | [**StreamPostRequest**](StreamPostRequest.md) |  | 

### Return type

[**Stream**](Stream.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamAssetByUuid

> StreamAsset DeleteStreamAssetByUuid(ctx, assetId, asset, streamId).Execute()

Detach Asset



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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | asset UUID
	asset := openapiclient.Asset("ports") // Asset | asset
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.DeleteStreamAssetByUuid(context.Background(), assetId, asset, streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamAssetByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStreamAssetByUuid`: StreamAsset
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.DeleteStreamAssetByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | asset UUID | 
**asset** | [**Asset**](.md) | asset | 
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamAssetByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamAsset**](StreamAsset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamByUuid

> Stream DeleteStreamByUuid(ctx, streamId).Execute()

Delete Stream



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.DeleteStreamByUuid(context.Background(), streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.DeleteStreamByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStreamByUuid`: Stream
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.DeleteStreamByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stream**](Stream.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamAssetByUuid

> StreamAsset GetStreamAssetByUuid(ctx, assetId, asset, streamId).Execute()

Get Asset



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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | asset UUID
	asset := openapiclient.Asset("ports") // Asset | asset
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.GetStreamAssetByUuid(context.Background(), assetId, asset, streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamAssetByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamAssetByUuid`: StreamAsset
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamAssetByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | asset UUID | 
**asset** | [**Asset**](.md) | asset | 
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamAssetByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamAsset**](StreamAsset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamByUuid

> Stream GetStreamByUuid(ctx, streamId).Execute()

Get Stream



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.GetStreamByUuid(context.Background(), streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamByUuid`: Stream
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stream**](Stream.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreams

> GetAllStreamResponse GetStreams(ctx).Offset(offset).Limit(limit).Execute()

Get Streams



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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.GetStreams(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreams`: GetAllStreamResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetAllStreamResponse**](GetAllStreamResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamsAssets

> GetAllStreamAssetResponse GetStreamsAssets(ctx).StreamAssetSearchRequest(streamAssetSearchRequest).Offset(offset).Limit(limit).Execute()

Get Assets



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
	streamAssetSearchRequest := *openapiclient.NewStreamAssetSearchRequest() // StreamAssetSearchRequest | 
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.GetStreamsAssets(context.Background()).StreamAssetSearchRequest(streamAssetSearchRequest).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetStreamsAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamsAssets`: GetAllStreamAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetStreamsAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamsAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamAssetSearchRequest** | [**StreamAssetSearchRequest**](StreamAssetSearchRequest.md) |  | 
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetAllStreamAssetResponse**](GetAllStreamAssetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionsInStream

> GetSubscriptionsInStreamResponse GetSubscriptionsInStream(ctx, streamId).Offset(offset).Limit(limit).Execute()

Get Stream's Subs



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
	resp, r, err := apiClient.StreamsApi.GetSubscriptionsInStream(context.Background(), streamId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetSubscriptionsInStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionsInStream`: GetSubscriptionsInStreamResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetSubscriptionsInStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsInStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetSubscriptionsInStreamResponse**](GetSubscriptionsInStreamResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamAssetByUuid

> StreamAsset UpdateStreamAssetByUuid(ctx, assetId, asset, streamId).StreamAssetPutRequest(streamAssetPutRequest).Execute()

Attach Asset



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
	assetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | asset UUID
	asset := openapiclient.Asset("ports") // Asset | asset
	streamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Stream UUID
	streamAssetPutRequest := *openapiclient.NewStreamAssetPutRequest() // StreamAssetPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.UpdateStreamAssetByUuid(context.Background(), assetId, asset, streamId).StreamAssetPutRequest(streamAssetPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.UpdateStreamAssetByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStreamAssetByUuid`: StreamAsset
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.UpdateStreamAssetByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** | asset UUID | 
**asset** | [**Asset**](.md) | asset | 
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamAssetByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamAssetPutRequest** | [**StreamAssetPutRequest**](StreamAssetPutRequest.md) |  | 

### Return type

[**StreamAsset**](StreamAsset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamByUuid

> Stream UpdateStreamByUuid(ctx, streamId).StreamPutRequest(streamPutRequest).Execute()

Update Stream



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
	streamPutRequest := *openapiclient.NewStreamPutRequest() // StreamPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsApi.UpdateStreamByUuid(context.Background(), streamId).StreamPutRequest(streamPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.UpdateStreamByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStreamByUuid`: Stream
	fmt.Fprintf(os.Stdout, "Response from `StreamsApi.UpdateStreamByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamPutRequest** | [**StreamPutRequest**](StreamPutRequest.md) |  | 

### Return type

[**Stream**](Stream.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


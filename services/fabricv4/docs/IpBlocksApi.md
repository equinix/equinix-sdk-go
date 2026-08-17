# \IPBlocksApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpBlockById**](IPBlocksApi.md#DeleteIpBlockById) | **Delete** /fabric/v4/ipBlocks/{uuid} | Delete IP Block by UUID
[**GetIpBlock**](IPBlocksApi.md#GetIpBlock) | **Get** /fabric/v4/ipBlocks/{uuid} | Retrieve IP Block by UUID
[**PatchIpBlockById**](IPBlocksApi.md#PatchIpBlockById) | **Patch** /fabric/v4/ipBlocks/{uuid} | patch IP Block by UUID
[**SearchIpBlocks**](IPBlocksApi.md#SearchIpBlocks) | **Post** /fabric/v4/ipBlocks/search | Search for IP Blocks
[**SubmitIpBlock**](IPBlocksApi.md#SubmitIpBlock) | **Post** /fabric/v4/ipBlocks | Submits new Equinix owned or customer owned IP Block request



## DeleteIpBlockById

> IpBlock DeleteIpBlockById(ctx, uuid).Execute()

Delete IP Block by UUID



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the IP Block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBlocksApi.DeleteIpBlockById(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.DeleteIpBlockById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIpBlockById`: IpBlock
	fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.DeleteIpBlockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the IP Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpBlockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpBlock

> IpBlock GetIpBlock(ctx, uuid).Execute()

Retrieve IP Block by UUID



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the IP Block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBlocksApi.GetIpBlock(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.GetIpBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpBlock`: IpBlock
	fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.GetIpBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the IP Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIpBlockById

> IpBlock PatchIpBlockById(ctx, uuid).PatchIpBlockRequestBodyItem(patchIpBlockRequestBodyItem).Execute()

patch IP Block by UUID



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the IP Block
	patchIpBlockRequestBodyItem := []openapiclient.PatchIpBlockRequestBodyItem{*openapiclient.NewPatchIpBlockRequestBodyItem(openapiclient.PatchIpBlockRequestBodyItem_op("add"), "Path_example")} // []PatchIpBlockRequestBodyItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBlocksApi.PatchIpBlockById(context.Background(), uuid).PatchIpBlockRequestBodyItem(patchIpBlockRequestBodyItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.PatchIpBlockById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIpBlockById`: IpBlock
	fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.PatchIpBlockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the IP Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIpBlockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchIpBlockRequestBodyItem** | [**[]PatchIpBlockRequestBodyItem**](PatchIpBlockRequestBodyItem.md) |  | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIpBlocks

> IpBlockSearchResponseBody SearchIpBlocks(ctx).IpBlocksSearchRequestBody(ipBlocksSearchRequestBody).Execute()

Search for IP Blocks



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
	ipBlocksSearchRequestBody := *openapiclient.NewIpBlocksSearchRequestBody() // IpBlocksSearchRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBlocksApi.SearchIpBlocks(context.Background()).IpBlocksSearchRequestBody(ipBlocksSearchRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.SearchIpBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIpBlocks`: IpBlockSearchResponseBody
	fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.SearchIpBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIpBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipBlocksSearchRequestBody** | [**IpBlocksSearchRequestBody**](IpBlocksSearchRequestBody.md) |  | 

### Return type

[**IpBlockSearchResponseBody**](IpBlockSearchResponseBody.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitIpBlock

> IpBlock SubmitIpBlock(ctx).SubmitIpBlockRequestBody(submitIpBlockRequestBody).Execute()

Submits new Equinix owned or customer owned IP Block request



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
	submitIpBlockRequestBody := *openapiclient.NewSubmitIpBlockRequestBody(openapiclient.TypeOfIpBlockProduct("IPV4_IP_BLOCK"), *openapiclient.NewIpBlockProjectRequest("ProjectId_example")) // SubmitIpBlockRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBlocksApi.SubmitIpBlock(context.Background()).SubmitIpBlockRequestBody(submitIpBlockRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.SubmitIpBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitIpBlock`: IpBlock
	fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.SubmitIpBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitIpBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitIpBlockRequestBody** | [**SubmitIpBlockRequestBody**](SubmitIpBlockRequestBody.md) |  | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


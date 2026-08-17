# \LoasApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoa**](LoasApi.md#CreateLoa) | **Post** /fabric/v4/loas | Create Loa
[**CreateLoaNoteByLoaId**](LoasApi.md#CreateLoaNoteByLoaId) | **Post** /fabric/v4/loas/{loaId}/notes | Create Loa Note
[**GetLoaActionsByUuid**](LoasApi.md#GetLoaActionsByUuid) | **Get** /fabric/v4/loas/{loaId}/actions/{actionId} | Get Loa Action by Action ID
[**GetLoaByUuid**](LoasApi.md#GetLoaByUuid) | **Get** /fabric/v4/loas/{loaId} | Get Loa
[**GetLoaConsumersByLoaId**](LoasApi.md#GetLoaConsumersByLoaId) | **Get** /fabric/v4/loas/{loaId}/consumers | Get Loa Consumers
[**GetLoaNotesByUuid**](LoasApi.md#GetLoaNotesByUuid) | **Get** /fabric/v4/loas/{loaId}/notes | Get Loa Notes
[**PerformLoaAction**](LoasApi.md#PerformLoaAction) | **Post** /fabric/v4/loas/{loaId}/actions | Loa Actions
[**SearchLoa**](LoasApi.md#SearchLoa) | **Post** /fabric/v4/loas/search | Search Loas
[**SearchLoaAction**](LoasApi.md#SearchLoaAction) | **Post** /fabric/v4/loas/{loaId}/actions/search | Search Loa Actions
[**UpdateLoaByUuid**](LoasApi.md#UpdateLoaByUuid) | **Patch** /fabric/v4/loas/{loaId} | Update Loa



## CreateLoa

> LoaResponse CreateLoa(ctx).CreateLoa(createLoa).Execute()

Create Loa



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
	createLoa := openapiclient.CreateLoa{IssueLoa: openapiclient.NewIssueLoa(openapiclient.LoaType("CAGE_LOA"), "LOA-Equinix-SV1-2026", openapiclient.LoaProductType("XC"), *openapiclient.NewLoaDemarcationPoint())} // CreateLoa | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.CreateLoa(context.Background()).CreateLoa(createLoa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.CreateLoa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoa`: LoaResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.CreateLoa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoa** | [**CreateLoa**](CreateLoa.md) |  | 

### Return type

[**LoaResponse**](LoaResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoaNoteByLoaId

> LoaNoteDetails CreateLoaNoteByLoaId(ctx, loaId).CreateLoaNote(createLoaNote).Execute()

Create Loa Note



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID
	createLoaNote := *openapiclient.NewCreateLoaNote("LOA under review by our team. Expected response within 2 business days") // CreateLoaNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.CreateLoaNoteByLoaId(context.Background(), loaId).CreateLoaNote(createLoaNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.CreateLoaNoteByLoaId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoaNoteByLoaId`: LoaNoteDetails
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.CreateLoaNoteByLoaId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoaNoteByLoaIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLoaNote** | [**CreateLoaNote**](CreateLoaNote.md) |  | 

### Return type

[**LoaNoteDetails**](LoaNoteDetails.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoaActionsByUuid

> LoaActionResponse GetLoaActionsByUuid(ctx, loaId, actionId).Execute()

Get Loa Action by Action ID



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID
	actionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Action UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.GetLoaActionsByUuid(context.Background(), loaId, actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.GetLoaActionsByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoaActionsByUuid`: LoaActionResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.GetLoaActionsByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 
**actionId** | **string** | Action UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoaActionsByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoaActionResponse**](LoaActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoaByUuid

> LoaResponse GetLoaByUuid(ctx, loaId).Execute()

Get Loa



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.GetLoaByUuid(context.Background(), loaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.GetLoaByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoaByUuid`: LoaResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.GetLoaByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoaByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoaResponse**](LoaResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoaConsumersByLoaId

> LoaConsumersResponse GetLoaConsumersByLoaId(ctx, loaId).Execute()

Get Loa Consumers



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.GetLoaConsumersByLoaId(context.Background(), loaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.GetLoaConsumersByLoaId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoaConsumersByLoaId`: LoaConsumersResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.GetLoaConsumersByLoaId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoaConsumersByLoaIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoaConsumersResponse**](LoaConsumersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoaNotesByUuid

> LoaNotesResponse GetLoaNotesByUuid(ctx, loaId).Execute()

Get Loa Notes



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.GetLoaNotesByUuid(context.Background(), loaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.GetLoaNotesByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoaNotesByUuid`: LoaNotesResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.GetLoaNotesByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoaNotesByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoaNotesResponse**](LoaNotesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformLoaAction

> LoaActionResponse PerformLoaAction(ctx, loaId).LoaActionRequest(loaActionRequest).Execute()

Loa Actions



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID
	loaActionRequest := *openapiclient.NewLoaActionRequest(openapiclient.LoaActionType("LOA_ISSUER_AUTHORIZATION")) // LoaActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.PerformLoaAction(context.Background(), loaId).LoaActionRequest(loaActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.PerformLoaAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformLoaAction`: LoaActionResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.PerformLoaAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformLoaActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loaActionRequest** | [**LoaActionRequest**](LoaActionRequest.md) |  | 

### Return type

[**LoaActionResponse**](LoaActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLoa

> LoaSearchResponse SearchLoa(ctx).LoaSearchRequest(loaSearchRequest).Execute()

Search Loas



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
	loaSearchRequest := *openapiclient.NewLoaSearchRequest() // LoaSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.SearchLoa(context.Background()).LoaSearchRequest(loaSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.SearchLoa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLoa`: LoaSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.SearchLoa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLoaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loaSearchRequest** | [**LoaSearchRequest**](LoaSearchRequest.md) |  | 

### Return type

[**LoaSearchResponse**](LoaSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLoaAction

> LoaActionSearchResponse SearchLoaAction(ctx, loaId).LoaActionSearchRequest(loaActionSearchRequest).Execute()

Search Loa Actions



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID
	loaActionSearchRequest := *openapiclient.NewLoaActionSearchRequest() // LoaActionSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.SearchLoaAction(context.Background(), loaId).LoaActionSearchRequest(loaActionSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.SearchLoaAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLoaAction`: LoaActionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.SearchLoaAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchLoaActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loaActionSearchRequest** | [**LoaActionSearchRequest**](LoaActionSearchRequest.md) |  | 

### Return type

[**LoaActionSearchResponse**](LoaActionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoaByUuid

> LoaResponse UpdateLoaByUuid(ctx, loaId).LoaReplaceOperation(loaReplaceOperation).Execute()

Update Loa



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
	loaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Loa UUID
	loaReplaceOperation := []openapiclient.LoaReplaceOperation{*openapiclient.NewLoaReplaceOperation(openapiclient.LoaOpEnum("replace"), "Path_example", map[string]interface{}(123))} // []LoaReplaceOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoasApi.UpdateLoaByUuid(context.Background(), loaId).LoaReplaceOperation(loaReplaceOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoasApi.UpdateLoaByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoaByUuid`: LoaResponse
	fmt.Fprintf(os.Stdout, "Response from `LoasApi.UpdateLoaByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loaId** | **string** | Loa UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoaByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loaReplaceOperation** | [**[]LoaReplaceOperation**](LoaReplaceOperation.md) |  | 

### Return type

[**LoaResponse**](LoaResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


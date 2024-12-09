# \ServiceTokensApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceToken**](ServiceTokensApi.md#CreateServiceToken) | **Post** /fabric/v4/serviceTokens | Create Service Token
[**CreateServiceTokenAction**](ServiceTokensApi.md#CreateServiceTokenAction) | **Post** /fabric/v4/serviceTokens/{serviceTokenId}/actions | ServiceToken Actions
[**DeleteServiceTokenByUuid**](ServiceTokensApi.md#DeleteServiceTokenByUuid) | **Delete** /fabric/v4/serviceTokens/{serviceTokenId} | Delete Token by uuid
[**GetServiceTokenByUuid**](ServiceTokensApi.md#GetServiceTokenByUuid) | **Get** /fabric/v4/serviceTokens/{serviceTokenId} | Get Token by uuid
[**GetServiceTokens**](ServiceTokensApi.md#GetServiceTokens) | **Get** /fabric/v4/serviceTokens | Get All Tokens
[**SearchServiceTokens**](ServiceTokensApi.md#SearchServiceTokens) | **Post** /fabric/v4/serviceTokens/search | Search servicetokens
[**UpdateServiceTokenByUuid**](ServiceTokensApi.md#UpdateServiceTokenByUuid) | **Patch** /fabric/v4/serviceTokens/{serviceTokenId} | Update Token By ID



## CreateServiceToken

> ServiceToken CreateServiceToken(ctx).ServiceToken(serviceToken).DryRun(dryRun).Execute()

Create Service Token



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
	serviceToken := *openapiclient.NewServiceToken() // ServiceToken | 
	dryRun := true // bool | option to verify that API calls will succeed (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.CreateServiceToken(context.Background()).ServiceToken(serviceToken).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.CreateServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceToken`: ServiceToken
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.CreateServiceToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceToken** | [**ServiceToken**](ServiceToken.md) |  | 
 **dryRun** | **bool** | option to verify that API calls will succeed | [default to false]

### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceTokenAction

> ServiceToken CreateServiceTokenAction(ctx, serviceTokenId).ServiceTokenActionRequest(serviceTokenActionRequest).Execute()

ServiceToken Actions



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
	serviceTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Token UUID
	serviceTokenActionRequest := *openapiclient.NewServiceTokenActionRequest(openapiclient.ServiceTokenActions("RESEND_EMAIL_NOTIFICATION")) // ServiceTokenActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.CreateServiceTokenAction(context.Background(), serviceTokenId).ServiceTokenActionRequest(serviceTokenActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.CreateServiceTokenAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceTokenAction`: ServiceToken
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.CreateServiceTokenAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceTokenId** | **string** | Service Token UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceTokenActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceTokenActionRequest** | [**ServiceTokenActionRequest**](ServiceTokenActionRequest.md) |  | 

### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceTokenByUuid

> ServiceToken DeleteServiceTokenByUuid(ctx, serviceTokenId).Execute()

Delete Token by uuid



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
	serviceTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Token UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.DeleteServiceTokenByUuid(context.Background(), serviceTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.DeleteServiceTokenByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceTokenByUuid`: ServiceToken
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.DeleteServiceTokenByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceTokenId** | **string** | Service Token UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceTokenByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTokenByUuid

> ServiceToken GetServiceTokenByUuid(ctx, serviceTokenId).Execute()

Get Token by uuid



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
	serviceTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Token UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.GetServiceTokenByUuid(context.Background(), serviceTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.GetServiceTokenByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTokenByUuid`: ServiceToken
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.GetServiceTokenByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceTokenId** | **string** | Service Token UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTokenByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTokens

> ServiceTokens GetServiceTokens(ctx).Offset(offset).Limit(limit).Execute()

Get All Tokens



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
	offset := float32(8.14) // float32 | offset (optional)
	limit := float32(8.14) // float32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.GetServiceTokens(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.GetServiceTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTokens`: ServiceTokens
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.GetServiceTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float32** | offset | 
 **limit** | **float32** | number of records to fetch | 

### Return type

[**ServiceTokens**](ServiceTokens.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServiceTokens

> ServiceTokens SearchServiceTokens(ctx).ServiceTokenSearchRequest(serviceTokenSearchRequest).Offset(offset).Limit(limit).Execute()

Search servicetokens



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
	serviceTokenSearchRequest := *openapiclient.NewServiceTokenSearchRequest() // ServiceTokenSearchRequest | 
	offset := float32(8.14) // float32 | offset (optional)
	limit := float32(8.14) // float32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.SearchServiceTokens(context.Background()).ServiceTokenSearchRequest(serviceTokenSearchRequest).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.SearchServiceTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchServiceTokens`: ServiceTokens
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.SearchServiceTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServiceTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceTokenSearchRequest** | [**ServiceTokenSearchRequest**](ServiceTokenSearchRequest.md) |  | 
 **offset** | **float32** | offset | 
 **limit** | **float32** | number of records to fetch | 

### Return type

[**ServiceTokens**](ServiceTokens.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceTokenByUuid

> ServiceToken UpdateServiceTokenByUuid(ctx, serviceTokenId).ServiceTokenChangeOperation(serviceTokenChangeOperation).Execute()

Update Token By ID



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
	serviceTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Token UUID
	serviceTokenChangeOperation := []openapiclient.ServiceTokenChangeOperation{*openapiclient.NewServiceTokenChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), "/expirationDateTime", interface{}(123))} // []ServiceTokenChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokensApi.UpdateServiceTokenByUuid(context.Background(), serviceTokenId).ServiceTokenChangeOperation(serviceTokenChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokensApi.UpdateServiceTokenByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceTokenByUuid`: ServiceToken
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokensApi.UpdateServiceTokenByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceTokenId** | **string** | Service Token UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceTokenByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceTokenChangeOperation** | [**[]ServiceTokenChangeOperation**](ServiceTokenChangeOperation.md) |  | 

### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


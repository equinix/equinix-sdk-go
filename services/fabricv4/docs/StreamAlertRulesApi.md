# \StreamAlertRulesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamAlertRules**](StreamAlertRulesApi.md#CreateStreamAlertRules) | **Post** /fabric/v4/streams/{streamId}/alertRules | Create Stream Alert Rules
[**DeleteStreamAlertRuleByUuid**](StreamAlertRulesApi.md#DeleteStreamAlertRuleByUuid) | **Delete** /fabric/v4/streams/{streamId}/alertRules/{alertRuleId} | Update Stream Alert Rules
[**GetStreamAlertRuleByUuid**](StreamAlertRulesApi.md#GetStreamAlertRuleByUuid) | **Get** /fabric/v4/streams/{streamId}/alertRules/{alertRuleId} | Get Stream Alert Rules
[**GetStreamAlertRules**](StreamAlertRulesApi.md#GetStreamAlertRules) | **Get** /fabric/v4/streams/{streamId}/alertRules | Get Stream Alert Rules
[**UpdateStreamAlertRuleByUuid**](StreamAlertRulesApi.md#UpdateStreamAlertRuleByUuid) | **Put** /fabric/v4/streams/{streamId}/alertRules/{alertRuleId} | Update Stream Alert Rules



## CreateStreamAlertRules

> StreamAlertRule CreateStreamAlertRules(ctx, streamId).AlertRulePostRequest(alertRulePostRequest).Execute()

Create Stream Alert Rules



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
	alertRulePostRequest := *openapiclient.NewAlertRulePostRequest() // AlertRulePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamAlertRulesApi.CreateStreamAlertRules(context.Background(), streamId).AlertRulePostRequest(alertRulePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAlertRulesApi.CreateStreamAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStreamAlertRules`: StreamAlertRule
	fmt.Fprintf(os.Stdout, "Response from `StreamAlertRulesApi.CreateStreamAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertRulePostRequest** | [**AlertRulePostRequest**](AlertRulePostRequest.md) |  | 

### Return type

[**StreamAlertRule**](StreamAlertRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamAlertRuleByUuid

> StreamAlertRule DeleteStreamAlertRuleByUuid(ctx, streamId, alertRuleId).Execute()

Update Stream Alert Rules



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
	alertRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | alert rule UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamAlertRulesApi.DeleteStreamAlertRuleByUuid(context.Background(), streamId, alertRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAlertRulesApi.DeleteStreamAlertRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStreamAlertRuleByUuid`: StreamAlertRule
	fmt.Fprintf(os.Stdout, "Response from `StreamAlertRulesApi.DeleteStreamAlertRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 
**alertRuleId** | **string** | alert rule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamAlertRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamAlertRule**](StreamAlertRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamAlertRuleByUuid

> StreamAlertRule GetStreamAlertRuleByUuid(ctx, streamId, alertRuleId).Execute()

Get Stream Alert Rules



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
	alertRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | alert rule UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamAlertRulesApi.GetStreamAlertRuleByUuid(context.Background(), streamId, alertRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAlertRulesApi.GetStreamAlertRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamAlertRuleByUuid`: StreamAlertRule
	fmt.Fprintf(os.Stdout, "Response from `StreamAlertRulesApi.GetStreamAlertRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 
**alertRuleId** | **string** | alert rule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamAlertRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamAlertRule**](StreamAlertRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamAlertRules

> GetAllStreamAlertRuleResponse GetStreamAlertRules(ctx, streamId).Offset(offset).Limit(limit).Execute()

Get Stream Alert Rules



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
	resp, r, err := apiClient.StreamAlertRulesApi.GetStreamAlertRules(context.Background(), streamId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAlertRulesApi.GetStreamAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamAlertRules`: GetAllStreamAlertRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamAlertRulesApi.GetStreamAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**GetAllStreamAlertRuleResponse**](GetAllStreamAlertRuleResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamAlertRuleByUuid

> StreamAlertRule UpdateStreamAlertRuleByUuid(ctx, streamId, alertRuleId).AlertRulePutRequest(alertRulePutRequest).Execute()

Update Stream Alert Rules



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
	alertRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | alert rule UUID
	alertRulePutRequest := *openapiclient.NewAlertRulePutRequest() // AlertRulePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamAlertRulesApi.UpdateStreamAlertRuleByUuid(context.Background(), streamId, alertRuleId).AlertRulePutRequest(alertRulePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAlertRulesApi.UpdateStreamAlertRuleByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStreamAlertRuleByUuid`: StreamAlertRule
	fmt.Fprintf(os.Stdout, "Response from `StreamAlertRulesApi.UpdateStreamAlertRuleByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream UUID | 
**alertRuleId** | **string** | alert rule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamAlertRuleByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertRulePutRequest** | [**AlertRulePutRequest**](AlertRulePutRequest.md) |  | 

### Return type

[**StreamAlertRule**](StreamAlertRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


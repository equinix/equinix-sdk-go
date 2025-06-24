# \UseApi

All URIs are relative to *https://sts.eqix.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UseCreateOidcProviderPost**](UseApi.md#UseCreateOidcProviderPost) | **Post** /use/createOidcProvider | /use/createOidcProvider
[**UseDeleteOidcProviderPost**](UseApi.md#UseDeleteOidcProviderPost) | **Post** /use/deleteOidcProvider | /use/deleteOidcProvider
[**UseListOidcProvidersPost**](UseApi.md#UseListOidcProvidersPost) | **Post** /use/listOidcProviders | /use/listOidcProviders
[**UseListSuspendedOidcProvidersPost**](UseApi.md#UseListSuspendedOidcProvidersPost) | **Post** /use/listSuspendedOidcProviders | /use/listSuspendedOidcProviders
[**UsePatchOidcProviderPost**](UseApi.md#UsePatchOidcProviderPost) | **Post** /use/patchOidcProvider | /use/patchOidcProvider
[**UseResumeOidcProviderPost**](UseApi.md#UseResumeOidcProviderPost) | **Post** /use/resumeOidcProvider | /use/resumeOidcProvider
[**UseSuspendOidcProviderPost**](UseApi.md#UseSuspendOidcProviderPost) | **Post** /use/suspendOidcProvider | /use/suspendOidcProvider



## UseCreateOidcProviderPost

> OpsCreateOidcProviderPost200Response UseCreateOidcProviderPost(ctx).UseCreateOidcProviderPostRequest(useCreateOidcProviderPostRequest).Execute()

/use/createOidcProvider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	useCreateOidcProviderPostRequest := *openapiclient.NewUseCreateOidcProviderPostRequest("Name_example", []string{"TrustedClientIds_example"}, "IssuerLocation_example", "IdpPrefix_example") // UseCreateOidcProviderPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseCreateOidcProviderPost(context.Background()).UseCreateOidcProviderPostRequest(useCreateOidcProviderPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseCreateOidcProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseCreateOidcProviderPost`: OpsCreateOidcProviderPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseCreateOidcProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseCreateOidcProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useCreateOidcProviderPostRequest** | [**UseCreateOidcProviderPostRequest**](UseCreateOidcProviderPostRequest.md) |  | 

### Return type

[**OpsCreateOidcProviderPost200Response**](OpsCreateOidcProviderPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseDeleteOidcProviderPost

> UseDeleteOidcProviderPost200Response UseDeleteOidcProviderPost(ctx).UseDeleteOidcProviderPostRequest(useDeleteOidcProviderPostRequest).Execute()

/use/deleteOidcProvider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	useDeleteOidcProviderPostRequest := *openapiclient.NewUseDeleteOidcProviderPostRequest("IdpId_example") // UseDeleteOidcProviderPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseDeleteOidcProviderPost(context.Background()).UseDeleteOidcProviderPostRequest(useDeleteOidcProviderPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseDeleteOidcProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseDeleteOidcProviderPost`: UseDeleteOidcProviderPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseDeleteOidcProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseDeleteOidcProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useDeleteOidcProviderPostRequest** | [**UseDeleteOidcProviderPostRequest**](UseDeleteOidcProviderPostRequest.md) |  | 

### Return type

[**UseDeleteOidcProviderPost200Response**](UseDeleteOidcProviderPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseListOidcProvidersPost

> OpsListOidcProvidersPost200Response UseListOidcProvidersPost(ctx).RequestBody(requestBody).Execute()

/use/listOidcProviders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseListOidcProvidersPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseListOidcProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseListOidcProvidersPost`: OpsListOidcProvidersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseListOidcProvidersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseListOidcProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**OpsListOidcProvidersPost200Response**](OpsListOidcProvidersPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseListSuspendedOidcProvidersPost

> OpsListOidcProvidersPost200Response UseListSuspendedOidcProvidersPost(ctx).RequestBody(requestBody).Execute()

/use/listSuspendedOidcProviders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseListSuspendedOidcProvidersPost(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseListSuspendedOidcProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseListSuspendedOidcProvidersPost`: OpsListOidcProvidersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseListSuspendedOidcProvidersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseListSuspendedOidcProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**OpsListOidcProvidersPost200Response**](OpsListOidcProvidersPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsePatchOidcProviderPost

> OpsCreateOidcProviderPost200Response UsePatchOidcProviderPost(ctx).UsePatchOidcProviderPostRequest(usePatchOidcProviderPostRequest).Execute()

/use/patchOidcProvider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	usePatchOidcProviderPostRequest := *openapiclient.NewUsePatchOidcProviderPostRequest("IdpId_example", "LastRev_example") // UsePatchOidcProviderPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UsePatchOidcProviderPost(context.Background()).UsePatchOidcProviderPostRequest(usePatchOidcProviderPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UsePatchOidcProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsePatchOidcProviderPost`: OpsCreateOidcProviderPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UsePatchOidcProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsePatchOidcProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usePatchOidcProviderPostRequest** | [**UsePatchOidcProviderPostRequest**](UsePatchOidcProviderPostRequest.md) |  | 

### Return type

[**OpsCreateOidcProviderPost200Response**](OpsCreateOidcProviderPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseResumeOidcProviderPost

> UseDeleteOidcProviderPost200Response UseResumeOidcProviderPost(ctx).UseDeleteOidcProviderPostRequest(useDeleteOidcProviderPostRequest).Execute()

/use/resumeOidcProvider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	useDeleteOidcProviderPostRequest := *openapiclient.NewUseDeleteOidcProviderPostRequest("IdpId_example") // UseDeleteOidcProviderPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseResumeOidcProviderPost(context.Background()).UseDeleteOidcProviderPostRequest(useDeleteOidcProviderPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseResumeOidcProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseResumeOidcProviderPost`: UseDeleteOidcProviderPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseResumeOidcProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseResumeOidcProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useDeleteOidcProviderPostRequest** | [**UseDeleteOidcProviderPostRequest**](UseDeleteOidcProviderPostRequest.md) |  | 

### Return type

[**UseDeleteOidcProviderPost200Response**](UseDeleteOidcProviderPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseSuspendOidcProviderPost

> UseDeleteOidcProviderPost200Response UseSuspendOidcProviderPost(ctx).UseDeleteOidcProviderPostRequest(useDeleteOidcProviderPostRequest).Execute()

/use/suspendOidcProvider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1"
)

func main() {
	useDeleteOidcProviderPostRequest := *openapiclient.NewUseDeleteOidcProviderPostRequest("IdpId_example") // UseDeleteOidcProviderPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseSuspendOidcProviderPost(context.Background()).UseDeleteOidcProviderPostRequest(useDeleteOidcProviderPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseSuspendOidcProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseSuspendOidcProviderPost`: UseDeleteOidcProviderPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseSuspendOidcProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseSuspendOidcProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useDeleteOidcProviderPostRequest** | [**UseDeleteOidcProviderPostRequest**](UseDeleteOidcProviderPostRequest.md) |  | 

### Return type

[**UseDeleteOidcProviderPost200Response**](UseDeleteOidcProviderPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


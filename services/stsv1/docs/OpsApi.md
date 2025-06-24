# \OpsApi

All URIs are relative to *https://sts.eqix.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpsCreateOidcProviderPost**](OpsApi.md#OpsCreateOidcProviderPost) | **Post** /ops/createOidcProvider | /ops/createOidcProvider
[**OpsListOidcProvidersPost**](OpsApi.md#OpsListOidcProvidersPost) | **Post** /ops/listOidcProviders | /ops/listOidcProviders



## OpsCreateOidcProviderPost

> OpsCreateOidcProviderPost200Response OpsCreateOidcProviderPost(ctx).OpsCreateOidcProviderPostRequest(opsCreateOidcProviderPostRequest).Execute()

/ops/createOidcProvider



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
	opsCreateOidcProviderPostRequest := *openapiclient.NewOpsCreateOidcProviderPostRequest("Name_example", []string{"TrustedClientIds_example"}, "IssuerUri_example", "IssuerLocation_example", *openapiclient.NewOpsCreateOidcProviderPostRequestJwks([]map[string]interface{}{map[string]interface{}(123)}), "IdpPrefix_example", "ProjectId_example") // OpsCreateOidcProviderPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsApi.OpsCreateOidcProviderPost(context.Background()).OpsCreateOidcProviderPostRequest(opsCreateOidcProviderPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsApi.OpsCreateOidcProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsCreateOidcProviderPost`: OpsCreateOidcProviderPost200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsApi.OpsCreateOidcProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsCreateOidcProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsCreateOidcProviderPostRequest** | [**OpsCreateOidcProviderPostRequest**](OpsCreateOidcProviderPostRequest.md) |  | 

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


## OpsListOidcProvidersPost

> OpsListOidcProvidersPost200Response OpsListOidcProvidersPost(ctx).OpsListOidcProvidersPostRequest(opsListOidcProvidersPostRequest).Execute()

/ops/listOidcProviders



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
	opsListOidcProvidersPostRequest := *openapiclient.NewOpsListOidcProvidersPostRequest("ProjectId_example") // OpsListOidcProvidersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsApi.OpsListOidcProvidersPost(context.Background()).OpsListOidcProvidersPostRequest(opsListOidcProvidersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsApi.OpsListOidcProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpsListOidcProvidersPost`: OpsListOidcProvidersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `OpsApi.OpsListOidcProvidersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpsListOidcProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opsListOidcProvidersPostRequest** | [**OpsListOidcProvidersPostRequest**](OpsListOidcProvidersPostRequest.md) |  | 

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


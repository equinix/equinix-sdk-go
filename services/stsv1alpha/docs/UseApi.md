# \UseApi

All URIs are relative to *https://sts.eqix.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UseTokenPost**](UseApi.md#UseTokenPost) | **Post** /use/token | 



## UseTokenPost

> TokenExchangeResponse UseTokenPost(ctx).GrantType(grantType).Scope(scope).SubjectToken(subjectToken).SubjectTokenType(subjectTokenType).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/stsv1alpha"
)

func main() {
	grantType := openapiclient._use_token_post_request_grant_type("client_credentials") // UseTokenPostRequestGrantType | 
	scope := "scope_example" // string |  (optional)
	subjectToken := "subjectToken_example" // string |  (optional)
	subjectTokenType := openapiclient._use_token_post_request_subject_token_type("urn:ietf:params:oauth:token-type:id_token") // UseTokenPostRequestSubjectTokenType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseApi.UseTokenPost(context.Background()).GrantType(grantType).Scope(scope).SubjectToken(subjectToken).SubjectTokenType(subjectTokenType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseApi.UseTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseTokenPost`: TokenExchangeResponse
	fmt.Fprintf(os.Stdout, "Response from `UseApi.UseTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | [**UseTokenPostRequestGrantType**](UseTokenPostRequestGrantType.md) |  | 
 **scope** | **string** |  | 
 **subjectToken** | **string** |  | 
 **subjectTokenType** | [**UseTokenPostRequestSubjectTokenType**](UseTokenPostRequestSubjectTokenType.md) |  | 

### Return type

[**TokenExchangeResponse**](TokenExchangeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


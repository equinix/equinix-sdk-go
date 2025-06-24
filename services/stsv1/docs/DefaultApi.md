# \DefaultApi

All URIs are relative to *https://sts.eqix.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UseTokenPost**](DefaultApi.md#UseTokenPost) | **Post** /use/token | 



## UseTokenPost

> UseTokenPost200Response UseTokenPost(ctx).GrantType(grantType).Scope(scope).SubjectToken(subjectToken).SubjectTokenType(subjectTokenType).Execute()





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
	grantType := openapiclient._use_token_post_request_grant_type("client_credentials") // UseTokenPostRequestGrantType | 
	scope := "scope_example" // string |  (optional)
	subjectToken := "subjectToken_example" // string |  (optional)
	subjectTokenType := openapiclient._use_token_post_request_subject_token_type("urn:ietf:params:oauth:token-type:id_token") // UseTokenPostRequestSubjectTokenType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.UseTokenPost(context.Background()).GrantType(grantType).Scope(scope).SubjectToken(subjectToken).SubjectTokenType(subjectTokenType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UseTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseTokenPost`: UseTokenPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UseTokenPost`: %v\n", resp)
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

[**UseTokenPost200Response**](UseTokenPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


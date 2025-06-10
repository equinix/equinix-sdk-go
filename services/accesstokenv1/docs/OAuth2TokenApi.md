# \OAuth2TokenApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOAuth2AccessToken**](OAuth2TokenApi.md#GetOAuth2AccessToken) | **Post** /oauth2/v1/token | Generate New Access Token
[**RefreshOAuth2AccessToken**](OAuth2TokenApi.md#RefreshOAuth2AccessToken) | **Post** /oauth2/v1/refreshaccesstoken | Renew Access Tokens



## GetOAuth2AccessToken

> Oauth2TokenResponse GetOAuth2AccessToken(ctx).Payload(payload).Execute()

Generate New Access Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/accesstokenv1"
)

func main() {
	payload := *openapiclient.NewOauth2TokenRequest("ClientId_example", "ClientSecret_example") // Oauth2TokenRequest | Request token json payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2TokenApi.GetOAuth2AccessToken(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2TokenApi.GetOAuth2AccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuth2AccessToken`: Oauth2TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2TokenApi.GetOAuth2AccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2AccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**Oauth2TokenRequest**](Oauth2TokenRequest.md) | Request token json payload | 

### Return type

[**Oauth2TokenResponse**](Oauth2TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshOAuth2AccessToken

> Oauth2TokenResponse RefreshOAuth2AccessToken(ctx).Payload(payload).Execute()

Renew Access Tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/accesstokenv1"
)

func main() {
	payload := *openapiclient.NewOauth2RefreshTokenRequest("ClientId_example", "ClientSecret_example", "RefreshToken_example") // Oauth2RefreshTokenRequest | Refresh token json payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2TokenApi.RefreshOAuth2AccessToken(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2TokenApi.RefreshOAuth2AccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshOAuth2AccessToken`: Oauth2TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2TokenApi.RefreshOAuth2AccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshOAuth2AccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**Oauth2RefreshTokenRequest**](Oauth2RefreshTokenRequest.md) | Refresh token json payload. | 

### Return type

[**Oauth2TokenResponse**](Oauth2TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


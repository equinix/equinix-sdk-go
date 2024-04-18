# \Layer2ConnectionsServiceProfilesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionUsingPOST**](Layer2ConnectionsServiceProfilesApi.md#CreateConnectionUsingPOST) | **Post** /ne/v1/l2/connections | Create L2 Connection
[**DeleteConnectionUsingDELETE**](Layer2ConnectionsServiceProfilesApi.md#DeleteConnectionUsingDELETE) | **Delete** /ne/v1/l2/connections/{uuid} | Delete Connection {uuid}
[**GetConnectionByUuidUsingGET**](Layer2ConnectionsServiceProfilesApi.md#GetConnectionByUuidUsingGET) | **Get** /ne/v1/l2/connections/{uuid} | Get L2 Connection {uuid}
[**GetConnectionsByGET**](Layer2ConnectionsServiceProfilesApi.md#GetConnectionsByGET) | **Get** /ne/v1/l2/buyer/connections | Get L2 Buyer Connections
[**GetProfileByIdUsingGET**](Layer2ConnectionsServiceProfilesApi.md#GetProfileByIdUsingGET) | **Get** /ne/v1/l2/serviceprofiles/services/{uuid} | Get L2 Service Profile {uuid}
[**GetProfilesByMetroUsingGET**](Layer2ConnectionsServiceProfilesApi.md#GetProfilesByMetroUsingGET) | **Get** /ne/v1/l2/serviceprofiles/services | Get L2 Service Profiles
[**PerformUserActionUsingPATCH**](Layer2ConnectionsServiceProfilesApi.md#PerformUserActionUsingPATCH) | **Patch** /ne/v1/l2/connections/{uuid} | Update Connection {uuid}
[**ValidateAuthorizationKeyUsingGET**](Layer2ConnectionsServiceProfilesApi.md#ValidateAuthorizationKeyUsingGET) | **Get** /ne/v1/l2/connections/validateAuthorizationKey | Get L2 Validate Authorization Key



## CreateConnectionUsingPOST

> PostConnectionResponse CreateConnectionUsingPOST(ctx).Authorization(authorization).Request(request).Execute()

Create L2 Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewPostConnectionRequest() // PostConnectionRequest | request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.CreateConnectionUsingPOST(context.Background()).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.CreateConnectionUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionUsingPOST`: PostConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.CreateConnectionUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**PostConnectionRequest**](PostConnectionRequest.md) | request | 

### Return type

[**PostConnectionResponse**](PostConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionUsingDELETE

> DeleteConnectionResponse DeleteConnectionUsingDELETE(ctx, uuid).Authorization(authorization).Execute()

Delete Connection {uuid}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	uuid := "uuid_example" // string | connId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.DeleteConnectionUsingDELETE(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.DeleteConnectionUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnectionUsingDELETE`: DeleteConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.DeleteConnectionUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | connId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 


### Return type

[**DeleteConnectionResponse**](DeleteConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionByUuidUsingGET

> GETConnectionByUuidResponse GetConnectionByUuidUsingGET(ctx, uuid).Authorization(authorization).Execute()

Get L2 Connection {uuid}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	uuid := "uuid_example" // string | Connection UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.GetConnectionByUuidUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.GetConnectionByUuidUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionByUuidUsingGET`: GETConnectionByUuidResponse
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.GetConnectionByUuidUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionByUuidUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 


### Return type

[**GETConnectionByUuidResponse**](GETConnectionByUuidResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsByGET

> GETConnectionsPageResponse GetConnectionsByGET(ctx).Authorization(authorization).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get L2 Buyer Connections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	pageNumber := int32(56) // int32 | page number (optional) (default to 0)
	pageSize := int32(56) // int32 | total number of records (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.GetConnectionsByGET(context.Background()).Authorization(authorization).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.GetConnectionsByGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsByGET`: GETConnectionsPageResponse
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.GetConnectionsByGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsByGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **pageNumber** | **int32** | page number | [default to 0]
 **pageSize** | **int32** | total number of records | [default to 20]

### Return type

[**GETConnectionsPageResponse**](GETConnectionsPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileByIdUsingGET

> GetServProfServicesResp GetProfileByIdUsingGET(ctx, uuid).Authorization(authorization).Execute()

Get L2 Service Profile {uuid}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	uuid := "uuid_example" // string | uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.GetProfileByIdUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.GetProfileByIdUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileByIdUsingGET`: GetServProfServicesResp
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.GetProfileByIdUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileByIdUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 


### Return type

[**GetServProfServicesResp**](GetServProfServicesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfilesByMetroUsingGET

> GetServProfServicesResp GetProfilesByMetroUsingGET(ctx).Authorization(authorization).MetroCode(metroCode).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get L2 Service Profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	metroCode := []string{"Inner_example"} // []string | metroCode (optional)
	pageNumber := int32(56) // int32 | page number (optional) (default to 0)
	pageSize := int32(56) // int32 | total number of records (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.GetProfilesByMetroUsingGET(context.Background()).Authorization(authorization).MetroCode(metroCode).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.GetProfilesByMetroUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfilesByMetroUsingGET`: GetServProfServicesResp
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.GetProfilesByMetroUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfilesByMetroUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **metroCode** | **[]string** | metroCode | 
 **pageNumber** | **int32** | page number | [default to 0]
 **pageSize** | **int32** | total number of records | [default to 20]

### Return type

[**GetServProfServicesResp**](GetServProfServicesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformUserActionUsingPATCH

> DeleteConnectionResponse PerformUserActionUsingPATCH(ctx, uuid).Authorization(authorization).Action(action).Request(request).Execute()

Update Connection {uuid}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	action := "action_example" // string | action
	uuid := "uuid_example" // string | connId
	request := *openapiclient.NewPatchRequest() // PatchRequest | request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.PerformUserActionUsingPATCH(context.Background(), uuid).Authorization(authorization).Action(action).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.PerformUserActionUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformUserActionUsingPATCH`: DeleteConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.PerformUserActionUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | connId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformUserActionUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **action** | **string** | action | 

 **request** | [**PatchRequest**](PatchRequest.md) | request | 

### Return type

[**DeleteConnectionResponse**](DeleteConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAuthorizationKeyUsingGET

> GetValidateAuthKeyRes ValidateAuthorizationKeyUsingGET(ctx).Authorization(authorization).AuthorizationKey(authorizationKey).MetroCode(metroCode).ProfileId(profileId).Region(region).ContentType(contentType).Execute()

Get L2 Validate Authorization Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	authorizationKey := "authorizationKey_example" // string | Could be AWS Account ID or Azure Service Key or Google Pairing Key or Oracle connection Id or any other authorization key
	metroCode := "metroCode_example" // string | metroCode
	profileId := "profileId_example" // string | profileId
	region := "region_example" // string | region
	contentType := "contentType_example" // string | Content-Type (optional) (default to "application/json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Layer2ConnectionsServiceProfilesApi.ValidateAuthorizationKeyUsingGET(context.Background()).Authorization(authorization).AuthorizationKey(authorizationKey).MetroCode(metroCode).ProfileId(profileId).Region(region).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Layer2ConnectionsServiceProfilesApi.ValidateAuthorizationKeyUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAuthorizationKeyUsingGET`: GetValidateAuthKeyRes
	fmt.Fprintf(os.Stdout, "Response from `Layer2ConnectionsServiceProfilesApi.ValidateAuthorizationKeyUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAuthorizationKeyUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **authorizationKey** | **string** | Could be AWS Account ID or Azure Service Key or Google Pairing Key or Oracle connection Id or any other authorization key | 
 **metroCode** | **string** | metroCode | 
 **profileId** | **string** | profileId | 
 **region** | **string** | region | 
 **contentType** | **string** | Content-Type | [default to &quot;application/json&quot;]

### Return type

[**GetValidateAuthKeyRes**](GetValidateAuthKeyRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


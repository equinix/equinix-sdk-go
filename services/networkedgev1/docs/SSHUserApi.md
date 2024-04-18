# \SSHUserApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateDeviceUsingPOST**](SSHUserApi.md#AssociateDeviceUsingPOST) | **Post** /ne/v1/sshUsers/{sshUserUuid}/devices/{deviceUuid} | Update SSH User {uuid}
[**CreateSshUserUsingPOST**](SSHUserApi.md#CreateSshUserUsingPOST) | **Post** /ne/v1/sshUsers | Create SSH User
[**DissociateDeviceUsingDELETE**](SSHUserApi.md#DissociateDeviceUsingDELETE) | **Delete** /ne/v1/sshUsers/{sshUserUuid}/devices/{deviceUuid} | Delete SSH User
[**GetSshUserUsingGET**](SSHUserApi.md#GetSshUserUsingGET) | **Get** /ne/v1/sshUsers/{uuid} | Get SSH User {uuid}
[**GetSshUsersUsingGET**](SSHUserApi.md#GetSshUsersUsingGET) | **Get** /ne/v1/sshUsers | Get SSH Users
[**IsSshUserAvailableForCreationUsingGET**](SSHUserApi.md#IsSshUserAvailableForCreationUsingGET) | **Get** /ne/v1/sshUsers/availability | Check SSH Username Availability
[**UpdateSshUserUsingPUT**](SSHUserApi.md#UpdateSshUserUsingPUT) | **Put** /ne/v1/sshUsers/{uuid} | Update SSH User Password



## AssociateDeviceUsingPOST

> AssociateDeviceUsingPOST(ctx, sshUserUuid, deviceUuid).Authorization(authorization).Execute()

Update SSH User {uuid}



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
	sshUserUuid := "sshUserUuid_example" // string | The id of the SSH user
	deviceUuid := "deviceUuid_example" // string | The unique Id of the device
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SSHUserApi.AssociateDeviceUsingPOST(context.Background(), sshUserUuid, deviceUuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.AssociateDeviceUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshUserUuid** | **string** | The id of the SSH user | 
**deviceUuid** | **string** | The unique Id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateDeviceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSshUserUsingPOST

> SshUserCreateResponse CreateSshUserUsingPOST(ctx).Authorization(authorization).Request(request).Execute()

Create SSH User



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
	request := *openapiclient.NewSshUserCreateRequest("user1", "pass12", "3da0a663-20d9-4b8f-8c5d-d5cf706840c8") // SshUserCreateRequest | SSH user info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHUserApi.CreateSshUserUsingPOST(context.Background()).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.CreateSshUserUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSshUserUsingPOST`: SshUserCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `SSHUserApi.CreateSshUserUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSshUserUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**SshUserCreateRequest**](SshUserCreateRequest.md) | SSH user info | 

### Return type

[**SshUserCreateResponse**](SshUserCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DissociateDeviceUsingDELETE

> SshUserInfoDissociateResponse DissociateDeviceUsingDELETE(ctx, sshUserUuid, deviceUuid).Authorization(authorization).Execute()

Delete SSH User



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
	sshUserUuid := "sshUserUuid_example" // string | uuid
	deviceUuid := "deviceUuid_example" // string | device unique Id
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHUserApi.DissociateDeviceUsingDELETE(context.Background(), sshUserUuid, deviceUuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.DissociateDeviceUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DissociateDeviceUsingDELETE`: SshUserInfoDissociateResponse
	fmt.Fprintf(os.Stdout, "Response from `SSHUserApi.DissociateDeviceUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshUserUuid** | **string** | uuid | 
**deviceUuid** | **string** | device unique Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateDeviceUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**SshUserInfoDissociateResponse**](SshUserInfoDissociateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshUserUsingGET

> SshUserInfoVerbose GetSshUserUsingGET(ctx, uuid).Authorization(authorization).Execute()

Get SSH User {uuid}



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
	uuid := "uuid_example" // string | uuid
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHUserApi.GetSshUserUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.GetSshUserUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshUserUsingGET`: SshUserInfoVerbose
	fmt.Fprintf(os.Stdout, "Response from `SSHUserApi.GetSshUserUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshUserUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**SshUserInfoVerbose**](SshUserInfoVerbose.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshUsersUsingGET

> SshUserPageResponse GetSshUsersUsingGET(ctx).Authorization(authorization).Username(username).VirtualDeviceUuid(virtualDeviceUuid).Verbose(verbose).AccountUcmId(accountUcmId).Offset(offset).Limit(limit).Execute()

Get SSH Users



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
	username := "username_example" // string | SSH user name (optional)
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Virtual Device unique Id (optional)
	verbose := true // bool | Is detailed info required (optional)
	accountUcmId := "accountUcmId_example" // string | Unique ID of an account (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHUserApi.GetSshUsersUsingGET(context.Background()).Authorization(authorization).Username(username).VirtualDeviceUuid(virtualDeviceUuid).Verbose(verbose).AccountUcmId(accountUcmId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.GetSshUsersUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshUsersUsingGET`: SshUserPageResponse
	fmt.Fprintf(os.Stdout, "Response from `SSHUserApi.GetSshUsersUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSshUsersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **username** | **string** | SSH user name | 
 **virtualDeviceUuid** | **string** | Virtual Device unique Id | 
 **verbose** | **bool** | Is detailed info required | 
 **accountUcmId** | **string** | Unique ID of an account | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**SshUserPageResponse**](SshUserPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsSshUserAvailableForCreationUsingGET

> IsSshUserAvailableForCreationUsingGET(ctx).Username(username).Authorization(authorization).Execute()

Check SSH Username Availability



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
	username := "username_example" // string | user name
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SSHUserApi.IsSshUserAvailableForCreationUsingGET(context.Background()).Username(username).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.IsSshUserAvailableForCreationUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsSshUserAvailableForCreationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | user name | 
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSshUserUsingPUT

> UpdateSshUserUsingPUT(ctx, uuid).Authorization(authorization).Request(request).Execute()

Update SSH User Password



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
	uuid := "uuid_example" // string | uuid
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewSshUserUpdateRequest("pass12") // SshUserUpdateRequest | SSH user info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SSHUserApi.UpdateSshUserUsingPUT(context.Background(), uuid).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHUserApi.UpdateSshUserUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSshUserUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**SshUserUpdateRequest**](SshUserUpdateRequest.md) | SSH user info | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


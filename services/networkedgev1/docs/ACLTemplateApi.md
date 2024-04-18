# \ACLTemplateApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceACLTemplateUsingPost**](ACLTemplateApi.md#CreateDeviceACLTemplateUsingPost) | **Post** /ne/v1/aclTemplates | Create ACL Template
[**DeletedeviceACLUsingDELETE**](ACLTemplateApi.md#DeletedeviceACLUsingDELETE) | **Delete** /ne/v1/aclTemplates/{uuid} | Delete ACL template
[**GetDeviceACLTemplateUsingGET1**](ACLTemplateApi.md#GetDeviceACLTemplateUsingGET1) | **Get** /ne/v1/aclTemplates | Get ACL Templates
[**GetDeviceTemplatebyUuid**](ACLTemplateApi.md#GetDeviceTemplatebyUuid) | **Get** /ne/v1/aclTemplates/{uuid} | Get ACL Template {uuid}
[**GetDeviceTemplatesbyUuid**](ACLTemplateApi.md#GetDeviceTemplatesbyUuid) | **Get** /ne/v1/devices/{virtualDeviceUuid}/acl | Get ACL of Virtual Device
[**PatchDeviceTemplatesbyUuid**](ACLTemplateApi.md#PatchDeviceTemplatesbyUuid) | **Patch** /ne/v1/devices/{virtualDeviceUuid}/acl | Update ACL of Virtual Device
[**PostDeviceTemplatesbyUuid**](ACLTemplateApi.md#PostDeviceTemplatesbyUuid) | **Post** /ne/v1/devices/{virtualDeviceUuid}/acl | Add ACL to Virtual Device
[**SendDnsLookupPOST1**](ACLTemplateApi.md#SendDnsLookupPOST1) | **Post** /ne/v1/dnsLookup | Post DNS Lookup
[**UpdateDeviceACLTemplateUsingPUT**](ACLTemplateApi.md#UpdateDeviceACLTemplateUsingPUT) | **Put** /ne/v1/aclTemplates/{uuid} | Update ACL Template



## CreateDeviceACLTemplateUsingPost

> VirtualDeviceCreateResponse CreateDeviceACLTemplateUsingPost(ctx).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()

Create ACL Template



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
	aclTemplateRequest := *openapiclient.NewDeviceACLTemplateRequest("Test Template", "Test Template Description", []openapiclient.InboundRules{*openapiclient.NewInboundRules()}) // DeviceACLTemplateRequest | Creates an ACL template.
	accountUcmId := "accountUcmId_example" // string | A reseller creating an ACL template for a customer can pass the accountUcmId of the customer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateApi.CreateDeviceACLTemplateUsingPost(context.Background()).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.CreateDeviceACLTemplateUsingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceACLTemplateUsingPost`: VirtualDeviceCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateApi.CreateDeviceACLTemplateUsingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceACLTemplateUsingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **aclTemplateRequest** | [**DeviceACLTemplateRequest**](DeviceACLTemplateRequest.md) | Creates an ACL template. | 
 **accountUcmId** | **string** | A reseller creating an ACL template for a customer can pass the accountUcmId of the customer. | 

### Return type

[**VirtualDeviceCreateResponse**](VirtualDeviceCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletedeviceACLUsingDELETE

> DeletedeviceACLUsingDELETE(ctx, uuid).Authorization(authorization).AccountUcmId(accountUcmId).Execute()

Delete ACL template



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
	uuid := "uuid_example" // string | Unique ID of an ACL template
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	accountUcmId := "accountUcmId_example" // string | A reseller deleting an ACL template for a customer must pass the accountUcmId of the customer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ACLTemplateApi.DeletedeviceACLUsingDELETE(context.Background(), uuid).Authorization(authorization).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.DeletedeviceACLUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of an ACL template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletedeviceACLUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountUcmId** | **string** | A reseller deleting an ACL template for a customer must pass the accountUcmId of the customer. | 

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


## GetDeviceACLTemplateUsingGET1

> DeviceACLPageResponse GetDeviceACLTemplateUsingGET1(ctx).Authorization(authorization).Offset(offset).Limit(limit).AccountUcmId(accountUcmId).Execute()

Get ACL Templates



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
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")
	accountUcmId := "accountUcmId_example" // string | Unique ID of the account. A reseller querying for the ACLs of a customer should input the accountUcmId of the customer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateApi.GetDeviceACLTemplateUsingGET1(context.Background()).Authorization(authorization).Offset(offset).Limit(limit).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.GetDeviceACLTemplateUsingGET1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceACLTemplateUsingGET1`: DeviceACLPageResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateApi.GetDeviceACLTemplateUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceACLTemplateUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]
 **accountUcmId** | **string** | Unique ID of the account. A reseller querying for the ACLs of a customer should input the accountUcmId of the customer. | 

### Return type

[**DeviceACLPageResponse**](DeviceACLPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceTemplatebyUuid

> ACLTemplateDetailsResponse GetDeviceTemplatebyUuid(ctx, uuid).Authorization(authorization).Execute()

Get ACL Template {uuid}



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
	uuid := "uuid_example" // string | Unique Id of an ACL template
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateApi.GetDeviceTemplatebyUuid(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.GetDeviceTemplatebyUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceTemplatebyUuid`: ACLTemplateDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateApi.GetDeviceTemplatebyUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of an ACL template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTemplatebyUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**ACLTemplateDetailsResponse**](ACLTemplateDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceTemplatesbyUuid

> InitialDeviceACLResponse GetDeviceTemplatesbyUuid(ctx, virtualDeviceUuid).Authorization(authorization).Execute()

Get ACL of Virtual Device



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
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a virtual device.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateApi.GetDeviceTemplatesbyUuid(context.Background(), virtualDeviceUuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.GetDeviceTemplatesbyUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceTemplatesbyUuid`: InitialDeviceACLResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateApi.GetDeviceTemplatesbyUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUuid** | **string** | Unique Id of a virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTemplatesbyUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**InitialDeviceACLResponse**](InitialDeviceACLResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDeviceTemplatesbyUuid

> PatchDeviceTemplatesbyUuid(ctx, virtualDeviceUuid).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()

Update ACL of Virtual Device



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
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a virtual device.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	aclTemplateRequest := *openapiclient.NewUpdateDeviceACLTemplateRequest([]openapiclient.ACLDetails{*openapiclient.NewACLDetails()}) // UpdateDeviceACLTemplateRequest | Update the ACL of a device.
	accountUcmId := "accountUcmId_example" // string | A reseller updating a device ACL template for a customer can pass the accountUcmId of the customer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ACLTemplateApi.PatchDeviceTemplatesbyUuid(context.Background(), virtualDeviceUuid).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.PatchDeviceTemplatesbyUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUuid** | **string** | Unique Id of a virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDeviceTemplatesbyUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **aclTemplateRequest** | [**UpdateDeviceACLTemplateRequest**](UpdateDeviceACLTemplateRequest.md) | Update the ACL of a device. | 
 **accountUcmId** | **string** | A reseller updating a device ACL template for a customer can pass the accountUcmId of the customer. | 

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


## PostDeviceTemplatesbyUuid

> PostDeviceTemplatesbyUuid(ctx, virtualDeviceUuid).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()

Add ACL to Virtual Device



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
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a virtual device.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	aclTemplateRequest := *openapiclient.NewUpdateDeviceACLTemplateRequest([]openapiclient.ACLDetails{*openapiclient.NewACLDetails()}) // UpdateDeviceACLTemplateRequest | Update the ACL of a device.
	accountUcmId := "accountUcmId_example" // string | A reseller updating a device ACL template for a customer can pass the accountUcmId of the customer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ACLTemplateApi.PostDeviceTemplatesbyUuid(context.Background(), virtualDeviceUuid).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.PostDeviceTemplatesbyUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUuid** | **string** | Unique Id of a virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceTemplatesbyUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **aclTemplateRequest** | [**UpdateDeviceACLTemplateRequest**](UpdateDeviceACLTemplateRequest.md) | Update the ACL of a device. | 
 **accountUcmId** | **string** | A reseller updating a device ACL template for a customer can pass the accountUcmId of the customer. | 

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


## SendDnsLookupPOST1

> DNSLookupResponse SendDnsLookupPOST1(ctx).Authorization(authorization).DnsLookupRequest(dnsLookupRequest).Execute()

Post DNS Lookup



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
	dnsLookupRequest := *openapiclient.NewDNSLookupRequest() // DNSLookupRequest | dnsLookupRequest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateApi.SendDnsLookupPOST1(context.Background()).Authorization(authorization).DnsLookupRequest(dnsLookupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.SendDnsLookupPOST1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendDnsLookupPOST1`: DNSLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateApi.SendDnsLookupPOST1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDnsLookupPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **dnsLookupRequest** | [**DNSLookupRequest**](DNSLookupRequest.md) | dnsLookupRequest | 

### Return type

[**DNSLookupResponse**](DNSLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceACLTemplateUsingPUT

> UpdateDeviceACLTemplateUsingPUT(ctx, uuid).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()

Update ACL Template



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
	uuid := "uuid_example" // string | Unique ID of an ACL template
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	aclTemplateRequest := *openapiclient.NewDeviceACLTemplateRequest("Test Template", "Test Template Description", []openapiclient.InboundRules{*openapiclient.NewInboundRules()}) // DeviceACLTemplateRequest | Update an ACL template.
	accountUcmId := "accountUcmId_example" // string | A reseller updating an ACL template for a customer must pass the accountUcmId of the customer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ACLTemplateApi.UpdateDeviceACLTemplateUsingPUT(context.Background(), uuid).Authorization(authorization).AclTemplateRequest(aclTemplateRequest).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateApi.UpdateDeviceACLTemplateUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of an ACL template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceACLTemplateUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **aclTemplateRequest** | [**DeviceACLTemplateRequest**](DeviceACLTemplateRequest.md) | Update an ACL template. | 
 **accountUcmId** | **string** | A reseller updating an ACL template for a customer must pass the accountUcmId of the customer. | 

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


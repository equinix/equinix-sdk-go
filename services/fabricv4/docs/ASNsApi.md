# \ASNsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsn**](ASNsApi.md#CreateAsn) | **Post** /fabric/v4/asns | Create ASN
[**DeleteAsnById**](ASNsApi.md#DeleteAsnById) | **Delete** /fabric/v4/asns/{asnId} | Delete ASN
[**GetAsnById**](ASNsApi.md#GetAsnById) | **Get** /fabric/v4/asns/{asnId} | Get ASN
[**PatchAsnById**](ASNsApi.md#PatchAsnById) | **Patch** /fabric/v4/asns/{asnId} | Update ASN
[**SearchAsns**](ASNsApi.md#SearchAsns) | **Post** /fabric/v4/asns/search | Search ASNs



## CreateAsn

> Asn CreateAsn(ctx).CreateAsnRequest(createAsnRequest).Execute()

Create ASN



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
	createAsnRequest := *openapiclient.NewCreateAsnRequest(*openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5")) // CreateAsnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNsApi.CreateAsn(context.Background()).CreateAsnRequest(createAsnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNsApi.CreateAsn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAsn`: Asn
	fmt.Fprintf(os.Stdout, "Response from `ASNsApi.CreateAsn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAsnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAsnRequest** | [**CreateAsnRequest**](CreateAsnRequest.md) |  | 

### Return type

[**Asn**](Asn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAsnById

> Asn DeleteAsnById(ctx, asnId).Execute()

Delete ASN



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
	asnId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the ASN record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNsApi.DeleteAsnById(context.Background(), asnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNsApi.DeleteAsnById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAsnById`: Asn
	fmt.Fprintf(os.Stdout, "Response from `ASNsApi.DeleteAsnById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asnId** | **string** | UUID of the ASN record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAsnByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Asn**](Asn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsnById

> Asn GetAsnById(ctx, asnId).Execute()

Get ASN



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
	asnId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the ASN record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNsApi.GetAsnById(context.Background(), asnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNsApi.GetAsnById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsnById`: Asn
	fmt.Fprintf(os.Stdout, "Response from `ASNsApi.GetAsnById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asnId** | **string** | UUID of the ASN record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsnByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Asn**](Asn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAsnById

> Asn PatchAsnById(ctx, asnId).AsnPatchOperation(asnPatchOperation).Execute()

Update ASN



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
	asnId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the ASN record
	asnPatchOperation := []openapiclient.AsnPatchOperation{*openapiclient.NewAsnPatchOperation(openapiclient.AsnPatchOperation_op("add"), "/ownershipChallengeDetails/signature", interface{}(MEUCIQCrExampleBase64SignatureValue7h4Yg==))} // []AsnPatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNsApi.PatchAsnById(context.Background(), asnId).AsnPatchOperation(asnPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNsApi.PatchAsnById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAsnById`: Asn
	fmt.Fprintf(os.Stdout, "Response from `ASNsApi.PatchAsnById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asnId** | **string** | UUID of the ASN record | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAsnByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asnPatchOperation** | [**[]AsnPatchOperation**](AsnPatchOperation.md) |  | 

### Return type

[**Asn**](Asn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAsns

> AsnSearchResponseBody SearchAsns(ctx).AsnSearchRequestBody(asnSearchRequestBody).Execute()

Search ASNs



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
	asnSearchRequestBody := *openapiclient.NewAsnSearchRequestBody() // AsnSearchRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNsApi.SearchAsns(context.Background()).AsnSearchRequestBody(asnSearchRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNsApi.SearchAsns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAsns`: AsnSearchResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ASNsApi.SearchAsns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAsnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asnSearchRequestBody** | [**AsnSearchRequestBody**](AsnSearchRequestBody.md) |  | 

### Return type

[**AsnSearchResponseBody**](AsnSearchResponseBody.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \LogosApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogo**](LogosApi.md#CreateLogo) | **Post** /fabric/v4/logos | Create Logo
[**DeleteLogoByUuid**](LogosApi.md#DeleteLogoByUuid) | **Delete** /fabric/v4/logos/{uuid} | Delete Logo
[**GetLogoByUuid**](LogosApi.md#GetLogoByUuid) | **Get** /fabric/v4/logos/{uuid} | Get Logo



## CreateLogo

> LogoResponse CreateLogo(ctx).Logo(logo).Name(name).Description(description).Type_(type_).Execute()

Create Logo



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
	logo := os.NewFile(1234, "some_file") // *os.File | Logo image file
	name := "name_example" // string | Name of the Logo
	description := "description_example" // string | Description of the logo
	type_ := "type__example" // string | Type of logo

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogosApi.CreateLogo(context.Background()).Logo(logo).Name(name).Description(description).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogosApi.CreateLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogo`: LogoResponse
	fmt.Fprintf(os.Stdout, "Response from `LogosApi.CreateLogo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logo** | ***os.File** | Logo image file | 
 **name** | **string** | Name of the Logo | 
 **description** | **string** | Description of the logo | 
 **type_** | **string** | Type of logo | 

### Return type

[**LogoResponse**](LogoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogoByUuid

> LogoResponse DeleteLogoByUuid(ctx, uuid).Execute()

Delete Logo



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the Logo

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogosApi.DeleteLogoByUuid(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogosApi.DeleteLogoByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLogoByUuid`: LogoResponse
	fmt.Fprintf(os.Stdout, "Response from `LogosApi.DeleteLogoByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the Logo | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogoByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogoResponse**](LogoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogoByUuid

> *os.File GetLogoByUuid(ctx, uuid).Execute()

Get Logo



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the Logo

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogosApi.GetLogoByUuid(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogosApi.GetLogoByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogoByUuid`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LogosApi.GetLogoByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the Logo | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogoByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/mixed, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


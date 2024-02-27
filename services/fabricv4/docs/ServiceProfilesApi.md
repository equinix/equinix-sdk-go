# \ServiceProfilesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceProfile**](ServiceProfilesApi.md#CreateServiceProfile) | **Post** /fabric/v4/serviceProfiles | Create Profile
[**DeleteServiceProfileByUuid**](ServiceProfilesApi.md#DeleteServiceProfileByUuid) | **Delete** /fabric/v4/serviceProfiles/{serviceProfileId} | Delete Profile
[**GetServiceProfileByUuid**](ServiceProfilesApi.md#GetServiceProfileByUuid) | **Get** /fabric/v4/serviceProfiles/{serviceProfileId} | Get Profile
[**GetServiceProfileMetrosByUuid**](ServiceProfilesApi.md#GetServiceProfileMetrosByUuid) | **Get** /fabric/v4/serviceProfiles/{serviceProfileId}/metros | Get Profile Metros
[**GetServiceProfiles**](ServiceProfilesApi.md#GetServiceProfiles) | **Get** /fabric/v4/serviceProfiles | Get all Profiles
[**PutServiceProfileByUuid**](ServiceProfilesApi.md#PutServiceProfileByUuid) | **Put** /fabric/v4/serviceProfiles/{serviceProfileId} | Replace Profile
[**SearchServiceProfiles**](ServiceProfilesApi.md#SearchServiceProfiles) | **Post** /fabric/v4/serviceProfiles/search | Profile Search
[**UpdateServiceProfileByUuid**](ServiceProfilesApi.md#UpdateServiceProfileByUuid) | **Patch** /fabric/v4/serviceProfiles/{serviceProfileId} | Update Profile



## CreateServiceProfile

> ServiceProfile CreateServiceProfile(ctx).ServiceProfileRequest(serviceProfileRequest).Execute()

Create Profile



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
	serviceProfileRequest := *openapiclient.NewServiceProfileRequest(openapiclient.ServiceProfileTypeEnum("L2_PROFILE"), "Sample Service Profile", "offering connectivity to my-network") // ServiceProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.CreateServiceProfile(context.Background()).ServiceProfileRequest(serviceProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.CreateServiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceProfile`: ServiceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.CreateServiceProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceProfileRequest** | [**ServiceProfileRequest**](ServiceProfileRequest.md) |  | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceProfileByUuid

> ServiceProfile DeleteServiceProfileByUuid(ctx, serviceProfileId).Execute()

Delete Profile



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
	serviceProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.DeleteServiceProfileByUuid(context.Background(), serviceProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.DeleteServiceProfileByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceProfileByUuid`: ServiceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.DeleteServiceProfileByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceProfileByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceProfileByUuid

> ServiceProfile GetServiceProfileByUuid(ctx, serviceProfileId).ViewPoint(viewPoint).Execute()

Get Profile



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
	serviceProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Profile UUID
	viewPoint := openapiclient.getServiceProfiles_viewPoint_parameter("aSide") // GetServiceProfilesViewPointParameter | flips view between buyer and seller representation (optional) (default to "aSide")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.GetServiceProfileByUuid(context.Background(), serviceProfileId).ViewPoint(viewPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.GetServiceProfileByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceProfileByUuid`: ServiceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.GetServiceProfileByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceProfileByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewPoint** | [**GetServiceProfilesViewPointParameter**](GetServiceProfilesViewPointParameter.md) | flips view between buyer and seller representation | [default to &quot;aSide&quot;]

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceProfileMetrosByUuid

> ServiceMetros GetServiceProfileMetrosByUuid(ctx, serviceProfileId).Offset(offset).Limit(limit).Execute()

Get Profile Metros



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
	serviceProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Profile UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.GetServiceProfileMetrosByUuid(context.Background(), serviceProfileId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.GetServiceProfileMetrosByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceProfileMetrosByUuid`: ServiceMetros
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.GetServiceProfileMetrosByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceProfileMetrosByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**ServiceMetros**](ServiceMetros.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceProfiles

> ServiceProfiles GetServiceProfiles(ctx).Offset(offset).Limit(limit).ViewPoint(viewPoint).Execute()

Get all Profiles



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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	viewPoint := openapiclient.getServiceProfiles_viewPoint_parameter("aSide") // GetServiceProfilesViewPointParameter | flips view between buyer and seller representation (optional) (default to "aSide")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.GetServiceProfiles(context.Background()).Offset(offset).Limit(limit).ViewPoint(viewPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.GetServiceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceProfiles`: ServiceProfiles
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.GetServiceProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **viewPoint** | [**GetServiceProfilesViewPointParameter**](GetServiceProfilesViewPointParameter.md) | flips view between buyer and seller representation | [default to &quot;aSide&quot;]

### Return type

[**ServiceProfiles**](ServiceProfiles.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceProfileByUuid

> ServiceProfile PutServiceProfileByUuid(ctx, serviceProfileId).IfMatch(ifMatch).ServiceProfileRequest(serviceProfileRequest).Execute()

Replace Profile



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
	serviceProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Profile UUID
	ifMatch := "ifMatch_example" // string | conditional request
	serviceProfileRequest := *openapiclient.NewServiceProfileRequest(openapiclient.ServiceProfileTypeEnum("L2_PROFILE"), "Sample Service Profile", "offering connectivity to my-network") // ServiceProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.PutServiceProfileByUuid(context.Background(), serviceProfileId).IfMatch(ifMatch).ServiceProfileRequest(serviceProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.PutServiceProfileByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceProfileByUuid`: ServiceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.PutServiceProfileByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceProfileByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | conditional request | 
 **serviceProfileRequest** | [**ServiceProfileRequest**](ServiceProfileRequest.md) |  | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServiceProfiles

> ServiceProfiles SearchServiceProfiles(ctx).ServiceProfileSearchRequest(serviceProfileSearchRequest).ViewPoint(viewPoint).Execute()

Profile Search



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
	serviceProfileSearchRequest := *openapiclient.NewServiceProfileSearchRequest() // ServiceProfileSearchRequest | 
	viewPoint := openapiclient.getServiceProfiles_viewPoint_parameter("aSide") // GetServiceProfilesViewPointParameter | flips view between buyer and seller representation (optional) (default to "aSide")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.SearchServiceProfiles(context.Background()).ServiceProfileSearchRequest(serviceProfileSearchRequest).ViewPoint(viewPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.SearchServiceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchServiceProfiles`: ServiceProfiles
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.SearchServiceProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServiceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceProfileSearchRequest** | [**ServiceProfileSearchRequest**](ServiceProfileSearchRequest.md) |  | 
 **viewPoint** | [**GetServiceProfilesViewPointParameter**](GetServiceProfilesViewPointParameter.md) | flips view between buyer and seller representation | [default to &quot;aSide&quot;]

### Return type

[**ServiceProfiles**](ServiceProfiles.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceProfileByUuid

> ServiceProfile UpdateServiceProfileByUuid(ctx, serviceProfileId).IfMatch(ifMatch).JsonPatchOperation(jsonPatchOperation).Execute()

Update Profile



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
	serviceProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service Profile UUID
	ifMatch := "ifMatch_example" // string | conditional request
	jsonPatchOperation := []openapiclient.JsonPatchOperation{openapiclient.JsonPatchOperation{AddOperation: openapiclient.NewAddOperation(openapiclient.OpEnum("add"), "Path_example", map[string]interface{}(123))}} // []JsonPatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceProfilesApi.UpdateServiceProfileByUuid(context.Background(), serviceProfileId).IfMatch(ifMatch).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceProfilesApi.UpdateServiceProfileByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceProfileByUuid`: ServiceProfile
	fmt.Fprintf(os.Stdout, "Response from `ServiceProfilesApi.UpdateServiceProfileByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceProfileByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | conditional request | 
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


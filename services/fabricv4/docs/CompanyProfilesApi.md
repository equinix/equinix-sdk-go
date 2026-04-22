# \CompanyProfilesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachPrivateServiceToProfile**](CompanyProfilesApi.md#AttachPrivateServiceToProfile) | **Put** /fabric/v4/companyProfiles/{companyProfileId}/privateServices/{privateServiceId} | Attach Private Service
[**AttachServiceProfileToProfile**](CompanyProfilesApi.md#AttachServiceProfileToProfile) | **Put** /fabric/v4/companyProfiles/{companyProfileId}/serviceProfiles/{serviceProfileId} | Attach Service Profile
[**AttachTagToProfile**](CompanyProfilesApi.md#AttachTagToProfile) | **Put** /fabric/v4/companyProfiles/{companyProfileId}/tags/{tagId} | Attach Tag
[**CreateCompanyProfile**](CompanyProfilesApi.md#CreateCompanyProfile) | **Post** /fabric/v4/companyProfiles | Create Company Profile
[**DeleteCompanyProfile**](CompanyProfilesApi.md#DeleteCompanyProfile) | **Delete** /fabric/v4/companyProfiles/{companyProfileId} | Delete Company Profile
[**DetachPrivateServiceFromProfile**](CompanyProfilesApi.md#DetachPrivateServiceFromProfile) | **Delete** /fabric/v4/companyProfiles/{companyProfileId}/privateServices/{privateServiceId} | Detach Private Service
[**DetachServiceProfileFromProfile**](CompanyProfilesApi.md#DetachServiceProfileFromProfile) | **Delete** /fabric/v4/companyProfiles/{companyProfileId}/serviceProfiles/{serviceProfileId} | Detach Service Profile
[**DetachTagFromProfile**](CompanyProfilesApi.md#DetachTagFromProfile) | **Delete** /fabric/v4/companyProfiles/{companyProfileId}/tags/{tagId} | Detach Tag
[**GetCompanyProfile**](CompanyProfilesApi.md#GetCompanyProfile) | **Get** /fabric/v4/companyProfiles/{companyProfileId} | Get Company Profile by UUID
[**GetCompanyProfilePrivateServices**](CompanyProfilesApi.md#GetCompanyProfilePrivateServices) | **Get** /fabric/v4/companyProfiles/{companyProfileId}/privateServices | Get Private Services
[**GetCompanyProfileServiceProfiles**](CompanyProfilesApi.md#GetCompanyProfileServiceProfiles) | **Get** /fabric/v4/companyProfiles/{companyProfileId}/serviceProfiles | Get Service Profiles
[**GetTags**](CompanyProfilesApi.md#GetTags) | **Get** /fabric/v4/companyProfiles/{companyProfileId}/tags | Get Tags
[**SearchCompanyProfile**](CompanyProfilesApi.md#SearchCompanyProfile) | **Post** /fabric/v4/companyProfiles/search | Search Company Profiles



## AttachPrivateServiceToProfile

> AttachPrivateServiceResponse AttachPrivateServiceToProfile(ctx, companyProfileId, privateServiceId).Execute()

Attach Private Service



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID
	privateServiceId := "privateServiceId_example" // string | Private Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.AttachPrivateServiceToProfile(context.Background(), companyProfileId, privateServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.AttachPrivateServiceToProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachPrivateServiceToProfile`: AttachPrivateServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.AttachPrivateServiceToProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 
**privateServiceId** | **string** | Private Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachPrivateServiceToProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachPrivateServiceResponse**](AttachPrivateServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachServiceProfileToProfile

> AttachServiceProfileResponse AttachServiceProfileToProfile(ctx, companyProfileId, serviceProfileId).Execute()

Attach Service Profile



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID
	serviceProfileId := "serviceProfileId_example" // string | Service Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.AttachServiceProfileToProfile(context.Background(), companyProfileId, serviceProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.AttachServiceProfileToProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachServiceProfileToProfile`: AttachServiceProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.AttachServiceProfileToProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachServiceProfileToProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachServiceProfileResponse**](AttachServiceProfileResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachTagToProfile

> AttachTagResponse AttachTagToProfile(ctx, companyProfileId, tagId).Execute()

Attach Tag



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID
	tagId := "tagId_example" // string | Tag UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.AttachTagToProfile(context.Background(), companyProfileId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.AttachTagToProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachTagToProfile`: AttachTagResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.AttachTagToProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 
**tagId** | **string** | Tag UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachTagToProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachTagResponse**](AttachTagResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyProfile

> CompanyProfileResponse CreateCompanyProfile(ctx).CompanyProfileRequest(companyProfileRequest).Execute()

Create Company Profile



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
	companyProfileRequest := *openapiclient.NewCompanyProfileRequest("COMPANY_PROFILE", "Equinix", "Global interconnection and data center company", "Equinix, Inc. connects the world's leading businesses to their customers, employees and partners inside the most interconnected data centers.") // CompanyProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.CreateCompanyProfile(context.Background()).CompanyProfileRequest(companyProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.CreateCompanyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompanyProfile`: CompanyProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.CreateCompanyProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyProfileRequest** | [**CompanyProfileRequest**](CompanyProfileRequest.md) |  | 

### Return type

[**CompanyProfileResponse**](CompanyProfileResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyProfile

> CompanyProfileResponse DeleteCompanyProfile(ctx, companyProfileId).Execute()

Delete Company Profile



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.DeleteCompanyProfile(context.Background(), companyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.DeleteCompanyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyProfile`: CompanyProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.DeleteCompanyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyProfileResponse**](CompanyProfileResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachPrivateServiceFromProfile

> AttachPrivateServiceResponse DetachPrivateServiceFromProfile(ctx, companyProfileId, privateServiceId).Execute()

Detach Private Service



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID
	privateServiceId := "privateServiceId_example" // string | Private Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.DetachPrivateServiceFromProfile(context.Background(), companyProfileId, privateServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.DetachPrivateServiceFromProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachPrivateServiceFromProfile`: AttachPrivateServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.DetachPrivateServiceFromProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 
**privateServiceId** | **string** | Private Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachPrivateServiceFromProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachPrivateServiceResponse**](AttachPrivateServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachServiceProfileFromProfile

> AttachServiceProfileResponse DetachServiceProfileFromProfile(ctx, companyProfileId, serviceProfileId).Execute()

Detach Service Profile



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID
	serviceProfileId := "serviceProfileId_example" // string | Service Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.DetachServiceProfileFromProfile(context.Background(), companyProfileId, serviceProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.DetachServiceProfileFromProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachServiceProfileFromProfile`: AttachServiceProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.DetachServiceProfileFromProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 
**serviceProfileId** | **string** | Service Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachServiceProfileFromProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachServiceProfileResponse**](AttachServiceProfileResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachTagFromProfile

> AttachTagResponse DetachTagFromProfile(ctx, companyProfileId, tagId).Execute()

Detach Tag



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID
	tagId := "tagId_example" // string | Tag UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.DetachTagFromProfile(context.Background(), companyProfileId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.DetachTagFromProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachTagFromProfile`: AttachTagResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.DetachTagFromProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 
**tagId** | **string** | Tag UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachTagFromProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachTagResponse**](AttachTagResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyProfile

> CompanyProfileResponse GetCompanyProfile(ctx, companyProfileId).Execute()

Get Company Profile by UUID



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.GetCompanyProfile(context.Background(), companyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.GetCompanyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyProfile`: CompanyProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.GetCompanyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyProfileResponse**](CompanyProfileResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyProfilePrivateServices

> PrivateServiceListResponse GetCompanyProfilePrivateServices(ctx, companyProfileId).Execute()

Get Private Services



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.GetCompanyProfilePrivateServices(context.Background(), companyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.GetCompanyProfilePrivateServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyProfilePrivateServices`: PrivateServiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.GetCompanyProfilePrivateServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyProfilePrivateServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateServiceListResponse**](PrivateServiceListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyProfileServiceProfiles

> ServiceProfileListResponse GetCompanyProfileServiceProfiles(ctx, companyProfileId).Execute()

Get Service Profiles



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.GetCompanyProfileServiceProfiles(context.Background(), companyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.GetCompanyProfileServiceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyProfileServiceProfiles`: ServiceProfileListResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.GetCompanyProfileServiceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyProfileServiceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceProfileListResponse**](ServiceProfileListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> TagListResponse GetTags(ctx, companyProfileId).Execute()

Get Tags



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
	companyProfileId := "companyProfileId_example" // string | Company Profile UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.GetTags(context.Background(), companyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: TagListResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyProfileId** | **string** | Company Profile UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagListResponse**](TagListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCompanyProfile

> CompanyProfileSearchResponse SearchCompanyProfile(ctx).CompanyProfileSearchRequest(companyProfileSearchRequest).ViewPoint(viewPoint).Execute()

Search Company Profiles



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
	companyProfileSearchRequest := *openapiclient.NewCompanyProfileSearchRequest() // CompanyProfileSearchRequest | 
	viewPoint := "viewPoint_example" // string | Viewpoint for the request, either 'aSide' or 'zSide' (optional) (default to "aSide")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyProfilesApi.SearchCompanyProfile(context.Background()).CompanyProfileSearchRequest(companyProfileSearchRequest).ViewPoint(viewPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyProfilesApi.SearchCompanyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCompanyProfile`: CompanyProfileSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CompanyProfilesApi.SearchCompanyProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompanyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyProfileSearchRequest** | [**CompanyProfileSearchRequest**](CompanyProfileSearchRequest.md) |  | 
 **viewPoint** | **string** | Viewpoint for the request, either &#39;aSide&#39; or &#39;zSide&#39; | [default to &quot;aSide&quot;]

### Return type

[**CompanyProfileSearchResponse**](CompanyProfileSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


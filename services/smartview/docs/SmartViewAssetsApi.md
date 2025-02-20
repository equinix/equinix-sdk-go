# \SmartViewAssetsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAffectedAsset**](SmartViewAssetsApi.md#GetAffectedAsset) | **Get** /smartview/v1/asset/tagpoint/affected-assets | get affected customers assets hierarchy
[**GetAsset**](SmartViewAssetsApi.md#GetAsset) | **Get** /smartview/v1/asset/list | get assets list information
[**GetAssetDetails**](SmartViewAssetsApi.md#GetAssetDetails) | **Get** /smartview/v1/asset/details | get details for an asset.
[**GetCurrentTagPoint**](SmartViewAssetsApi.md#GetCurrentTagPoint) | **Get** /smartview/v1/asset/tagpoint/current | obtain latest tag point data
[**GetTagpointTrending**](SmartViewAssetsApi.md#GetTagpointTrending) | **Get** /smartview/v1/asset/tagpoint/trending | obtain trending tag point data
[**Mixin3**](SmartViewAssetsApi.md#Mixin3) | **Post** /smartview/v1/asset/details | get asset details
[**Mixin3_0**](SmartViewAssetsApi.md#Mixin3_0) | **Post** /smartview/v1/asset/tagpoint/current | get current tag points data
[**SearchAsset**](SmartViewAssetsApi.md#SearchAsset) | **Get** /smartview/v1/asset/search | Search for Assets matching identifiers



## GetAffectedAsset

> HierarchyNode GetAffectedAsset(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).AssetId(assetId).Classification(classification).Execute()

get affected customers assets hierarchy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | customer account number
	ibx := "ibx_example" // string | ibx code
	assetId := "assetId_example" // string | asset id
	classification := openapiclient.getAsset_classification_parameter("Electrical") // GetAssetClassificationParameter | asset classification(Electrical, Mechanical)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.GetAffectedAsset(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).AssetId(assetId).Classification(classification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.GetAffectedAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffectedAsset`: HierarchyNode
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.GetAffectedAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAffectedAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | customer account number | 
 **ibx** | **string** | ibx code | 
 **assetId** | **string** | asset id | 
 **classification** | [**GetAssetClassificationParameter**](GetAssetClassificationParameter.md) | asset classification(Electrical, Mechanical) | 

### Return type

[**HierarchyNode**](HierarchyNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> AssetsList GetAsset(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Classification(classification).Cages(cages).Execute()

get assets list information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | customer account number
	ibx := "ibx_example" // string | ibx code
	classification := openapiclient.getAsset_classification_parameter("Electrical") // GetAssetClassificationParameter | asset classification (Electrical, Mechanical)
	cages := []string{"Inner_example"} // []string | cage unique space id to be used to filter the assets list assumed to be all cage unique space id if no value is sent.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.GetAsset(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Classification(classification).Cages(cages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.GetAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsset`: AssetsList
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.GetAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | customer account number | 
 **ibx** | **string** | ibx code | 
 **classification** | [**GetAssetClassificationParameter**](GetAssetClassificationParameter.md) | asset classification (Electrical, Mechanical) | 
 **cages** | **[]string** | cage unique space id to be used to filter the assets list assumed to be all cage unique space id if no value is sent.  | 

### Return type

[**AssetsList**](AssetsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDetails

> AssetDetailsGetResponse GetAssetDetails(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Classification(classification).AssetId(assetId).Execute()

get details for an asset.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | customer account number
	ibx := "ibx_example" // string | ibx code
	classification := "classification_example" // string | asset classification (Electrical, Mechanical)
	assetId := "assetId_example" // string | asset id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.GetAssetDetails(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Classification(classification).AssetId(assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.GetAssetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetDetails`: AssetDetailsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.GetAssetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | customer account number | 
 **ibx** | **string** | ibx code | 
 **classification** | **string** | asset classification (Electrical, Mechanical) | 
 **assetId** | **string** | asset id | 

### Return type

[**AssetDetailsGetResponse**](AssetDetailsGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentTagPoint

> TagPointData GetCurrentTagPoint(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Tagid(tagid).Execute()

obtain latest tag point data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | customer account number
	ibx := "ibx_example" // string | ibx code
	tagid := "tagid_example" // string | tag id is the unique identifier for the tag point

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.GetCurrentTagPoint(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Tagid(tagid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.GetCurrentTagPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentTagPoint`: TagPointData
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.GetCurrentTagPoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentTagPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | customer account number | 
 **ibx** | **string** | ibx code | 
 **tagid** | **string** | tag id is the unique identifier for the tag point | 

### Return type

[**TagPointData**](TagPointData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagpointTrending

> TagPointTrendingResponse GetTagpointTrending(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Tagid(tagid).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()

obtain trending tag point data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | customer account number
	ibx := "ibx_example" // string | ibx code
	tagid := "tagid_example" // string | tag id is the unique identifier for the tag point
	interval := "interval_example" // string | tag point data interval (1h, 1d, reading)
	fromDate := int32(56) // int32 | from date UTC time (1494345600000)
	toDate := int32(56) // int32 | to date UTC time (1494432000000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.GetTagpointTrending(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).Tagid(tagid).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.GetTagpointTrending``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagpointTrending`: TagPointTrendingResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.GetTagpointTrending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagpointTrendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | customer account number | 
 **ibx** | **string** | ibx code | 
 **tagid** | **string** | tag id is the unique identifier for the tag point | 
 **interval** | **string** | tag point data interval (1h, 1d, reading) | 
 **fromDate** | **int32** | from date UTC time (1494345600000) | 
 **toDate** | **int32** | to date UTC time (1494432000000) | 

### Return type

[**TagPointTrendingResponse**](TagPointTrendingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Mixin3

> AssetDetailsResponse Mixin3(ctx).Authorization(authorization).Payload(payload).Execute()

get asset details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	payload := *openapiclient.NewAssetDetailsRequest() // AssetDetailsRequest | request schema (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.Mixin3(context.Background()).Authorization(authorization).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.Mixin3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin3`: AssetDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.Mixin3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **payload** | [**AssetDetailsRequest**](AssetDetailsRequest.md) | request schema | 

### Return type

[**AssetDetailsResponse**](AssetDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Mixin3_0

> TagPointData Mixin3_0(ctx).Authorization(authorization).Payload(payload).Execute()

get current tag points data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	payload := *openapiclient.NewCurrentTagPointRequest() // CurrentTagPointRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.Mixin3_0(context.Background()).Authorization(authorization).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.Mixin3_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin3_0`: TagPointData
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.Mixin3_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin3_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **payload** | [**CurrentTagPointRequest**](CurrentTagPointRequest.md) |  | 

### Return type

[**TagPointData**](TagPointData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAsset

> Assets SearchAsset(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).SearchString(searchString).Execute()

Search for Assets matching identifiers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | Customer Account Number
	ibx := "ibx_example" // string | IBX Code
	searchString := "searchString_example" // string | Search String

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewAssetsApi.SearchAsset(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).SearchString(searchString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewAssetsApi.SearchAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAsset`: Assets
	fmt.Fprintf(os.Stdout, "Response from `SmartViewAssetsApi.SearchAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **string** | IBX Code | 
 **searchString** | **string** | Search String | 

### Return type

[**Assets**](Assets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


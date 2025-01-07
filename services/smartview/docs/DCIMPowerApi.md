# \DCIMPowerApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Mixin0**](DCIMPowerApi.md#Mixin0) | **Post** /power/v1/current | Fetch current power consumption data
[**Mixin0_0**](DCIMPowerApi.md#Mixin0_0) | **Get** /power/v1/trending | Fetch Trending Power Data. 
[**PowerV1CurrentGet**](DCIMPowerApi.md#PowerV1CurrentGet) | **Get** /power/v1/current | Fetch current power consumption data



## Mixin0

> PowerDataResponseIBX Mixin0(ctx).Authorization(authorization).Body(body).Execute()

Fetch current power consumption data



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
	body := *openapiclient.NewPowerCurrentPostRequest() // PowerCurrentPostRequest | request payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DCIMPowerApi.Mixin0(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DCIMPowerApi.Mixin0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin0`: PowerDataResponseIBX
	fmt.Fprintf(os.Stdout, "Response from `DCIMPowerApi.Mixin0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **body** | [**PowerCurrentPostRequest**](PowerCurrentPostRequest.md) | request payload | 

### Return type

[**PowerDataResponseIBX**](PowerDataResponseIBX.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Mixin0_0

> TrendingPowerData Mixin0_0(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()

Fetch Trending Power Data. 



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
	ibx := float32(8.14) // float32 | IBX Code
	levelType := openapiclient._power_v1_current_get_levelType_parameter("ibx") // PowerV1CurrentGetLevelTypeParameter | [ibx|cage|cabinet|circuit]
	levelValue := "levelValue_example" // string | ibx code, cage unique space id, cabinet unique space id and serial number for level type ibx, cage, cabinet and circuit respectively. 
	interval := "interval_example" // string | [recording|1h|1d]
	fromDate := "fromDate_example" // string | timestamp expected to be epoch long ( milliseconds ).
	toDate := "toDate_example" // string | timestamp expected to be epoch long ( milliseconds ).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DCIMPowerApi.Mixin0_0(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DCIMPowerApi.Mixin0_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin0_0`: TrendingPowerData
	fmt.Fprintf(os.Stdout, "Response from `DCIMPowerApi.Mixin0_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin0_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **float32** | IBX Code | 
 **levelType** | [**PowerV1CurrentGetLevelTypeParameter**](PowerV1CurrentGetLevelTypeParameter.md) | [ibx|cage|cabinet|circuit] | 
 **levelValue** | **string** | ibx code, cage unique space id, cabinet unique space id and serial number for level type ibx, cage, cabinet and circuit respectively.  | 
 **interval** | **string** | [recording|1h|1d] | 
 **fromDate** | **string** | timestamp expected to be epoch long ( milliseconds ). | 
 **toDate** | **string** | timestamp expected to be epoch long ( milliseconds ). | 

### Return type

[**TrendingPowerData**](TrendingPowerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PowerV1CurrentGet

> PowerData PowerV1CurrentGet(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Execute()

Fetch current power consumption data



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
	levelType := openapiclient._power_v1_current_get_levelType_parameter("ibx") // PowerV1CurrentGetLevelTypeParameter | level type allowed value [ibx|cage|cabinet|circuit]
	levelValue := "levelValue_example" // string | level value - ibx code, cage unique space id, cabinet unique space id, serial number for level type ibx, cage, cabinet, circuit respectively. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DCIMPowerApi.PowerV1CurrentGet(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DCIMPowerApi.PowerV1CurrentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PowerV1CurrentGet`: PowerData
	fmt.Fprintf(os.Stdout, "Response from `DCIMPowerApi.PowerV1CurrentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPowerV1CurrentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **string** | IBX Code | 
 **levelType** | [**PowerV1CurrentGetLevelTypeParameter**](PowerV1CurrentGetLevelTypeParameter.md) | level type allowed value [ibx|cage|cabinet|circuit] | 
 **levelValue** | **string** | level value - ibx code, cage unique space id, cabinet unique space id, serial number for level type ibx, cage, cabinet, circuit respectively.  | 

### Return type

[**PowerData**](PowerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


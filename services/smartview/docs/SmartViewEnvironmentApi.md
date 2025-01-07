# \SmartViewEnvironmentApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Mixin0**](SmartViewEnvironmentApi.md#Mixin0) | **Get** /environment/v1/current | Get current environmental data
[**Mixin0_0**](SmartViewEnvironmentApi.md#Mixin0_0) | **Get** /environment/v1/listCurrent | Fetch list of environmental data
[**Mixin0_1**](SmartViewEnvironmentApi.md#Mixin0_1) | **Get** /environment/v1/trending | Fetch trending environmental data



## Mixin0

> EnvironmentData Mixin0(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Execute()

Get current environmental data



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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	accountNo := "accountNo_example" // string | Customer Account Number
	ibx := "ibx_example" // string | IBX Code
	levelType := openapiclient.Mixin0_levelType_parameter("IBX") // Mixin0LevelTypeParameter | Level Type
	levelValue := "levelValue_example" // string | Level Value is ibxCode, zoneUsID, cageUsID, sensorid for  levelType ibx, zone, cage, sensor resp. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentApi.Mixin0(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentApi.Mixin0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin0`: EnvironmentData
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentApi.Mixin0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **string** | IBX Code | 
 **levelType** | [**Mixin0LevelTypeParameter**](Mixin0LevelTypeParameter.md) | Level Type | 
 **levelValue** | **string** | Level Value is ibxCode, zoneUsID, cageUsID, sensorid for  levelType ibx, zone, cage, sensor resp.  | 

### Return type

[**EnvironmentData**](EnvironmentData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Mixin0_0

> EnvironmentDataResponse Mixin0_0(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).Execute()

Fetch list of environmental data



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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	accountNo := "accountNo_example" // string | Customer Account Number
	ibx := "ibx_example" // string | IBX Code
	levelType := openapiclient.Mixin0_levelType_parameter("IBX") // Mixin0LevelTypeParameter | Level Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentApi.Mixin0_0(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentApi.Mixin0_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin0_0`: EnvironmentDataResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentApi.Mixin0_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin0_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **string** | IBX Code | 
 **levelType** | [**Mixin0LevelTypeParameter**](Mixin0LevelTypeParameter.md) | Level Type | 

### Return type

[**EnvironmentDataResponse**](EnvironmentDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Mixin0_1

> TrendingEnvironmentData Mixin0_1(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).DataPoint(dataPoint).LevelType(levelType).LevelValue(levelValue).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()

Fetch trending environmental data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	accountNo := "accountNo_example" // string | Customer Account Number
	ibx := float32(8.14) // float32 | IBX Code
	dataPoint := openapiclient.Mixin0_dataPoint_parameter("temperature") // Mixin0DataPointParameter | data point
	levelType := openapiclient.Mixin0_levelType_parameter("IBX") // Mixin0LevelTypeParameter | Level Type
	levelValue := "levelValue_example" // string | Level Value is ibxCode, zone, cage, sensorid for  levelType ibx, zone, cage, sensor resp. 
	interval := openapiclient.Mixin0_interval_parameter("reading") // Mixin0IntervalParameter | 
	fromDate := time.Now() // time.Time | date in long
	toDate := time.Now() // time.Time | date in long

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentApi.Mixin0_1(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).DataPoint(dataPoint).LevelType(levelType).LevelValue(levelValue).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentApi.Mixin0_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin0_1`: TrendingEnvironmentData
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentApi.Mixin0_1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin0_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **float32** | IBX Code | 
 **dataPoint** | [**Mixin0DataPointParameter**](Mixin0DataPointParameter.md) | data point | 
 **levelType** | [**Mixin0LevelTypeParameter**](Mixin0LevelTypeParameter.md) | Level Type | 
 **levelValue** | **string** | Level Value is ibxCode, zone, cage, sensorid for  levelType ibx, zone, cage, sensor resp.  | 
 **interval** | [**Mixin0IntervalParameter**](Mixin0IntervalParameter.md) |  | 
 **fromDate** | **time.Time** | date in long | 
 **toDate** | **time.Time** | date in long | 

### Return type

[**TrendingEnvironmentData**](TrendingEnvironmentData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SmartViewEnvironmentApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Mixin1**](SmartViewEnvironmentApi.md#Mixin1) | **Get** /environment/v1/current | Get current environmental data
[**Mixin1_0**](SmartViewEnvironmentApi.md#Mixin1_0) | **Get** /environment/v1/listCurrent | Fetch list of environmental data
[**Mixin1_1**](SmartViewEnvironmentApi.md#Mixin1_1) | **Get** /environment/v1/trending | Fetch trending environmental data



## Mixin1

> EnvironmentData Mixin1(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Execute()

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
	levelType := openapiclient.Mixin1_levelType_parameter("IBX") // Mixin1LevelTypeParameter | Level Type
	levelValue := "levelValue_example" // string | Level Value is ibxCode, zoneUsID, cageUsID, sensorid for  levelType ibx, zone, cage, sensor resp. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentApi.Mixin1(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).LevelValue(levelValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentApi.Mixin1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin1`: EnvironmentData
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentApi.Mixin1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **string** | IBX Code | 
 **levelType** | [**Mixin1LevelTypeParameter**](Mixin1LevelTypeParameter.md) | Level Type | 
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


## Mixin1_0

> EnvironmentDataResponse Mixin1_0(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).Execute()

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
	levelType := openapiclient.Mixin1_levelType_parameter("IBX") // Mixin1LevelTypeParameter | Level Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentApi.Mixin1_0(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).LevelType(levelType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentApi.Mixin1_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin1_0`: EnvironmentDataResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentApi.Mixin1_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin1_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **string** | IBX Code | 
 **levelType** | [**Mixin1LevelTypeParameter**](Mixin1LevelTypeParameter.md) | Level Type | 

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


## Mixin1_1

> TrendingEnvironmentData Mixin1_1(ctx).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).DataPoint(dataPoint).LevelType(levelType).LevelValue(levelValue).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()

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
	dataPoint := openapiclient.Mixin1_dataPoint_parameter("temperature") // Mixin1DataPointParameter | data point
	levelType := openapiclient.Mixin1_levelType_parameter("IBX") // Mixin1LevelTypeParameter | Level Type
	levelValue := "levelValue_example" // string | Level Value is ibxCode, zone, cage, sensorid for  levelType ibx, zone, cage, sensor resp. 
	interval := openapiclient.Mixin1_interval_parameter("reading") // Mixin1IntervalParameter | 
	fromDate := time.Now() // time.Time | date in long
	toDate := time.Now() // time.Time | date in long

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentApi.Mixin1_1(context.Background()).Authorization(authorization).AccountNo(accountNo).Ibx(ibx).DataPoint(dataPoint).LevelType(levelType).LevelValue(levelValue).Interval(interval).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentApi.Mixin1_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Mixin1_1`: TrendingEnvironmentData
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentApi.Mixin1_1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMixin1_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **accountNo** | **string** | Customer Account Number | 
 **ibx** | **float32** | IBX Code | 
 **dataPoint** | [**Mixin1DataPointParameter**](Mixin1DataPointParameter.md) | data point | 
 **levelType** | [**Mixin1LevelTypeParameter**](Mixin1LevelTypeParameter.md) | Level Type | 
 **levelValue** | **string** | Level Value is ibxCode, zone, cage, sensorid for  levelType ibx, zone, cage, sensor resp.  | 
 **interval** | [**Mixin1IntervalParameter**](Mixin1IntervalParameter.md) |  | 
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


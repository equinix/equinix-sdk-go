# \SmartViewEnvironmentalApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSensorReadings**](SmartViewEnvironmentalApi.md#GetSensorReadings) | **Get** /smartview/v2/environmental/ibxs/{ibx}/sensors/readings | IBX sensors current readings
[**GetSingleSensorReadings**](SmartViewEnvironmentalApi.md#GetSingleSensorReadings) | **Get** /smartview/v2/environmental/ibxs/{ibx}/sensors/{sensorId}/readings | Single sensor current value(s)



## GetSensorReadings

> SensorReadingsResponse GetSensorReadings(ctx, ibx).Authorization(authorization).Type_(type_).Zone(zone).Offset(offset).Limit(limit).Execute()

IBX sensors current readings



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
	ibx := "CH2" // string | IBX where the sensor is located.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	type_ := "HUMIDITY" // string | Type of sensor data to display, e.g. specifying 'HUMIDITY' means only sensor readings with HUMIDITY value will be included (optional)
	zone := "CH2:1:06:ColoArea:1" // string | Zone name. Adding this parameter limits the query to sensors in the specified zone (optional)
	offset := int32(56) // int32 | Results offset you want to retrieve (0..N) (optional)
	limit := int32(56) // int32 | Number of records to retrieve per request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentalApi.GetSensorReadings(context.Background(), ibx).Authorization(authorization).Type_(type_).Zone(zone).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentalApi.GetSensorReadings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSensorReadings`: SensorReadingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentalApi.GetSensorReadings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ibx** | **string** | IBX where the sensor is located. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSensorReadingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **type_** | **string** | Type of sensor data to display, e.g. specifying &#39;HUMIDITY&#39; means only sensor readings with HUMIDITY value will be included | 
 **zone** | **string** | Zone name. Adding this parameter limits the query to sensors in the specified zone | 
 **offset** | **int32** | Results offset you want to retrieve (0..N) | 
 **limit** | **int32** | Number of records to retrieve per request. | 

### Return type

[**SensorReadingsResponse**](SensorReadingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleSensorReadings

> SensorReading GetSingleSensorReadings(ctx, ibx, sensorId).Authorization(authorization).Execute()

Single sensor current value(s)



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
	ibx := "SV2" // string | IBX where the sensor is located.
	sensorId := "CH2.Environmental.MbusColo3Mod5.MOD35.C3TS03" // string | Id of sensor to read.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewEnvironmentalApi.GetSingleSensorReadings(context.Background(), ibx, sensorId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewEnvironmentalApi.GetSingleSensorReadings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleSensorReadings`: SensorReading
	fmt.Fprintf(os.Stdout, "Response from `SmartViewEnvironmentalApi.GetSingleSensorReadings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ibx** | **string** | IBX where the sensor is located. | 
**sensorId** | **string** | Id of sensor to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleSensorReadingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**SensorReading**](SensorReading.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


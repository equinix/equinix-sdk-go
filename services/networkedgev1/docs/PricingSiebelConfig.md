# PricingSiebelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermLength** | Pointer to **string** | The termlength of the Siebel order. | [optional] 
**OrderNumber** | Pointer to **string** | The order number. | [optional] 
**Core** | Pointer to **int32** | The core selection on Siebel. | [optional] 
**Throughput** | Pointer to **string** | Throughput. | [optional] 
**ThroughputUnit** | Pointer to **string** | The throughput unit. | [optional] 
**PackageCode** | Pointer to **string** | The software package code. | [optional] 
**AdditionalBandwidth** | Pointer to **string** | The additional bandwidth selection on Siebel. | [optional] 
**Primary** | Pointer to [**PricingSiebelConfigPrimary**](PricingSiebelConfigPrimary.md) |  | [optional] 
**Secondary** | Pointer to [**PricingSiebelConfigSecondary**](PricingSiebelConfigSecondary.md) |  | [optional] 

## Methods

### NewPricingSiebelConfig

`func NewPricingSiebelConfig() *PricingSiebelConfig`

NewPricingSiebelConfig instantiates a new PricingSiebelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSiebelConfigWithDefaults

`func NewPricingSiebelConfigWithDefaults() *PricingSiebelConfig`

NewPricingSiebelConfigWithDefaults instantiates a new PricingSiebelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermLength

`func (o *PricingSiebelConfig) GetTermLength() string`

GetTermLength returns the TermLength field if non-nil, zero value otherwise.

### GetTermLengthOk

`func (o *PricingSiebelConfig) GetTermLengthOk() (*string, bool)`

GetTermLengthOk returns a tuple with the TermLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermLength

`func (o *PricingSiebelConfig) SetTermLength(v string)`

SetTermLength sets TermLength field to given value.

### HasTermLength

`func (o *PricingSiebelConfig) HasTermLength() bool`

HasTermLength returns a boolean if a field has been set.

### GetOrderNumber

`func (o *PricingSiebelConfig) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *PricingSiebelConfig) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *PricingSiebelConfig) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *PricingSiebelConfig) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetCore

`func (o *PricingSiebelConfig) GetCore() int32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *PricingSiebelConfig) GetCoreOk() (*int32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *PricingSiebelConfig) SetCore(v int32)`

SetCore sets Core field to given value.

### HasCore

`func (o *PricingSiebelConfig) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetThroughput

`func (o *PricingSiebelConfig) GetThroughput() string`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *PricingSiebelConfig) GetThroughputOk() (*string, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *PricingSiebelConfig) SetThroughput(v string)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *PricingSiebelConfig) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *PricingSiebelConfig) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *PricingSiebelConfig) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *PricingSiebelConfig) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *PricingSiebelConfig) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetPackageCode

`func (o *PricingSiebelConfig) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *PricingSiebelConfig) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *PricingSiebelConfig) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *PricingSiebelConfig) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetAdditionalBandwidth

`func (o *PricingSiebelConfig) GetAdditionalBandwidth() string`

GetAdditionalBandwidth returns the AdditionalBandwidth field if non-nil, zero value otherwise.

### GetAdditionalBandwidthOk

`func (o *PricingSiebelConfig) GetAdditionalBandwidthOk() (*string, bool)`

GetAdditionalBandwidthOk returns a tuple with the AdditionalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBandwidth

`func (o *PricingSiebelConfig) SetAdditionalBandwidth(v string)`

SetAdditionalBandwidth sets AdditionalBandwidth field to given value.

### HasAdditionalBandwidth

`func (o *PricingSiebelConfig) HasAdditionalBandwidth() bool`

HasAdditionalBandwidth returns a boolean if a field has been set.

### GetPrimary

`func (o *PricingSiebelConfig) GetPrimary() PricingSiebelConfigPrimary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PricingSiebelConfig) GetPrimaryOk() (*PricingSiebelConfigPrimary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PricingSiebelConfig) SetPrimary(v PricingSiebelConfigPrimary)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PricingSiebelConfig) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *PricingSiebelConfig) GetSecondary() PricingSiebelConfigSecondary`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *PricingSiebelConfig) GetSecondaryOk() (*PricingSiebelConfigSecondary, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *PricingSiebelConfig) SetSecondary(v PricingSiebelConfigSecondary)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *PricingSiebelConfig) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



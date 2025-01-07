# MeteredPower

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | account number | [optional] 
**Asset** | [**MeteredPowerAssetDetails**](MeteredPowerAssetDetails.md) |  | 
**Cage** | Pointer to **string** | cage | [optional] 
**CageSerialNo** | Pointer to **string** | cage serial number | [optional] 
**DataQuality** | Pointer to **string** | data quality: Good | Bad | Uncertain | [optional] 
**Ibx** | **string** | ibx | 
**Reading** | [**MeteredPowerValueWithUnit**](MeteredPowerValueWithUnit.md) |  | 
**ReadingTime** | Pointer to **string** | message reading time | [optional] 
**StreamId** | **string** | unique message id | 
**Tag** | [**MeteredPowerTagDetails**](MeteredPowerTagDetails.md) |  | 

## Methods

### NewMeteredPower

`func NewMeteredPower(asset MeteredPowerAssetDetails, ibx string, reading MeteredPowerValueWithUnit, streamId string, tag MeteredPowerTagDetails, ) *MeteredPower`

NewMeteredPower instantiates a new MeteredPower object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteredPowerWithDefaults

`func NewMeteredPowerWithDefaults() *MeteredPower`

NewMeteredPowerWithDefaults instantiates a new MeteredPower object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *MeteredPower) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *MeteredPower) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *MeteredPower) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *MeteredPower) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAsset

`func (o *MeteredPower) GetAsset() MeteredPowerAssetDetails`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *MeteredPower) GetAssetOk() (*MeteredPowerAssetDetails, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *MeteredPower) SetAsset(v MeteredPowerAssetDetails)`

SetAsset sets Asset field to given value.


### GetCage

`func (o *MeteredPower) GetCage() string`

GetCage returns the Cage field if non-nil, zero value otherwise.

### GetCageOk

`func (o *MeteredPower) GetCageOk() (*string, bool)`

GetCageOk returns a tuple with the Cage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCage

`func (o *MeteredPower) SetCage(v string)`

SetCage sets Cage field to given value.

### HasCage

`func (o *MeteredPower) HasCage() bool`

HasCage returns a boolean if a field has been set.

### GetCageSerialNo

`func (o *MeteredPower) GetCageSerialNo() string`

GetCageSerialNo returns the CageSerialNo field if non-nil, zero value otherwise.

### GetCageSerialNoOk

`func (o *MeteredPower) GetCageSerialNoOk() (*string, bool)`

GetCageSerialNoOk returns a tuple with the CageSerialNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageSerialNo

`func (o *MeteredPower) SetCageSerialNo(v string)`

SetCageSerialNo sets CageSerialNo field to given value.

### HasCageSerialNo

`func (o *MeteredPower) HasCageSerialNo() bool`

HasCageSerialNo returns a boolean if a field has been set.

### GetDataQuality

`func (o *MeteredPower) GetDataQuality() string`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *MeteredPower) GetDataQualityOk() (*string, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *MeteredPower) SetDataQuality(v string)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *MeteredPower) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### GetIbx

`func (o *MeteredPower) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *MeteredPower) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *MeteredPower) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetReading

`func (o *MeteredPower) GetReading() MeteredPowerValueWithUnit`

GetReading returns the Reading field if non-nil, zero value otherwise.

### GetReadingOk

`func (o *MeteredPower) GetReadingOk() (*MeteredPowerValueWithUnit, bool)`

GetReadingOk returns a tuple with the Reading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReading

`func (o *MeteredPower) SetReading(v MeteredPowerValueWithUnit)`

SetReading sets Reading field to given value.


### GetReadingTime

`func (o *MeteredPower) GetReadingTime() string`

GetReadingTime returns the ReadingTime field if non-nil, zero value otherwise.

### GetReadingTimeOk

`func (o *MeteredPower) GetReadingTimeOk() (*string, bool)`

GetReadingTimeOk returns a tuple with the ReadingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingTime

`func (o *MeteredPower) SetReadingTime(v string)`

SetReadingTime sets ReadingTime field to given value.

### HasReadingTime

`func (o *MeteredPower) HasReadingTime() bool`

HasReadingTime returns a boolean if a field has been set.

### GetStreamId

`func (o *MeteredPower) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *MeteredPower) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *MeteredPower) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetTag

`func (o *MeteredPower) GetTag() MeteredPowerTagDetails`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MeteredPower) GetTagOk() (*MeteredPowerTagDetails, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MeteredPower) SetTag(v MeteredPowerTagDetails)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



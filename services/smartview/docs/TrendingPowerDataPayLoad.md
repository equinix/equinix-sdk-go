# TrendingPowerDataPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]ComparisonDataTrend**](ComparisonDataTrend.md) |  | [optional] 
**Ibx** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to [**TrendingPowerDataPayLoadInterval**](TrendingPowerDataPayLoadInterval.md) |  | [optional] 
**LevelType** | Pointer to [**TrendingPowerDataPayLoadLevelType**](TrendingPowerDataPayLoadLevelType.md) |  | [optional] 
**LevelValue** | Pointer to **string** | ibx code, cage unique space id, cabinet unique space id and serial number for levelType ibx, cage, cabinet, circuit resp. | [optional] 

## Methods

### NewTrendingPowerDataPayLoad

`func NewTrendingPowerDataPayLoad() *TrendingPowerDataPayLoad`

NewTrendingPowerDataPayLoad instantiates a new TrendingPowerDataPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendingPowerDataPayLoadWithDefaults

`func NewTrendingPowerDataPayLoadWithDefaults() *TrendingPowerDataPayLoad`

NewTrendingPowerDataPayLoadWithDefaults instantiates a new TrendingPowerDataPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *TrendingPowerDataPayLoad) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *TrendingPowerDataPayLoad) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *TrendingPowerDataPayLoad) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *TrendingPowerDataPayLoad) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetData

`func (o *TrendingPowerDataPayLoad) GetData() []ComparisonDataTrend`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrendingPowerDataPayLoad) GetDataOk() (*[]ComparisonDataTrend, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrendingPowerDataPayLoad) SetData(v []ComparisonDataTrend)`

SetData sets Data field to given value.

### HasData

`func (o *TrendingPowerDataPayLoad) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIbx

`func (o *TrendingPowerDataPayLoad) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *TrendingPowerDataPayLoad) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *TrendingPowerDataPayLoad) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *TrendingPowerDataPayLoad) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetInterval

`func (o *TrendingPowerDataPayLoad) GetInterval() TrendingPowerDataPayLoadInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TrendingPowerDataPayLoad) GetIntervalOk() (*TrendingPowerDataPayLoadInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TrendingPowerDataPayLoad) SetInterval(v TrendingPowerDataPayLoadInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TrendingPowerDataPayLoad) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLevelType

`func (o *TrendingPowerDataPayLoad) GetLevelType() TrendingPowerDataPayLoadLevelType`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *TrendingPowerDataPayLoad) GetLevelTypeOk() (*TrendingPowerDataPayLoadLevelType, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *TrendingPowerDataPayLoad) SetLevelType(v TrendingPowerDataPayLoadLevelType)`

SetLevelType sets LevelType field to given value.

### HasLevelType

`func (o *TrendingPowerDataPayLoad) HasLevelType() bool`

HasLevelType returns a boolean if a field has been set.

### GetLevelValue

`func (o *TrendingPowerDataPayLoad) GetLevelValue() string`

GetLevelValue returns the LevelValue field if non-nil, zero value otherwise.

### GetLevelValueOk

`func (o *TrendingPowerDataPayLoad) GetLevelValueOk() (*string, bool)`

GetLevelValueOk returns a tuple with the LevelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelValue

`func (o *TrendingPowerDataPayLoad) SetLevelValue(v string)`

SetLevelValue sets LevelValue field to given value.

### HasLevelValue

`func (o *TrendingPowerDataPayLoad) HasLevelValue() bool`

HasLevelValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



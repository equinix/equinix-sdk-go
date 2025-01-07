# TagPointTrendingResponsePayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]TagPointTrendingData**](TagPointTrendingData.md) | trend data of tag | [optional] 
**DataType** | Pointer to **string** | data type of trend data vlaues | [optional] 
**Ibx** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to [**TagPointTrendingResponsePayLoadInterval**](TagPointTrendingResponsePayLoadInterval.md) |  | [optional] 
**TagDisplayName** | Pointer to **string** | the   | [optional] 
**TagId** | Pointer to **string** | the unique identifiers for the tag point ids for which the trending point is requested.  | [optional] 
**Uom** | Pointer to **string** |  | [optional] 

## Methods

### NewTagPointTrendingResponsePayLoad

`func NewTagPointTrendingResponsePayLoad() *TagPointTrendingResponsePayLoad`

NewTagPointTrendingResponsePayLoad instantiates a new TagPointTrendingResponsePayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPointTrendingResponsePayLoadWithDefaults

`func NewTagPointTrendingResponsePayLoadWithDefaults() *TagPointTrendingResponsePayLoad`

NewTagPointTrendingResponsePayLoadWithDefaults instantiates a new TagPointTrendingResponsePayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *TagPointTrendingResponsePayLoad) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *TagPointTrendingResponsePayLoad) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *TagPointTrendingResponsePayLoad) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *TagPointTrendingResponsePayLoad) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetData

`func (o *TagPointTrendingResponsePayLoad) GetData() []TagPointTrendingData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TagPointTrendingResponsePayLoad) GetDataOk() (*[]TagPointTrendingData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TagPointTrendingResponsePayLoad) SetData(v []TagPointTrendingData)`

SetData sets Data field to given value.

### HasData

`func (o *TagPointTrendingResponsePayLoad) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDataType

`func (o *TagPointTrendingResponsePayLoad) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TagPointTrendingResponsePayLoad) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TagPointTrendingResponsePayLoad) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *TagPointTrendingResponsePayLoad) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIbx

`func (o *TagPointTrendingResponsePayLoad) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *TagPointTrendingResponsePayLoad) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *TagPointTrendingResponsePayLoad) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *TagPointTrendingResponsePayLoad) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetInterval

`func (o *TagPointTrendingResponsePayLoad) GetInterval() TagPointTrendingResponsePayLoadInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TagPointTrendingResponsePayLoad) GetIntervalOk() (*TagPointTrendingResponsePayLoadInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TagPointTrendingResponsePayLoad) SetInterval(v TagPointTrendingResponsePayLoadInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TagPointTrendingResponsePayLoad) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTagDisplayName

`func (o *TagPointTrendingResponsePayLoad) GetTagDisplayName() string`

GetTagDisplayName returns the TagDisplayName field if non-nil, zero value otherwise.

### GetTagDisplayNameOk

`func (o *TagPointTrendingResponsePayLoad) GetTagDisplayNameOk() (*string, bool)`

GetTagDisplayNameOk returns a tuple with the TagDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagDisplayName

`func (o *TagPointTrendingResponsePayLoad) SetTagDisplayName(v string)`

SetTagDisplayName sets TagDisplayName field to given value.

### HasTagDisplayName

`func (o *TagPointTrendingResponsePayLoad) HasTagDisplayName() bool`

HasTagDisplayName returns a boolean if a field has been set.

### GetTagId

`func (o *TagPointTrendingResponsePayLoad) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *TagPointTrendingResponsePayLoad) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *TagPointTrendingResponsePayLoad) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *TagPointTrendingResponsePayLoad) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetUom

`func (o *TagPointTrendingResponsePayLoad) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *TagPointTrendingResponsePayLoad) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *TagPointTrendingResponsePayLoad) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *TagPointTrendingResponsePayLoad) HasUom() bool`

HasUom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



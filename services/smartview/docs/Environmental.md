# Environmental

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**Ibx** | **[]string** |  | 
**Level** | Pointer to [**[]EnvironmentalLevelInner**](EnvironmentalLevelInner.md) | This field is not required. If not provided in the request, subscription will include all environmental messages at the IBX, zone, and cage levels. If granularity is specified, at least one level must be provided. | [optional] 

## Methods

### NewEnvironmental

`func NewEnvironmental(accountNumber string, ibx []string, ) *Environmental`

NewEnvironmental instantiates a new Environmental object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentalWithDefaults

`func NewEnvironmentalWithDefaults() *Environmental`

NewEnvironmentalWithDefaults instantiates a new Environmental object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Environmental) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Environmental) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Environmental) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetIbx

`func (o *Environmental) GetIbx() []string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Environmental) GetIbxOk() (*[]string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Environmental) SetIbx(v []string)`

SetIbx sets Ibx field to given value.


### GetLevel

`func (o *Environmental) GetLevel() []EnvironmentalLevelInner`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Environmental) GetLevelOk() (*[]EnvironmentalLevelInner, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Environmental) SetLevel(v []EnvironmentalLevelInner)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Environmental) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CompanyProfileChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCompanyProfileChange

`func NewCompanyProfileChange() *CompanyProfileChange`

NewCompanyProfileChange instantiates a new CompanyProfileChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileChangeWithDefaults

`func NewCompanyProfileChangeWithDefaults() *CompanyProfileChange`

NewCompanyProfileChangeWithDefaults instantiates a new CompanyProfileChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompanyProfileChange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyProfileChange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyProfileChange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyProfileChange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyProfileChange) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyProfileChange) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyProfileChange) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyProfileChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *CompanyProfileChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *CompanyProfileChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *CompanyProfileChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *CompanyProfileChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetData

`func (o *CompanyProfileChange) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompanyProfileChange) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompanyProfileChange) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CompanyProfileChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



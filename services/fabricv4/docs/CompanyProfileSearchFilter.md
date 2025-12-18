# CompanyProfileSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CompanyProfileSearchFilter**](CompanyProfileSearchFilter.md) |  | [optional] 
**Or** | Pointer to [**[]CompanyProfileSearchFilter**](CompanyProfileSearchFilter.md) |  | [optional] 
**Property** | Pointer to **string** | Searchable field names in company profile | [optional] 
**Operator** | Pointer to **string** | Comparison operators for filtering | [optional] 
**Values** | Pointer to **[]string** | Values to compare against | [optional] 

## Methods

### NewCompanyProfileSearchFilter

`func NewCompanyProfileSearchFilter() *CompanyProfileSearchFilter`

NewCompanyProfileSearchFilter instantiates a new CompanyProfileSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileSearchFilterWithDefaults

`func NewCompanyProfileSearchFilterWithDefaults() *CompanyProfileSearchFilter`

NewCompanyProfileSearchFilterWithDefaults instantiates a new CompanyProfileSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *CompanyProfileSearchFilter) GetAnd() []CompanyProfileSearchFilter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *CompanyProfileSearchFilter) GetAndOk() (*[]CompanyProfileSearchFilter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *CompanyProfileSearchFilter) SetAnd(v []CompanyProfileSearchFilter)`

SetAnd sets And field to given value.

### HasAnd

`func (o *CompanyProfileSearchFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *CompanyProfileSearchFilter) GetOr() []CompanyProfileSearchFilter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CompanyProfileSearchFilter) GetOrOk() (*[]CompanyProfileSearchFilter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CompanyProfileSearchFilter) SetOr(v []CompanyProfileSearchFilter)`

SetOr sets Or field to given value.

### HasOr

`func (o *CompanyProfileSearchFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *CompanyProfileSearchFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CompanyProfileSearchFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CompanyProfileSearchFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CompanyProfileSearchFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CompanyProfileSearchFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CompanyProfileSearchFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CompanyProfileSearchFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CompanyProfileSearchFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CompanyProfileSearchFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CompanyProfileSearchFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CompanyProfileSearchFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CompanyProfileSearchFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



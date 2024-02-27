# ServiceProfileFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/name&#x60; - Service Profile name  * &#x60;/uuid&#x60; - Service Profile uuid  * &#x60;/state&#x60; - Service Profile status  * &#x60;/metros/code&#x60; - Service Profile metro code  * &#x60;/visibility&#x60; - Service Profile package  * &#x60;/type&#x60; - Service Profile package  * &#x60;/project/projectId&#x60; - Service Profile project id  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**And** | Pointer to [**[]ServiceProfileSimpleExpression**](ServiceProfileSimpleExpression.md) |  | [optional] 

## Methods

### NewServiceProfileFilter

`func NewServiceProfileFilter() *ServiceProfileFilter`

NewServiceProfileFilter instantiates a new ServiceProfileFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileFilterWithDefaults

`func NewServiceProfileFilterWithDefaults() *ServiceProfileFilter`

NewServiceProfileFilterWithDefaults instantiates a new ServiceProfileFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ServiceProfileFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ServiceProfileFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ServiceProfileFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ServiceProfileFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *ServiceProfileFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ServiceProfileFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ServiceProfileFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ServiceProfileFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ServiceProfileFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ServiceProfileFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ServiceProfileFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ServiceProfileFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetAnd

`func (o *ServiceProfileFilter) GetAnd() []ServiceProfileSimpleExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ServiceProfileFilter) GetAndOk() (*[]ServiceProfileSimpleExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ServiceProfileFilter) SetAnd(v []ServiceProfileSimpleExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ServiceProfileFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



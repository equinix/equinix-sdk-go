# ServiceProfileSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/name&#x60; - Service Profile name  * &#x60;/uuid&#x60; - Service Profile uuid  * &#x60;/state&#x60; - Service Profile status  * &#x60;/metros/code&#x60; - Service Profile metro code  * &#x60;/visibility&#x60; - Service Profile package  * &#x60;/type&#x60; - Service Profile package  * &#x60;/project/projectId&#x60; - Service Profile project id  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceProfileSimpleExpression

`func NewServiceProfileSimpleExpression() *ServiceProfileSimpleExpression`

NewServiceProfileSimpleExpression instantiates a new ServiceProfileSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileSimpleExpressionWithDefaults

`func NewServiceProfileSimpleExpressionWithDefaults() *ServiceProfileSimpleExpression`

NewServiceProfileSimpleExpressionWithDefaults instantiates a new ServiceProfileSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ServiceProfileSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ServiceProfileSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ServiceProfileSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ServiceProfileSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *ServiceProfileSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ServiceProfileSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ServiceProfileSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ServiceProfileSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ServiceProfileSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ServiceProfileSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ServiceProfileSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ServiceProfileSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



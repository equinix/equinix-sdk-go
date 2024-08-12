# TimeServiceSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/project/projectId&#x60; - project id (mandatory)  * &#x60;/name&#x60; - Precision Time Service name  * &#x60;/uuid&#x60; - Precision Time Service uuid  * &#x60;/type&#x60; - Precision Time Service protocol  * &#x60;/state&#x60; - Precision Time Service status  * &#x60;/account/accountNumber&#x60; - Precision Time Service account number  * &#x60;/package/code&#x60; - Precision Time Service package  * &#x60;/_*&#x60; - all-category search  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTimeServiceSimpleExpression

`func NewTimeServiceSimpleExpression() *TimeServiceSimpleExpression`

NewTimeServiceSimpleExpression instantiates a new TimeServiceSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeServiceSimpleExpressionWithDefaults

`func NewTimeServiceSimpleExpressionWithDefaults() *TimeServiceSimpleExpression`

NewTimeServiceSimpleExpressionWithDefaults instantiates a new TimeServiceSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *TimeServiceSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *TimeServiceSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *TimeServiceSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *TimeServiceSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *TimeServiceSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TimeServiceSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TimeServiceSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *TimeServiceSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *TimeServiceSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TimeServiceSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TimeServiceSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *TimeServiceSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



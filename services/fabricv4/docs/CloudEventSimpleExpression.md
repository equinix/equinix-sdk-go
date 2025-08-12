# CloudEventSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/subject&#x60; - Cloud Event subject description  * &#x60;/type&#x60; - Cloud Event type  * &#x60;/time&#x60; - Time of Cloud Events  * &#x60;/equinixproject&#x60; - Equinix Project of Cloud Events  * &#x60;/equinixorganization&#x60; - Equinix Organization of Cloud Events  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;BETWEEN&#x60; - between  * &#x60;IN&#x60; - in  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudEventSimpleExpression

`func NewCloudEventSimpleExpression() *CloudEventSimpleExpression`

NewCloudEventSimpleExpression instantiates a new CloudEventSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEventSimpleExpressionWithDefaults

`func NewCloudEventSimpleExpressionWithDefaults() *CloudEventSimpleExpression`

NewCloudEventSimpleExpressionWithDefaults instantiates a new CloudEventSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *CloudEventSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudEventSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudEventSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudEventSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CloudEventSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudEventSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudEventSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CloudEventSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CloudEventSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CloudEventSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CloudEventSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CloudEventSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



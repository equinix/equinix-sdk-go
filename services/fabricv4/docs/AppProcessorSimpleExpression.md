# AppProcessorSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:   * &#x60;/project/projectId&#x60; - project id   * &#x60;/uuid&#x60; - App Processor uuid   * &#x60;/name&#x60; - App Processor name   * &#x60;/description&#x60; - App Processor description   * &#x60;/state&#x60; - App Processor status   * &#x60;/ipv4Subnet&#x60; - App Processor ipv4 subnet   * &#x60;/bandwidth&#x60; - App Processor bandwidth   * &#x60;/changeLog/createdDateTime&#x60; - App Processor creation timestamp   * &#x60;/changeLog/updatedDateTime&#x60; - App Processor last updated timestamp   * &#x60;/changeLog/deletedDateTime&#x60; - App Processor deletion timestamp  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:   * &#x60;&#x3D;&#x60; - equal   * &#x60;!&#x3D;&#x60; - not equal   * &#x60;&gt;&#x60; - greater than   * &#x60;&lt;&#x60; - less than   * &#x60;IN&#x60; - in   * &#x60;NOT IN&#x60; - not in   * &#x60;LIKE&#x60; - like   * &#x60;ILIKE&#x60; - case-insensitive like   * &#x60;BETWEEN&#x60; - between   * &#x60;NOT BETWEEN&#x60; - not between  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAppProcessorSimpleExpression

`func NewAppProcessorSimpleExpression() *AppProcessorSimpleExpression`

NewAppProcessorSimpleExpression instantiates a new AppProcessorSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProcessorSimpleExpressionWithDefaults

`func NewAppProcessorSimpleExpressionWithDefaults() *AppProcessorSimpleExpression`

NewAppProcessorSimpleExpressionWithDefaults instantiates a new AppProcessorSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AppProcessorSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppProcessorSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppProcessorSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppProcessorSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *AppProcessorSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppProcessorSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppProcessorSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppProcessorSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *AppProcessorSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppProcessorSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppProcessorSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AppProcessorSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



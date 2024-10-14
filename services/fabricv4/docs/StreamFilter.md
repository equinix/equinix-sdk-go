# StreamFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/subject&#x60; - subject  * &#x60;/type&#x60; - type  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;in&#x60; - in  * &#x60;LIKE&#x60; - case-sensitive like  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]StreamFilterSimpleExpression**](StreamFilterSimpleExpression.md) |  | [optional] 

## Methods

### NewStreamFilter

`func NewStreamFilter() *StreamFilter`

NewStreamFilter instantiates a new StreamFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamFilterWithDefaults

`func NewStreamFilterWithDefaults() *StreamFilter`

NewStreamFilterWithDefaults instantiates a new StreamFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StreamFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *StreamFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StreamFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StreamFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StreamFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *StreamFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StreamFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StreamFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StreamFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *StreamFilter) GetOr() []StreamFilterSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *StreamFilter) GetOrOk() (*[]StreamFilterSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *StreamFilter) SetOr(v []StreamFilterSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *StreamFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



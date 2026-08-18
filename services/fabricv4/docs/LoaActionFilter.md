# LoaActionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**LoaActionFieldName**](LoaActionFieldName.md) |  | [optional] 
**Operator** | Pointer to [**LoaActionSearchSimpleExpressionsOperator**](LoaActionSearchSimpleExpressionsOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]LoaActionSearchSimpleExpressions**](LoaActionSearchSimpleExpressions.md) |  | [optional] 

## Methods

### NewLoaActionFilter

`func NewLoaActionFilter() *LoaActionFilter`

NewLoaActionFilter instantiates a new LoaActionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionFilterWithDefaults

`func NewLoaActionFilterWithDefaults() *LoaActionFilter`

NewLoaActionFilterWithDefaults instantiates a new LoaActionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *LoaActionFilter) GetProperty() LoaActionFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *LoaActionFilter) GetPropertyOk() (*LoaActionFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *LoaActionFilter) SetProperty(v LoaActionFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *LoaActionFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *LoaActionFilter) GetOperator() LoaActionSearchSimpleExpressionsOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LoaActionFilter) GetOperatorOk() (*LoaActionSearchSimpleExpressionsOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LoaActionFilter) SetOperator(v LoaActionSearchSimpleExpressionsOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LoaActionFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *LoaActionFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LoaActionFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LoaActionFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *LoaActionFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *LoaActionFilter) GetOr() []LoaActionSearchSimpleExpressions`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *LoaActionFilter) GetOrOk() (*[]LoaActionSearchSimpleExpressions, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *LoaActionFilter) SetOr(v []LoaActionSearchSimpleExpressions)`

SetOr sets Or field to given value.

### HasOr

`func (o *LoaActionFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



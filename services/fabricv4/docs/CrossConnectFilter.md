# CrossConnectFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**CrossConnectFieldName**](CrossConnectFieldName.md) |  | [optional] 
**Operator** | Pointer to [**LoaSimpleExpressionOperator**](LoaSimpleExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]CrossConnectSearchSimpleExpressions**](CrossConnectSearchSimpleExpressions.md) |  | [optional] 

## Methods

### NewCrossConnectFilter

`func NewCrossConnectFilter() *CrossConnectFilter`

NewCrossConnectFilter instantiates a new CrossConnectFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectFilterWithDefaults

`func NewCrossConnectFilterWithDefaults() *CrossConnectFilter`

NewCrossConnectFilterWithDefaults instantiates a new CrossConnectFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *CrossConnectFilter) GetProperty() CrossConnectFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CrossConnectFilter) GetPropertyOk() (*CrossConnectFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CrossConnectFilter) SetProperty(v CrossConnectFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CrossConnectFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CrossConnectFilter) GetOperator() LoaSimpleExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CrossConnectFilter) GetOperatorOk() (*LoaSimpleExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CrossConnectFilter) SetOperator(v LoaSimpleExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CrossConnectFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CrossConnectFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CrossConnectFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CrossConnectFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CrossConnectFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *CrossConnectFilter) GetOr() []CrossConnectSearchSimpleExpressions`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CrossConnectFilter) GetOrOk() (*[]CrossConnectSearchSimpleExpressions, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CrossConnectFilter) SetOr(v []CrossConnectSearchSimpleExpressions)`

SetOr sets Or field to given value.

### HasOr

`func (o *CrossConnectFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



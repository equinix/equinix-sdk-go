# StreamAssetFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/uuid&#x60; - Asset uuid  * &#x60;/streamUuid&#x60; - Stream uuid  * &#x60;/projectId&#x60; - Asset projectId  * &#x60;/_*&#x60; - all-category search  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]StreamAssetSimpleExpression**](StreamAssetSimpleExpression.md) |  | [optional] 

## Methods

### NewStreamAssetFilter

`func NewStreamAssetFilter() *StreamAssetFilter`

NewStreamAssetFilter instantiates a new StreamAssetFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamAssetFilterWithDefaults

`func NewStreamAssetFilterWithDefaults() *StreamAssetFilter`

NewStreamAssetFilterWithDefaults instantiates a new StreamAssetFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StreamAssetFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamAssetFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamAssetFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamAssetFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *StreamAssetFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StreamAssetFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StreamAssetFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StreamAssetFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *StreamAssetFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StreamAssetFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StreamAssetFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StreamAssetFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *StreamAssetFilter) GetOr() []StreamAssetSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *StreamAssetFilter) GetOrOk() (*[]StreamAssetSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *StreamAssetFilter) SetOr(v []StreamAssetSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *StreamAssetFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



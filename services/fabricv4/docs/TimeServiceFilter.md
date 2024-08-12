# TimeServiceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/project/projectId&#x60; - project id (mandatory)  * &#x60;/name&#x60; - Precision Time Service name  * &#x60;/uuid&#x60; - Precision Time Service uuid  * &#x60;/type&#x60; - Precision Time Service protocol  * &#x60;/state&#x60; - Precision Time Service status  * &#x60;/account/accountNumber&#x60; - Precision Time Service account number  * &#x60;/package/code&#x60; - Precision Time Service package  * &#x60;/_*&#x60; - all-category search  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]TimeServiceSimpleExpression**](TimeServiceSimpleExpression.md) |  | [optional] 

## Methods

### NewTimeServiceFilter

`func NewTimeServiceFilter() *TimeServiceFilter`

NewTimeServiceFilter instantiates a new TimeServiceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeServiceFilterWithDefaults

`func NewTimeServiceFilterWithDefaults() *TimeServiceFilter`

NewTimeServiceFilterWithDefaults instantiates a new TimeServiceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *TimeServiceFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *TimeServiceFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *TimeServiceFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *TimeServiceFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *TimeServiceFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TimeServiceFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TimeServiceFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *TimeServiceFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *TimeServiceFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TimeServiceFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TimeServiceFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *TimeServiceFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *TimeServiceFilter) GetOr() []TimeServiceSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *TimeServiceFilter) GetOrOk() (*[]TimeServiceSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *TimeServiceFilter) SetOr(v []TimeServiceSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *TimeServiceFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



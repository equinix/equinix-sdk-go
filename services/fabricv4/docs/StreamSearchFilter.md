# StreamSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/uuid&#x60; - Stream uuid  * &#x60;/name&#x60; - Stream name  * &#x60;/state&#x60; - Stream state (&#x60;DEPROVISIONED&#x60; returned only when explicitly requested)  * &#x60;/type&#x60; - Stream type  * &#x60;/description&#x60; - Stream description  * &#x60;/project/projectId&#x60; - Stream project id  * &#x60;/changeLog/createdDateTime&#x60; - Stream created date time  * &#x60;/changeLog/updatedDateTime&#x60; - Stream updated date time  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;BETWEEN&#x60; - between  * &#x60;NOT BETWEEN&#x60; - not between  * &#x60;LIKE&#x60; - like  * &#x60;ILIKE&#x60; - like case-insensitive  * &#x60;IN&#x60; - in  * &#x60;NOT IN&#x60; - not in  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]StreamSearchSimpleExpression**](StreamSearchSimpleExpression.md) |  | [optional] 

## Methods

### NewStreamSearchFilter

`func NewStreamSearchFilter() *StreamSearchFilter`

NewStreamSearchFilter instantiates a new StreamSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSearchFilterWithDefaults

`func NewStreamSearchFilterWithDefaults() *StreamSearchFilter`

NewStreamSearchFilterWithDefaults instantiates a new StreamSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StreamSearchFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamSearchFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamSearchFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamSearchFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *StreamSearchFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StreamSearchFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StreamSearchFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StreamSearchFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *StreamSearchFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StreamSearchFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StreamSearchFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StreamSearchFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *StreamSearchFilter) GetOr() []StreamSearchSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *StreamSearchFilter) GetOrOk() (*[]StreamSearchSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *StreamSearchFilter) SetOr(v []StreamSearchSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *StreamSearchFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



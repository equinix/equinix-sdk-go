# StreamSubscriptionSearchSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/uuid&#x60; - Subscription uuid  * &#x60;/streamUuid&#x60; - Stream uuid the subscription belongs to  * &#x60;/state&#x60; - Subscription state (&#x60;DEPROVISIONED&#x60; returned only when explicitly requested)  * &#x60;/name&#x60; - Subscription name  * &#x60;/project/projectId&#x60; - Subscription project id  * &#x60;/description&#x60; - Subscription description  * &#x60;/changeLog/createdDateTime&#x60; - Subscription created date time  * &#x60;/changeLog/updatedDateTime&#x60; - Subscription updated date time  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;BETWEEN&#x60; - between  * &#x60;NOT BETWEEN&#x60; - not between  * &#x60;LIKE&#x60; - like  * &#x60;ILIKE&#x60; - like case-insensitive  * &#x60;IN&#x60; - in  * &#x60;NOT IN&#x60; - not in  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStreamSubscriptionSearchSimpleExpression

`func NewStreamSubscriptionSearchSimpleExpression() *StreamSubscriptionSearchSimpleExpression`

NewStreamSubscriptionSearchSimpleExpression instantiates a new StreamSubscriptionSearchSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSearchSimpleExpressionWithDefaults

`func NewStreamSubscriptionSearchSimpleExpressionWithDefaults() *StreamSubscriptionSearchSimpleExpression`

NewStreamSubscriptionSearchSimpleExpressionWithDefaults instantiates a new StreamSubscriptionSearchSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StreamSubscriptionSearchSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamSubscriptionSearchSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamSubscriptionSearchSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamSubscriptionSearchSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *StreamSubscriptionSearchSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StreamSubscriptionSearchSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StreamSubscriptionSearchSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StreamSubscriptionSearchSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *StreamSubscriptionSearchSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StreamSubscriptionSearchSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StreamSubscriptionSearchSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StreamSubscriptionSearchSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



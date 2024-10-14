# CloudRouterActionsSearchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/type&#x60; - type of update  * &#x60;/state&#x60; - action state  * &#x60;/connection/uuid&#x60; - connection uuid associated  * &#x60;/_*&#x60; - all-category search  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;~*&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudRouterActionsSearchExpression

`func NewCloudRouterActionsSearchExpression() *CloudRouterActionsSearchExpression`

NewCloudRouterActionsSearchExpression instantiates a new CloudRouterActionsSearchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterActionsSearchExpressionWithDefaults

`func NewCloudRouterActionsSearchExpressionWithDefaults() *CloudRouterActionsSearchExpression`

NewCloudRouterActionsSearchExpressionWithDefaults instantiates a new CloudRouterActionsSearchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *CloudRouterActionsSearchExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterActionsSearchExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterActionsSearchExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterActionsSearchExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CloudRouterActionsSearchExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudRouterActionsSearchExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudRouterActionsSearchExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CloudRouterActionsSearchExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CloudRouterActionsSearchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CloudRouterActionsSearchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CloudRouterActionsSearchExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CloudRouterActionsSearchExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StreamFilterSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/subject&#x60; - subject  * &#x60;/type&#x60; - type  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;in&#x60; - in  * &#x60;LIKE&#x60; - case-sensitive like  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** | ### Supported event or metric names to use on filters with property /type:   * &#x60;*&#x60; - all events or metrics   * &#x60;equinix.fabric.port.*&#x60; - port events or metrics   * &#x60;equinix.fabric.connection.*&#x60; - connection events or metrics   * &#x60;equinix.fabric.router.*&#x60; - cloud router events   * &#x60;equinix.fabric.metro.*&#x60; - metro metrics   * &#x60;equinix.fabric.network.*&#x60; - network events   * &#x60;equinix.fabric.service_token.*&#x60; - service token events   * &#x60;equinix.network_edge.*&#x60; - network edge events   * &#x60;equinix.network_edge.acl.*&#x60; - network edge acl events   * &#x60;equinix.network_edge.device.*&#x60; - network edge device events   * &#x60;equinix.access_manager.*&#x60; - identity access manager events   * &#x60;equinix.access_manager.user.role.*&#x60; - identity access manager user role events ### Supported event or metric names to use on filters with property /subject:   * &#x60;*&#x60; - all events or metrics   * &#x60;/fabric/v4/ports/&lt;uuid&gt;&#x60; - port events or metrics   * &#x60;/fabric/v4/connections/&lt;uuid&gt;&#x60; - connection events or metrics   * &#x60;/fabric/v4/routers/&lt;uuid&gt;&#x60; - cloud router events   * &#x60;/fabric/v4/metros/&lt;metroCode&gt;&#x60; - metro metrics   * &#x60;/fabric/v4/networks/&lt;uuid&gt;&#x60; - network events   * &#x60;/fabric/v4/tokens/&lt;uuid&gt;&#x60; - service token events   * &#x60;/ne/v1/acl/&lt;uuid&gt;&#x60; - network edge acl events   * &#x60;/ne/v1/devices/&lt;uuid&gt;&#x60; - network edge device events   * &#x60;/am/v2/users/&lt;uuid&gt;/roleAssignments/&lt;uuid&gt;&#x60; - identity access manager events  | [optional] 

## Methods

### NewStreamFilterSimpleExpression

`func NewStreamFilterSimpleExpression() *StreamFilterSimpleExpression`

NewStreamFilterSimpleExpression instantiates a new StreamFilterSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamFilterSimpleExpressionWithDefaults

`func NewStreamFilterSimpleExpressionWithDefaults() *StreamFilterSimpleExpression`

NewStreamFilterSimpleExpressionWithDefaults instantiates a new StreamFilterSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StreamFilterSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamFilterSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamFilterSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamFilterSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *StreamFilterSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StreamFilterSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StreamFilterSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StreamFilterSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *StreamFilterSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StreamFilterSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StreamFilterSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StreamFilterSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



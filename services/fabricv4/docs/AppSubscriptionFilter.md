# AppSubscriptionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:   * &#x60;/project/projectId&#x60; - project id   * &#x60;/uuid&#x60; - App Subscription uuid   * &#x60;/state&#x60; - App Subscription status   * &#x60;/source/appLink/uuid&#x60; - Source App Link uuid   * &#x60;/source/ipSubnets&#x60; - Source App Link ip subnets   * &#x60;/target/appService/uuid&#x60; - Target App Service uuid   * &#x60;/target/geoScope&#x60; - Target App Service geo scope   * &#x60;/target/prioritization&#x60; - Target App Service prioritization   * &#x60;/changeLog/createdDateTime&#x60; - App Subscription creation timestamp   * &#x60;/changeLog/updatedDateTime&#x60; - App Subscription last updated timestamp   * &#x60;/changeLog/deletedDateTime&#x60; - App Subscription deletion timestamp  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:   * &#x60;&#x3D;&#x60; - equal   * &#x60;!&#x3D;&#x60; - not equal   * &#x60;&gt;&#x60; - greater than   * &#x60;&lt;&#x60; - less than   * &#x60;IN&#x60; - in   * &#x60;NOT IN&#x60; - not in   * &#x60;LIKE&#x60; - like   * &#x60;ILIKE&#x60; - case-insensitive like   * &#x60;BETWEEN&#x60; - between   * &#x60;NOT BETWEEN&#x60; - not between  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]AppSubscriptionSimpleExpression**](AppSubscriptionSimpleExpression.md) |  | [optional] 

## Methods

### NewAppSubscriptionFilter

`func NewAppSubscriptionFilter() *AppSubscriptionFilter`

NewAppSubscriptionFilter instantiates a new AppSubscriptionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionFilterWithDefaults

`func NewAppSubscriptionFilterWithDefaults() *AppSubscriptionFilter`

NewAppSubscriptionFilterWithDefaults instantiates a new AppSubscriptionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AppSubscriptionFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppSubscriptionFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppSubscriptionFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppSubscriptionFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *AppSubscriptionFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppSubscriptionFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppSubscriptionFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppSubscriptionFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *AppSubscriptionFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppSubscriptionFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppSubscriptionFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AppSubscriptionFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *AppSubscriptionFilter) GetOr() []AppSubscriptionSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *AppSubscriptionFilter) GetOrOk() (*[]AppSubscriptionSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *AppSubscriptionFilter) SetOr(v []AppSubscriptionSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *AppSubscriptionFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



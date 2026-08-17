# AppLinkFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:   * &#x60;/project/projectId&#x60; - project id   * &#x60;/uuid&#x60; - App Link uuid   * &#x60;/name&#x60; - App Link name   * &#x60;/description&#x60; - App Link description   * &#x60;/state&#x60; - App Link status   * &#x60;/ipv4Address&#x60; - App Link ipv4 address   * &#x60;/bandwidth&#x60; - App Link bandwidth   * &#x60;/changeLog/createdDateTime&#x60; - App Link creation timestamp   * &#x60;/changeLog/updatedDateTime&#x60; - App Link last updated timestamp   * &#x60;/changeLog/deletedDateTime&#x60; - App Link deletion timestamp  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:   * &#x60;&#x3D;&#x60; - equal   * &#x60;!&#x3D;&#x60; - not equal   * &#x60;&gt;&#x60; - greater than   * &#x60;&lt;&#x60; - less than   * &#x60;IN&#x60; - in   * &#x60;NOT IN&#x60; - not in   * &#x60;LIKE&#x60; - like   * &#x60;ILIKE&#x60; - case-insensitive like   * &#x60;BETWEEN&#x60; - between   * &#x60;NOT BETWEEN&#x60; - not between  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]AppLinkSimpleExpression**](AppLinkSimpleExpression.md) |  | [optional] 

## Methods

### NewAppLinkFilter

`func NewAppLinkFilter() *AppLinkFilter`

NewAppLinkFilter instantiates a new AppLinkFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkFilterWithDefaults

`func NewAppLinkFilterWithDefaults() *AppLinkFilter`

NewAppLinkFilterWithDefaults instantiates a new AppLinkFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AppLinkFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppLinkFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppLinkFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppLinkFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *AppLinkFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppLinkFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppLinkFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppLinkFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *AppLinkFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppLinkFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppLinkFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AppLinkFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *AppLinkFilter) GetOr() []AppLinkSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *AppLinkFilter) GetOrOk() (*[]AppLinkSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *AppLinkFilter) SetOr(v []AppLinkSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *AppLinkFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



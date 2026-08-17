# AppLinkAttachServiceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:   * &#x60;/uuid&#x60; - App Service attach to App Link uuid   * &#x60;/attachmentStatus&#x60; - App Service attach to App Link status  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:   * &#x60;&#x3D;&#x60; - equal   * &#x60;!&#x3D;&#x60; - not equal   * &#x60;IN&#x60; - in   * &#x60;NOT IN&#x60; - not in  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]AppLinkAttachServiceSimpleExpression**](AppLinkAttachServiceSimpleExpression.md) |  | [optional] 

## Methods

### NewAppLinkAttachServiceFilter

`func NewAppLinkAttachServiceFilter() *AppLinkAttachServiceFilter`

NewAppLinkAttachServiceFilter instantiates a new AppLinkAttachServiceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachServiceFilterWithDefaults

`func NewAppLinkAttachServiceFilterWithDefaults() *AppLinkAttachServiceFilter`

NewAppLinkAttachServiceFilterWithDefaults instantiates a new AppLinkAttachServiceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AppLinkAttachServiceFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppLinkAttachServiceFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppLinkAttachServiceFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppLinkAttachServiceFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *AppLinkAttachServiceFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppLinkAttachServiceFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppLinkAttachServiceFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppLinkAttachServiceFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *AppLinkAttachServiceFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppLinkAttachServiceFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppLinkAttachServiceFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AppLinkAttachServiceFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *AppLinkAttachServiceFilter) GetOr() []AppLinkAttachServiceSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *AppLinkAttachServiceFilter) GetOrOk() (*[]AppLinkAttachServiceSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *AppLinkAttachServiceFilter) SetOr(v []AppLinkAttachServiceSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *AppLinkAttachServiceFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AppLinkAttachDomainFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:   * &#x60;/uuid&#x60; - App Domain attach to App Link uuid   * &#x60;/attachmentStatus&#x60; - App Domain attach to App Link status  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:   * &#x60;&#x3D;&#x60; - equal   * &#x60;!&#x3D;&#x60; - not equal   * &#x60;IN&#x60; - in   * &#x60;NOT IN&#x60; - not in  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]AppLinkAttachDomainSimpleExpression**](AppLinkAttachDomainSimpleExpression.md) |  | [optional] 

## Methods

### NewAppLinkAttachDomainFilter

`func NewAppLinkAttachDomainFilter() *AppLinkAttachDomainFilter`

NewAppLinkAttachDomainFilter instantiates a new AppLinkAttachDomainFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachDomainFilterWithDefaults

`func NewAppLinkAttachDomainFilterWithDefaults() *AppLinkAttachDomainFilter`

NewAppLinkAttachDomainFilterWithDefaults instantiates a new AppLinkAttachDomainFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AppLinkAttachDomainFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppLinkAttachDomainFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppLinkAttachDomainFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppLinkAttachDomainFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *AppLinkAttachDomainFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AppLinkAttachDomainFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AppLinkAttachDomainFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AppLinkAttachDomainFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *AppLinkAttachDomainFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppLinkAttachDomainFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppLinkAttachDomainFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AppLinkAttachDomainFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *AppLinkAttachDomainFilter) GetOr() []AppLinkAttachDomainSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *AppLinkAttachDomainFilter) GetOrOk() (*[]AppLinkAttachDomainSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *AppLinkAttachDomainFilter) SetOr(v []AppLinkAttachDomainSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *AppLinkAttachDomainFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



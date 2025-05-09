# CloudRouterCommandSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/type&#x60; - type of command  * &#x60;/name&#x60; - name of command  * &#x60;/state&#x60; - state of command  * &#x60;/request/destination&#x60; - destination of command request  * &#x60;/request/sourceConnection/uuid&#x60; - source connection uuid  * &#x60;/_*&#x60; - all-category search  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;~*&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]CloudRouterCommandSearchExpression**](CloudRouterCommandSearchExpression.md) |  | [optional] 

## Methods

### NewCloudRouterCommandSearchFilter

`func NewCloudRouterCommandSearchFilter() *CloudRouterCommandSearchFilter`

NewCloudRouterCommandSearchFilter instantiates a new CloudRouterCommandSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandSearchFilterWithDefaults

`func NewCloudRouterCommandSearchFilterWithDefaults() *CloudRouterCommandSearchFilter`

NewCloudRouterCommandSearchFilterWithDefaults instantiates a new CloudRouterCommandSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *CloudRouterCommandSearchFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterCommandSearchFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterCommandSearchFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterCommandSearchFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CloudRouterCommandSearchFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudRouterCommandSearchFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudRouterCommandSearchFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CloudRouterCommandSearchFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CloudRouterCommandSearchFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CloudRouterCommandSearchFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CloudRouterCommandSearchFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CloudRouterCommandSearchFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *CloudRouterCommandSearchFilter) GetOr() []CloudRouterCommandSearchExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CloudRouterCommandSearchFilter) GetOrOk() (*[]CloudRouterCommandSearchExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CloudRouterCommandSearchFilter) SetOr(v []CloudRouterCommandSearchExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *CloudRouterCommandSearchFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



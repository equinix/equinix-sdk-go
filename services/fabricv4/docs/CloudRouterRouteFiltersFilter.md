# CloudRouterRouteFiltersFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CloudRouterRouteFilterExpression**](CloudRouterRouteFilterExpression.md) |  | [optional] 
**Or** | Pointer to [**[]CloudRouterRouteFilterExpression**](CloudRouterRouteFilterExpression.md) |  | [optional] 

## Methods

### NewCloudRouterRouteFiltersFilter

`func NewCloudRouterRouteFiltersFilter() *CloudRouterRouteFiltersFilter`

NewCloudRouterRouteFiltersFilter instantiates a new CloudRouterRouteFiltersFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterRouteFiltersFilterWithDefaults

`func NewCloudRouterRouteFiltersFilterWithDefaults() *CloudRouterRouteFiltersFilter`

NewCloudRouterRouteFiltersFilterWithDefaults instantiates a new CloudRouterRouteFiltersFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *CloudRouterRouteFiltersFilter) GetAnd() []CloudRouterRouteFilterExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *CloudRouterRouteFiltersFilter) GetAndOk() (*[]CloudRouterRouteFilterExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *CloudRouterRouteFiltersFilter) SetAnd(v []CloudRouterRouteFilterExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *CloudRouterRouteFiltersFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *CloudRouterRouteFiltersFilter) GetOr() []CloudRouterRouteFilterExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CloudRouterRouteFiltersFilter) GetOrOk() (*[]CloudRouterRouteFilterExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CloudRouterRouteFiltersFilter) SetOr(v []CloudRouterRouteFilterExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *CloudRouterRouteFiltersFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



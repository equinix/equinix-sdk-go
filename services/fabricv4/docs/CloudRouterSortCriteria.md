# CloudRouterSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**CloudRouterSortDirection**](CloudRouterSortDirection.md) |  | [optional] [default to CLOUDROUTERSORTDIRECTION_DESC]
**Property** | Pointer to [**CloudRouterSortBy**](CloudRouterSortBy.md) |  | [optional] [default to CLOUDROUTERSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewCloudRouterSortCriteria

`func NewCloudRouterSortCriteria() *CloudRouterSortCriteria`

NewCloudRouterSortCriteria instantiates a new CloudRouterSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterSortCriteriaWithDefaults

`func NewCloudRouterSortCriteriaWithDefaults() *CloudRouterSortCriteria`

NewCloudRouterSortCriteriaWithDefaults instantiates a new CloudRouterSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CloudRouterSortCriteria) GetDirection() CloudRouterSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudRouterSortCriteria) GetDirectionOk() (*CloudRouterSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudRouterSortCriteria) SetDirection(v CloudRouterSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudRouterSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *CloudRouterSortCriteria) GetProperty() CloudRouterSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterSortCriteria) GetPropertyOk() (*CloudRouterSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterSortCriteria) SetProperty(v CloudRouterSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



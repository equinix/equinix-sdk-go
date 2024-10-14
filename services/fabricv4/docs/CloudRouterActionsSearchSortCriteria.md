# CloudRouterActionsSearchSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**CloudRouterActionsSearchSortDirection**](CloudRouterActionsSearchSortDirection.md) |  | [optional] [default to CLOUDROUTERACTIONSSEARCHSORTDIRECTION_DESC]
**Property** | Pointer to [**CloudRouterActionsSearchSortBy**](CloudRouterActionsSearchSortBy.md) |  | [optional] [default to CLOUDROUTERACTIONSSEARCHSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewCloudRouterActionsSearchSortCriteria

`func NewCloudRouterActionsSearchSortCriteria() *CloudRouterActionsSearchSortCriteria`

NewCloudRouterActionsSearchSortCriteria instantiates a new CloudRouterActionsSearchSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterActionsSearchSortCriteriaWithDefaults

`func NewCloudRouterActionsSearchSortCriteriaWithDefaults() *CloudRouterActionsSearchSortCriteria`

NewCloudRouterActionsSearchSortCriteriaWithDefaults instantiates a new CloudRouterActionsSearchSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CloudRouterActionsSearchSortCriteria) GetDirection() CloudRouterActionsSearchSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudRouterActionsSearchSortCriteria) GetDirectionOk() (*CloudRouterActionsSearchSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudRouterActionsSearchSortCriteria) SetDirection(v CloudRouterActionsSearchSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudRouterActionsSearchSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *CloudRouterActionsSearchSortCriteria) GetProperty() CloudRouterActionsSearchSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterActionsSearchSortCriteria) GetPropertyOk() (*CloudRouterActionsSearchSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterActionsSearchSortCriteria) SetProperty(v CloudRouterActionsSearchSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterActionsSearchSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



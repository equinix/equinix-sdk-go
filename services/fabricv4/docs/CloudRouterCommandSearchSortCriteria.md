# CloudRouterCommandSearchSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**CloudRouterCommandSearchSortDirection**](CloudRouterCommandSearchSortDirection.md) |  | [optional] [default to CLOUDROUTERCOMMANDSEARCHSORTDIRECTION_DESC]
**Property** | Pointer to [**CloudRouterCommandSearchSortBy**](CloudRouterCommandSearchSortBy.md) |  | [optional] [default to CLOUDROUTERCOMMANDSEARCHSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewCloudRouterCommandSearchSortCriteria

`func NewCloudRouterCommandSearchSortCriteria() *CloudRouterCommandSearchSortCriteria`

NewCloudRouterCommandSearchSortCriteria instantiates a new CloudRouterCommandSearchSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandSearchSortCriteriaWithDefaults

`func NewCloudRouterCommandSearchSortCriteriaWithDefaults() *CloudRouterCommandSearchSortCriteria`

NewCloudRouterCommandSearchSortCriteriaWithDefaults instantiates a new CloudRouterCommandSearchSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CloudRouterCommandSearchSortCriteria) GetDirection() CloudRouterCommandSearchSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudRouterCommandSearchSortCriteria) GetDirectionOk() (*CloudRouterCommandSearchSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudRouterCommandSearchSortCriteria) SetDirection(v CloudRouterCommandSearchSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudRouterCommandSearchSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *CloudRouterCommandSearchSortCriteria) GetProperty() CloudRouterCommandSearchSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterCommandSearchSortCriteria) GetPropertyOk() (*CloudRouterCommandSearchSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterCommandSearchSortCriteria) SetProperty(v CloudRouterCommandSearchSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterCommandSearchSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



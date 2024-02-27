# ServiceProfileSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ServiceProfileSortDirection**](ServiceProfileSortDirection.md) |  | [optional] [default to SERVICEPROFILESORTDIRECTION_DESC]
**Property** | Pointer to [**ServiceProfileSortBy**](ServiceProfileSortBy.md) |  | [optional] [default to SERVICEPROFILESORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewServiceProfileSortCriteria

`func NewServiceProfileSortCriteria() *ServiceProfileSortCriteria`

NewServiceProfileSortCriteria instantiates a new ServiceProfileSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileSortCriteriaWithDefaults

`func NewServiceProfileSortCriteriaWithDefaults() *ServiceProfileSortCriteria`

NewServiceProfileSortCriteriaWithDefaults instantiates a new ServiceProfileSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ServiceProfileSortCriteria) GetDirection() ServiceProfileSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ServiceProfileSortCriteria) GetDirectionOk() (*ServiceProfileSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ServiceProfileSortCriteria) SetDirection(v ServiceProfileSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ServiceProfileSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *ServiceProfileSortCriteria) GetProperty() ServiceProfileSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ServiceProfileSortCriteria) GetPropertyOk() (*ServiceProfileSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ServiceProfileSortCriteria) SetProperty(v ServiceProfileSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ServiceProfileSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



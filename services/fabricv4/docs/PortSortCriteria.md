# PortSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**PortSortDirection**](PortSortDirection.md) |  | [optional] [default to PORTSORTDIRECTION_DESC]
**Property** | Pointer to [**PortSortBy**](PortSortBy.md) |  | [optional] [default to PORTSORTBY_DEVICE_NAME]

## Methods

### NewPortSortCriteria

`func NewPortSortCriteria() *PortSortCriteria`

NewPortSortCriteria instantiates a new PortSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortSortCriteriaWithDefaults

`func NewPortSortCriteriaWithDefaults() *PortSortCriteria`

NewPortSortCriteriaWithDefaults instantiates a new PortSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *PortSortCriteria) GetDirection() PortSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PortSortCriteria) GetDirectionOk() (*PortSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PortSortCriteria) SetDirection(v PortSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PortSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *PortSortCriteria) GetProperty() PortSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PortSortCriteria) GetPropertyOk() (*PortSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PortSortCriteria) SetProperty(v PortSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PortSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RfAttachmentSortItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**RfAttachmentSortItemProperty**](RfAttachmentSortItemProperty.md) |  | [optional] [default to RFATTACHMENTSORTITEMPROPERTY_CHANGE_LOG_UPDATED_DATE_TIME]
**Direction** | Pointer to [**SortItemDirection**](SortItemDirection.md) |  | [optional] [default to SORTITEMDIRECTION_DESC]

## Methods

### NewRfAttachmentSortItem

`func NewRfAttachmentSortItem() *RfAttachmentSortItem`

NewRfAttachmentSortItem instantiates a new RfAttachmentSortItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfAttachmentSortItemWithDefaults

`func NewRfAttachmentSortItemWithDefaults() *RfAttachmentSortItem`

NewRfAttachmentSortItemWithDefaults instantiates a new RfAttachmentSortItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *RfAttachmentSortItem) GetProperty() RfAttachmentSortItemProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RfAttachmentSortItem) GetPropertyOk() (*RfAttachmentSortItemProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RfAttachmentSortItem) SetProperty(v RfAttachmentSortItemProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RfAttachmentSortItem) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetDirection

`func (o *RfAttachmentSortItem) GetDirection() SortItemDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RfAttachmentSortItem) GetDirectionOk() (*SortItemDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RfAttachmentSortItem) SetDirection(v SortItemDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RfAttachmentSortItem) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



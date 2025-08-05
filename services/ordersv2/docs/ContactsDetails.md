# ContactsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ContactsDetailsType**](ContactsDetailsType.md) |  | [optional] 
**Value** | Pointer to **string** | Value of contact details &#x60;type&#x60; | [optional] 

## Methods

### NewContactsDetails

`func NewContactsDetails() *ContactsDetails`

NewContactsDetails instantiates a new ContactsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsDetailsWithDefaults

`func NewContactsDetailsWithDefaults() *ContactsDetails`

NewContactsDetailsWithDefaults instantiates a new ContactsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContactsDetails) GetType() ContactsDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactsDetails) GetTypeOk() (*ContactsDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactsDetails) SetType(v ContactsDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *ContactsDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ContactsDetails) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactsDetails) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactsDetails) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactsDetails) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



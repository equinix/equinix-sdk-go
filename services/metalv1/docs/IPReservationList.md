# IPReservationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | Pointer to [**[]IPReservationListIpAddressesInner**](IPReservationListIpAddressesInner.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewIPReservationList

`func NewIPReservationList() *IPReservationList`

NewIPReservationList instantiates a new IPReservationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPReservationListWithDefaults

`func NewIPReservationListWithDefaults() *IPReservationList`

NewIPReservationListWithDefaults instantiates a new IPReservationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *IPReservationList) GetIpAddresses() []IPReservationListIpAddressesInner`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *IPReservationList) GetIpAddressesOk() (*[]IPReservationListIpAddressesInner, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *IPReservationList) SetIpAddresses(v []IPReservationListIpAddressesInner)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *IPReservationList) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMeta

`func (o *IPReservationList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IPReservationList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IPReservationList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IPReservationList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



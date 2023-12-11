# IPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to [**IPAddressAddressFamily**](IPAddressAddressFamily.md) |  | [optional] 
**Cidr** | Pointer to **int32** | Cidr Size for the IP Block created. Valid values depends on the operating system being provisioned. (28..32 for IPv4 addresses, 124..127 for IPv6 addresses) | [optional] 
**IpReservations** | Pointer to **[]string** | UUIDs of any IP reservations to use when assigning IPs | [optional] 
**Public** | Pointer to **bool** | Address Type for IP Address | [optional] [default to true]

## Methods

### NewIPAddress

`func NewIPAddress() *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *IPAddress) GetAddressFamily() IPAddressAddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *IPAddress) GetAddressFamilyOk() (*IPAddressAddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *IPAddress) SetAddressFamily(v IPAddressAddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *IPAddress) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetCidr

`func (o *IPAddress) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IPAddress) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IPAddress) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IPAddress) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetIpReservations

`func (o *IPAddress) GetIpReservations() []string`

GetIpReservations returns the IpReservations field if non-nil, zero value otherwise.

### GetIpReservationsOk

`func (o *IPAddress) GetIpReservationsOk() (*[]string, bool)`

GetIpReservationsOk returns a tuple with the IpReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservations

`func (o *IPAddress) SetIpReservations(v []string)`

SetIpReservations sets IpReservations field to given value.

### HasIpReservations

`func (o *IPAddress) HasIpReservations() bool`

HasIpReservations returns a boolean if a field has been set.

### GetPublic

`func (o *IPAddress) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPAddress) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPAddress) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPAddress) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BGPSessionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to [**BGPSessionInputAddressFamily**](BGPSessionInputAddressFamily.md) |  | [optional] 
**DefaultRoute** | Pointer to **bool** | Set the default route policy. | [optional] [default to false]

## Methods

### NewBGPSessionInput

`func NewBGPSessionInput() *BGPSessionInput`

NewBGPSessionInput instantiates a new BGPSessionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPSessionInputWithDefaults

`func NewBGPSessionInputWithDefaults() *BGPSessionInput`

NewBGPSessionInputWithDefaults instantiates a new BGPSessionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *BGPSessionInput) GetAddressFamily() BGPSessionInputAddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *BGPSessionInput) GetAddressFamilyOk() (*BGPSessionInputAddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *BGPSessionInput) SetAddressFamily(v BGPSessionInputAddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *BGPSessionInput) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *BGPSessionInput) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *BGPSessionInput) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *BGPSessionInput) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *BGPSessionInput) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



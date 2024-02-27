# RoutingProtocolDirectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RoutingProtocolDirectTypeType**](RoutingProtocolDirectTypeType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**DirectIpv4** | Pointer to [**DirectConnectionIpv4**](DirectConnectionIpv4.md) |  | [optional] 
**DirectIpv6** | Pointer to [**DirectConnectionIpv6**](DirectConnectionIpv6.md) |  | [optional] 

## Methods

### NewRoutingProtocolDirectType

`func NewRoutingProtocolDirectType(type_ RoutingProtocolDirectTypeType, ) *RoutingProtocolDirectType`

NewRoutingProtocolDirectType instantiates a new RoutingProtocolDirectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolDirectTypeWithDefaults

`func NewRoutingProtocolDirectTypeWithDefaults() *RoutingProtocolDirectType`

NewRoutingProtocolDirectTypeWithDefaults instantiates a new RoutingProtocolDirectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocolDirectType) GetType() RoutingProtocolDirectTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolDirectType) GetTypeOk() (*RoutingProtocolDirectTypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolDirectType) SetType(v RoutingProtocolDirectTypeType)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolDirectType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolDirectType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolDirectType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolDirectType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirectIpv4

`func (o *RoutingProtocolDirectType) GetDirectIpv4() DirectConnectionIpv4`

GetDirectIpv4 returns the DirectIpv4 field if non-nil, zero value otherwise.

### GetDirectIpv4Ok

`func (o *RoutingProtocolDirectType) GetDirectIpv4Ok() (*DirectConnectionIpv4, bool)`

GetDirectIpv4Ok returns a tuple with the DirectIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv4

`func (o *RoutingProtocolDirectType) SetDirectIpv4(v DirectConnectionIpv4)`

SetDirectIpv4 sets DirectIpv4 field to given value.

### HasDirectIpv4

`func (o *RoutingProtocolDirectType) HasDirectIpv4() bool`

HasDirectIpv4 returns a boolean if a field has been set.

### GetDirectIpv6

`func (o *RoutingProtocolDirectType) GetDirectIpv6() DirectConnectionIpv6`

GetDirectIpv6 returns the DirectIpv6 field if non-nil, zero value otherwise.

### GetDirectIpv6Ok

`func (o *RoutingProtocolDirectType) GetDirectIpv6Ok() (*DirectConnectionIpv6, bool)`

GetDirectIpv6Ok returns a tuple with the DirectIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv6

`func (o *RoutingProtocolDirectType) SetDirectIpv6(v DirectConnectionIpv6)`

SetDirectIpv6 sets DirectIpv6 field to given value.

### HasDirectIpv6

`func (o *RoutingProtocolDirectType) HasDirectIpv6() bool`

HasDirectIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



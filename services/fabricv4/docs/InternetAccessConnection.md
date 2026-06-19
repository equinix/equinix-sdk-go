# InternetAccessConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | URI of the connection | 
**Uuid** | **string** | Unique identifier for the connection | 
**PeeringIpv4** | Pointer to [**InternetAccessPeeringIpv4**](InternetAccessPeeringIpv4.md) |  | [optional] 
**PeeringIpv6** | Pointer to [**InternetAccessPeeringIpv6**](InternetAccessPeeringIpv6.md) |  | [optional] 

## Methods

### NewInternetAccessConnection

`func NewInternetAccessConnection(href string, uuid string, ) *InternetAccessConnection`

NewInternetAccessConnection instantiates a new InternetAccessConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessConnectionWithDefaults

`func NewInternetAccessConnectionWithDefaults() *InternetAccessConnection`

NewInternetAccessConnectionWithDefaults instantiates a new InternetAccessConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *InternetAccessConnection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InternetAccessConnection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InternetAccessConnection) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *InternetAccessConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InternetAccessConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InternetAccessConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetPeeringIpv4

`func (o *InternetAccessConnection) GetPeeringIpv4() InternetAccessPeeringIpv4`

GetPeeringIpv4 returns the PeeringIpv4 field if non-nil, zero value otherwise.

### GetPeeringIpv4Ok

`func (o *InternetAccessConnection) GetPeeringIpv4Ok() (*InternetAccessPeeringIpv4, bool)`

GetPeeringIpv4Ok returns a tuple with the PeeringIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringIpv4

`func (o *InternetAccessConnection) SetPeeringIpv4(v InternetAccessPeeringIpv4)`

SetPeeringIpv4 sets PeeringIpv4 field to given value.

### HasPeeringIpv4

`func (o *InternetAccessConnection) HasPeeringIpv4() bool`

HasPeeringIpv4 returns a boolean if a field has been set.

### GetPeeringIpv6

`func (o *InternetAccessConnection) GetPeeringIpv6() InternetAccessPeeringIpv6`

GetPeeringIpv6 returns the PeeringIpv6 field if non-nil, zero value otherwise.

### GetPeeringIpv6Ok

`func (o *InternetAccessConnection) GetPeeringIpv6Ok() (*InternetAccessPeeringIpv6, bool)`

GetPeeringIpv6Ok returns a tuple with the PeeringIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringIpv6

`func (o *InternetAccessConnection) SetPeeringIpv6(v InternetAccessPeeringIpv6)`

SetPeeringIpv6 sets PeeringIpv6 field to given value.

### HasPeeringIpv6

`func (o *InternetAccessConnection) HasPeeringIpv6() bool`

HasPeeringIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



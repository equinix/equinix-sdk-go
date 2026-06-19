# InternetAccessConnectionDirectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Unique identifier for the connection | 
**PeeringIpv4** | Pointer to [**InternetAccessPeeringIpv4Request**](InternetAccessPeeringIpv4Request.md) |  | [optional] 
**PeeringIpv6** | Pointer to [**InternetAccessPeeringIpv6Request**](InternetAccessPeeringIpv6Request.md) |  | [optional] 

## Methods

### NewInternetAccessConnectionDirectRequest

`func NewInternetAccessConnectionDirectRequest(uuid string, ) *InternetAccessConnectionDirectRequest`

NewInternetAccessConnectionDirectRequest instantiates a new InternetAccessConnectionDirectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessConnectionDirectRequestWithDefaults

`func NewInternetAccessConnectionDirectRequestWithDefaults() *InternetAccessConnectionDirectRequest`

NewInternetAccessConnectionDirectRequestWithDefaults instantiates a new InternetAccessConnectionDirectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *InternetAccessConnectionDirectRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InternetAccessConnectionDirectRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InternetAccessConnectionDirectRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetPeeringIpv4

`func (o *InternetAccessConnectionDirectRequest) GetPeeringIpv4() InternetAccessPeeringIpv4Request`

GetPeeringIpv4 returns the PeeringIpv4 field if non-nil, zero value otherwise.

### GetPeeringIpv4Ok

`func (o *InternetAccessConnectionDirectRequest) GetPeeringIpv4Ok() (*InternetAccessPeeringIpv4Request, bool)`

GetPeeringIpv4Ok returns a tuple with the PeeringIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringIpv4

`func (o *InternetAccessConnectionDirectRequest) SetPeeringIpv4(v InternetAccessPeeringIpv4Request)`

SetPeeringIpv4 sets PeeringIpv4 field to given value.

### HasPeeringIpv4

`func (o *InternetAccessConnectionDirectRequest) HasPeeringIpv4() bool`

HasPeeringIpv4 returns a boolean if a field has been set.

### GetPeeringIpv6

`func (o *InternetAccessConnectionDirectRequest) GetPeeringIpv6() InternetAccessPeeringIpv6Request`

GetPeeringIpv6 returns the PeeringIpv6 field if non-nil, zero value otherwise.

### GetPeeringIpv6Ok

`func (o *InternetAccessConnectionDirectRequest) GetPeeringIpv6Ok() (*InternetAccessPeeringIpv6Request, bool)`

GetPeeringIpv6Ok returns a tuple with the PeeringIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringIpv6

`func (o *InternetAccessConnectionDirectRequest) SetPeeringIpv6(v InternetAccessPeeringIpv6Request)`

SetPeeringIpv6 sets PeeringIpv6 field to given value.

### HasPeeringIpv6

`func (o *InternetAccessConnectionDirectRequest) HasPeeringIpv6() bool`

HasPeeringIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



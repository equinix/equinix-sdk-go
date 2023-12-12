# PortData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address is set for NetworkPort ports | [optional] 
**Bonded** | Pointer to **bool** | Bonded is true for NetworkPort ports in a bond and NetworkBondPort ports that are active | [optional] 

## Methods

### NewPortData

`func NewPortData() *PortData`

NewPortData instantiates a new PortData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDataWithDefaults

`func NewPortDataWithDefaults() *PortData`

NewPortDataWithDefaults instantiates a new PortData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *PortData) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PortData) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PortData) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PortData) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetBonded

`func (o *PortData) GetBonded() bool`

GetBonded returns the Bonded field if non-nil, zero value otherwise.

### GetBondedOk

`func (o *PortData) GetBondedOk() (*bool, bool)`

GetBondedOk returns a tuple with the Bonded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonded

`func (o *PortData) SetBonded(v bool)`

SetBonded sets Bonded field to given value.

### HasBonded

`func (o *PortData) HasBonded() bool`

HasBonded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



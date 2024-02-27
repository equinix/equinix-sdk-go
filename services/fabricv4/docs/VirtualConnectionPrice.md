# VirtualConnectionPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Either uuid or rest of attributes are required | [optional] 
**Type** | Pointer to [**VirtualConnectionPriceConnectionType**](VirtualConnectionPriceConnectionType.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** |  | [optional] 
**ASide** | Pointer to [**VirtualConnectionPriceASide**](VirtualConnectionPriceASide.md) |  | [optional] 
**ZSide** | Pointer to [**VirtualConnectionPriceZSide**](VirtualConnectionPriceZSide.md) |  | [optional] 

## Methods

### NewVirtualConnectionPrice

`func NewVirtualConnectionPrice() *VirtualConnectionPrice`

NewVirtualConnectionPrice instantiates a new VirtualConnectionPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualConnectionPriceWithDefaults

`func NewVirtualConnectionPriceWithDefaults() *VirtualConnectionPrice`

NewVirtualConnectionPriceWithDefaults instantiates a new VirtualConnectionPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *VirtualConnectionPrice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualConnectionPrice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualConnectionPrice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualConnectionPrice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *VirtualConnectionPrice) GetType() VirtualConnectionPriceConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualConnectionPrice) GetTypeOk() (*VirtualConnectionPriceConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualConnectionPrice) SetType(v VirtualConnectionPriceConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualConnectionPrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBandwidth

`func (o *VirtualConnectionPrice) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *VirtualConnectionPrice) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *VirtualConnectionPrice) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *VirtualConnectionPrice) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetASide

`func (o *VirtualConnectionPrice) GetASide() VirtualConnectionPriceASide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *VirtualConnectionPrice) GetASideOk() (*VirtualConnectionPriceASide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *VirtualConnectionPrice) SetASide(v VirtualConnectionPriceASide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *VirtualConnectionPrice) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *VirtualConnectionPrice) GetZSide() VirtualConnectionPriceZSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *VirtualConnectionPrice) GetZSideOk() (*VirtualConnectionPriceZSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *VirtualConnectionPrice) SetZSide(v VirtualConnectionPriceZSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *VirtualConnectionPrice) HasZSide() bool`

HasZSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bond** | Pointer to [**BondPortData**](BondPortData.md) |  | [optional] 
**Data** | Pointer to [**PortData**](PortData.md) |  | [optional] 
**DisbondOperationSupported** | Pointer to **bool** | Indicates whether or not the bond can be broken on the port (when applicable). | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PortType**](PortType.md) |  | [optional] 
**NetworkType** | Pointer to [**PortNetworkType**](PortNetworkType.md) |  | [optional] 
**NativeVirtualNetwork** | Pointer to [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 
**VirtualNetworks** | Pointer to [**[]VirtualNetwork**](VirtualNetwork.md) |  | [optional] 

## Methods

### NewPort

`func NewPort() *Port`

NewPort instantiates a new Port object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortWithDefaults

`func NewPortWithDefaults() *Port`

NewPortWithDefaults instantiates a new Port object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBond

`func (o *Port) GetBond() BondPortData`

GetBond returns the Bond field if non-nil, zero value otherwise.

### GetBondOk

`func (o *Port) GetBondOk() (*BondPortData, bool)`

GetBondOk returns a tuple with the Bond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBond

`func (o *Port) SetBond(v BondPortData)`

SetBond sets Bond field to given value.

### HasBond

`func (o *Port) HasBond() bool`

HasBond returns a boolean if a field has been set.

### GetData

`func (o *Port) GetData() PortData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Port) GetDataOk() (*PortData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Port) SetData(v PortData)`

SetData sets Data field to given value.

### HasData

`func (o *Port) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisbondOperationSupported

`func (o *Port) GetDisbondOperationSupported() bool`

GetDisbondOperationSupported returns the DisbondOperationSupported field if non-nil, zero value otherwise.

### GetDisbondOperationSupportedOk

`func (o *Port) GetDisbondOperationSupportedOk() (*bool, bool)`

GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbondOperationSupported

`func (o *Port) SetDisbondOperationSupported(v bool)`

SetDisbondOperationSupported sets DisbondOperationSupported field to given value.

### HasDisbondOperationSupported

`func (o *Port) HasDisbondOperationSupported() bool`

HasDisbondOperationSupported returns a boolean if a field has been set.

### GetHref

`func (o *Port) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Port) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Port) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Port) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Port) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Port) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Port) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Port) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Port) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Port) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Port) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Port) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Port) GetType() PortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Port) GetTypeOk() (*PortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Port) SetType(v PortType)`

SetType sets Type field to given value.

### HasType

`func (o *Port) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkType

`func (o *Port) GetNetworkType() PortNetworkType`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Port) GetNetworkTypeOk() (*PortNetworkType, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Port) SetNetworkType(v PortNetworkType)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Port) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNativeVirtualNetwork

`func (o *Port) GetNativeVirtualNetwork() VirtualNetwork`

GetNativeVirtualNetwork returns the NativeVirtualNetwork field if non-nil, zero value otherwise.

### GetNativeVirtualNetworkOk

`func (o *Port) GetNativeVirtualNetworkOk() (*VirtualNetwork, bool)`

GetNativeVirtualNetworkOk returns a tuple with the NativeVirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVirtualNetwork

`func (o *Port) SetNativeVirtualNetwork(v VirtualNetwork)`

SetNativeVirtualNetwork sets NativeVirtualNetwork field to given value.

### HasNativeVirtualNetwork

`func (o *Port) HasNativeVirtualNetwork() bool`

HasNativeVirtualNetwork returns a boolean if a field has been set.

### GetVirtualNetworks

`func (o *Port) GetVirtualNetworks() []VirtualNetwork`

GetVirtualNetworks returns the VirtualNetworks field if non-nil, zero value otherwise.

### GetVirtualNetworksOk

`func (o *Port) GetVirtualNetworksOk() (*[]VirtualNetwork, bool)`

GetVirtualNetworksOk returns a tuple with the VirtualNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworks

`func (o *Port) SetVirtualNetworks(v []VirtualNetwork)`

SetVirtualNetworks sets VirtualNetworks field to given value.

### HasVirtualNetworks

`func (o *Port) HasVirtualNetworks() bool`

HasVirtualNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



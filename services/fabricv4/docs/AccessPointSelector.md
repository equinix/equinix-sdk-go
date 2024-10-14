# AccessPointSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AccessPointSelectorType**](AccessPointSelectorType.md) |  | [optional] 
**Port** | Pointer to [**SimplifiedMetadataEntity**](SimplifiedMetadataEntity.md) |  | [optional] 
**LinkProtocol** | Pointer to [**SimplifiedLinkProtocol**](SimplifiedLinkProtocol.md) |  | [optional] 
**VirtualDevice** | Pointer to [**SimplifiedVirtualDevice**](SimplifiedVirtualDevice.md) |  | [optional] 
**Interface** | Pointer to [**VirtualDeviceInterface**](VirtualDeviceInterface.md) |  | [optional] 
**Network** | Pointer to [**SimplifiedTokenNetwork**](SimplifiedTokenNetwork.md) |  | [optional] 

## Methods

### NewAccessPointSelector

`func NewAccessPointSelector() *AccessPointSelector`

NewAccessPointSelector instantiates a new AccessPointSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPointSelectorWithDefaults

`func NewAccessPointSelectorWithDefaults() *AccessPointSelector`

NewAccessPointSelectorWithDefaults instantiates a new AccessPointSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessPointSelector) GetType() AccessPointSelectorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessPointSelector) GetTypeOk() (*AccessPointSelectorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessPointSelector) SetType(v AccessPointSelectorType)`

SetType sets Type field to given value.

### HasType

`func (o *AccessPointSelector) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPort

`func (o *AccessPointSelector) GetPort() SimplifiedMetadataEntity`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AccessPointSelector) GetPortOk() (*SimplifiedMetadataEntity, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AccessPointSelector) SetPort(v SimplifiedMetadataEntity)`

SetPort sets Port field to given value.

### HasPort

`func (o *AccessPointSelector) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLinkProtocol

`func (o *AccessPointSelector) GetLinkProtocol() SimplifiedLinkProtocol`

GetLinkProtocol returns the LinkProtocol field if non-nil, zero value otherwise.

### GetLinkProtocolOk

`func (o *AccessPointSelector) GetLinkProtocolOk() (*SimplifiedLinkProtocol, bool)`

GetLinkProtocolOk returns a tuple with the LinkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkProtocol

`func (o *AccessPointSelector) SetLinkProtocol(v SimplifiedLinkProtocol)`

SetLinkProtocol sets LinkProtocol field to given value.

### HasLinkProtocol

`func (o *AccessPointSelector) HasLinkProtocol() bool`

HasLinkProtocol returns a boolean if a field has been set.

### GetVirtualDevice

`func (o *AccessPointSelector) GetVirtualDevice() SimplifiedVirtualDevice`

GetVirtualDevice returns the VirtualDevice field if non-nil, zero value otherwise.

### GetVirtualDeviceOk

`func (o *AccessPointSelector) GetVirtualDeviceOk() (*SimplifiedVirtualDevice, bool)`

GetVirtualDeviceOk returns a tuple with the VirtualDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDevice

`func (o *AccessPointSelector) SetVirtualDevice(v SimplifiedVirtualDevice)`

SetVirtualDevice sets VirtualDevice field to given value.

### HasVirtualDevice

`func (o *AccessPointSelector) HasVirtualDevice() bool`

HasVirtualDevice returns a boolean if a field has been set.

### GetInterface

`func (o *AccessPointSelector) GetInterface() VirtualDeviceInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *AccessPointSelector) GetInterfaceOk() (*VirtualDeviceInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *AccessPointSelector) SetInterface(v VirtualDeviceInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *AccessPointSelector) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *AccessPointSelector) GetNetwork() SimplifiedTokenNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AccessPointSelector) GetNetworkOk() (*SimplifiedTokenNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AccessPointSelector) SetNetwork(v SimplifiedTokenNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AccessPointSelector) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



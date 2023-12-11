# MetadataNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** |  | [optional] 
**Interfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Network** | Pointer to [**MetadataNetworkNetwork**](MetadataNetworkNetwork.md) |  | [optional] 

## Methods

### NewMetadataNetwork

`func NewMetadataNetwork() *MetadataNetwork`

NewMetadataNetwork instantiates a new MetadataNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataNetworkWithDefaults

`func NewMetadataNetworkWithDefaults() *MetadataNetwork`

NewMetadataNetworkWithDefaults instantiates a new MetadataNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *MetadataNetwork) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *MetadataNetwork) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *MetadataNetwork) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *MetadataNetwork) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetInterfaces

`func (o *MetadataNetwork) GetInterfaces() []map[string]interface{}`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *MetadataNetwork) GetInterfacesOk() (*[]map[string]interface{}, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *MetadataNetwork) SetInterfaces(v []map[string]interface{})`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *MetadataNetwork) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetNetwork

`func (o *MetadataNetwork) GetNetwork() MetadataNetworkNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MetadataNetwork) GetNetworkOk() (*MetadataNetworkNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MetadataNetwork) SetNetwork(v MetadataNetworkNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MetadataNetwork) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



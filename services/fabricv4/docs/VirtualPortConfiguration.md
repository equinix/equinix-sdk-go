# VirtualPortConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyout** | Pointer to **bool** | Buyout (true) or standard (false) configuration of the port at this access point. &lt;br&gt; Buyout ports offer free, unlimited connections. Standard ports do not. The default is false. | [optional] [default to false]

## Methods

### NewVirtualPortConfiguration

`func NewVirtualPortConfiguration() *VirtualPortConfiguration`

NewVirtualPortConfiguration instantiates a new VirtualPortConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualPortConfigurationWithDefaults

`func NewVirtualPortConfigurationWithDefaults() *VirtualPortConfiguration`

NewVirtualPortConfigurationWithDefaults instantiates a new VirtualPortConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyout

`func (o *VirtualPortConfiguration) GetBuyout() bool`

GetBuyout returns the Buyout field if non-nil, zero value otherwise.

### GetBuyoutOk

`func (o *VirtualPortConfiguration) GetBuyoutOk() (*bool, bool)`

GetBuyoutOk returns a tuple with the Buyout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyout

`func (o *VirtualPortConfiguration) SetBuyout(v bool)`

SetBuyout sets Buyout field to given value.

### HasBuyout

`func (o *VirtualPortConfiguration) HasBuyout() bool`

HasBuyout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



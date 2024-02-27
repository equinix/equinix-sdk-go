# ServiceProfileLinkProtocolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncapsulationStrategy** | Pointer to [**ServiceProfileLinkProtocolConfigEncapsulationStrategy**](ServiceProfileLinkProtocolConfigEncapsulationStrategy.md) |  | [optional] 
**NamedTags** | Pointer to **[]string** |  | [optional] 
**VlanCTagLabel** | Pointer to **string** | was ctagLabel | [optional] 
**ReuseVlanSTag** | Pointer to **bool** |  | [optional] [default to false]
**Encapsulation** | Pointer to [**ServiceProfileLinkProtocolConfigEncapsulation**](ServiceProfileLinkProtocolConfigEncapsulation.md) |  | [optional] 

## Methods

### NewServiceProfileLinkProtocolConfig

`func NewServiceProfileLinkProtocolConfig() *ServiceProfileLinkProtocolConfig`

NewServiceProfileLinkProtocolConfig instantiates a new ServiceProfileLinkProtocolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileLinkProtocolConfigWithDefaults

`func NewServiceProfileLinkProtocolConfigWithDefaults() *ServiceProfileLinkProtocolConfig`

NewServiceProfileLinkProtocolConfigWithDefaults instantiates a new ServiceProfileLinkProtocolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncapsulationStrategy

`func (o *ServiceProfileLinkProtocolConfig) GetEncapsulationStrategy() ServiceProfileLinkProtocolConfigEncapsulationStrategy`

GetEncapsulationStrategy returns the EncapsulationStrategy field if non-nil, zero value otherwise.

### GetEncapsulationStrategyOk

`func (o *ServiceProfileLinkProtocolConfig) GetEncapsulationStrategyOk() (*ServiceProfileLinkProtocolConfigEncapsulationStrategy, bool)`

GetEncapsulationStrategyOk returns a tuple with the EncapsulationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulationStrategy

`func (o *ServiceProfileLinkProtocolConfig) SetEncapsulationStrategy(v ServiceProfileLinkProtocolConfigEncapsulationStrategy)`

SetEncapsulationStrategy sets EncapsulationStrategy field to given value.

### HasEncapsulationStrategy

`func (o *ServiceProfileLinkProtocolConfig) HasEncapsulationStrategy() bool`

HasEncapsulationStrategy returns a boolean if a field has been set.

### GetNamedTags

`func (o *ServiceProfileLinkProtocolConfig) GetNamedTags() []string`

GetNamedTags returns the NamedTags field if non-nil, zero value otherwise.

### GetNamedTagsOk

`func (o *ServiceProfileLinkProtocolConfig) GetNamedTagsOk() (*[]string, bool)`

GetNamedTagsOk returns a tuple with the NamedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedTags

`func (o *ServiceProfileLinkProtocolConfig) SetNamedTags(v []string)`

SetNamedTags sets NamedTags field to given value.

### HasNamedTags

`func (o *ServiceProfileLinkProtocolConfig) HasNamedTags() bool`

HasNamedTags returns a boolean if a field has been set.

### GetVlanCTagLabel

`func (o *ServiceProfileLinkProtocolConfig) GetVlanCTagLabel() string`

GetVlanCTagLabel returns the VlanCTagLabel field if non-nil, zero value otherwise.

### GetVlanCTagLabelOk

`func (o *ServiceProfileLinkProtocolConfig) GetVlanCTagLabelOk() (*string, bool)`

GetVlanCTagLabelOk returns a tuple with the VlanCTagLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagLabel

`func (o *ServiceProfileLinkProtocolConfig) SetVlanCTagLabel(v string)`

SetVlanCTagLabel sets VlanCTagLabel field to given value.

### HasVlanCTagLabel

`func (o *ServiceProfileLinkProtocolConfig) HasVlanCTagLabel() bool`

HasVlanCTagLabel returns a boolean if a field has been set.

### GetReuseVlanSTag

`func (o *ServiceProfileLinkProtocolConfig) GetReuseVlanSTag() bool`

GetReuseVlanSTag returns the ReuseVlanSTag field if non-nil, zero value otherwise.

### GetReuseVlanSTagOk

`func (o *ServiceProfileLinkProtocolConfig) GetReuseVlanSTagOk() (*bool, bool)`

GetReuseVlanSTagOk returns a tuple with the ReuseVlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseVlanSTag

`func (o *ServiceProfileLinkProtocolConfig) SetReuseVlanSTag(v bool)`

SetReuseVlanSTag sets ReuseVlanSTag field to given value.

### HasReuseVlanSTag

`func (o *ServiceProfileLinkProtocolConfig) HasReuseVlanSTag() bool`

HasReuseVlanSTag returns a boolean if a field has been set.

### GetEncapsulation

`func (o *ServiceProfileLinkProtocolConfig) GetEncapsulation() ServiceProfileLinkProtocolConfigEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *ServiceProfileLinkProtocolConfig) GetEncapsulationOk() (*ServiceProfileLinkProtocolConfigEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *ServiceProfileLinkProtocolConfig) SetEncapsulation(v ServiceProfileLinkProtocolConfigEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *ServiceProfileLinkProtocolConfig) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InboundRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol. | [optional] 
**SrcPort** | Pointer to **string** | Source port. | [optional] 
**DstPort** | Pointer to **string** | Destination port. | [optional] 
**Subnet** | Pointer to **string** | An array of subnets. | [optional] 
**SeqNo** | Pointer to **int32** | The sequence number of the inbound rule. | [optional] 
**Description** | Pointer to **string** | Description of the inboundRule. | [optional] 

## Methods

### NewInboundRules

`func NewInboundRules() *InboundRules`

NewInboundRules instantiates a new InboundRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundRulesWithDefaults

`func NewInboundRulesWithDefaults() *InboundRules`

NewInboundRulesWithDefaults instantiates a new InboundRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InboundRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InboundRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InboundRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InboundRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *InboundRules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InboundRules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InboundRules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InboundRules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstPort

`func (o *InboundRules) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *InboundRules) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *InboundRules) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *InboundRules) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetSubnet

`func (o *InboundRules) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InboundRules) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InboundRules) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InboundRules) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSeqNo

`func (o *InboundRules) GetSeqNo() int32`

GetSeqNo returns the SeqNo field if non-nil, zero value otherwise.

### GetSeqNoOk

`func (o *InboundRules) GetSeqNoOk() (*int32, bool)`

GetSeqNoOk returns a tuple with the SeqNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNo

`func (o *InboundRules) SetSeqNo(v int32)`

SetSeqNo sets SeqNo field to given value.

### HasSeqNo

`func (o *InboundRules) HasSeqNo() bool`

HasSeqNo returns a boolean if a field has been set.

### GetDescription

`func (o *InboundRules) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InboundRules) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InboundRules) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InboundRules) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



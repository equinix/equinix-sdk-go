# IpBlockIpv4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**AddressingPlans** | Pointer to [**[]IpBlockAddressingPlans**](IpBlockAddressingPlans.md) | Collection of addressing plans  | [optional] 
**Questions** | Pointer to [**[]IpBlockQuestions**](IpBlockQuestions.md) | Connection of questions  | [optional] 
**PrefixLength** | **int32** | Length of the IP block, number after the / (slash) | 

## Methods

### NewIpBlockIpv4Request

`func NewIpBlockIpv4Request(prefixLength int32, ) *IpBlockIpv4Request`

NewIpBlockIpv4Request instantiates a new IpBlockIpv4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockIpv4RequestWithDefaults

`func NewIpBlockIpv4RequestWithDefaults() *IpBlockIpv4Request`

NewIpBlockIpv4RequestWithDefaults instantiates a new IpBlockIpv4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *IpBlockIpv4Request) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IpBlockIpv4Request) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IpBlockIpv4Request) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IpBlockIpv4Request) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAddressingPlans

`func (o *IpBlockIpv4Request) GetAddressingPlans() []IpBlockAddressingPlans`

GetAddressingPlans returns the AddressingPlans field if non-nil, zero value otherwise.

### GetAddressingPlansOk

`func (o *IpBlockIpv4Request) GetAddressingPlansOk() (*[]IpBlockAddressingPlans, bool)`

GetAddressingPlansOk returns a tuple with the AddressingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingPlans

`func (o *IpBlockIpv4Request) SetAddressingPlans(v []IpBlockAddressingPlans)`

SetAddressingPlans sets AddressingPlans field to given value.

### HasAddressingPlans

`func (o *IpBlockIpv4Request) HasAddressingPlans() bool`

HasAddressingPlans returns a boolean if a field has been set.

### GetQuestions

`func (o *IpBlockIpv4Request) GetQuestions() []IpBlockQuestions`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *IpBlockIpv4Request) GetQuestionsOk() (*[]IpBlockQuestions, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *IpBlockIpv4Request) SetQuestions(v []IpBlockQuestions)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *IpBlockIpv4Request) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetPrefixLength

`func (o *IpBlockIpv4Request) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpBlockIpv4Request) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpBlockIpv4Request) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



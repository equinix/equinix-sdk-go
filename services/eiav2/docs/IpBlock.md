# IpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**AddressingPlans** | Pointer to [**[]IpBlockAddressingPlans**](IpBlockAddressingPlans.md) | Collection of addressing plans  | [optional] 
**Questions** | Pointer to [**[]IpBlockQuestions**](IpBlockQuestions.md) | Connection of questions  | [optional] 

## Methods

### NewIpBlock

`func NewIpBlock() *IpBlock`

NewIpBlock instantiates a new IpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockWithDefaults

`func NewIpBlockWithDefaults() *IpBlock`

NewIpBlockWithDefaults instantiates a new IpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *IpBlock) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IpBlock) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IpBlock) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IpBlock) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAddressingPlans

`func (o *IpBlock) GetAddressingPlans() []IpBlockAddressingPlans`

GetAddressingPlans returns the AddressingPlans field if non-nil, zero value otherwise.

### GetAddressingPlansOk

`func (o *IpBlock) GetAddressingPlansOk() (*[]IpBlockAddressingPlans, bool)`

GetAddressingPlansOk returns a tuple with the AddressingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingPlans

`func (o *IpBlock) SetAddressingPlans(v []IpBlockAddressingPlans)`

SetAddressingPlans sets AddressingPlans field to given value.

### HasAddressingPlans

`func (o *IpBlock) HasAddressingPlans() bool`

HasAddressingPlans returns a boolean if a field has been set.

### GetQuestions

`func (o *IpBlock) GetQuestions() []IpBlockQuestions`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *IpBlock) GetQuestionsOk() (*[]IpBlockQuestions, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *IpBlock) SetQuestions(v []IpBlockQuestions)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *IpBlock) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



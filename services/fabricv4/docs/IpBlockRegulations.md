# IpBlockRegulations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressingPlans** | [**[]IpBlockAddressingPlan**](IpBlockAddressingPlan.md) | List of addressing plans | 
**Questions** | [**IpBlockRegulationsQuestions**](IpBlockRegulationsQuestions.md) |  | 

## Methods

### NewIpBlockRegulations

`func NewIpBlockRegulations(addressingPlans []IpBlockAddressingPlan, questions IpBlockRegulationsQuestions, ) *IpBlockRegulations`

NewIpBlockRegulations instantiates a new IpBlockRegulations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockRegulationsWithDefaults

`func NewIpBlockRegulationsWithDefaults() *IpBlockRegulations`

NewIpBlockRegulationsWithDefaults instantiates a new IpBlockRegulations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressingPlans

`func (o *IpBlockRegulations) GetAddressingPlans() []IpBlockAddressingPlan`

GetAddressingPlans returns the AddressingPlans field if non-nil, zero value otherwise.

### GetAddressingPlansOk

`func (o *IpBlockRegulations) GetAddressingPlansOk() (*[]IpBlockAddressingPlan, bool)`

GetAddressingPlansOk returns a tuple with the AddressingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingPlans

`func (o *IpBlockRegulations) SetAddressingPlans(v []IpBlockAddressingPlan)`

SetAddressingPlans sets AddressingPlans field to given value.


### GetQuestions

`func (o *IpBlockRegulations) GetQuestions() IpBlockRegulationsQuestions`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *IpBlockRegulations) GetQuestionsOk() (*IpBlockRegulationsQuestions, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *IpBlockRegulations) SetQuestions(v IpBlockRegulationsQuestions)`

SetQuestions sets Questions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



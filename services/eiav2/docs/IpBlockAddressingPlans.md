# IpBlockAddressingPlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** |  | 
**Purpose** | **string** | The purpose of IP Subnet  | 
**Immediate** | **int32** | Number of ip addresses to be used immediatelly  | 
**AfterThreeMonths** | **int32** | Number of ip addresses to be used after 3 months  | 

## Methods

### NewIpBlockAddressingPlans

`func NewIpBlockAddressingPlans(size int32, purpose string, immediate int32, afterThreeMonths int32, ) *IpBlockAddressingPlans`

NewIpBlockAddressingPlans instantiates a new IpBlockAddressingPlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockAddressingPlansWithDefaults

`func NewIpBlockAddressingPlansWithDefaults() *IpBlockAddressingPlans`

NewIpBlockAddressingPlansWithDefaults instantiates a new IpBlockAddressingPlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *IpBlockAddressingPlans) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IpBlockAddressingPlans) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IpBlockAddressingPlans) SetSize(v int32)`

SetSize sets Size field to given value.


### GetPurpose

`func (o *IpBlockAddressingPlans) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *IpBlockAddressingPlans) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *IpBlockAddressingPlans) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetImmediate

`func (o *IpBlockAddressingPlans) GetImmediate() int32`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *IpBlockAddressingPlans) GetImmediateOk() (*int32, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *IpBlockAddressingPlans) SetImmediate(v int32)`

SetImmediate sets Immediate field to given value.


### GetAfterThreeMonths

`func (o *IpBlockAddressingPlans) GetAfterThreeMonths() int32`

GetAfterThreeMonths returns the AfterThreeMonths field if non-nil, zero value otherwise.

### GetAfterThreeMonthsOk

`func (o *IpBlockAddressingPlans) GetAfterThreeMonthsOk() (*int32, bool)`

GetAfterThreeMonthsOk returns a tuple with the AfterThreeMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterThreeMonths

`func (o *IpBlockAddressingPlans) SetAfterThreeMonths(v int32)`

SetAfterThreeMonths sets AfterThreeMonths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



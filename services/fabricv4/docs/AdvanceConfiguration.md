# AdvanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ntp** | Pointer to [**[]Md5**](Md5.md) |  | [optional] 
**Ptp** | Pointer to [**PtpAdvanceConfiguration**](PtpAdvanceConfiguration.md) |  | [optional] 

## Methods

### NewAdvanceConfiguration

`func NewAdvanceConfiguration() *AdvanceConfiguration`

NewAdvanceConfiguration instantiates a new AdvanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvanceConfigurationWithDefaults

`func NewAdvanceConfigurationWithDefaults() *AdvanceConfiguration`

NewAdvanceConfigurationWithDefaults instantiates a new AdvanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNtp

`func (o *AdvanceConfiguration) GetNtp() []Md5`

GetNtp returns the Ntp field if non-nil, zero value otherwise.

### GetNtpOk

`func (o *AdvanceConfiguration) GetNtpOk() (*[]Md5, bool)`

GetNtpOk returns a tuple with the Ntp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtp

`func (o *AdvanceConfiguration) SetNtp(v []Md5)`

SetNtp sets Ntp field to given value.

### HasNtp

`func (o *AdvanceConfiguration) HasNtp() bool`

HasNtp returns a boolean if a field has been set.

### GetPtp

`func (o *AdvanceConfiguration) GetPtp() PtpAdvanceConfiguration`

GetPtp returns the Ptp field if non-nil, zero value otherwise.

### GetPtpOk

`func (o *AdvanceConfiguration) GetPtpOk() (*PtpAdvanceConfiguration, bool)`

GetPtpOk returns a tuple with the Ptp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtp

`func (o *AdvanceConfiguration) SetPtp(v PtpAdvanceConfiguration)`

SetPtp sets Ptp field to given value.

### HasPtp

`func (o *AdvanceConfiguration) HasPtp() bool`

HasPtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



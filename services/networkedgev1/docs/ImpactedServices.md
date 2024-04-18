# ImpactedServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | Pointer to **string** | The name of the impacted service. | [optional] 
**Impact** | Pointer to **string** | The type of impact, whether the impacted service is down or delayed. | [optional] 
**ServiceStartTime** | Pointer to **string** | Start of the downtime of the service. | [optional] 
**ServiceEndTime** | Pointer to **string** | End of the downtime of the service. | [optional] 
**ErrorMessage** | Pointer to **string** | Downtime message of the service. | [optional] 

## Methods

### NewImpactedServices

`func NewImpactedServices() *ImpactedServices`

NewImpactedServices instantiates a new ImpactedServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpactedServicesWithDefaults

`func NewImpactedServicesWithDefaults() *ImpactedServices`

NewImpactedServicesWithDefaults instantiates a new ImpactedServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ImpactedServices) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ImpactedServices) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ImpactedServices) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ImpactedServices) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetImpact

`func (o *ImpactedServices) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ImpactedServices) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ImpactedServices) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ImpactedServices) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetServiceStartTime

`func (o *ImpactedServices) GetServiceStartTime() string`

GetServiceStartTime returns the ServiceStartTime field if non-nil, zero value otherwise.

### GetServiceStartTimeOk

`func (o *ImpactedServices) GetServiceStartTimeOk() (*string, bool)`

GetServiceStartTimeOk returns a tuple with the ServiceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStartTime

`func (o *ImpactedServices) SetServiceStartTime(v string)`

SetServiceStartTime sets ServiceStartTime field to given value.

### HasServiceStartTime

`func (o *ImpactedServices) HasServiceStartTime() bool`

HasServiceStartTime returns a boolean if a field has been set.

### GetServiceEndTime

`func (o *ImpactedServices) GetServiceEndTime() string`

GetServiceEndTime returns the ServiceEndTime field if non-nil, zero value otherwise.

### GetServiceEndTimeOk

`func (o *ImpactedServices) GetServiceEndTimeOk() (*string, bool)`

GetServiceEndTimeOk returns a tuple with the ServiceEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndTime

`func (o *ImpactedServices) SetServiceEndTime(v string)`

SetServiceEndTime sets ServiceEndTime field to given value.

### HasServiceEndTime

`func (o *ImpactedServices) HasServiceEndTime() bool`

HasServiceEndTime returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ImpactedServices) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ImpactedServices) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ImpactedServices) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ImpactedServices) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



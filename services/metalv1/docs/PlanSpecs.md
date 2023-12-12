# PlanSpecs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to [**[]PlanSpecsCpusInner**](PlanSpecsCpusInner.md) |  | [optional] 
**Memory** | Pointer to [**PlanSpecsMemory**](PlanSpecsMemory.md) |  | [optional] 
**Drives** | Pointer to [**[]PlanSpecsDrivesInner**](PlanSpecsDrivesInner.md) |  | [optional] 
**Nics** | Pointer to [**[]PlanSpecsNicsInner**](PlanSpecsNicsInner.md) |  | [optional] 
**Features** | Pointer to [**PlanSpecsFeatures**](PlanSpecsFeatures.md) |  | [optional] 

## Methods

### NewPlanSpecs

`func NewPlanSpecs() *PlanSpecs`

NewPlanSpecs instantiates a new PlanSpecs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanSpecsWithDefaults

`func NewPlanSpecsWithDefaults() *PlanSpecs`

NewPlanSpecsWithDefaults instantiates a new PlanSpecs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *PlanSpecs) GetCpus() []PlanSpecsCpusInner`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *PlanSpecs) GetCpusOk() (*[]PlanSpecsCpusInner, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *PlanSpecs) SetCpus(v []PlanSpecsCpusInner)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *PlanSpecs) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemory

`func (o *PlanSpecs) GetMemory() PlanSpecsMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PlanSpecs) GetMemoryOk() (*PlanSpecsMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PlanSpecs) SetMemory(v PlanSpecsMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PlanSpecs) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDrives

`func (o *PlanSpecs) GetDrives() []PlanSpecsDrivesInner`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *PlanSpecs) GetDrivesOk() (*[]PlanSpecsDrivesInner, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *PlanSpecs) SetDrives(v []PlanSpecsDrivesInner)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *PlanSpecs) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetNics

`func (o *PlanSpecs) GetNics() []PlanSpecsNicsInner`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *PlanSpecs) GetNicsOk() (*[]PlanSpecsNicsInner, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *PlanSpecs) SetNics(v []PlanSpecsNicsInner)`

SetNics sets Nics field to given value.

### HasNics

`func (o *PlanSpecs) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetFeatures

`func (o *PlanSpecs) GetFeatures() PlanSpecsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PlanSpecs) GetFeaturesOk() (*PlanSpecsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PlanSpecs) SetFeatures(v PlanSpecsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PlanSpecs) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



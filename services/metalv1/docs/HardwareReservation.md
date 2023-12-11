# HardwareReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomRate** | Pointer to **float32** | Amount that will be charged for every billing_cycle. | [optional] 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**Facility** | Pointer to [**Facility**](Facility.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NeedOfService** | Pointer to **bool** | Whether this Device requires assistance from Equinix Metal. | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Provisionable** | Pointer to **bool** | Whether the reserved server is provisionable or not. Spare devices can&#39;t be provisioned unless they are activated first. | [optional] 
**ShortId** | Pointer to **string** | Short version of the ID. | [optional] 
**Spare** | Pointer to **bool** | Whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Equinix Metal | [optional] 
**SwitchUuid** | Pointer to **string** | Switch short id. This can be used to determine if two devices are connected to the same switch, for example. | [optional] 
**TerminationTime** | Pointer to **time.Time** | Expiration date for the reservation. | [optional] 

## Methods

### NewHardwareReservation

`func NewHardwareReservation() *HardwareReservation`

NewHardwareReservation instantiates a new HardwareReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareReservationWithDefaults

`func NewHardwareReservationWithDefaults() *HardwareReservation`

NewHardwareReservationWithDefaults instantiates a new HardwareReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HardwareReservation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HardwareReservation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HardwareReservation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HardwareReservation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomRate

`func (o *HardwareReservation) GetCustomRate() float32`

GetCustomRate returns the CustomRate field if non-nil, zero value otherwise.

### GetCustomRateOk

`func (o *HardwareReservation) GetCustomRateOk() (*float32, bool)`

GetCustomRateOk returns a tuple with the CustomRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRate

`func (o *HardwareReservation) SetCustomRate(v float32)`

SetCustomRate sets CustomRate field to given value.

### HasCustomRate

`func (o *HardwareReservation) HasCustomRate() bool`

HasCustomRate returns a boolean if a field has been set.

### GetDevice

`func (o *HardwareReservation) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *HardwareReservation) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *HardwareReservation) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *HardwareReservation) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFacility

`func (o *HardwareReservation) GetFacility() Facility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *HardwareReservation) GetFacilityOk() (*Facility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *HardwareReservation) SetFacility(v Facility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *HardwareReservation) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHref

`func (o *HardwareReservation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HardwareReservation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HardwareReservation) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *HardwareReservation) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *HardwareReservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HardwareReservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HardwareReservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HardwareReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNeedOfService

`func (o *HardwareReservation) GetNeedOfService() bool`

GetNeedOfService returns the NeedOfService field if non-nil, zero value otherwise.

### GetNeedOfServiceOk

`func (o *HardwareReservation) GetNeedOfServiceOk() (*bool, bool)`

GetNeedOfServiceOk returns a tuple with the NeedOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedOfService

`func (o *HardwareReservation) SetNeedOfService(v bool)`

SetNeedOfService sets NeedOfService field to given value.

### HasNeedOfService

`func (o *HardwareReservation) HasNeedOfService() bool`

HasNeedOfService returns a boolean if a field has been set.

### GetPlan

`func (o *HardwareReservation) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *HardwareReservation) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *HardwareReservation) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *HardwareReservation) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProject

`func (o *HardwareReservation) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *HardwareReservation) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *HardwareReservation) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *HardwareReservation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvisionable

`func (o *HardwareReservation) GetProvisionable() bool`

GetProvisionable returns the Provisionable field if non-nil, zero value otherwise.

### GetProvisionableOk

`func (o *HardwareReservation) GetProvisionableOk() (*bool, bool)`

GetProvisionableOk returns a tuple with the Provisionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionable

`func (o *HardwareReservation) SetProvisionable(v bool)`

SetProvisionable sets Provisionable field to given value.

### HasProvisionable

`func (o *HardwareReservation) HasProvisionable() bool`

HasProvisionable returns a boolean if a field has been set.

### GetShortId

`func (o *HardwareReservation) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *HardwareReservation) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *HardwareReservation) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *HardwareReservation) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetSpare

`func (o *HardwareReservation) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *HardwareReservation) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *HardwareReservation) SetSpare(v bool)`

SetSpare sets Spare field to given value.

### HasSpare

`func (o *HardwareReservation) HasSpare() bool`

HasSpare returns a boolean if a field has been set.

### GetSwitchUuid

`func (o *HardwareReservation) GetSwitchUuid() string`

GetSwitchUuid returns the SwitchUuid field if non-nil, zero value otherwise.

### GetSwitchUuidOk

`func (o *HardwareReservation) GetSwitchUuidOk() (*string, bool)`

GetSwitchUuidOk returns a tuple with the SwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUuid

`func (o *HardwareReservation) SetSwitchUuid(v string)`

SetSwitchUuid sets SwitchUuid field to given value.

### HasSwitchUuid

`func (o *HardwareReservation) HasSwitchUuid() bool`

HasSwitchUuid returns a boolean if a field has been set.

### GetTerminationTime

`func (o *HardwareReservation) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *HardwareReservation) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *HardwareReservation) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *HardwareReservation) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



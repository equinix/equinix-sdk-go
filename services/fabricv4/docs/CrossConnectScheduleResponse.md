# CrossConnectScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Schedule URI. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Schedule Identifier. | [optional] [readonly] 
**Type** | Pointer to [**CrossConnectScheduleResponseType**](CrossConnectScheduleResponseType.md) |  | [optional] 
**Code** | Pointer to [**ScheduleType**](ScheduleType.md) |  | [optional] [readonly] 
**State** | Pointer to [**CrossConnectScheduleState**](CrossConnectScheduleState.md) |  | [optional] [readonly] 
**AutoConvertToStandard** | Pointer to **bool** | Indicates whether the Cross Connect is scheduled as STANDARD when no EXPEDITE scheduling is available. Returned only for EXPEDITE schedule.  | [optional] [readonly] 
**ReadinessDateTime** | Pointer to **time.Time** | Carrier Circuit delivery date is the date when the Carrier commits to enable the circuit. Your order will be placed on hold, pending the date you enter. Equinix will wait for the Carrier to deliver the service and then will complete the installation.  | [optional] [readonly] 
**RequestedDateTime** | Pointer to **time.Time** | Requested Installation Date and Time. | [optional] [readonly] 
**EquinixProposedDateTime** | Pointer to **time.Time** | Date and Time proposed by Equinix when the EXPEDITE schedule is not available.  | [optional] [readonly] 
**InstallationDateTime** | Pointer to **time.Time** | Installation Date and Time. | [optional] [readonly] 
**DeinstallationDateTime** | Pointer to **time.Time** | Deinstallation Date and Time. Returned for a removal schedule. | [optional] [readonly] 
**ChangeLog** | Pointer to [**CrossConnectChangelog**](CrossConnectChangelog.md) |  | [optional] 

## Methods

### NewCrossConnectScheduleResponse

`func NewCrossConnectScheduleResponse() *CrossConnectScheduleResponse`

NewCrossConnectScheduleResponse instantiates a new CrossConnectScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectScheduleResponseWithDefaults

`func NewCrossConnectScheduleResponseWithDefaults() *CrossConnectScheduleResponse`

NewCrossConnectScheduleResponseWithDefaults instantiates a new CrossConnectScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectScheduleResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectScheduleResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectScheduleResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectScheduleResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectScheduleResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectScheduleResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectScheduleResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectScheduleResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectScheduleResponse) GetType() CrossConnectScheduleResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectScheduleResponse) GetTypeOk() (*CrossConnectScheduleResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectScheduleResponse) SetType(v CrossConnectScheduleResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectScheduleResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *CrossConnectScheduleResponse) GetCode() ScheduleType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CrossConnectScheduleResponse) GetCodeOk() (*ScheduleType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CrossConnectScheduleResponse) SetCode(v ScheduleType)`

SetCode sets Code field to given value.

### HasCode

`func (o *CrossConnectScheduleResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetState

`func (o *CrossConnectScheduleResponse) GetState() CrossConnectScheduleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CrossConnectScheduleResponse) GetStateOk() (*CrossConnectScheduleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CrossConnectScheduleResponse) SetState(v CrossConnectScheduleState)`

SetState sets State field to given value.

### HasState

`func (o *CrossConnectScheduleResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAutoConvertToStandard

`func (o *CrossConnectScheduleResponse) GetAutoConvertToStandard() bool`

GetAutoConvertToStandard returns the AutoConvertToStandard field if non-nil, zero value otherwise.

### GetAutoConvertToStandardOk

`func (o *CrossConnectScheduleResponse) GetAutoConvertToStandardOk() (*bool, bool)`

GetAutoConvertToStandardOk returns a tuple with the AutoConvertToStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConvertToStandard

`func (o *CrossConnectScheduleResponse) SetAutoConvertToStandard(v bool)`

SetAutoConvertToStandard sets AutoConvertToStandard field to given value.

### HasAutoConvertToStandard

`func (o *CrossConnectScheduleResponse) HasAutoConvertToStandard() bool`

HasAutoConvertToStandard returns a boolean if a field has been set.

### GetReadinessDateTime

`func (o *CrossConnectScheduleResponse) GetReadinessDateTime() time.Time`

GetReadinessDateTime returns the ReadinessDateTime field if non-nil, zero value otherwise.

### GetReadinessDateTimeOk

`func (o *CrossConnectScheduleResponse) GetReadinessDateTimeOk() (*time.Time, bool)`

GetReadinessDateTimeOk returns a tuple with the ReadinessDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessDateTime

`func (o *CrossConnectScheduleResponse) SetReadinessDateTime(v time.Time)`

SetReadinessDateTime sets ReadinessDateTime field to given value.

### HasReadinessDateTime

`func (o *CrossConnectScheduleResponse) HasReadinessDateTime() bool`

HasReadinessDateTime returns a boolean if a field has been set.

### GetRequestedDateTime

`func (o *CrossConnectScheduleResponse) GetRequestedDateTime() time.Time`

GetRequestedDateTime returns the RequestedDateTime field if non-nil, zero value otherwise.

### GetRequestedDateTimeOk

`func (o *CrossConnectScheduleResponse) GetRequestedDateTimeOk() (*time.Time, bool)`

GetRequestedDateTimeOk returns a tuple with the RequestedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDateTime

`func (o *CrossConnectScheduleResponse) SetRequestedDateTime(v time.Time)`

SetRequestedDateTime sets RequestedDateTime field to given value.

### HasRequestedDateTime

`func (o *CrossConnectScheduleResponse) HasRequestedDateTime() bool`

HasRequestedDateTime returns a boolean if a field has been set.

### GetEquinixProposedDateTime

`func (o *CrossConnectScheduleResponse) GetEquinixProposedDateTime() time.Time`

GetEquinixProposedDateTime returns the EquinixProposedDateTime field if non-nil, zero value otherwise.

### GetEquinixProposedDateTimeOk

`func (o *CrossConnectScheduleResponse) GetEquinixProposedDateTimeOk() (*time.Time, bool)`

GetEquinixProposedDateTimeOk returns a tuple with the EquinixProposedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixProposedDateTime

`func (o *CrossConnectScheduleResponse) SetEquinixProposedDateTime(v time.Time)`

SetEquinixProposedDateTime sets EquinixProposedDateTime field to given value.

### HasEquinixProposedDateTime

`func (o *CrossConnectScheduleResponse) HasEquinixProposedDateTime() bool`

HasEquinixProposedDateTime returns a boolean if a field has been set.

### GetInstallationDateTime

`func (o *CrossConnectScheduleResponse) GetInstallationDateTime() time.Time`

GetInstallationDateTime returns the InstallationDateTime field if non-nil, zero value otherwise.

### GetInstallationDateTimeOk

`func (o *CrossConnectScheduleResponse) GetInstallationDateTimeOk() (*time.Time, bool)`

GetInstallationDateTimeOk returns a tuple with the InstallationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationDateTime

`func (o *CrossConnectScheduleResponse) SetInstallationDateTime(v time.Time)`

SetInstallationDateTime sets InstallationDateTime field to given value.

### HasInstallationDateTime

`func (o *CrossConnectScheduleResponse) HasInstallationDateTime() bool`

HasInstallationDateTime returns a boolean if a field has been set.

### GetDeinstallationDateTime

`func (o *CrossConnectScheduleResponse) GetDeinstallationDateTime() time.Time`

GetDeinstallationDateTime returns the DeinstallationDateTime field if non-nil, zero value otherwise.

### GetDeinstallationDateTimeOk

`func (o *CrossConnectScheduleResponse) GetDeinstallationDateTimeOk() (*time.Time, bool)`

GetDeinstallationDateTimeOk returns a tuple with the DeinstallationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinstallationDateTime

`func (o *CrossConnectScheduleResponse) SetDeinstallationDateTime(v time.Time)`

SetDeinstallationDateTime sets DeinstallationDateTime field to given value.

### HasDeinstallationDateTime

`func (o *CrossConnectScheduleResponse) HasDeinstallationDateTime() bool`

HasDeinstallationDateTime returns a boolean if a field has been set.

### GetChangeLog

`func (o *CrossConnectScheduleResponse) GetChangeLog() CrossConnectChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CrossConnectScheduleResponse) GetChangeLogOk() (*CrossConnectChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CrossConnectScheduleResponse) SetChangeLog(v CrossConnectChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CrossConnectScheduleResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



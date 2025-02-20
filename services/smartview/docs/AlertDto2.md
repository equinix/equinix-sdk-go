# AlertDto2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** |  | [optional] 
**Acknowledge** | Pointer to **bool** |  | [optional] 
**AffectedCustomerAsset** | Pointer to **string** |  | [optional] 
**AlertType** | Pointer to [**AlertType**](AlertType.md) |  | [optional] 
**AlertTypeName** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**Assetclassification** | Pointer to **string** |  | [optional] 
**Assetname** | Pointer to **string** |  | [optional] 
**Assettype** | Pointer to **string** |  | [optional] 
**ConditionalAlert** | Pointer to [**[]ConditionalAlert**](ConditionalAlert.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **float32** | createdOn | [optional] 
**Currentvalue** | Pointer to **string** |  | [optional] 
**Eventtype** | Pointer to **string** |  | [optional] 
**Ibx** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Lastmaintenance** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **float32** | modifiedOn | [optional] 
**NotificationType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Relatedincidents** | Pointer to **string** |  | [optional] 
**Resiliency** | Pointer to **string** |  | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Tagid** | Pointer to **string** |  | [optional] 
**ThresholdUnit** | Pointer to **string** |  | [optional] 
**ThresholdValue** | Pointer to **string** |  | [optional] 
**ThresholdValueMax** | Pointer to **string** |  | [optional] 
**ThresholdValueMin** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Timeacknowledged** | Pointer to **string** |  | [optional] 
**Timeprocessed** | Pointer to **string** |  | [optional] 
**TimetriggeredMilisec** | Pointer to **string** |  | [optional] 
**TriggeredOn** | Pointer to **float32** | lastTriggeredOn | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uom** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertDto2

`func NewAlertDto2() *AlertDto2`

NewAlertDto2 instantiates a new AlertDto2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDto2WithDefaults

`func NewAlertDto2WithDefaults() *AlertDto2`

NewAlertDto2WithDefaults instantiates a new AlertDto2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *AlertDto2) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *AlertDto2) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *AlertDto2) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *AlertDto2) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetAcknowledge

`func (o *AlertDto2) GetAcknowledge() bool`

GetAcknowledge returns the Acknowledge field if non-nil, zero value otherwise.

### GetAcknowledgeOk

`func (o *AlertDto2) GetAcknowledgeOk() (*bool, bool)`

GetAcknowledgeOk returns a tuple with the Acknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledge

`func (o *AlertDto2) SetAcknowledge(v bool)`

SetAcknowledge sets Acknowledge field to given value.

### HasAcknowledge

`func (o *AlertDto2) HasAcknowledge() bool`

HasAcknowledge returns a boolean if a field has been set.

### GetAffectedCustomerAsset

`func (o *AlertDto2) GetAffectedCustomerAsset() string`

GetAffectedCustomerAsset returns the AffectedCustomerAsset field if non-nil, zero value otherwise.

### GetAffectedCustomerAssetOk

`func (o *AlertDto2) GetAffectedCustomerAssetOk() (*string, bool)`

GetAffectedCustomerAssetOk returns a tuple with the AffectedCustomerAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedCustomerAsset

`func (o *AlertDto2) SetAffectedCustomerAsset(v string)`

SetAffectedCustomerAsset sets AffectedCustomerAsset field to given value.

### HasAffectedCustomerAsset

`func (o *AlertDto2) HasAffectedCustomerAsset() bool`

HasAffectedCustomerAsset returns a boolean if a field has been set.

### GetAlertType

`func (o *AlertDto2) GetAlertType() AlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertDto2) GetAlertTypeOk() (*AlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertDto2) SetAlertType(v AlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *AlertDto2) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetAlertTypeName

`func (o *AlertDto2) GetAlertTypeName() string`

GetAlertTypeName returns the AlertTypeName field if non-nil, zero value otherwise.

### GetAlertTypeNameOk

`func (o *AlertDto2) GetAlertTypeNameOk() (*string, bool)`

GetAlertTypeNameOk returns a tuple with the AlertTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeName

`func (o *AlertDto2) SetAlertTypeName(v string)`

SetAlertTypeName sets AlertTypeName field to given value.

### HasAlertTypeName

`func (o *AlertDto2) HasAlertTypeName() bool`

HasAlertTypeName returns a boolean if a field has been set.

### GetAsset

`func (o *AlertDto2) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AlertDto2) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AlertDto2) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AlertDto2) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetclassification

`func (o *AlertDto2) GetAssetclassification() string`

GetAssetclassification returns the Assetclassification field if non-nil, zero value otherwise.

### GetAssetclassificationOk

`func (o *AlertDto2) GetAssetclassificationOk() (*string, bool)`

GetAssetclassificationOk returns a tuple with the Assetclassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetclassification

`func (o *AlertDto2) SetAssetclassification(v string)`

SetAssetclassification sets Assetclassification field to given value.

### HasAssetclassification

`func (o *AlertDto2) HasAssetclassification() bool`

HasAssetclassification returns a boolean if a field has been set.

### GetAssetname

`func (o *AlertDto2) GetAssetname() string`

GetAssetname returns the Assetname field if non-nil, zero value otherwise.

### GetAssetnameOk

`func (o *AlertDto2) GetAssetnameOk() (*string, bool)`

GetAssetnameOk returns a tuple with the Assetname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetname

`func (o *AlertDto2) SetAssetname(v string)`

SetAssetname sets Assetname field to given value.

### HasAssetname

`func (o *AlertDto2) HasAssetname() bool`

HasAssetname returns a boolean if a field has been set.

### GetAssettype

`func (o *AlertDto2) GetAssettype() string`

GetAssettype returns the Assettype field if non-nil, zero value otherwise.

### GetAssettypeOk

`func (o *AlertDto2) GetAssettypeOk() (*string, bool)`

GetAssettypeOk returns a tuple with the Assettype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssettype

`func (o *AlertDto2) SetAssettype(v string)`

SetAssettype sets Assettype field to given value.

### HasAssettype

`func (o *AlertDto2) HasAssettype() bool`

HasAssettype returns a boolean if a field has been set.

### GetConditionalAlert

`func (o *AlertDto2) GetConditionalAlert() []ConditionalAlert`

GetConditionalAlert returns the ConditionalAlert field if non-nil, zero value otherwise.

### GetConditionalAlertOk

`func (o *AlertDto2) GetConditionalAlertOk() (*[]ConditionalAlert, bool)`

GetConditionalAlertOk returns a tuple with the ConditionalAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAlert

`func (o *AlertDto2) SetConditionalAlert(v []ConditionalAlert)`

SetConditionalAlert sets ConditionalAlert field to given value.

### HasConditionalAlert

`func (o *AlertDto2) HasConditionalAlert() bool`

HasConditionalAlert returns a boolean if a field has been set.

### GetCountry

`func (o *AlertDto2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AlertDto2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AlertDto2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AlertDto2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedOn

`func (o *AlertDto2) GetCreatedOn() float32`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *AlertDto2) GetCreatedOnOk() (*float32, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *AlertDto2) SetCreatedOn(v float32)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *AlertDto2) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCurrentvalue

`func (o *AlertDto2) GetCurrentvalue() string`

GetCurrentvalue returns the Currentvalue field if non-nil, zero value otherwise.

### GetCurrentvalueOk

`func (o *AlertDto2) GetCurrentvalueOk() (*string, bool)`

GetCurrentvalueOk returns a tuple with the Currentvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentvalue

`func (o *AlertDto2) SetCurrentvalue(v string)`

SetCurrentvalue sets Currentvalue field to given value.

### HasCurrentvalue

`func (o *AlertDto2) HasCurrentvalue() bool`

HasCurrentvalue returns a boolean if a field has been set.

### GetEventtype

`func (o *AlertDto2) GetEventtype() string`

GetEventtype returns the Eventtype field if non-nil, zero value otherwise.

### GetEventtypeOk

`func (o *AlertDto2) GetEventtypeOk() (*string, bool)`

GetEventtypeOk returns a tuple with the Eventtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventtype

`func (o *AlertDto2) SetEventtype(v string)`

SetEventtype sets Eventtype field to given value.

### HasEventtype

`func (o *AlertDto2) HasEventtype() bool`

HasEventtype returns a boolean if a field has been set.

### GetIbx

`func (o *AlertDto2) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *AlertDto2) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *AlertDto2) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *AlertDto2) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetId

`func (o *AlertDto2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertDto2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertDto2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertDto2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastmaintenance

`func (o *AlertDto2) GetLastmaintenance() string`

GetLastmaintenance returns the Lastmaintenance field if non-nil, zero value otherwise.

### GetLastmaintenanceOk

`func (o *AlertDto2) GetLastmaintenanceOk() (*string, bool)`

GetLastmaintenanceOk returns a tuple with the Lastmaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastmaintenance

`func (o *AlertDto2) SetLastmaintenance(v string)`

SetLastmaintenance sets Lastmaintenance field to given value.

### HasLastmaintenance

`func (o *AlertDto2) HasLastmaintenance() bool`

HasLastmaintenance returns a boolean if a field has been set.

### GetMetro

`func (o *AlertDto2) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *AlertDto2) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *AlertDto2) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *AlertDto2) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetModifiedOn

`func (o *AlertDto2) GetModifiedOn() float32`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *AlertDto2) GetModifiedOnOk() (*float32, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *AlertDto2) SetModifiedOn(v float32)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *AlertDto2) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetNotificationType

`func (o *AlertDto2) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *AlertDto2) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *AlertDto2) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *AlertDto2) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetRegion

`func (o *AlertDto2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AlertDto2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AlertDto2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AlertDto2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRelatedincidents

`func (o *AlertDto2) GetRelatedincidents() string`

GetRelatedincidents returns the Relatedincidents field if non-nil, zero value otherwise.

### GetRelatedincidentsOk

`func (o *AlertDto2) GetRelatedincidentsOk() (*string, bool)`

GetRelatedincidentsOk returns a tuple with the Relatedincidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedincidents

`func (o *AlertDto2) SetRelatedincidents(v string)`

SetRelatedincidents sets Relatedincidents field to given value.

### HasRelatedincidents

`func (o *AlertDto2) HasRelatedincidents() bool`

HasRelatedincidents returns a boolean if a field has been set.

### GetResiliency

`func (o *AlertDto2) GetResiliency() string`

GetResiliency returns the Resiliency field if non-nil, zero value otherwise.

### GetResiliencyOk

`func (o *AlertDto2) GetResiliencyOk() (*string, bool)`

GetResiliencyOk returns a tuple with the Resiliency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliency

`func (o *AlertDto2) SetResiliency(v string)`

SetResiliency sets Resiliency field to given value.

### HasResiliency

`func (o *AlertDto2) HasResiliency() bool`

HasResiliency returns a boolean if a field has been set.

### GetSection

`func (o *AlertDto2) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *AlertDto2) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *AlertDto2) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *AlertDto2) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertDto2) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertDto2) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertDto2) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertDto2) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagid

`func (o *AlertDto2) GetTagid() string`

GetTagid returns the Tagid field if non-nil, zero value otherwise.

### GetTagidOk

`func (o *AlertDto2) GetTagidOk() (*string, bool)`

GetTagidOk returns a tuple with the Tagid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagid

`func (o *AlertDto2) SetTagid(v string)`

SetTagid sets Tagid field to given value.

### HasTagid

`func (o *AlertDto2) HasTagid() bool`

HasTagid returns a boolean if a field has been set.

### GetThresholdUnit

`func (o *AlertDto2) GetThresholdUnit() string`

GetThresholdUnit returns the ThresholdUnit field if non-nil, zero value otherwise.

### GetThresholdUnitOk

`func (o *AlertDto2) GetThresholdUnitOk() (*string, bool)`

GetThresholdUnitOk returns a tuple with the ThresholdUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUnit

`func (o *AlertDto2) SetThresholdUnit(v string)`

SetThresholdUnit sets ThresholdUnit field to given value.

### HasThresholdUnit

`func (o *AlertDto2) HasThresholdUnit() bool`

HasThresholdUnit returns a boolean if a field has been set.

### GetThresholdValue

`func (o *AlertDto2) GetThresholdValue() string`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *AlertDto2) GetThresholdValueOk() (*string, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *AlertDto2) SetThresholdValue(v string)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *AlertDto2) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetThresholdValueMax

`func (o *AlertDto2) GetThresholdValueMax() string`

GetThresholdValueMax returns the ThresholdValueMax field if non-nil, zero value otherwise.

### GetThresholdValueMaxOk

`func (o *AlertDto2) GetThresholdValueMaxOk() (*string, bool)`

GetThresholdValueMaxOk returns a tuple with the ThresholdValueMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValueMax

`func (o *AlertDto2) SetThresholdValueMax(v string)`

SetThresholdValueMax sets ThresholdValueMax field to given value.

### HasThresholdValueMax

`func (o *AlertDto2) HasThresholdValueMax() bool`

HasThresholdValueMax returns a boolean if a field has been set.

### GetThresholdValueMin

`func (o *AlertDto2) GetThresholdValueMin() string`

GetThresholdValueMin returns the ThresholdValueMin field if non-nil, zero value otherwise.

### GetThresholdValueMinOk

`func (o *AlertDto2) GetThresholdValueMinOk() (*string, bool)`

GetThresholdValueMinOk returns a tuple with the ThresholdValueMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValueMin

`func (o *AlertDto2) SetThresholdValueMin(v string)`

SetThresholdValueMin sets ThresholdValueMin field to given value.

### HasThresholdValueMin

`func (o *AlertDto2) HasThresholdValueMin() bool`

HasThresholdValueMin returns a boolean if a field has been set.

### GetTimeZone

`func (o *AlertDto2) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AlertDto2) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AlertDto2) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AlertDto2) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTimeacknowledged

`func (o *AlertDto2) GetTimeacknowledged() string`

GetTimeacknowledged returns the Timeacknowledged field if non-nil, zero value otherwise.

### GetTimeacknowledgedOk

`func (o *AlertDto2) GetTimeacknowledgedOk() (*string, bool)`

GetTimeacknowledgedOk returns a tuple with the Timeacknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeacknowledged

`func (o *AlertDto2) SetTimeacknowledged(v string)`

SetTimeacknowledged sets Timeacknowledged field to given value.

### HasTimeacknowledged

`func (o *AlertDto2) HasTimeacknowledged() bool`

HasTimeacknowledged returns a boolean if a field has been set.

### GetTimeprocessed

`func (o *AlertDto2) GetTimeprocessed() string`

GetTimeprocessed returns the Timeprocessed field if non-nil, zero value otherwise.

### GetTimeprocessedOk

`func (o *AlertDto2) GetTimeprocessedOk() (*string, bool)`

GetTimeprocessedOk returns a tuple with the Timeprocessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeprocessed

`func (o *AlertDto2) SetTimeprocessed(v string)`

SetTimeprocessed sets Timeprocessed field to given value.

### HasTimeprocessed

`func (o *AlertDto2) HasTimeprocessed() bool`

HasTimeprocessed returns a boolean if a field has been set.

### GetTimetriggeredMilisec

`func (o *AlertDto2) GetTimetriggeredMilisec() string`

GetTimetriggeredMilisec returns the TimetriggeredMilisec field if non-nil, zero value otherwise.

### GetTimetriggeredMilisecOk

`func (o *AlertDto2) GetTimetriggeredMilisecOk() (*string, bool)`

GetTimetriggeredMilisecOk returns a tuple with the TimetriggeredMilisec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetriggeredMilisec

`func (o *AlertDto2) SetTimetriggeredMilisec(v string)`

SetTimetriggeredMilisec sets TimetriggeredMilisec field to given value.

### HasTimetriggeredMilisec

`func (o *AlertDto2) HasTimetriggeredMilisec() bool`

HasTimetriggeredMilisec returns a boolean if a field has been set.

### GetTriggeredOn

`func (o *AlertDto2) GetTriggeredOn() float32`

GetTriggeredOn returns the TriggeredOn field if non-nil, zero value otherwise.

### GetTriggeredOnOk

`func (o *AlertDto2) GetTriggeredOnOk() (*float32, bool)`

GetTriggeredOnOk returns a tuple with the TriggeredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredOn

`func (o *AlertDto2) SetTriggeredOn(v float32)`

SetTriggeredOn sets TriggeredOn field to given value.

### HasTriggeredOn

`func (o *AlertDto2) HasTriggeredOn() bool`

HasTriggeredOn returns a boolean if a field has been set.

### GetType

`func (o *AlertDto2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertDto2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertDto2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertDto2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUom

`func (o *AlertDto2) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *AlertDto2) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *AlertDto2) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *AlertDto2) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetYear

`func (o *AlertDto2) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AlertDto2) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AlertDto2) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *AlertDto2) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



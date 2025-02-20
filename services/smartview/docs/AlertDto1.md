# AlertDto1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** |  | [optional] 
**AlertPaused** | Pointer to **bool** | isAlertPaused | [optional] 
**AlertType** | Pointer to [**AlertType**](AlertType.md) |  | [optional] 
**ConditionalAlert** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **float32** | createdOn | [optional] 
**CustomerAssets** | Pointer to [**[]CustomerAssets**](CustomerAssets.md) |  | [optional] 
**EnabledAction** | Pointer to **string** |  | [optional] 
**FormatedSection** | Pointer to **string** |  | [optional] 
**HeartbeatType** | Pointer to **string** |  | [optional] 
**Ibx** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InfraAssets** | Pointer to [**[]CustomerAssets**](CustomerAssets.md) |  | [optional] 
**IsDuplicate** | Pointer to **bool** |  | [optional] 
**LastTriggeredOn** | Pointer to **float32** | lastTriggeredOn | [optional] 
**Metro** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **float32** | modifiedOn | [optional] 
**Recipients** | Pointer to [**[]RecipientsArray**](RecipientsArray.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ThresholdUnit** | Pointer to **string** |  | [optional] 
**ThresholdValue** | Pointer to **string** |  | [optional] 
**ThresholdValueMax** | Pointer to **string** |  | [optional] 
**ThresholdValueMin** | Pointer to **string** |  | [optional] 
**Ucmid** | Pointer to **string** |  | [optional] 
**Uom** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertDto1

`func NewAlertDto1() *AlertDto1`

NewAlertDto1 instantiates a new AlertDto1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDto1WithDefaults

`func NewAlertDto1WithDefaults() *AlertDto1`

NewAlertDto1WithDefaults instantiates a new AlertDto1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *AlertDto1) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *AlertDto1) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *AlertDto1) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *AlertDto1) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetAlertPaused

`func (o *AlertDto1) GetAlertPaused() bool`

GetAlertPaused returns the AlertPaused field if non-nil, zero value otherwise.

### GetAlertPausedOk

`func (o *AlertDto1) GetAlertPausedOk() (*bool, bool)`

GetAlertPausedOk returns a tuple with the AlertPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPaused

`func (o *AlertDto1) SetAlertPaused(v bool)`

SetAlertPaused sets AlertPaused field to given value.

### HasAlertPaused

`func (o *AlertDto1) HasAlertPaused() bool`

HasAlertPaused returns a boolean if a field has been set.

### GetAlertType

`func (o *AlertDto1) GetAlertType() AlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertDto1) GetAlertTypeOk() (*AlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertDto1) SetAlertType(v AlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *AlertDto1) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetConditionalAlert

`func (o *AlertDto1) GetConditionalAlert() string`

GetConditionalAlert returns the ConditionalAlert field if non-nil, zero value otherwise.

### GetConditionalAlertOk

`func (o *AlertDto1) GetConditionalAlertOk() (*string, bool)`

GetConditionalAlertOk returns a tuple with the ConditionalAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAlert

`func (o *AlertDto1) SetConditionalAlert(v string)`

SetConditionalAlert sets ConditionalAlert field to given value.

### HasConditionalAlert

`func (o *AlertDto1) HasConditionalAlert() bool`

HasConditionalAlert returns a boolean if a field has been set.

### GetCountry

`func (o *AlertDto1) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AlertDto1) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AlertDto1) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AlertDto1) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AlertDto1) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertDto1) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertDto1) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AlertDto1) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *AlertDto1) GetCreatedOn() float32`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *AlertDto1) GetCreatedOnOk() (*float32, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *AlertDto1) SetCreatedOn(v float32)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *AlertDto1) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCustomerAssets

`func (o *AlertDto1) GetCustomerAssets() []CustomerAssets`

GetCustomerAssets returns the CustomerAssets field if non-nil, zero value otherwise.

### GetCustomerAssetsOk

`func (o *AlertDto1) GetCustomerAssetsOk() (*[]CustomerAssets, bool)`

GetCustomerAssetsOk returns a tuple with the CustomerAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAssets

`func (o *AlertDto1) SetCustomerAssets(v []CustomerAssets)`

SetCustomerAssets sets CustomerAssets field to given value.

### HasCustomerAssets

`func (o *AlertDto1) HasCustomerAssets() bool`

HasCustomerAssets returns a boolean if a field has been set.

### GetEnabledAction

`func (o *AlertDto1) GetEnabledAction() string`

GetEnabledAction returns the EnabledAction field if non-nil, zero value otherwise.

### GetEnabledActionOk

`func (o *AlertDto1) GetEnabledActionOk() (*string, bool)`

GetEnabledActionOk returns a tuple with the EnabledAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAction

`func (o *AlertDto1) SetEnabledAction(v string)`

SetEnabledAction sets EnabledAction field to given value.

### HasEnabledAction

`func (o *AlertDto1) HasEnabledAction() bool`

HasEnabledAction returns a boolean if a field has been set.

### GetFormatedSection

`func (o *AlertDto1) GetFormatedSection() string`

GetFormatedSection returns the FormatedSection field if non-nil, zero value otherwise.

### GetFormatedSectionOk

`func (o *AlertDto1) GetFormatedSectionOk() (*string, bool)`

GetFormatedSectionOk returns a tuple with the FormatedSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatedSection

`func (o *AlertDto1) SetFormatedSection(v string)`

SetFormatedSection sets FormatedSection field to given value.

### HasFormatedSection

`func (o *AlertDto1) HasFormatedSection() bool`

HasFormatedSection returns a boolean if a field has been set.

### GetHeartbeatType

`func (o *AlertDto1) GetHeartbeatType() string`

GetHeartbeatType returns the HeartbeatType field if non-nil, zero value otherwise.

### GetHeartbeatTypeOk

`func (o *AlertDto1) GetHeartbeatTypeOk() (*string, bool)`

GetHeartbeatTypeOk returns a tuple with the HeartbeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatType

`func (o *AlertDto1) SetHeartbeatType(v string)`

SetHeartbeatType sets HeartbeatType field to given value.

### HasHeartbeatType

`func (o *AlertDto1) HasHeartbeatType() bool`

HasHeartbeatType returns a boolean if a field has been set.

### GetIbx

`func (o *AlertDto1) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *AlertDto1) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *AlertDto1) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *AlertDto1) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetId

`func (o *AlertDto1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertDto1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertDto1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertDto1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfraAssets

`func (o *AlertDto1) GetInfraAssets() []CustomerAssets`

GetInfraAssets returns the InfraAssets field if non-nil, zero value otherwise.

### GetInfraAssetsOk

`func (o *AlertDto1) GetInfraAssetsOk() (*[]CustomerAssets, bool)`

GetInfraAssetsOk returns a tuple with the InfraAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraAssets

`func (o *AlertDto1) SetInfraAssets(v []CustomerAssets)`

SetInfraAssets sets InfraAssets field to given value.

### HasInfraAssets

`func (o *AlertDto1) HasInfraAssets() bool`

HasInfraAssets returns a boolean if a field has been set.

### GetIsDuplicate

`func (o *AlertDto1) GetIsDuplicate() bool`

GetIsDuplicate returns the IsDuplicate field if non-nil, zero value otherwise.

### GetIsDuplicateOk

`func (o *AlertDto1) GetIsDuplicateOk() (*bool, bool)`

GetIsDuplicateOk returns a tuple with the IsDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDuplicate

`func (o *AlertDto1) SetIsDuplicate(v bool)`

SetIsDuplicate sets IsDuplicate field to given value.

### HasIsDuplicate

`func (o *AlertDto1) HasIsDuplicate() bool`

HasIsDuplicate returns a boolean if a field has been set.

### GetLastTriggeredOn

`func (o *AlertDto1) GetLastTriggeredOn() float32`

GetLastTriggeredOn returns the LastTriggeredOn field if non-nil, zero value otherwise.

### GetLastTriggeredOnOk

`func (o *AlertDto1) GetLastTriggeredOnOk() (*float32, bool)`

GetLastTriggeredOnOk returns a tuple with the LastTriggeredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredOn

`func (o *AlertDto1) SetLastTriggeredOn(v float32)`

SetLastTriggeredOn sets LastTriggeredOn field to given value.

### HasLastTriggeredOn

`func (o *AlertDto1) HasLastTriggeredOn() bool`

HasLastTriggeredOn returns a boolean if a field has been set.

### GetMetro

`func (o *AlertDto1) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *AlertDto1) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *AlertDto1) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *AlertDto1) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetModifiedOn

`func (o *AlertDto1) GetModifiedOn() float32`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *AlertDto1) GetModifiedOnOk() (*float32, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *AlertDto1) SetModifiedOn(v float32)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *AlertDto1) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertDto1) GetRecipients() []RecipientsArray`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertDto1) GetRecipientsOk() (*[]RecipientsArray, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertDto1) SetRecipients(v []RecipientsArray)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertDto1) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRegion

`func (o *AlertDto1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AlertDto1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AlertDto1) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AlertDto1) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSection

`func (o *AlertDto1) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *AlertDto1) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *AlertDto1) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *AlertDto1) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetStatus

`func (o *AlertDto1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertDto1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertDto1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertDto1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThresholdUnit

`func (o *AlertDto1) GetThresholdUnit() string`

GetThresholdUnit returns the ThresholdUnit field if non-nil, zero value otherwise.

### GetThresholdUnitOk

`func (o *AlertDto1) GetThresholdUnitOk() (*string, bool)`

GetThresholdUnitOk returns a tuple with the ThresholdUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUnit

`func (o *AlertDto1) SetThresholdUnit(v string)`

SetThresholdUnit sets ThresholdUnit field to given value.

### HasThresholdUnit

`func (o *AlertDto1) HasThresholdUnit() bool`

HasThresholdUnit returns a boolean if a field has been set.

### GetThresholdValue

`func (o *AlertDto1) GetThresholdValue() string`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *AlertDto1) GetThresholdValueOk() (*string, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *AlertDto1) SetThresholdValue(v string)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *AlertDto1) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetThresholdValueMax

`func (o *AlertDto1) GetThresholdValueMax() string`

GetThresholdValueMax returns the ThresholdValueMax field if non-nil, zero value otherwise.

### GetThresholdValueMaxOk

`func (o *AlertDto1) GetThresholdValueMaxOk() (*string, bool)`

GetThresholdValueMaxOk returns a tuple with the ThresholdValueMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValueMax

`func (o *AlertDto1) SetThresholdValueMax(v string)`

SetThresholdValueMax sets ThresholdValueMax field to given value.

### HasThresholdValueMax

`func (o *AlertDto1) HasThresholdValueMax() bool`

HasThresholdValueMax returns a boolean if a field has been set.

### GetThresholdValueMin

`func (o *AlertDto1) GetThresholdValueMin() string`

GetThresholdValueMin returns the ThresholdValueMin field if non-nil, zero value otherwise.

### GetThresholdValueMinOk

`func (o *AlertDto1) GetThresholdValueMinOk() (*string, bool)`

GetThresholdValueMinOk returns a tuple with the ThresholdValueMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValueMin

`func (o *AlertDto1) SetThresholdValueMin(v string)`

SetThresholdValueMin sets ThresholdValueMin field to given value.

### HasThresholdValueMin

`func (o *AlertDto1) HasThresholdValueMin() bool`

HasThresholdValueMin returns a boolean if a field has been set.

### GetUcmid

`func (o *AlertDto1) GetUcmid() string`

GetUcmid returns the Ucmid field if non-nil, zero value otherwise.

### GetUcmidOk

`func (o *AlertDto1) GetUcmidOk() (*string, bool)`

GetUcmidOk returns a tuple with the Ucmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcmid

`func (o *AlertDto1) SetUcmid(v string)`

SetUcmid sets Ucmid field to given value.

### HasUcmid

`func (o *AlertDto1) HasUcmid() bool`

HasUcmid returns a boolean if a field has been set.

### GetUom

`func (o *AlertDto1) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *AlertDto1) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *AlertDto1) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *AlertDto1) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetUserId

`func (o *AlertDto1) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AlertDto1) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AlertDto1) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AlertDto1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



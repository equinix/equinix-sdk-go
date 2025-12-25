# GetServProfServicesRespContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AuthKeyLabel** | Pointer to **string** |  | [optional] 
**ConnectionNameLabel** | Pointer to **string** |  | [optional] 
**RequiredRedundancy** | Pointer to **bool** |  | [optional] 
**AllowCustomSpeed** | Pointer to **bool** |  | [optional] 
**SpeedBands** | Pointer to [**[]SpeedBand**](SpeedBand.md) |  | [optional] 
**Metros** | Pointer to [**GetServProfServicesRespContentMetros**](GetServProfServicesRespContentMetros.md) |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**VlanSameAsPrimary** | Pointer to **bool** |  | [optional] 
**TagType** | Pointer to **string** |  | [optional] 
**CtagLabel** | Pointer to **string** |  | [optional] 
**ApiAvailable** | Pointer to **bool** |  | [optional] 
**SelfProfile** | Pointer to **bool** |  | [optional] 
**ProfileEncapsulation** | Pointer to **string** |  | [optional] 
**AuthorizationKey** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**GetServProfServicesRespContentfeatures**](GetServProfServicesRespContentfeatures.md) |  | [optional] 

## Methods

### NewGetServProfServicesRespContent

`func NewGetServProfServicesRespContent() *GetServProfServicesRespContent`

NewGetServProfServicesRespContent instantiates a new GetServProfServicesRespContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServProfServicesRespContentWithDefaults

`func NewGetServProfServicesRespContentWithDefaults() *GetServProfServicesRespContent`

NewGetServProfServicesRespContentWithDefaults instantiates a new GetServProfServicesRespContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *GetServProfServicesRespContent) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetServProfServicesRespContent) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetServProfServicesRespContent) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetServProfServicesRespContent) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *GetServProfServicesRespContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetServProfServicesRespContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetServProfServicesRespContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetServProfServicesRespContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthKeyLabel

`func (o *GetServProfServicesRespContent) GetAuthKeyLabel() string`

GetAuthKeyLabel returns the AuthKeyLabel field if non-nil, zero value otherwise.

### GetAuthKeyLabelOk

`func (o *GetServProfServicesRespContent) GetAuthKeyLabelOk() (*string, bool)`

GetAuthKeyLabelOk returns a tuple with the AuthKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeyLabel

`func (o *GetServProfServicesRespContent) SetAuthKeyLabel(v string)`

SetAuthKeyLabel sets AuthKeyLabel field to given value.

### HasAuthKeyLabel

`func (o *GetServProfServicesRespContent) HasAuthKeyLabel() bool`

HasAuthKeyLabel returns a boolean if a field has been set.

### GetConnectionNameLabel

`func (o *GetServProfServicesRespContent) GetConnectionNameLabel() string`

GetConnectionNameLabel returns the ConnectionNameLabel field if non-nil, zero value otherwise.

### GetConnectionNameLabelOk

`func (o *GetServProfServicesRespContent) GetConnectionNameLabelOk() (*string, bool)`

GetConnectionNameLabelOk returns a tuple with the ConnectionNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionNameLabel

`func (o *GetServProfServicesRespContent) SetConnectionNameLabel(v string)`

SetConnectionNameLabel sets ConnectionNameLabel field to given value.

### HasConnectionNameLabel

`func (o *GetServProfServicesRespContent) HasConnectionNameLabel() bool`

HasConnectionNameLabel returns a boolean if a field has been set.

### GetRequiredRedundancy

`func (o *GetServProfServicesRespContent) GetRequiredRedundancy() bool`

GetRequiredRedundancy returns the RequiredRedundancy field if non-nil, zero value otherwise.

### GetRequiredRedundancyOk

`func (o *GetServProfServicesRespContent) GetRequiredRedundancyOk() (*bool, bool)`

GetRequiredRedundancyOk returns a tuple with the RequiredRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRedundancy

`func (o *GetServProfServicesRespContent) SetRequiredRedundancy(v bool)`

SetRequiredRedundancy sets RequiredRedundancy field to given value.

### HasRequiredRedundancy

`func (o *GetServProfServicesRespContent) HasRequiredRedundancy() bool`

HasRequiredRedundancy returns a boolean if a field has been set.

### GetAllowCustomSpeed

`func (o *GetServProfServicesRespContent) GetAllowCustomSpeed() bool`

GetAllowCustomSpeed returns the AllowCustomSpeed field if non-nil, zero value otherwise.

### GetAllowCustomSpeedOk

`func (o *GetServProfServicesRespContent) GetAllowCustomSpeedOk() (*bool, bool)`

GetAllowCustomSpeedOk returns a tuple with the AllowCustomSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomSpeed

`func (o *GetServProfServicesRespContent) SetAllowCustomSpeed(v bool)`

SetAllowCustomSpeed sets AllowCustomSpeed field to given value.

### HasAllowCustomSpeed

`func (o *GetServProfServicesRespContent) HasAllowCustomSpeed() bool`

HasAllowCustomSpeed returns a boolean if a field has been set.

### GetSpeedBands

`func (o *GetServProfServicesRespContent) GetSpeedBands() []SpeedBand`

GetSpeedBands returns the SpeedBands field if non-nil, zero value otherwise.

### GetSpeedBandsOk

`func (o *GetServProfServicesRespContent) GetSpeedBandsOk() (*[]SpeedBand, bool)`

GetSpeedBandsOk returns a tuple with the SpeedBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedBands

`func (o *GetServProfServicesRespContent) SetSpeedBands(v []SpeedBand)`

SetSpeedBands sets SpeedBands field to given value.

### HasSpeedBands

`func (o *GetServProfServicesRespContent) HasSpeedBands() bool`

HasSpeedBands returns a boolean if a field has been set.

### GetMetros

`func (o *GetServProfServicesRespContent) GetMetros() GetServProfServicesRespContentMetros`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *GetServProfServicesRespContent) GetMetrosOk() (*GetServProfServicesRespContentMetros, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *GetServProfServicesRespContent) SetMetros(v GetServProfServicesRespContentMetros)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *GetServProfServicesRespContent) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetCreatedDate

`func (o *GetServProfServicesRespContent) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *GetServProfServicesRespContent) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *GetServProfServicesRespContent) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *GetServProfServicesRespContent) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetServProfServicesRespContent) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetServProfServicesRespContent) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetServProfServicesRespContent) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetServProfServicesRespContent) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *GetServProfServicesRespContent) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *GetServProfServicesRespContent) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *GetServProfServicesRespContent) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *GetServProfServicesRespContent) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *GetServProfServicesRespContent) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetServProfServicesRespContent) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetServProfServicesRespContent) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *GetServProfServicesRespContent) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetVlanSameAsPrimary

`func (o *GetServProfServicesRespContent) GetVlanSameAsPrimary() bool`

GetVlanSameAsPrimary returns the VlanSameAsPrimary field if non-nil, zero value otherwise.

### GetVlanSameAsPrimaryOk

`func (o *GetServProfServicesRespContent) GetVlanSameAsPrimaryOk() (*bool, bool)`

GetVlanSameAsPrimaryOk returns a tuple with the VlanSameAsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSameAsPrimary

`func (o *GetServProfServicesRespContent) SetVlanSameAsPrimary(v bool)`

SetVlanSameAsPrimary sets VlanSameAsPrimary field to given value.

### HasVlanSameAsPrimary

`func (o *GetServProfServicesRespContent) HasVlanSameAsPrimary() bool`

HasVlanSameAsPrimary returns a boolean if a field has been set.

### GetTagType

`func (o *GetServProfServicesRespContent) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *GetServProfServicesRespContent) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *GetServProfServicesRespContent) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *GetServProfServicesRespContent) HasTagType() bool`

HasTagType returns a boolean if a field has been set.

### GetCtagLabel

`func (o *GetServProfServicesRespContent) GetCtagLabel() string`

GetCtagLabel returns the CtagLabel field if non-nil, zero value otherwise.

### GetCtagLabelOk

`func (o *GetServProfServicesRespContent) GetCtagLabelOk() (*string, bool)`

GetCtagLabelOk returns a tuple with the CtagLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtagLabel

`func (o *GetServProfServicesRespContent) SetCtagLabel(v string)`

SetCtagLabel sets CtagLabel field to given value.

### HasCtagLabel

`func (o *GetServProfServicesRespContent) HasCtagLabel() bool`

HasCtagLabel returns a boolean if a field has been set.

### GetApiAvailable

`func (o *GetServProfServicesRespContent) GetApiAvailable() bool`

GetApiAvailable returns the ApiAvailable field if non-nil, zero value otherwise.

### GetApiAvailableOk

`func (o *GetServProfServicesRespContent) GetApiAvailableOk() (*bool, bool)`

GetApiAvailableOk returns a tuple with the ApiAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAvailable

`func (o *GetServProfServicesRespContent) SetApiAvailable(v bool)`

SetApiAvailable sets ApiAvailable field to given value.

### HasApiAvailable

`func (o *GetServProfServicesRespContent) HasApiAvailable() bool`

HasApiAvailable returns a boolean if a field has been set.

### GetSelfProfile

`func (o *GetServProfServicesRespContent) GetSelfProfile() bool`

GetSelfProfile returns the SelfProfile field if non-nil, zero value otherwise.

### GetSelfProfileOk

`func (o *GetServProfServicesRespContent) GetSelfProfileOk() (*bool, bool)`

GetSelfProfileOk returns a tuple with the SelfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfProfile

`func (o *GetServProfServicesRespContent) SetSelfProfile(v bool)`

SetSelfProfile sets SelfProfile field to given value.

### HasSelfProfile

`func (o *GetServProfServicesRespContent) HasSelfProfile() bool`

HasSelfProfile returns a boolean if a field has been set.

### GetProfileEncapsulation

`func (o *GetServProfServicesRespContent) GetProfileEncapsulation() string`

GetProfileEncapsulation returns the ProfileEncapsulation field if non-nil, zero value otherwise.

### GetProfileEncapsulationOk

`func (o *GetServProfServicesRespContent) GetProfileEncapsulationOk() (*string, bool)`

GetProfileEncapsulationOk returns a tuple with the ProfileEncapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEncapsulation

`func (o *GetServProfServicesRespContent) SetProfileEncapsulation(v string)`

SetProfileEncapsulation sets ProfileEncapsulation field to given value.

### HasProfileEncapsulation

`func (o *GetServProfServicesRespContent) HasProfileEncapsulation() bool`

HasProfileEncapsulation returns a boolean if a field has been set.

### GetAuthorizationKey

`func (o *GetServProfServicesRespContent) GetAuthorizationKey() string`

GetAuthorizationKey returns the AuthorizationKey field if non-nil, zero value otherwise.

### GetAuthorizationKeyOk

`func (o *GetServProfServicesRespContent) GetAuthorizationKeyOk() (*string, bool)`

GetAuthorizationKeyOk returns a tuple with the AuthorizationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationKey

`func (o *GetServProfServicesRespContent) SetAuthorizationKey(v string)`

SetAuthorizationKey sets AuthorizationKey field to given value.

### HasAuthorizationKey

`func (o *GetServProfServicesRespContent) HasAuthorizationKey() bool`

HasAuthorizationKey returns a boolean if a field has been set.

### GetOrganizationName

`func (o *GetServProfServicesRespContent) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GetServProfServicesRespContent) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GetServProfServicesRespContent) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *GetServProfServicesRespContent) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPrivate

`func (o *GetServProfServicesRespContent) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GetServProfServicesRespContent) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GetServProfServicesRespContent) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GetServProfServicesRespContent) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetFeatures

`func (o *GetServProfServicesRespContent) GetFeatures() GetServProfServicesRespContentfeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GetServProfServicesRespContent) GetFeaturesOk() (*GetServProfServicesRespContentfeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GetServProfServicesRespContent) SetFeatures(v GetServProfServicesRespContentfeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GetServProfServicesRespContent) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



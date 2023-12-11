# Entitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**FeatureAccess** | Pointer to **map[string]interface{}** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**InstanceQuota** | Pointer to **map[string]interface{}** |  | [optional] 
**IpQuota** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectQuota** | Pointer to **int32** |  | [optional] [default to 0]
**Slug** | **string** |  | 
**VolumeLimits** | Pointer to **map[string]interface{}** |  | [optional] 
**VolumeQuota** | Pointer to **map[string]interface{}** |  | [optional] 
**Weight** | **int32** |  | 

## Methods

### NewEntitlement

`func NewEntitlement(id string, slug string, weight int32, ) *Entitlement`

NewEntitlement instantiates a new Entitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementWithDefaults

`func NewEntitlementWithDefaults() *Entitlement`

NewEntitlementWithDefaults instantiates a new Entitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Entitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Entitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Entitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Entitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatureAccess

`func (o *Entitlement) GetFeatureAccess() map[string]interface{}`

GetFeatureAccess returns the FeatureAccess field if non-nil, zero value otherwise.

### GetFeatureAccessOk

`func (o *Entitlement) GetFeatureAccessOk() (*map[string]interface{}, bool)`

GetFeatureAccessOk returns a tuple with the FeatureAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureAccess

`func (o *Entitlement) SetFeatureAccess(v map[string]interface{})`

SetFeatureAccess sets FeatureAccess field to given value.

### HasFeatureAccess

`func (o *Entitlement) HasFeatureAccess() bool`

HasFeatureAccess returns a boolean if a field has been set.

### GetHref

`func (o *Entitlement) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Entitlement) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Entitlement) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Entitlement) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Entitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlement) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceQuota

`func (o *Entitlement) GetInstanceQuota() map[string]interface{}`

GetInstanceQuota returns the InstanceQuota field if non-nil, zero value otherwise.

### GetInstanceQuotaOk

`func (o *Entitlement) GetInstanceQuotaOk() (*map[string]interface{}, bool)`

GetInstanceQuotaOk returns a tuple with the InstanceQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceQuota

`func (o *Entitlement) SetInstanceQuota(v map[string]interface{})`

SetInstanceQuota sets InstanceQuota field to given value.

### HasInstanceQuota

`func (o *Entitlement) HasInstanceQuota() bool`

HasInstanceQuota returns a boolean if a field has been set.

### GetIpQuota

`func (o *Entitlement) GetIpQuota() map[string]interface{}`

GetIpQuota returns the IpQuota field if non-nil, zero value otherwise.

### GetIpQuotaOk

`func (o *Entitlement) GetIpQuotaOk() (*map[string]interface{}, bool)`

GetIpQuotaOk returns a tuple with the IpQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpQuota

`func (o *Entitlement) SetIpQuota(v map[string]interface{})`

SetIpQuota sets IpQuota field to given value.

### HasIpQuota

`func (o *Entitlement) HasIpQuota() bool`

HasIpQuota returns a boolean if a field has been set.

### GetName

`func (o *Entitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectQuota

`func (o *Entitlement) GetProjectQuota() int32`

GetProjectQuota returns the ProjectQuota field if non-nil, zero value otherwise.

### GetProjectQuotaOk

`func (o *Entitlement) GetProjectQuotaOk() (*int32, bool)`

GetProjectQuotaOk returns a tuple with the ProjectQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectQuota

`func (o *Entitlement) SetProjectQuota(v int32)`

SetProjectQuota sets ProjectQuota field to given value.

### HasProjectQuota

`func (o *Entitlement) HasProjectQuota() bool`

HasProjectQuota returns a boolean if a field has been set.

### GetSlug

`func (o *Entitlement) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Entitlement) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Entitlement) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetVolumeLimits

`func (o *Entitlement) GetVolumeLimits() map[string]interface{}`

GetVolumeLimits returns the VolumeLimits field if non-nil, zero value otherwise.

### GetVolumeLimitsOk

`func (o *Entitlement) GetVolumeLimitsOk() (*map[string]interface{}, bool)`

GetVolumeLimitsOk returns a tuple with the VolumeLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLimits

`func (o *Entitlement) SetVolumeLimits(v map[string]interface{})`

SetVolumeLimits sets VolumeLimits field to given value.

### HasVolumeLimits

`func (o *Entitlement) HasVolumeLimits() bool`

HasVolumeLimits returns a boolean if a field has been set.

### GetVolumeQuota

`func (o *Entitlement) GetVolumeQuota() map[string]interface{}`

GetVolumeQuota returns the VolumeQuota field if non-nil, zero value otherwise.

### GetVolumeQuotaOk

`func (o *Entitlement) GetVolumeQuotaOk() (*map[string]interface{}, bool)`

GetVolumeQuotaOk returns a tuple with the VolumeQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeQuota

`func (o *Entitlement) SetVolumeQuota(v map[string]interface{})`

SetVolumeQuota sets VolumeQuota field to given value.

### HasVolumeQuota

`func (o *Entitlement) HasVolumeQuota() bool`

HasVolumeQuota returns a boolean if a field has been set.

### GetWeight

`func (o *Entitlement) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Entitlement) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Entitlement) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildDate** | Pointer to **string** | The date when Metal&#39;s current OS image was built and released. | [optional] 
**DeprecationDate** | Pointer to **string** | The date when the OS is deprecated. This is the date set by the upstream operating system maintainer. | [optional] 
**Distro** | Pointer to **string** |  | [optional] 
**DistroLabel** | Pointer to **string** |  | [optional] 
**EndOfLifeDate** | Pointer to **string** | The date when the Metal OS image no longer receives any updates and may be disabled at any time. Typically the same as the deprecation date set by the upstream OS maintainers. | [optional] 
**EndOfServiceDate** | Pointer to **string** | Metal-set date for when the OS is nearing end of life, typically 30 days before end of life. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Licensed** | Pointer to **bool** | Indicates if the OS is licensed or not. Licensed operating systems are priced according to pricing property. | [optional] 
**LifecycleState** | Pointer to **string** | Where in the lifecycle the OS image is. Possible states are - &#x60;\&quot;testing\&quot;&#x60;, &#x60;\&quot;pre_release\&quot;&#x60;, &#x60;\&quot;active\&quot;&#x60;, &#x60;\&quot;deprecated\&quot;&#x60;, &#x60;\&quot;end_of_service\&quot;&#x60;, or &#x60;\&quot;end_of_life\&quot;&#x60;. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Preinstallable** | Pointer to **bool** | Indicates whether servers can be preinstalled with OS image in order to shorten provision time. | [optional] 
**Pricing** | Pointer to **map[string]interface{}** | This object contains price per time unit and optional multiplier value if license price depends on hardware plan or components (e.g. number of cores). | [optional] 
**ProvisionableOn** | Pointer to **[]string** |  | [optional] 
**ReleaseDate** | Pointer to **string** | The date the upstream operating system maintainer released this version of the OS. | [optional] 
**ReleaseNotes** | Pointer to **string** | The current release notes for this OS image, typically in Markdown format. | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**DefaultOperatingSystem** | Pointer to **bool** | Default operating system for the distro. | [optional] [readonly] 

## Methods

### NewOperatingSystem

`func NewOperatingSystem() *OperatingSystem`

NewOperatingSystem instantiates a new OperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemWithDefaults

`func NewOperatingSystemWithDefaults() *OperatingSystem`

NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildDate

`func (o *OperatingSystem) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *OperatingSystem) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *OperatingSystem) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *OperatingSystem) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetDeprecationDate

`func (o *OperatingSystem) GetDeprecationDate() string`

GetDeprecationDate returns the DeprecationDate field if non-nil, zero value otherwise.

### GetDeprecationDateOk

`func (o *OperatingSystem) GetDeprecationDateOk() (*string, bool)`

GetDeprecationDateOk returns a tuple with the DeprecationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationDate

`func (o *OperatingSystem) SetDeprecationDate(v string)`

SetDeprecationDate sets DeprecationDate field to given value.

### HasDeprecationDate

`func (o *OperatingSystem) HasDeprecationDate() bool`

HasDeprecationDate returns a boolean if a field has been set.

### GetDistro

`func (o *OperatingSystem) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *OperatingSystem) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *OperatingSystem) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *OperatingSystem) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetDistroLabel

`func (o *OperatingSystem) GetDistroLabel() string`

GetDistroLabel returns the DistroLabel field if non-nil, zero value otherwise.

### GetDistroLabelOk

`func (o *OperatingSystem) GetDistroLabelOk() (*string, bool)`

GetDistroLabelOk returns a tuple with the DistroLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroLabel

`func (o *OperatingSystem) SetDistroLabel(v string)`

SetDistroLabel sets DistroLabel field to given value.

### HasDistroLabel

`func (o *OperatingSystem) HasDistroLabel() bool`

HasDistroLabel returns a boolean if a field has been set.

### GetEndOfLifeDate

`func (o *OperatingSystem) GetEndOfLifeDate() string`

GetEndOfLifeDate returns the EndOfLifeDate field if non-nil, zero value otherwise.

### GetEndOfLifeDateOk

`func (o *OperatingSystem) GetEndOfLifeDateOk() (*string, bool)`

GetEndOfLifeDateOk returns a tuple with the EndOfLifeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeDate

`func (o *OperatingSystem) SetEndOfLifeDate(v string)`

SetEndOfLifeDate sets EndOfLifeDate field to given value.

### HasEndOfLifeDate

`func (o *OperatingSystem) HasEndOfLifeDate() bool`

HasEndOfLifeDate returns a boolean if a field has been set.

### GetEndOfServiceDate

`func (o *OperatingSystem) GetEndOfServiceDate() string`

GetEndOfServiceDate returns the EndOfServiceDate field if non-nil, zero value otherwise.

### GetEndOfServiceDateOk

`func (o *OperatingSystem) GetEndOfServiceDateOk() (*string, bool)`

GetEndOfServiceDateOk returns a tuple with the EndOfServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfServiceDate

`func (o *OperatingSystem) SetEndOfServiceDate(v string)`

SetEndOfServiceDate sets EndOfServiceDate field to given value.

### HasEndOfServiceDate

`func (o *OperatingSystem) HasEndOfServiceDate() bool`

HasEndOfServiceDate returns a boolean if a field has been set.

### GetId

`func (o *OperatingSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatingSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatingSystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OperatingSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicensed

`func (o *OperatingSystem) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *OperatingSystem) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *OperatingSystem) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *OperatingSystem) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetLifecycleState

`func (o *OperatingSystem) GetLifecycleState() string`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *OperatingSystem) GetLifecycleStateOk() (*string, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *OperatingSystem) SetLifecycleState(v string)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *OperatingSystem) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetName

`func (o *OperatingSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatingSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatingSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OperatingSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreinstallable

`func (o *OperatingSystem) GetPreinstallable() bool`

GetPreinstallable returns the Preinstallable field if non-nil, zero value otherwise.

### GetPreinstallableOk

`func (o *OperatingSystem) GetPreinstallableOk() (*bool, bool)`

GetPreinstallableOk returns a tuple with the Preinstallable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreinstallable

`func (o *OperatingSystem) SetPreinstallable(v bool)`

SetPreinstallable sets Preinstallable field to given value.

### HasPreinstallable

`func (o *OperatingSystem) HasPreinstallable() bool`

HasPreinstallable returns a boolean if a field has been set.

### GetPricing

`func (o *OperatingSystem) GetPricing() map[string]interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *OperatingSystem) GetPricingOk() (*map[string]interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *OperatingSystem) SetPricing(v map[string]interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *OperatingSystem) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvisionableOn

`func (o *OperatingSystem) GetProvisionableOn() []string`

GetProvisionableOn returns the ProvisionableOn field if non-nil, zero value otherwise.

### GetProvisionableOnOk

`func (o *OperatingSystem) GetProvisionableOnOk() (*[]string, bool)`

GetProvisionableOnOk returns a tuple with the ProvisionableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionableOn

`func (o *OperatingSystem) SetProvisionableOn(v []string)`

SetProvisionableOn sets ProvisionableOn field to given value.

### HasProvisionableOn

`func (o *OperatingSystem) HasProvisionableOn() bool`

HasProvisionableOn returns a boolean if a field has been set.

### GetReleaseDate

`func (o *OperatingSystem) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *OperatingSystem) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *OperatingSystem) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *OperatingSystem) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleaseNotes

`func (o *OperatingSystem) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *OperatingSystem) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *OperatingSystem) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *OperatingSystem) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetSlug

`func (o *OperatingSystem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OperatingSystem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OperatingSystem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OperatingSystem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVersion

`func (o *OperatingSystem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OperatingSystem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OperatingSystem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OperatingSystem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefaultOperatingSystem

`func (o *OperatingSystem) GetDefaultOperatingSystem() bool`

GetDefaultOperatingSystem returns the DefaultOperatingSystem field if non-nil, zero value otherwise.

### GetDefaultOperatingSystemOk

`func (o *OperatingSystem) GetDefaultOperatingSystemOk() (*bool, bool)`

GetDefaultOperatingSystemOk returns a tuple with the DefaultOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOperatingSystem

`func (o *OperatingSystem) SetDefaultOperatingSystem(v bool)`

SetDefaultOperatingSystem sets DefaultOperatingSystem field to given value.

### HasDefaultOperatingSystem

`func (o *OperatingSystem) HasDefaultOperatingSystem() bool`

HasDefaultOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



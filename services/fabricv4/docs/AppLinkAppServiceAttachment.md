# AppLinkAppServiceAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | Pointer to [**AppServiceType**](AppServiceType.md) |  | [optional] [default to APPSERVICETYPE_APP_SERVICE]
**Uuid** | **string** | Equinix-assigned access point identifier | 
**GeoScope** | **string** | Geo scope for the App Service | 
**DestinationIp** | **string** | Target IP for forwarding API requests | 
**AttachmentStatus** | Pointer to [**AppLinkAttachState**](AppLinkAttachState.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**AppLinkAppServiceAttachmentChange**](AppLinkAppServiceAttachmentChange.md) |  | [optional] 

## Methods

### NewAppLinkAppServiceAttachment

`func NewAppLinkAppServiceAttachment(uuid string, geoScope string, destinationIp string, ) *AppLinkAppServiceAttachment`

NewAppLinkAppServiceAttachment instantiates a new AppLinkAppServiceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAppServiceAttachmentWithDefaults

`func NewAppLinkAppServiceAttachmentWithDefaults() *AppLinkAppServiceAttachment`

NewAppLinkAppServiceAttachmentWithDefaults instantiates a new AppLinkAppServiceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppLinkAppServiceAttachment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppLinkAppServiceAttachment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppLinkAppServiceAttachment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppLinkAppServiceAttachment) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppLinkAppServiceAttachment) GetType() AppServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkAppServiceAttachment) GetTypeOk() (*AppServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkAppServiceAttachment) SetType(v AppServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *AppLinkAppServiceAttachment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AppLinkAppServiceAttachment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkAppServiceAttachment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkAppServiceAttachment) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetGeoScope

`func (o *AppLinkAppServiceAttachment) GetGeoScope() string`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *AppLinkAppServiceAttachment) GetGeoScopeOk() (*string, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *AppLinkAppServiceAttachment) SetGeoScope(v string)`

SetGeoScope sets GeoScope field to given value.


### GetDestinationIp

`func (o *AppLinkAppServiceAttachment) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *AppLinkAppServiceAttachment) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *AppLinkAppServiceAttachment) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.


### GetAttachmentStatus

`func (o *AppLinkAppServiceAttachment) GetAttachmentStatus() AppLinkAttachState`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AppLinkAppServiceAttachment) GetAttachmentStatusOk() (*AppLinkAttachState, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AppLinkAppServiceAttachment) SetAttachmentStatus(v AppLinkAttachState)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AppLinkAppServiceAttachment) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.

### GetChangeLog

`func (o *AppLinkAppServiceAttachment) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppLinkAppServiceAttachment) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppLinkAppServiceAttachment) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppLinkAppServiceAttachment) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *AppLinkAppServiceAttachment) GetChange() AppLinkAppServiceAttachmentChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AppLinkAppServiceAttachment) GetChangeOk() (*AppLinkAppServiceAttachmentChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AppLinkAppServiceAttachment) SetChange(v AppLinkAppServiceAttachmentChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *AppLinkAppServiceAttachment) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



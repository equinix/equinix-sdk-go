# AppLinkAttachedAppService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppServiceType**](AppServiceType.md) |  | [default to APPSERVICETYPE_APP_SERVICE]
**Uuid** | **string** | Equinix-assigned access point identifier | 
**Name** | Pointer to **string** | Customer-provided App Service name | [optional] 
**Description** | Pointer to **string** | Customer-provided App Service description | [optional] 
**State** | Pointer to [**AppServiceState**](AppServiceState.md) |  | [optional] 
**Endpoint** | Pointer to **string** | Accessible endpoint through this service | [optional] 
**SourceDomains** | Pointer to **[]string** | List of source domains from where traffic is allowed | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**GeoScope** | **string** | Geo scope for the App Service | 
**DestinationIp** | **string** | Target IP for forwarding API requests | 
**AttachmentStatus** | Pointer to [**AppLinkAttachState**](AppLinkAttachState.md) |  | [optional] 

## Methods

### NewAppLinkAttachedAppService

`func NewAppLinkAttachedAppService(type_ AppServiceType, uuid string, geoScope string, destinationIp string, ) *AppLinkAttachedAppService`

NewAppLinkAttachedAppService instantiates a new AppLinkAttachedAppService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachedAppServiceWithDefaults

`func NewAppLinkAttachedAppServiceWithDefaults() *AppLinkAttachedAppService`

NewAppLinkAttachedAppServiceWithDefaults instantiates a new AppLinkAttachedAppService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppLinkAttachedAppService) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppLinkAttachedAppService) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppLinkAttachedAppService) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppLinkAttachedAppService) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppLinkAttachedAppService) GetType() AppServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkAttachedAppService) GetTypeOk() (*AppServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkAttachedAppService) SetType(v AppServiceType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppLinkAttachedAppService) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkAttachedAppService) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkAttachedAppService) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppLinkAttachedAppService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLinkAttachedAppService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLinkAttachedAppService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppLinkAttachedAppService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AppLinkAttachedAppService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppLinkAttachedAppService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppLinkAttachedAppService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppLinkAttachedAppService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppLinkAttachedAppService) GetState() AppServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppLinkAttachedAppService) GetStateOk() (*AppServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppLinkAttachedAppService) SetState(v AppServiceState)`

SetState sets State field to given value.

### HasState

`func (o *AppLinkAttachedAppService) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEndpoint

`func (o *AppLinkAttachedAppService) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AppLinkAttachedAppService) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AppLinkAttachedAppService) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AppLinkAttachedAppService) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetSourceDomains

`func (o *AppLinkAttachedAppService) GetSourceDomains() []string`

GetSourceDomains returns the SourceDomains field if non-nil, zero value otherwise.

### GetSourceDomainsOk

`func (o *AppLinkAttachedAppService) GetSourceDomainsOk() (*[]string, bool)`

GetSourceDomainsOk returns a tuple with the SourceDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDomains

`func (o *AppLinkAttachedAppService) SetSourceDomains(v []string)`

SetSourceDomains sets SourceDomains field to given value.

### HasSourceDomains

`func (o *AppLinkAttachedAppService) HasSourceDomains() bool`

HasSourceDomains returns a boolean if a field has been set.

### GetProject

`func (o *AppLinkAttachedAppService) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppLinkAttachedAppService) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppLinkAttachedAppService) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *AppLinkAttachedAppService) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetGeoScope

`func (o *AppLinkAttachedAppService) GetGeoScope() string`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *AppLinkAttachedAppService) GetGeoScopeOk() (*string, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *AppLinkAttachedAppService) SetGeoScope(v string)`

SetGeoScope sets GeoScope field to given value.


### GetDestinationIp

`func (o *AppLinkAttachedAppService) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *AppLinkAttachedAppService) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *AppLinkAttachedAppService) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.


### GetAttachmentStatus

`func (o *AppLinkAttachedAppService) GetAttachmentStatus() AppLinkAttachState`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AppLinkAttachedAppService) GetAttachmentStatusOk() (*AppLinkAttachState, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AppLinkAttachedAppService) SetAttachmentStatus(v AppLinkAttachState)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AppLinkAttachedAppService) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



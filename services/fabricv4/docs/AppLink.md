# AppLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppLinkType**](AppLinkType.md) |  | [default to APPLINKTYPE_APP_LINK]
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | **string** | Customer-provided App Link name | 
**Description** | Pointer to **string** | Customer-provided App Link description | [optional] 
**State** | Pointer to [**AppLinkState**](AppLinkState.md) |  | [optional] 
**Router** | [**AppLinkCloudRouter**](AppLinkCloudRouter.md) |  | 
**Ipv4Address** | Pointer to **string** | App Link IP address | [optional] 
**Bandwidth** | **int32** | App Link aggregated data transfer capacity in Mbps | 
**Project** | [**Project**](Project.md) |  | 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**AppLinkChange**](AppLinkChange.md) |  | [optional] 

## Methods

### NewAppLink

`func NewAppLink(type_ AppLinkType, name string, router AppLinkCloudRouter, bandwidth int32, project Project, ) *AppLink`

NewAppLink instantiates a new AppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkWithDefaults

`func NewAppLinkWithDefaults() *AppLink`

NewAppLinkWithDefaults instantiates a new AppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppLink) GetType() AppLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLink) GetTypeOk() (*AppLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLink) SetType(v AppLinkType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppLink) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLink) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLink) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppLink) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AppLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLink) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppLink) GetState() AppLinkState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppLink) GetStateOk() (*AppLinkState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppLink) SetState(v AppLinkState)`

SetState sets State field to given value.

### HasState

`func (o *AppLink) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRouter

`func (o *AppLink) GetRouter() AppLinkCloudRouter`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *AppLink) GetRouterOk() (*AppLinkCloudRouter, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *AppLink) SetRouter(v AppLinkCloudRouter)`

SetRouter sets Router field to given value.


### GetIpv4Address

`func (o *AppLink) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *AppLink) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *AppLink) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *AppLink) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetBandwidth

`func (o *AppLink) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *AppLink) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *AppLink) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetProject

`func (o *AppLink) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppLink) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppLink) SetProject(v Project)`

SetProject sets Project field to given value.


### GetChangeLog

`func (o *AppLink) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppLink) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppLink) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppLink) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *AppLink) GetChange() AppLinkChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AppLink) GetChangeOk() (*AppLinkChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AppLink) SetChange(v AppLinkChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *AppLink) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



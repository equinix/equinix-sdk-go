# AppDomainAttachedAppLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppLinkType**](AppLinkType.md) |  | [default to APPLINKTYPE_APP_LINK]
**Uuid** | **string** | Equinix-assigned access point identifier | 
**Name** | Pointer to **string** | Customer-provided App Link name | [optional] 
**Description** | Pointer to **string** | Customer-provided App Link description | [optional] 
**State** | Pointer to [**AppLinkState**](AppLinkState.md) |  | [optional] 
**Router** | Pointer to [**AppLinkCloudRouter**](AppLinkCloudRouter.md) |  | [optional] 
**Ipv4Address** | Pointer to **string** | App Link IP address | [optional] 
**Bandwidth** | Pointer to **int32** | App Link aggregated data transfer capacity in Mbps | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewAppDomainAttachedAppLink

`func NewAppDomainAttachedAppLink(type_ AppLinkType, uuid string, ) *AppDomainAttachedAppLink`

NewAppDomainAttachedAppLink instantiates a new AppDomainAttachedAppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDomainAttachedAppLinkWithDefaults

`func NewAppDomainAttachedAppLinkWithDefaults() *AppDomainAttachedAppLink`

NewAppDomainAttachedAppLinkWithDefaults instantiates a new AppDomainAttachedAppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppDomainAttachedAppLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppDomainAttachedAppLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppDomainAttachedAppLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppDomainAttachedAppLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppDomainAttachedAppLink) GetType() AppLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppDomainAttachedAppLink) GetTypeOk() (*AppLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppDomainAttachedAppLink) SetType(v AppLinkType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppDomainAttachedAppLink) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppDomainAttachedAppLink) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppDomainAttachedAppLink) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppDomainAttachedAppLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDomainAttachedAppLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDomainAttachedAppLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppDomainAttachedAppLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AppDomainAttachedAppLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDomainAttachedAppLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDomainAttachedAppLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDomainAttachedAppLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppDomainAttachedAppLink) GetState() AppLinkState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppDomainAttachedAppLink) GetStateOk() (*AppLinkState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppDomainAttachedAppLink) SetState(v AppLinkState)`

SetState sets State field to given value.

### HasState

`func (o *AppDomainAttachedAppLink) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRouter

`func (o *AppDomainAttachedAppLink) GetRouter() AppLinkCloudRouter`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *AppDomainAttachedAppLink) GetRouterOk() (*AppLinkCloudRouter, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *AppDomainAttachedAppLink) SetRouter(v AppLinkCloudRouter)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *AppDomainAttachedAppLink) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetIpv4Address

`func (o *AppDomainAttachedAppLink) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *AppDomainAttachedAppLink) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *AppDomainAttachedAppLink) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *AppDomainAttachedAppLink) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetBandwidth

`func (o *AppDomainAttachedAppLink) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *AppDomainAttachedAppLink) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *AppDomainAttachedAppLink) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *AppDomainAttachedAppLink) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetProject

`func (o *AppDomainAttachedAppLink) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppDomainAttachedAppLink) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppDomainAttachedAppLink) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *AppDomainAttachedAppLink) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



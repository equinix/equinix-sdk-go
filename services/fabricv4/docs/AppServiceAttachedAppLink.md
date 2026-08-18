# AppServiceAttachedAppLink

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

### NewAppServiceAttachedAppLink

`func NewAppServiceAttachedAppLink(type_ AppLinkType, uuid string, ) *AppServiceAttachedAppLink`

NewAppServiceAttachedAppLink instantiates a new AppServiceAttachedAppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAttachedAppLinkWithDefaults

`func NewAppServiceAttachedAppLinkWithDefaults() *AppServiceAttachedAppLink`

NewAppServiceAttachedAppLinkWithDefaults instantiates a new AppServiceAttachedAppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppServiceAttachedAppLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppServiceAttachedAppLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppServiceAttachedAppLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppServiceAttachedAppLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppServiceAttachedAppLink) GetType() AppLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppServiceAttachedAppLink) GetTypeOk() (*AppLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppServiceAttachedAppLink) SetType(v AppLinkType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppServiceAttachedAppLink) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppServiceAttachedAppLink) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppServiceAttachedAppLink) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppServiceAttachedAppLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceAttachedAppLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceAttachedAppLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppServiceAttachedAppLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AppServiceAttachedAppLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceAttachedAppLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceAttachedAppLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceAttachedAppLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppServiceAttachedAppLink) GetState() AppLinkState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppServiceAttachedAppLink) GetStateOk() (*AppLinkState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppServiceAttachedAppLink) SetState(v AppLinkState)`

SetState sets State field to given value.

### HasState

`func (o *AppServiceAttachedAppLink) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRouter

`func (o *AppServiceAttachedAppLink) GetRouter() AppLinkCloudRouter`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *AppServiceAttachedAppLink) GetRouterOk() (*AppLinkCloudRouter, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *AppServiceAttachedAppLink) SetRouter(v AppLinkCloudRouter)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *AppServiceAttachedAppLink) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetIpv4Address

`func (o *AppServiceAttachedAppLink) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *AppServiceAttachedAppLink) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *AppServiceAttachedAppLink) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *AppServiceAttachedAppLink) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetBandwidth

`func (o *AppServiceAttachedAppLink) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *AppServiceAttachedAppLink) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *AppServiceAttachedAppLink) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *AppServiceAttachedAppLink) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetProject

`func (o *AppServiceAttachedAppLink) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppServiceAttachedAppLink) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppServiceAttachedAppLink) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *AppServiceAttachedAppLink) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



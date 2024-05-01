# RoutingProtocolReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Routing protocol URI | [optional] 
**Uuid** | Pointer to **string** | Routing protocol identifier | [optional] 
**Type** | [**RoutingProtocolType**](RoutingProtocolType.md) |  | 
**Name** | **string** | Name of the routing protocol instance. | 
**Description** | Pointer to **string** | Description of the routing protocol instance | [optional] 
**Ipv4** | Pointer to [**RoutingProtocolIpv4**](RoutingProtocolIpv4.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6**](RoutingProtocolIpv6.md) |  | [optional] 
**ChangeLog** | [**ChangeLog**](ChangeLog.md) |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Links** | [**[]Link**](Link.md) |  | 

## Methods

### NewRoutingProtocolReadModel

`func NewRoutingProtocolReadModel(type_ RoutingProtocolType, name string, changeLog ChangeLog, links []Link, ) *RoutingProtocolReadModel`

NewRoutingProtocolReadModel instantiates a new RoutingProtocolReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolReadModelWithDefaults

`func NewRoutingProtocolReadModelWithDefaults() *RoutingProtocolReadModel`

NewRoutingProtocolReadModelWithDefaults instantiates a new RoutingProtocolReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RoutingProtocolReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolReadModel) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolReadModel) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolReadModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolReadModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolReadModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolReadModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolReadModel) GetType() RoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolReadModel) GetTypeOk() (*RoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolReadModel) SetType(v RoutingProtocolType)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolReadModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolReadModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolReadModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoutingProtocolReadModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolReadModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolReadModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolReadModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *RoutingProtocolReadModel) GetIpv4() RoutingProtocolIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolReadModel) GetIpv4Ok() (*RoutingProtocolIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolReadModel) SetIpv4(v RoutingProtocolIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolReadModel) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolReadModel) GetIpv6() RoutingProtocolIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolReadModel) GetIpv6Ok() (*RoutingProtocolIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolReadModel) SetIpv6(v RoutingProtocolIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolReadModel) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetChangeLog

`func (o *RoutingProtocolReadModel) GetChangeLog() ChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RoutingProtocolReadModel) GetChangeLogOk() (*ChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RoutingProtocolReadModel) SetChangeLog(v ChangeLog)`

SetChangeLog sets ChangeLog field to given value.


### GetTags

`func (o *RoutingProtocolReadModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolReadModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolReadModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolReadModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLinks

`func (o *RoutingProtocolReadModel) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoutingProtocolReadModel) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoutingProtocolReadModel) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



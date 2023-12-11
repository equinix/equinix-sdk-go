# BgpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int32** | Autonomous System Number. ASN is required with Global BGP. With Local BGP the private ASN, 65000, is assigned. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeploymentType** | Pointer to [**BgpConfigDeploymentType**](BgpConfigDeploymentType.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxPrefix** | Pointer to **int32** | The maximum number of route filters allowed per server | [optional] [default to 10]
**Md5** | Pointer to **NullableString** | (Optional) Password for BGP session in plaintext (not a checksum) | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**Ranges** | Pointer to [**[]GlobalBgpRange**](GlobalBgpRange.md) | The IP block ranges associated to the ASN (Populated in Global BGP only) | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**RouteObject** | Pointer to **string** | Specifies AS-MACRO (aka AS-SET) to use when building client route filters | [optional] 
**Sessions** | Pointer to [**[]BgpSession**](BgpSession.md) | The direct connections between neighboring routers that want to exchange routing information. | [optional] 
**Status** | Pointer to [**BgpConfigStatus**](BgpConfigStatus.md) |  | [optional] 

## Methods

### NewBgpConfig

`func NewBgpConfig() *BgpConfig`

NewBgpConfig instantiates a new BgpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigWithDefaults

`func NewBgpConfigWithDefaults() *BgpConfig`

NewBgpConfigWithDefaults instantiates a new BgpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *BgpConfig) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BgpConfig) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BgpConfig) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *BgpConfig) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BgpConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BgpConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BgpConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BgpConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeploymentType

`func (o *BgpConfig) GetDeploymentType() BgpConfigDeploymentType`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *BgpConfig) GetDeploymentTypeOk() (*BgpConfigDeploymentType, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *BgpConfig) SetDeploymentType(v BgpConfigDeploymentType)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *BgpConfig) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetHref

`func (o *BgpConfig) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BgpConfig) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BgpConfig) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BgpConfig) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BgpConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BgpConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BgpConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BgpConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxPrefix

`func (o *BgpConfig) GetMaxPrefix() int32`

GetMaxPrefix returns the MaxPrefix field if non-nil, zero value otherwise.

### GetMaxPrefixOk

`func (o *BgpConfig) GetMaxPrefixOk() (*int32, bool)`

GetMaxPrefixOk returns a tuple with the MaxPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefix

`func (o *BgpConfig) SetMaxPrefix(v int32)`

SetMaxPrefix sets MaxPrefix field to given value.

### HasMaxPrefix

`func (o *BgpConfig) HasMaxPrefix() bool`

HasMaxPrefix returns a boolean if a field has been set.

### GetMd5

`func (o *BgpConfig) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *BgpConfig) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *BgpConfig) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *BgpConfig) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *BgpConfig) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *BgpConfig) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetProject

`func (o *BgpConfig) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BgpConfig) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BgpConfig) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *BgpConfig) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRanges

`func (o *BgpConfig) GetRanges() []GlobalBgpRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *BgpConfig) GetRangesOk() (*[]GlobalBgpRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *BgpConfig) SetRanges(v []GlobalBgpRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *BgpConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BgpConfig) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BgpConfig) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BgpConfig) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BgpConfig) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetRouteObject

`func (o *BgpConfig) GetRouteObject() string`

GetRouteObject returns the RouteObject field if non-nil, zero value otherwise.

### GetRouteObjectOk

`func (o *BgpConfig) GetRouteObjectOk() (*string, bool)`

GetRouteObjectOk returns a tuple with the RouteObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteObject

`func (o *BgpConfig) SetRouteObject(v string)`

SetRouteObject sets RouteObject field to given value.

### HasRouteObject

`func (o *BgpConfig) HasRouteObject() bool`

HasRouteObject returns a boolean if a field has been set.

### GetSessions

`func (o *BgpConfig) GetSessions() []BgpSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *BgpConfig) GetSessionsOk() (*[]BgpSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *BgpConfig) SetSessions(v []BgpSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *BgpConfig) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetStatus

`func (o *BgpConfig) GetStatus() BgpConfigStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BgpConfig) GetStatusOk() (*BgpConfigStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BgpConfig) SetStatus(v BgpConfigStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BgpConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



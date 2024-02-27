# LinkProtocolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | LinkProtocol URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned network identifier | [optional] 
**State** | Pointer to [**LinkProtocolState**](LinkProtocolState.md) |  | [optional] 
**Type** | Pointer to [**LinkProtocolRequestType**](LinkProtocolRequestType.md) |  | [optional] 
**VlanTag** | Pointer to **int32** |  | [optional] 
**Vni** | Pointer to **int32** |  | [optional] 
**VlanTagMin** | Pointer to **int32** |  | [optional] 
**VlanTagMax** | Pointer to **int32** |  | [optional] 
**VlanSTag** | Pointer to **int32** |  | [optional] 
**VlanCTag** | Pointer to **int32** |  | [optional] 
**VlanCTagMin** | Pointer to **int32** |  | [optional] 
**VlanCTagMax** | Pointer to **int32** |  | [optional] 
**SubInterface** | Pointer to [**SubInterface**](SubInterface.md) |  | [optional] 
**Asset** | Pointer to [**LinkProtocolConnection**](LinkProtocolConnection.md) |  | [optional] 
**ServiceToken** | Pointer to [**LinkProtocolServiceToken**](LinkProtocolServiceToken.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewLinkProtocolResponse

`func NewLinkProtocolResponse() *LinkProtocolResponse`

NewLinkProtocolResponse instantiates a new LinkProtocolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolResponseWithDefaults

`func NewLinkProtocolResponseWithDefaults() *LinkProtocolResponse`

NewLinkProtocolResponseWithDefaults instantiates a new LinkProtocolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LinkProtocolResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkProtocolResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkProtocolResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LinkProtocolResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *LinkProtocolResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LinkProtocolResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LinkProtocolResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LinkProtocolResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *LinkProtocolResponse) GetState() LinkProtocolState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkProtocolResponse) GetStateOk() (*LinkProtocolState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LinkProtocolResponse) SetState(v LinkProtocolState)`

SetState sets State field to given value.

### HasState

`func (o *LinkProtocolResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *LinkProtocolResponse) GetType() LinkProtocolRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocolResponse) GetTypeOk() (*LinkProtocolRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocolResponse) SetType(v LinkProtocolRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkProtocolResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlanTag

`func (o *LinkProtocolResponse) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *LinkProtocolResponse) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *LinkProtocolResponse) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *LinkProtocolResponse) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.

### GetVni

`func (o *LinkProtocolResponse) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *LinkProtocolResponse) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *LinkProtocolResponse) SetVni(v int32)`

SetVni sets Vni field to given value.

### HasVni

`func (o *LinkProtocolResponse) HasVni() bool`

HasVni returns a boolean if a field has been set.

### GetVlanTagMin

`func (o *LinkProtocolResponse) GetVlanTagMin() int32`

GetVlanTagMin returns the VlanTagMin field if non-nil, zero value otherwise.

### GetVlanTagMinOk

`func (o *LinkProtocolResponse) GetVlanTagMinOk() (*int32, bool)`

GetVlanTagMinOk returns a tuple with the VlanTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMin

`func (o *LinkProtocolResponse) SetVlanTagMin(v int32)`

SetVlanTagMin sets VlanTagMin field to given value.

### HasVlanTagMin

`func (o *LinkProtocolResponse) HasVlanTagMin() bool`

HasVlanTagMin returns a boolean if a field has been set.

### GetVlanTagMax

`func (o *LinkProtocolResponse) GetVlanTagMax() int32`

GetVlanTagMax returns the VlanTagMax field if non-nil, zero value otherwise.

### GetVlanTagMaxOk

`func (o *LinkProtocolResponse) GetVlanTagMaxOk() (*int32, bool)`

GetVlanTagMaxOk returns a tuple with the VlanTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMax

`func (o *LinkProtocolResponse) SetVlanTagMax(v int32)`

SetVlanTagMax sets VlanTagMax field to given value.

### HasVlanTagMax

`func (o *LinkProtocolResponse) HasVlanTagMax() bool`

HasVlanTagMax returns a boolean if a field has been set.

### GetVlanSTag

`func (o *LinkProtocolResponse) GetVlanSTag() int32`

GetVlanSTag returns the VlanSTag field if non-nil, zero value otherwise.

### GetVlanSTagOk

`func (o *LinkProtocolResponse) GetVlanSTagOk() (*int32, bool)`

GetVlanSTagOk returns a tuple with the VlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSTag

`func (o *LinkProtocolResponse) SetVlanSTag(v int32)`

SetVlanSTag sets VlanSTag field to given value.

### HasVlanSTag

`func (o *LinkProtocolResponse) HasVlanSTag() bool`

HasVlanSTag returns a boolean if a field has been set.

### GetVlanCTag

`func (o *LinkProtocolResponse) GetVlanCTag() int32`

GetVlanCTag returns the VlanCTag field if non-nil, zero value otherwise.

### GetVlanCTagOk

`func (o *LinkProtocolResponse) GetVlanCTagOk() (*int32, bool)`

GetVlanCTagOk returns a tuple with the VlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTag

`func (o *LinkProtocolResponse) SetVlanCTag(v int32)`

SetVlanCTag sets VlanCTag field to given value.

### HasVlanCTag

`func (o *LinkProtocolResponse) HasVlanCTag() bool`

HasVlanCTag returns a boolean if a field has been set.

### GetVlanCTagMin

`func (o *LinkProtocolResponse) GetVlanCTagMin() int32`

GetVlanCTagMin returns the VlanCTagMin field if non-nil, zero value otherwise.

### GetVlanCTagMinOk

`func (o *LinkProtocolResponse) GetVlanCTagMinOk() (*int32, bool)`

GetVlanCTagMinOk returns a tuple with the VlanCTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMin

`func (o *LinkProtocolResponse) SetVlanCTagMin(v int32)`

SetVlanCTagMin sets VlanCTagMin field to given value.

### HasVlanCTagMin

`func (o *LinkProtocolResponse) HasVlanCTagMin() bool`

HasVlanCTagMin returns a boolean if a field has been set.

### GetVlanCTagMax

`func (o *LinkProtocolResponse) GetVlanCTagMax() int32`

GetVlanCTagMax returns the VlanCTagMax field if non-nil, zero value otherwise.

### GetVlanCTagMaxOk

`func (o *LinkProtocolResponse) GetVlanCTagMaxOk() (*int32, bool)`

GetVlanCTagMaxOk returns a tuple with the VlanCTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMax

`func (o *LinkProtocolResponse) SetVlanCTagMax(v int32)`

SetVlanCTagMax sets VlanCTagMax field to given value.

### HasVlanCTagMax

`func (o *LinkProtocolResponse) HasVlanCTagMax() bool`

HasVlanCTagMax returns a boolean if a field has been set.

### GetSubInterface

`func (o *LinkProtocolResponse) GetSubInterface() SubInterface`

GetSubInterface returns the SubInterface field if non-nil, zero value otherwise.

### GetSubInterfaceOk

`func (o *LinkProtocolResponse) GetSubInterfaceOk() (*SubInterface, bool)`

GetSubInterfaceOk returns a tuple with the SubInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubInterface

`func (o *LinkProtocolResponse) SetSubInterface(v SubInterface)`

SetSubInterface sets SubInterface field to given value.

### HasSubInterface

`func (o *LinkProtocolResponse) HasSubInterface() bool`

HasSubInterface returns a boolean if a field has been set.

### GetAsset

`func (o *LinkProtocolResponse) GetAsset() LinkProtocolConnection`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *LinkProtocolResponse) GetAssetOk() (*LinkProtocolConnection, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *LinkProtocolResponse) SetAsset(v LinkProtocolConnection)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *LinkProtocolResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetServiceToken

`func (o *LinkProtocolResponse) GetServiceToken() LinkProtocolServiceToken`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *LinkProtocolResponse) GetServiceTokenOk() (*LinkProtocolServiceToken, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *LinkProtocolResponse) SetServiceToken(v LinkProtocolServiceToken)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *LinkProtocolResponse) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetChangeLog

`func (o *LinkProtocolResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *LinkProtocolResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *LinkProtocolResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *LinkProtocolResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



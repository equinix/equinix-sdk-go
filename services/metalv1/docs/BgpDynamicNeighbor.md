# BgpDynamicNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the resource | [optional] [readonly] 
**BgpNeighborAsn** | Pointer to **int32** | The ASN of the dynamic BGP neighbor | [optional] 
**BgpNeighborRange** | Pointer to **string** | Network range of the dynamic BGP neighbor in CIDR format | [optional] 
**MetalGateway** | Pointer to [**VrfMetalGateway**](VrfMetalGateway.md) |  | [optional] 
**State** | Pointer to [**BgpDynamicNeighborState**](BgpDynamicNeighborState.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserLimited**](UserLimited.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBgpDynamicNeighbor

`func NewBgpDynamicNeighbor() *BgpDynamicNeighbor`

NewBgpDynamicNeighbor instantiates a new BgpDynamicNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpDynamicNeighborWithDefaults

`func NewBgpDynamicNeighborWithDefaults() *BgpDynamicNeighbor`

NewBgpDynamicNeighborWithDefaults instantiates a new BgpDynamicNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BgpDynamicNeighbor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BgpDynamicNeighbor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BgpDynamicNeighbor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BgpDynamicNeighbor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBgpNeighborAsn

`func (o *BgpDynamicNeighbor) GetBgpNeighborAsn() int32`

GetBgpNeighborAsn returns the BgpNeighborAsn field if non-nil, zero value otherwise.

### GetBgpNeighborAsnOk

`func (o *BgpDynamicNeighbor) GetBgpNeighborAsnOk() (*int32, bool)`

GetBgpNeighborAsnOk returns a tuple with the BgpNeighborAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborAsn

`func (o *BgpDynamicNeighbor) SetBgpNeighborAsn(v int32)`

SetBgpNeighborAsn sets BgpNeighborAsn field to given value.

### HasBgpNeighborAsn

`func (o *BgpDynamicNeighbor) HasBgpNeighborAsn() bool`

HasBgpNeighborAsn returns a boolean if a field has been set.

### GetBgpNeighborRange

`func (o *BgpDynamicNeighbor) GetBgpNeighborRange() string`

GetBgpNeighborRange returns the BgpNeighborRange field if non-nil, zero value otherwise.

### GetBgpNeighborRangeOk

`func (o *BgpDynamicNeighbor) GetBgpNeighborRangeOk() (*string, bool)`

GetBgpNeighborRangeOk returns a tuple with the BgpNeighborRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborRange

`func (o *BgpDynamicNeighbor) SetBgpNeighborRange(v string)`

SetBgpNeighborRange sets BgpNeighborRange field to given value.

### HasBgpNeighborRange

`func (o *BgpDynamicNeighbor) HasBgpNeighborRange() bool`

HasBgpNeighborRange returns a boolean if a field has been set.

### GetMetalGateway

`func (o *BgpDynamicNeighbor) GetMetalGateway() VrfMetalGateway`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *BgpDynamicNeighbor) GetMetalGatewayOk() (*VrfMetalGateway, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *BgpDynamicNeighbor) SetMetalGateway(v VrfMetalGateway)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *BgpDynamicNeighbor) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetState

`func (o *BgpDynamicNeighbor) GetState() BgpDynamicNeighborState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BgpDynamicNeighbor) GetStateOk() (*BgpDynamicNeighborState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BgpDynamicNeighbor) SetState(v BgpDynamicNeighborState)`

SetState sets State field to given value.

### HasState

`func (o *BgpDynamicNeighbor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHref

`func (o *BgpDynamicNeighbor) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BgpDynamicNeighbor) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BgpDynamicNeighbor) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BgpDynamicNeighbor) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BgpDynamicNeighbor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BgpDynamicNeighbor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BgpDynamicNeighbor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BgpDynamicNeighbor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BgpDynamicNeighbor) GetCreatedBy() UserLimited`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BgpDynamicNeighbor) GetCreatedByOk() (*UserLimited, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BgpDynamicNeighbor) SetCreatedBy(v UserLimited)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BgpDynamicNeighbor) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BgpDynamicNeighbor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BgpDynamicNeighbor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BgpDynamicNeighbor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BgpDynamicNeighbor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *BgpDynamicNeighbor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BgpDynamicNeighbor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BgpDynamicNeighbor) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BgpDynamicNeighbor) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BgpDynamicNeighborCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpNeighborRange** | **string** | Network range of the dynamic BGP neighbor in CIDR format | 
**BgpNeighborAsn** | **int32** | The ASN of the dynamic BGP neighbor | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBgpDynamicNeighborCreateInput

`func NewBgpDynamicNeighborCreateInput(bgpNeighborRange string, bgpNeighborAsn int32, ) *BgpDynamicNeighborCreateInput`

NewBgpDynamicNeighborCreateInput instantiates a new BgpDynamicNeighborCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpDynamicNeighborCreateInputWithDefaults

`func NewBgpDynamicNeighborCreateInputWithDefaults() *BgpDynamicNeighborCreateInput`

NewBgpDynamicNeighborCreateInputWithDefaults instantiates a new BgpDynamicNeighborCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpNeighborRange

`func (o *BgpDynamicNeighborCreateInput) GetBgpNeighborRange() string`

GetBgpNeighborRange returns the BgpNeighborRange field if non-nil, zero value otherwise.

### GetBgpNeighborRangeOk

`func (o *BgpDynamicNeighborCreateInput) GetBgpNeighborRangeOk() (*string, bool)`

GetBgpNeighborRangeOk returns a tuple with the BgpNeighborRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborRange

`func (o *BgpDynamicNeighborCreateInput) SetBgpNeighborRange(v string)`

SetBgpNeighborRange sets BgpNeighborRange field to given value.


### GetBgpNeighborAsn

`func (o *BgpDynamicNeighborCreateInput) GetBgpNeighborAsn() int32`

GetBgpNeighborAsn returns the BgpNeighborAsn field if non-nil, zero value otherwise.

### GetBgpNeighborAsnOk

`func (o *BgpDynamicNeighborCreateInput) GetBgpNeighborAsnOk() (*int32, bool)`

GetBgpNeighborAsnOk returns a tuple with the BgpNeighborAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborAsn

`func (o *BgpDynamicNeighborCreateInput) SetBgpNeighborAsn(v int32)`

SetBgpNeighborAsn sets BgpNeighborAsn field to given value.


### GetTags

`func (o *BgpDynamicNeighborCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BgpDynamicNeighborCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BgpDynamicNeighborCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BgpDynamicNeighborCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



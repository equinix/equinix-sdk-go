# GlobalBgpRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobalBgpRange

`func NewGlobalBgpRange() *GlobalBgpRange`

NewGlobalBgpRange instantiates a new GlobalBgpRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalBgpRangeWithDefaults

`func NewGlobalBgpRangeWithDefaults() *GlobalBgpRange`

NewGlobalBgpRangeWithDefaults instantiates a new GlobalBgpRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *GlobalBgpRange) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *GlobalBgpRange) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *GlobalBgpRange) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *GlobalBgpRange) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetHref

`func (o *GlobalBgpRange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GlobalBgpRange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GlobalBgpRange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GlobalBgpRange) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *GlobalBgpRange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalBgpRange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalBgpRange) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalBgpRange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *GlobalBgpRange) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GlobalBgpRange) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GlobalBgpRange) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *GlobalBgpRange) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRange

`func (o *GlobalBgpRange) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GlobalBgpRange) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GlobalBgpRange) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GlobalBgpRange) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



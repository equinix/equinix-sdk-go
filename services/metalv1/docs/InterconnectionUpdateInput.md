# InterconnectionUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**InterconnectionMode**](InterconnectionMode.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Redundancy** | Pointer to **string** | Updating from &#39;redundant&#39; to &#39;primary&#39; will remove a secondary port, while updating from &#39;primary&#39; to &#39;redundant&#39; will add one. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInterconnectionUpdateInput

`func NewInterconnectionUpdateInput() *InterconnectionUpdateInput`

NewInterconnectionUpdateInput instantiates a new InterconnectionUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionUpdateInputWithDefaults

`func NewInterconnectionUpdateInputWithDefaults() *InterconnectionUpdateInput`

NewInterconnectionUpdateInputWithDefaults instantiates a new InterconnectionUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *InterconnectionUpdateInput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *InterconnectionUpdateInput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *InterconnectionUpdateInput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *InterconnectionUpdateInput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *InterconnectionUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterconnectionUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterconnectionUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterconnectionUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *InterconnectionUpdateInput) GetMode() InterconnectionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InterconnectionUpdateInput) GetModeOk() (*InterconnectionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InterconnectionUpdateInput) SetMode(v InterconnectionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InterconnectionUpdateInput) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *InterconnectionUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterconnectionUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterconnectionUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterconnectionUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedundancy

`func (o *InterconnectionUpdateInput) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *InterconnectionUpdateInput) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *InterconnectionUpdateInput) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *InterconnectionUpdateInput) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetTags

`func (o *InterconnectionUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InterconnectionUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InterconnectionUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InterconnectionUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



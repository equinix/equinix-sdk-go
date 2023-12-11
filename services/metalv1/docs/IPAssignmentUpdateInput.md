# IPAssignmentUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIPAssignmentUpdateInput

`func NewIPAssignmentUpdateInput() *IPAssignmentUpdateInput`

NewIPAssignmentUpdateInput instantiates a new IPAssignmentUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAssignmentUpdateInputWithDefaults

`func NewIPAssignmentUpdateInputWithDefaults() *IPAssignmentUpdateInput`

NewIPAssignmentUpdateInputWithDefaults instantiates a new IPAssignmentUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *IPAssignmentUpdateInput) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IPAssignmentUpdateInput) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IPAssignmentUpdateInput) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IPAssignmentUpdateInput) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCustomdata

`func (o *IPAssignmentUpdateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *IPAssignmentUpdateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *IPAssignmentUpdateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *IPAssignmentUpdateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetTags

`func (o *IPAssignmentUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAssignmentUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAssignmentUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAssignmentUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



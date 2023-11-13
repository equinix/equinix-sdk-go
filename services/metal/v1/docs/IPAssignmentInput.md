# IPAssignmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIPAssignmentInput

`func NewIPAssignmentInput(address string, ) *IPAssignmentInput`

NewIPAssignmentInput instantiates a new IPAssignmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAssignmentInputWithDefaults

`func NewIPAssignmentInputWithDefaults() *IPAssignmentInput`

NewIPAssignmentInputWithDefaults instantiates a new IPAssignmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPAssignmentInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAssignmentInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAssignmentInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCustomdata

`func (o *IPAssignmentInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *IPAssignmentInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *IPAssignmentInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *IPAssignmentInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



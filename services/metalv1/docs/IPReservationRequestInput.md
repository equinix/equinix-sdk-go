# IPReservationRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **string** |  | [optional] 
**FailOnApprovalRequired** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to **string** | The code of the metro you are requesting the IP reservation in. | [optional] 
**Quantity** | **int32** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewIPReservationRequestInput

`func NewIPReservationRequestInput(quantity int32, type_ string, ) *IPReservationRequestInput`

NewIPReservationRequestInput instantiates a new IPReservationRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPReservationRequestInputWithDefaults

`func NewIPReservationRequestInputWithDefaults() *IPReservationRequestInput`

NewIPReservationRequestInputWithDefaults instantiates a new IPReservationRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *IPReservationRequestInput) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IPReservationRequestInput) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IPReservationRequestInput) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IPReservationRequestInput) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomdata

`func (o *IPReservationRequestInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *IPReservationRequestInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *IPReservationRequestInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *IPReservationRequestInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDetails

`func (o *IPReservationRequestInput) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IPReservationRequestInput) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IPReservationRequestInput) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IPReservationRequestInput) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *IPReservationRequestInput) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *IPReservationRequestInput) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *IPReservationRequestInput) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *IPReservationRequestInput) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetFailOnApprovalRequired

`func (o *IPReservationRequestInput) GetFailOnApprovalRequired() bool`

GetFailOnApprovalRequired returns the FailOnApprovalRequired field if non-nil, zero value otherwise.

### GetFailOnApprovalRequiredOk

`func (o *IPReservationRequestInput) GetFailOnApprovalRequiredOk() (*bool, bool)`

GetFailOnApprovalRequiredOk returns a tuple with the FailOnApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnApprovalRequired

`func (o *IPReservationRequestInput) SetFailOnApprovalRequired(v bool)`

SetFailOnApprovalRequired sets FailOnApprovalRequired field to given value.

### HasFailOnApprovalRequired

`func (o *IPReservationRequestInput) HasFailOnApprovalRequired() bool`

HasFailOnApprovalRequired returns a boolean if a field has been set.

### GetMetro

`func (o *IPReservationRequestInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *IPReservationRequestInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *IPReservationRequestInput) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *IPReservationRequestInput) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetQuantity

`func (o *IPReservationRequestInput) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IPReservationRequestInput) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IPReservationRequestInput) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTags

`func (o *IPReservationRequestInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPReservationRequestInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPReservationRequestInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPReservationRequestInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *IPReservationRequestInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPReservationRequestInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPReservationRequestInput) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



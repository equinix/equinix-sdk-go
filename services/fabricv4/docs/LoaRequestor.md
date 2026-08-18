# LoaRequestor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address of the requestor contact. | [optional] 
**OrgId** | Pointer to **int32** | Organization ID of the requestor. | [optional] 
**OrgName** | Pointer to **string** | Organization name of the requestor. | [optional] 

## Methods

### NewLoaRequestor

`func NewLoaRequestor() *LoaRequestor`

NewLoaRequestor instantiates a new LoaRequestor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaRequestorWithDefaults

`func NewLoaRequestorWithDefaults() *LoaRequestor`

NewLoaRequestorWithDefaults instantiates a new LoaRequestor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoaRequestor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoaRequestor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoaRequestor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoaRequestor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgId

`func (o *LoaRequestor) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LoaRequestor) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LoaRequestor) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *LoaRequestor) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *LoaRequestor) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *LoaRequestor) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *LoaRequestor) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *LoaRequestor) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



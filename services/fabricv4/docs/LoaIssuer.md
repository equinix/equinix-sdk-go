# LoaIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address of the issuer contact. | [optional] 
**OrgId** | Pointer to **int32** | Organization ID of the issuer. | [optional] 
**OrgName** | Pointer to **string** | Organization name of the issuer. | [optional] 

## Methods

### NewLoaIssuer

`func NewLoaIssuer() *LoaIssuer`

NewLoaIssuer instantiates a new LoaIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaIssuerWithDefaults

`func NewLoaIssuerWithDefaults() *LoaIssuer`

NewLoaIssuerWithDefaults instantiates a new LoaIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoaIssuer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoaIssuer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoaIssuer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoaIssuer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgId

`func (o *LoaIssuer) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LoaIssuer) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LoaIssuer) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *LoaIssuer) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *LoaIssuer) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *LoaIssuer) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *LoaIssuer) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *LoaIssuer) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



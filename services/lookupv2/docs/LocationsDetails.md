# LocationsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibx** | Pointer to **string** | IBX locations code | [optional] 
**AccessRestricted** | Pointer to **bool** | IBX access is restricted when the value is &#x60;true&#x60; | [optional] [default to false]
**SpecialPrivilege** | Pointer to **bool** | When set to &#x60;true&#x60;, user will be allowed to access IBXs with restricted access. This only applies for specific partners | [optional] [default to false]
**Accounts** | Pointer to [**[]Accounts**](Accounts.md) |  | [optional] 
**Cages** | Pointer to [**[]CageDetails**](CageDetails.md) |  | [optional] 

## Methods

### NewLocationsDetails

`func NewLocationsDetails() *LocationsDetails`

NewLocationsDetails instantiates a new LocationsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationsDetailsWithDefaults

`func NewLocationsDetailsWithDefaults() *LocationsDetails`

NewLocationsDetailsWithDefaults instantiates a new LocationsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbx

`func (o *LocationsDetails) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *LocationsDetails) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *LocationsDetails) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *LocationsDetails) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetAccessRestricted

`func (o *LocationsDetails) GetAccessRestricted() bool`

GetAccessRestricted returns the AccessRestricted field if non-nil, zero value otherwise.

### GetAccessRestrictedOk

`func (o *LocationsDetails) GetAccessRestrictedOk() (*bool, bool)`

GetAccessRestrictedOk returns a tuple with the AccessRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRestricted

`func (o *LocationsDetails) SetAccessRestricted(v bool)`

SetAccessRestricted sets AccessRestricted field to given value.

### HasAccessRestricted

`func (o *LocationsDetails) HasAccessRestricted() bool`

HasAccessRestricted returns a boolean if a field has been set.

### GetSpecialPrivilege

`func (o *LocationsDetails) GetSpecialPrivilege() bool`

GetSpecialPrivilege returns the SpecialPrivilege field if non-nil, zero value otherwise.

### GetSpecialPrivilegeOk

`func (o *LocationsDetails) GetSpecialPrivilegeOk() (*bool, bool)`

GetSpecialPrivilegeOk returns a tuple with the SpecialPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrivilege

`func (o *LocationsDetails) SetSpecialPrivilege(v bool)`

SetSpecialPrivilege sets SpecialPrivilege field to given value.

### HasSpecialPrivilege

`func (o *LocationsDetails) HasSpecialPrivilege() bool`

HasSpecialPrivilege returns a boolean if a field has been set.

### GetAccounts

`func (o *LocationsDetails) GetAccounts() []Accounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *LocationsDetails) GetAccountsOk() (*[]Accounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *LocationsDetails) SetAccounts(v []Accounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *LocationsDetails) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCages

`func (o *LocationsDetails) GetCages() []CageDetails`

GetCages returns the Cages field if non-nil, zero value otherwise.

### GetCagesOk

`func (o *LocationsDetails) GetCagesOk() (*[]CageDetails, bool)`

GetCagesOk returns a tuple with the Cages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCages

`func (o *LocationsDetails) SetCages(v []CageDetails)`

SetCages sets Cages field to given value.

### HasCages

`func (o *LocationsDetails) HasCages() bool`

HasCages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



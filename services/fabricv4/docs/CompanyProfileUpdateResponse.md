# CompanyProfileUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyProfile** | Pointer to [**CompanyProfileResponse**](CompanyProfileResponse.md) |  | [optional] 
**Change** | Pointer to [**CompanyProfileChange**](CompanyProfileChange.md) |  | [optional] 

## Methods

### NewCompanyProfileUpdateResponse

`func NewCompanyProfileUpdateResponse() *CompanyProfileUpdateResponse`

NewCompanyProfileUpdateResponse instantiates a new CompanyProfileUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileUpdateResponseWithDefaults

`func NewCompanyProfileUpdateResponseWithDefaults() *CompanyProfileUpdateResponse`

NewCompanyProfileUpdateResponseWithDefaults instantiates a new CompanyProfileUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyProfile

`func (o *CompanyProfileUpdateResponse) GetCompanyProfile() CompanyProfileResponse`

GetCompanyProfile returns the CompanyProfile field if non-nil, zero value otherwise.

### GetCompanyProfileOk

`func (o *CompanyProfileUpdateResponse) GetCompanyProfileOk() (*CompanyProfileResponse, bool)`

GetCompanyProfileOk returns a tuple with the CompanyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyProfile

`func (o *CompanyProfileUpdateResponse) SetCompanyProfile(v CompanyProfileResponse)`

SetCompanyProfile sets CompanyProfile field to given value.

### HasCompanyProfile

`func (o *CompanyProfileUpdateResponse) HasCompanyProfile() bool`

HasCompanyProfile returns a boolean if a field has been set.

### GetChange

`func (o *CompanyProfileUpdateResponse) GetChange() CompanyProfileChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *CompanyProfileUpdateResponse) GetChangeOk() (*CompanyProfileChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *CompanyProfileUpdateResponse) SetChange(v CompanyProfileChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *CompanyProfileUpdateResponse) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



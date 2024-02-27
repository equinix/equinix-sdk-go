# ConnectionSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceToken** | Pointer to [**ServiceToken**](ServiceToken.md) |  | [optional] 
**AccessPoint** | Pointer to [**AccessPoint**](AccessPoint.md) |  | [optional] 
**CompanyProfile** | Pointer to [**ConnectionCompanyProfile**](ConnectionCompanyProfile.md) |  | [optional] 
**Invitation** | Pointer to [**ConnectionInvitation**](ConnectionInvitation.md) |  | [optional] 
**AdditionalInfo** | Pointer to [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Any additional information, which is not part of connection metadata or configuration | [optional] 

## Methods

### NewConnectionSide

`func NewConnectionSide() *ConnectionSide`

NewConnectionSide instantiates a new ConnectionSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionSideWithDefaults

`func NewConnectionSideWithDefaults() *ConnectionSide`

NewConnectionSideWithDefaults instantiates a new ConnectionSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceToken

`func (o *ConnectionSide) GetServiceToken() ServiceToken`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *ConnectionSide) GetServiceTokenOk() (*ServiceToken, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *ConnectionSide) SetServiceToken(v ServiceToken)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *ConnectionSide) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetAccessPoint

`func (o *ConnectionSide) GetAccessPoint() AccessPoint`

GetAccessPoint returns the AccessPoint field if non-nil, zero value otherwise.

### GetAccessPointOk

`func (o *ConnectionSide) GetAccessPointOk() (*AccessPoint, bool)`

GetAccessPointOk returns a tuple with the AccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoint

`func (o *ConnectionSide) SetAccessPoint(v AccessPoint)`

SetAccessPoint sets AccessPoint field to given value.

### HasAccessPoint

`func (o *ConnectionSide) HasAccessPoint() bool`

HasAccessPoint returns a boolean if a field has been set.

### GetCompanyProfile

`func (o *ConnectionSide) GetCompanyProfile() ConnectionCompanyProfile`

GetCompanyProfile returns the CompanyProfile field if non-nil, zero value otherwise.

### GetCompanyProfileOk

`func (o *ConnectionSide) GetCompanyProfileOk() (*ConnectionCompanyProfile, bool)`

GetCompanyProfileOk returns a tuple with the CompanyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyProfile

`func (o *ConnectionSide) SetCompanyProfile(v ConnectionCompanyProfile)`

SetCompanyProfile sets CompanyProfile field to given value.

### HasCompanyProfile

`func (o *ConnectionSide) HasCompanyProfile() bool`

HasCompanyProfile returns a boolean if a field has been set.

### GetInvitation

`func (o *ConnectionSide) GetInvitation() ConnectionInvitation`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *ConnectionSide) GetInvitationOk() (*ConnectionInvitation, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *ConnectionSide) SetInvitation(v ConnectionInvitation)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *ConnectionSide) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ConnectionSide) GetAdditionalInfo() []ConnectionSideAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ConnectionSide) GetAdditionalInfoOk() (*[]ConnectionSideAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ConnectionSide) SetAdditionalInfo(v []ConnectionSideAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ConnectionSide) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



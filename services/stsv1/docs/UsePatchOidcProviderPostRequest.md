# UsePatchOidcProviderPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-friendly name for the identity provider. | [optional] 
**TrustedClientIds** | Pointer to **[]string** | List of OAuth 2.0 client ids for this provider that are permitted to exchange ID tokens for access tokens. The value of the &#x60;aud&#x60; claim in an ID token is checked against this list during token exchange. | [optional] 
**GroupMembershipClaim** | Pointer to [**UsePatchOidcProviderPostRequestGroupMembershipClaim**](UsePatchOidcProviderPostRequestGroupMembershipClaim.md) |  | [optional] 
**IdpId** | **string** | Uniquely identifies a trusted Identity Provider within a root project. | 
**LastRev** | **string** | An opaque string that represents the expected revision of a given resource. This is provided when a resource is   updatd so that if a concurrent update has occured since the resource was read, then the collision will be detected. | 

## Methods

### NewUsePatchOidcProviderPostRequest

`func NewUsePatchOidcProviderPostRequest(idpId string, lastRev string, ) *UsePatchOidcProviderPostRequest`

NewUsePatchOidcProviderPostRequest instantiates a new UsePatchOidcProviderPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsePatchOidcProviderPostRequestWithDefaults

`func NewUsePatchOidcProviderPostRequestWithDefaults() *UsePatchOidcProviderPostRequest`

NewUsePatchOidcProviderPostRequestWithDefaults instantiates a new UsePatchOidcProviderPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsePatchOidcProviderPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsePatchOidcProviderPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsePatchOidcProviderPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsePatchOidcProviderPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrustedClientIds

`func (o *UsePatchOidcProviderPostRequest) GetTrustedClientIds() []string`

GetTrustedClientIds returns the TrustedClientIds field if non-nil, zero value otherwise.

### GetTrustedClientIdsOk

`func (o *UsePatchOidcProviderPostRequest) GetTrustedClientIdsOk() (*[]string, bool)`

GetTrustedClientIdsOk returns a tuple with the TrustedClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedClientIds

`func (o *UsePatchOidcProviderPostRequest) SetTrustedClientIds(v []string)`

SetTrustedClientIds sets TrustedClientIds field to given value.

### HasTrustedClientIds

`func (o *UsePatchOidcProviderPostRequest) HasTrustedClientIds() bool`

HasTrustedClientIds returns a boolean if a field has been set.

### GetGroupMembershipClaim

`func (o *UsePatchOidcProviderPostRequest) GetGroupMembershipClaim() UsePatchOidcProviderPostRequestGroupMembershipClaim`

GetGroupMembershipClaim returns the GroupMembershipClaim field if non-nil, zero value otherwise.

### GetGroupMembershipClaimOk

`func (o *UsePatchOidcProviderPostRequest) GetGroupMembershipClaimOk() (*UsePatchOidcProviderPostRequestGroupMembershipClaim, bool)`

GetGroupMembershipClaimOk returns a tuple with the GroupMembershipClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipClaim

`func (o *UsePatchOidcProviderPostRequest) SetGroupMembershipClaim(v UsePatchOidcProviderPostRequestGroupMembershipClaim)`

SetGroupMembershipClaim sets GroupMembershipClaim field to given value.

### HasGroupMembershipClaim

`func (o *UsePatchOidcProviderPostRequest) HasGroupMembershipClaim() bool`

HasGroupMembershipClaim returns a boolean if a field has been set.

### GetIdpId

`func (o *UsePatchOidcProviderPostRequest) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *UsePatchOidcProviderPostRequest) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *UsePatchOidcProviderPostRequest) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetLastRev

`func (o *UsePatchOidcProviderPostRequest) GetLastRev() string`

GetLastRev returns the LastRev field if non-nil, zero value otherwise.

### GetLastRevOk

`func (o *UsePatchOidcProviderPostRequest) GetLastRevOk() (*string, bool)`

GetLastRevOk returns a tuple with the LastRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRev

`func (o *UsePatchOidcProviderPostRequest) SetLastRev(v string)`

SetLastRev sets LastRev field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



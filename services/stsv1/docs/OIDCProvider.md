# OIDCProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpId** | **string** | Uniquely identifies a trusted Identity Provider within a root project. | 
**UpdatedAt** | Pointer to **string** | A string timestamp indicating when the resource was last updated. Formatted like: \&quot;2025-02-12T17:24:19.033772087Z\&quot; | [optional] 
**UpdatedBy** | Pointer to **string** | A string indicating the principal who last invoked an operation to update the resource. | [optional] 
**CreatedAt** | **string** | A string timestamp indicating when the resource was created. Formatted like: \&quot;2025-02-12T17:24:19.033772087Z\&quot; | 
**Rev** | **string** | An opaque string that represents the revision of a given resource. Each time the resource is updated, this value   changes. | 
**CreatedBy** | **string** | A string indicating the principal who invoked an operation to create the resource. | 
**IssuerLocation** | **string** | The OIDC issuer location URL. | 
**Name** | **string** | A human-friendly name for the identity provider. | 
**JwksRetrievedAt** | **string** | Timestamp string formatted like: \&quot;2025-02-13T17:10:00.864707507Z\&quot;. | 
**Jwks** | [**OIDCProviderJwks**](OIDCProviderJwks.md) |  | 
**Status** | [**ProviderStatus**](ProviderStatus.md) |  | 
**IssuerUri** | **string** | An OIDC Issuer URI | 
**TrustedClientIds** | **[]string** | List of OAuth 2.0 client ids for this provider that are permitted to exchange ID tokens for access tokens. The value of the &#x60;aud&#x60; claim in an ID token is checked against this list during token exchange. | 
**GroupMembershipClaim** | Pointer to **string** | Name of the claim in the ID tokens provided by this OIDC issuer whose value the STS should interpret as containing a user&#39;s group memberships, for authorization purposes. The value of the group membership claim in an ID token must be an array of strings, where each string is a unique, non-reassignable identifier for a group. When this property is not set, the STS does not interpret any claim from this provider as a group membership claim. | [optional] 

## Methods

### NewOIDCProvider

`func NewOIDCProvider(idpId string, createdAt string, rev string, createdBy string, issuerLocation string, name string, jwksRetrievedAt string, jwks OIDCProviderJwks, status ProviderStatus, issuerUri string, trustedClientIds []string, ) *OIDCProvider`

NewOIDCProvider instantiates a new OIDCProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCProviderWithDefaults

`func NewOIDCProviderWithDefaults() *OIDCProvider`

NewOIDCProviderWithDefaults instantiates a new OIDCProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpId

`func (o *OIDCProvider) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *OIDCProvider) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *OIDCProvider) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetUpdatedAt

`func (o *OIDCProvider) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OIDCProvider) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OIDCProvider) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OIDCProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *OIDCProvider) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OIDCProvider) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OIDCProvider) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *OIDCProvider) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OIDCProvider) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OIDCProvider) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OIDCProvider) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetRev

`func (o *OIDCProvider) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *OIDCProvider) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *OIDCProvider) SetRev(v string)`

SetRev sets Rev field to given value.


### GetCreatedBy

`func (o *OIDCProvider) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OIDCProvider) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OIDCProvider) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetIssuerLocation

`func (o *OIDCProvider) GetIssuerLocation() string`

GetIssuerLocation returns the IssuerLocation field if non-nil, zero value otherwise.

### GetIssuerLocationOk

`func (o *OIDCProvider) GetIssuerLocationOk() (*string, bool)`

GetIssuerLocationOk returns a tuple with the IssuerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerLocation

`func (o *OIDCProvider) SetIssuerLocation(v string)`

SetIssuerLocation sets IssuerLocation field to given value.


### GetName

`func (o *OIDCProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OIDCProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OIDCProvider) SetName(v string)`

SetName sets Name field to given value.


### GetJwksRetrievedAt

`func (o *OIDCProvider) GetJwksRetrievedAt() string`

GetJwksRetrievedAt returns the JwksRetrievedAt field if non-nil, zero value otherwise.

### GetJwksRetrievedAtOk

`func (o *OIDCProvider) GetJwksRetrievedAtOk() (*string, bool)`

GetJwksRetrievedAtOk returns a tuple with the JwksRetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksRetrievedAt

`func (o *OIDCProvider) SetJwksRetrievedAt(v string)`

SetJwksRetrievedAt sets JwksRetrievedAt field to given value.


### GetJwks

`func (o *OIDCProvider) GetJwks() OIDCProviderJwks`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *OIDCProvider) GetJwksOk() (*OIDCProviderJwks, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *OIDCProvider) SetJwks(v OIDCProviderJwks)`

SetJwks sets Jwks field to given value.


### GetStatus

`func (o *OIDCProvider) GetStatus() ProviderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OIDCProvider) GetStatusOk() (*ProviderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OIDCProvider) SetStatus(v ProviderStatus)`

SetStatus sets Status field to given value.


### GetIssuerUri

`func (o *OIDCProvider) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *OIDCProvider) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *OIDCProvider) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.


### GetTrustedClientIds

`func (o *OIDCProvider) GetTrustedClientIds() []string`

GetTrustedClientIds returns the TrustedClientIds field if non-nil, zero value otherwise.

### GetTrustedClientIdsOk

`func (o *OIDCProvider) GetTrustedClientIdsOk() (*[]string, bool)`

GetTrustedClientIdsOk returns a tuple with the TrustedClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedClientIds

`func (o *OIDCProvider) SetTrustedClientIds(v []string)`

SetTrustedClientIds sets TrustedClientIds field to given value.


### GetGroupMembershipClaim

`func (o *OIDCProvider) GetGroupMembershipClaim() string`

GetGroupMembershipClaim returns the GroupMembershipClaim field if non-nil, zero value otherwise.

### GetGroupMembershipClaimOk

`func (o *OIDCProvider) GetGroupMembershipClaimOk() (*string, bool)`

GetGroupMembershipClaimOk returns a tuple with the GroupMembershipClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipClaim

`func (o *OIDCProvider) SetGroupMembershipClaim(v string)`

SetGroupMembershipClaim sets GroupMembershipClaim field to given value.

### HasGroupMembershipClaim

`func (o *OIDCProvider) HasGroupMembershipClaim() bool`

HasGroupMembershipClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



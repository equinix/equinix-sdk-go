# OpsCreateOidcProviderPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the identity provider. | 
**TrustedClientIds** | **[]string** | List of OAuth 2.0 client ids for this provider that are permitted to exchange ID tokens for access tokens. The value of the &#x60;aud&#x60; claim in an ID token is checked against this list during token exchange. | 
**GroupMembershipClaim** | Pointer to **string** | Name of the claim in the ID tokens provided by this OIDC issuer whose value the STS should interpret as containing a user&#39;s group memberships, for authorization purposes. The value of the group membership claim in an ID token must be an array of strings, where each string is a unique, non-reassignable identifier for a group. When this property is not set, the STS does not interpret any claim from this provider as a group membership claim. | [optional] 
**IssuerUri** | **string** | An OIDC Issuer URI | 
**IssuerLocation** | **string** | The OIDC issuer location URL. | 
**Jwks** | [**OpsCreateOidcProviderPostRequestJwks**](OpsCreateOidcProviderPostRequestJwks.md) |  | 
**IdpPrefix** | **string** |  | 
**ProjectId** | **string** | Globally unique identifier of a project. | 

## Methods

### NewOpsCreateOidcProviderPostRequest

`func NewOpsCreateOidcProviderPostRequest(name string, trustedClientIds []string, issuerUri string, issuerLocation string, jwks OpsCreateOidcProviderPostRequestJwks, idpPrefix string, projectId string, ) *OpsCreateOidcProviderPostRequest`

NewOpsCreateOidcProviderPostRequest instantiates a new OpsCreateOidcProviderPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsCreateOidcProviderPostRequestWithDefaults

`func NewOpsCreateOidcProviderPostRequestWithDefaults() *OpsCreateOidcProviderPostRequest`

NewOpsCreateOidcProviderPostRequestWithDefaults instantiates a new OpsCreateOidcProviderPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpsCreateOidcProviderPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpsCreateOidcProviderPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpsCreateOidcProviderPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTrustedClientIds

`func (o *OpsCreateOidcProviderPostRequest) GetTrustedClientIds() []string`

GetTrustedClientIds returns the TrustedClientIds field if non-nil, zero value otherwise.

### GetTrustedClientIdsOk

`func (o *OpsCreateOidcProviderPostRequest) GetTrustedClientIdsOk() (*[]string, bool)`

GetTrustedClientIdsOk returns a tuple with the TrustedClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedClientIds

`func (o *OpsCreateOidcProviderPostRequest) SetTrustedClientIds(v []string)`

SetTrustedClientIds sets TrustedClientIds field to given value.


### GetGroupMembershipClaim

`func (o *OpsCreateOidcProviderPostRequest) GetGroupMembershipClaim() string`

GetGroupMembershipClaim returns the GroupMembershipClaim field if non-nil, zero value otherwise.

### GetGroupMembershipClaimOk

`func (o *OpsCreateOidcProviderPostRequest) GetGroupMembershipClaimOk() (*string, bool)`

GetGroupMembershipClaimOk returns a tuple with the GroupMembershipClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipClaim

`func (o *OpsCreateOidcProviderPostRequest) SetGroupMembershipClaim(v string)`

SetGroupMembershipClaim sets GroupMembershipClaim field to given value.

### HasGroupMembershipClaim

`func (o *OpsCreateOidcProviderPostRequest) HasGroupMembershipClaim() bool`

HasGroupMembershipClaim returns a boolean if a field has been set.

### GetIssuerUri

`func (o *OpsCreateOidcProviderPostRequest) GetIssuerUri() string`

GetIssuerUri returns the IssuerUri field if non-nil, zero value otherwise.

### GetIssuerUriOk

`func (o *OpsCreateOidcProviderPostRequest) GetIssuerUriOk() (*string, bool)`

GetIssuerUriOk returns a tuple with the IssuerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUri

`func (o *OpsCreateOidcProviderPostRequest) SetIssuerUri(v string)`

SetIssuerUri sets IssuerUri field to given value.


### GetIssuerLocation

`func (o *OpsCreateOidcProviderPostRequest) GetIssuerLocation() string`

GetIssuerLocation returns the IssuerLocation field if non-nil, zero value otherwise.

### GetIssuerLocationOk

`func (o *OpsCreateOidcProviderPostRequest) GetIssuerLocationOk() (*string, bool)`

GetIssuerLocationOk returns a tuple with the IssuerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerLocation

`func (o *OpsCreateOidcProviderPostRequest) SetIssuerLocation(v string)`

SetIssuerLocation sets IssuerLocation field to given value.


### GetJwks

`func (o *OpsCreateOidcProviderPostRequest) GetJwks() OpsCreateOidcProviderPostRequestJwks`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *OpsCreateOidcProviderPostRequest) GetJwksOk() (*OpsCreateOidcProviderPostRequestJwks, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *OpsCreateOidcProviderPostRequest) SetJwks(v OpsCreateOidcProviderPostRequestJwks)`

SetJwks sets Jwks field to given value.


### GetIdpPrefix

`func (o *OpsCreateOidcProviderPostRequest) GetIdpPrefix() string`

GetIdpPrefix returns the IdpPrefix field if non-nil, zero value otherwise.

### GetIdpPrefixOk

`func (o *OpsCreateOidcProviderPostRequest) GetIdpPrefixOk() (*string, bool)`

GetIdpPrefixOk returns a tuple with the IdpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpPrefix

`func (o *OpsCreateOidcProviderPostRequest) SetIdpPrefix(v string)`

SetIdpPrefix sets IdpPrefix field to given value.


### GetProjectId

`func (o *OpsCreateOidcProviderPostRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpsCreateOidcProviderPostRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpsCreateOidcProviderPostRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



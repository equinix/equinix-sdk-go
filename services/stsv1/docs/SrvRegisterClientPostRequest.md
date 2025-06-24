# SrvRegisterClientPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenEndpointAuthMethod** | **interface{}** |  | [default to client_secret_basic]
**GrantTypes** | **[]interface{}** |  | 
**ResponseTypes** | **[]interface{}** |  | 
**ServiceAccessPolicyId** | **string** | Uniquely identifies an access policy within a project. | 
**ServiceId** | Pointer to **string** | Fully qualified, universally unique id of a service. Starts with the NamespaceId. Formatted like   \&quot;service:&amp;lt;namespace&amp;gt;/&amp;lt;service&amp;gt;\&quot;. | [optional] 
**ServiceErn** | Pointer to **string** | Equinix resource name, a universally unique identifier for a resource across all clouds, regions, and   services. Formatted as   \&quot;ern:&amp;lt;cloudId&amp;gt;:&amp;lt;serviceId&amp;gt;:&amp;lt;regionId&amp;gt;:&amp;lt;projectId&amp;gt;:&amp;lt;resourceType&amp;gt;:&amp;lt;resourceId&amp;gt;\&quot;. | [optional] 

## Methods

### NewSrvRegisterClientPostRequest

`func NewSrvRegisterClientPostRequest(tokenEndpointAuthMethod interface{}, grantTypes []interface{}, responseTypes []interface{}, serviceAccessPolicyId string, ) *SrvRegisterClientPostRequest`

NewSrvRegisterClientPostRequest instantiates a new SrvRegisterClientPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrvRegisterClientPostRequestWithDefaults

`func NewSrvRegisterClientPostRequestWithDefaults() *SrvRegisterClientPostRequest`

NewSrvRegisterClientPostRequestWithDefaults instantiates a new SrvRegisterClientPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenEndpointAuthMethod

`func (o *SrvRegisterClientPostRequest) GetTokenEndpointAuthMethod() interface{}`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *SrvRegisterClientPostRequest) GetTokenEndpointAuthMethodOk() (*interface{}, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *SrvRegisterClientPostRequest) SetTokenEndpointAuthMethod(v interface{})`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### SetTokenEndpointAuthMethodNil

`func (o *SrvRegisterClientPostRequest) SetTokenEndpointAuthMethodNil(b bool)`

 SetTokenEndpointAuthMethodNil sets the value for TokenEndpointAuthMethod to be an explicit nil

### UnsetTokenEndpointAuthMethod
`func (o *SrvRegisterClientPostRequest) UnsetTokenEndpointAuthMethod()`

UnsetTokenEndpointAuthMethod ensures that no value is present for TokenEndpointAuthMethod, not even an explicit nil
### GetGrantTypes

`func (o *SrvRegisterClientPostRequest) GetGrantTypes() []interface{}`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *SrvRegisterClientPostRequest) GetGrantTypesOk() (*[]interface{}, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *SrvRegisterClientPostRequest) SetGrantTypes(v []interface{})`

SetGrantTypes sets GrantTypes field to given value.


### GetResponseTypes

`func (o *SrvRegisterClientPostRequest) GetResponseTypes() []interface{}`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *SrvRegisterClientPostRequest) GetResponseTypesOk() (*[]interface{}, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *SrvRegisterClientPostRequest) SetResponseTypes(v []interface{})`

SetResponseTypes sets ResponseTypes field to given value.


### GetServiceAccessPolicyId

`func (o *SrvRegisterClientPostRequest) GetServiceAccessPolicyId() string`

GetServiceAccessPolicyId returns the ServiceAccessPolicyId field if non-nil, zero value otherwise.

### GetServiceAccessPolicyIdOk

`func (o *SrvRegisterClientPostRequest) GetServiceAccessPolicyIdOk() (*string, bool)`

GetServiceAccessPolicyIdOk returns a tuple with the ServiceAccessPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessPolicyId

`func (o *SrvRegisterClientPostRequest) SetServiceAccessPolicyId(v string)`

SetServiceAccessPolicyId sets ServiceAccessPolicyId field to given value.


### GetServiceId

`func (o *SrvRegisterClientPostRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SrvRegisterClientPostRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SrvRegisterClientPostRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SrvRegisterClientPostRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceErn

`func (o *SrvRegisterClientPostRequest) GetServiceErn() string`

GetServiceErn returns the ServiceErn field if non-nil, zero value otherwise.

### GetServiceErnOk

`func (o *SrvRegisterClientPostRequest) GetServiceErnOk() (*string, bool)`

GetServiceErnOk returns a tuple with the ServiceErn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceErn

`func (o *SrvRegisterClientPostRequest) SetServiceErn(v string)`

SetServiceErn sets ServiceErn field to given value.

### HasServiceErn

`func (o *SrvRegisterClientPostRequest) HasServiceErn() bool`

HasServiceErn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



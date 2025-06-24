# ClientRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Fully qualified, universally unique id of a service. Starts with the NamespaceId. Formatted like   \&quot;service:&amp;lt;namespace&amp;gt;/&amp;lt;service&amp;gt;\&quot;. | 
**ServiceAccessPolicyId** | **string** | Uniquely identifies an access policy within a project. | 
**CreatedBy** | **string** | A string indicating the principal who invoked an operation to create the resource. | 
**CreatedAt** | **string** | A string timestamp indicating when the resource was created. Formatted like: \&quot;2025-02-12T17:24:19.033772087Z\&quot; | 
**UpdatedBy** | Pointer to **string** | A string indicating the principal who last invoked an operation to update the resource. | [optional] 
**UpdatedAt** | Pointer to **string** | A string timestamp indicating when the resource was last updated. Formatted like: \&quot;2025-02-12T17:24:19.033772087Z\&quot; | [optional] 
**ClientId** | **string** | An OAuth 2.0 client id for an EaaS service. | 

## Methods

### NewClientRegistration

`func NewClientRegistration(serviceId string, serviceAccessPolicyId string, createdBy string, createdAt string, clientId string, ) *ClientRegistration`

NewClientRegistration instantiates a new ClientRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRegistrationWithDefaults

`func NewClientRegistrationWithDefaults() *ClientRegistration`

NewClientRegistrationWithDefaults instantiates a new ClientRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ClientRegistration) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ClientRegistration) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ClientRegistration) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceAccessPolicyId

`func (o *ClientRegistration) GetServiceAccessPolicyId() string`

GetServiceAccessPolicyId returns the ServiceAccessPolicyId field if non-nil, zero value otherwise.

### GetServiceAccessPolicyIdOk

`func (o *ClientRegistration) GetServiceAccessPolicyIdOk() (*string, bool)`

GetServiceAccessPolicyIdOk returns a tuple with the ServiceAccessPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessPolicyId

`func (o *ClientRegistration) SetServiceAccessPolicyId(v string)`

SetServiceAccessPolicyId sets ServiceAccessPolicyId field to given value.


### GetCreatedBy

`func (o *ClientRegistration) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClientRegistration) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClientRegistration) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *ClientRegistration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClientRegistration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClientRegistration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedBy

`func (o *ClientRegistration) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClientRegistration) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClientRegistration) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ClientRegistration) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClientRegistration) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClientRegistration) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClientRegistration) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClientRegistration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClientId

`func (o *ClientRegistration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientRegistration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientRegistration) SetClientId(v string)`

SetClientId sets ClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



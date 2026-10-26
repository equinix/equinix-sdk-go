# CrossConnectLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Sub-resource URI. | [optional] [readonly] 
**Rel** | Pointer to **string** | Operation identifier of the linked sub-resource. | [optional] [readonly] 
**Method** | Pointer to **string** | HTTP method to use on the sub-resource. | [optional] [readonly] 
**ContentType** | Pointer to **string** | Content type returned by the sub-resource. | [optional] [readonly] 
**Authenticate** | Pointer to **bool** | Whether authentication is required. | [optional] [readonly] 

## Methods

### NewCrossConnectLink

`func NewCrossConnectLink() *CrossConnectLink`

NewCrossConnectLink instantiates a new CrossConnectLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectLinkWithDefaults

`func NewCrossConnectLinkWithDefaults() *CrossConnectLink`

NewCrossConnectLinkWithDefaults instantiates a new CrossConnectLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetRel

`func (o *CrossConnectLink) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *CrossConnectLink) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *CrossConnectLink) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *CrossConnectLink) HasRel() bool`

HasRel returns a boolean if a field has been set.

### GetMethod

`func (o *CrossConnectLink) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CrossConnectLink) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CrossConnectLink) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CrossConnectLink) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetContentType

`func (o *CrossConnectLink) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CrossConnectLink) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CrossConnectLink) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CrossConnectLink) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetAuthenticate

`func (o *CrossConnectLink) GetAuthenticate() bool`

GetAuthenticate returns the Authenticate field if non-nil, zero value otherwise.

### GetAuthenticateOk

`func (o *CrossConnectLink) GetAuthenticateOk() (*bool, bool)`

GetAuthenticateOk returns a tuple with the Authenticate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticate

`func (o *CrossConnectLink) SetAuthenticate(v bool)`

SetAuthenticate sets Authenticate field to given value.

### HasAuthenticate

`func (o *CrossConnectLink) HasAuthenticate() bool`

HasAuthenticate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Rel** | Pointer to **string** | OperationId from Swagger hub spec | [optional] 
**Method** | Pointer to **string** | Http method type | [optional] 
**ContentType** | Pointer to **string** | Content type for the response | [optional] 
**Authenticate** | Pointer to **bool** | Authentication required or not | [optional] 

## Methods

### NewLink

`func NewLink() *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Link) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Link) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Link) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Link) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetRel

`func (o *Link) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *Link) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *Link) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *Link) HasRel() bool`

HasRel returns a boolean if a field has been set.

### GetMethod

`func (o *Link) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Link) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Link) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Link) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetContentType

`func (o *Link) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Link) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Link) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Link) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetAuthenticate

`func (o *Link) GetAuthenticate() bool`

GetAuthenticate returns the Authenticate field if non-nil, zero value otherwise.

### GetAuthenticateOk

`func (o *Link) GetAuthenticateOk() (*bool, bool)`

GetAuthenticateOk returns a tuple with the Authenticate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticate

`func (o *Link) SetAuthenticate(v bool)`

SetAuthenticate sets Authenticate field to given value.

### HasAuthenticate

`func (o *Link) HasAuthenticate() bool`

HasAuthenticate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



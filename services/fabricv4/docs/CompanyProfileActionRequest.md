# CompanyProfileActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | CompanyProfile Action Type | 
**Comments** | Pointer to **string** | Optional comments for the action | [optional] 

## Methods

### NewCompanyProfileActionRequest

`func NewCompanyProfileActionRequest(type_ string, ) *CompanyProfileActionRequest`

NewCompanyProfileActionRequest instantiates a new CompanyProfileActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileActionRequestWithDefaults

`func NewCompanyProfileActionRequestWithDefaults() *CompanyProfileActionRequest`

NewCompanyProfileActionRequestWithDefaults instantiates a new CompanyProfileActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompanyProfileActionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyProfileActionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyProfileActionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetComments

`func (o *CompanyProfileActionRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CompanyProfileActionRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CompanyProfileActionRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CompanyProfileActionRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



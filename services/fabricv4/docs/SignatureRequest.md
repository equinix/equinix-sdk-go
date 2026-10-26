# SignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SignatureType**](SignatureType.md) |  | 
**Email** | Pointer to **string** | Signatory email address. | [optional] 

## Methods

### NewSignatureRequest

`func NewSignatureRequest(type_ SignatureType, ) *SignatureRequest`

NewSignatureRequest instantiates a new SignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestWithDefaults

`func NewSignatureRequestWithDefaults() *SignatureRequest`

NewSignatureRequestWithDefaults instantiates a new SignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SignatureRequest) GetType() SignatureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignatureRequest) GetTypeOk() (*SignatureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignatureRequest) SetType(v SignatureType)`

SetType sets Type field to given value.


### GetEmail

`func (o *SignatureRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignatureRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignatureRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignatureRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



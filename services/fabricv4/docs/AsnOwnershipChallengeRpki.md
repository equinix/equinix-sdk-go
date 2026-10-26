# AsnOwnershipChallengeRpki

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Challenge message to sign | [optional] 
**Signature** | Pointer to **string** | Signed challenge value | [optional] 

## Methods

### NewAsnOwnershipChallengeRpki

`func NewAsnOwnershipChallengeRpki() *AsnOwnershipChallengeRpki`

NewAsnOwnershipChallengeRpki instantiates a new AsnOwnershipChallengeRpki object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnOwnershipChallengeRpkiWithDefaults

`func NewAsnOwnershipChallengeRpkiWithDefaults() *AsnOwnershipChallengeRpki`

NewAsnOwnershipChallengeRpkiWithDefaults instantiates a new AsnOwnershipChallengeRpki object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AsnOwnershipChallengeRpki) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AsnOwnershipChallengeRpki) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AsnOwnershipChallengeRpki) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AsnOwnershipChallengeRpki) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSignature

`func (o *AsnOwnershipChallengeRpki) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AsnOwnershipChallengeRpki) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AsnOwnershipChallengeRpki) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AsnOwnershipChallengeRpki) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



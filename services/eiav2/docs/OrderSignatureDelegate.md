# OrderSignatureDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | **string** | Email address that the signature request should be sent to in case of DELEGATE signature  | 

## Methods

### NewOrderSignatureDelegate

`func NewOrderSignatureDelegate(email string, ) *OrderSignatureDelegate`

NewOrderSignatureDelegate instantiates a new OrderSignatureDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSignatureDelegateWithDefaults

`func NewOrderSignatureDelegateWithDefaults() *OrderSignatureDelegate`

NewOrderSignatureDelegateWithDefaults instantiates a new OrderSignatureDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *OrderSignatureDelegate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrderSignatureDelegate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrderSignatureDelegate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrderSignatureDelegate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *OrderSignatureDelegate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrderSignatureDelegate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrderSignatureDelegate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrderSignatureDelegate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *OrderSignatureDelegate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderSignatureDelegate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderSignatureDelegate) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



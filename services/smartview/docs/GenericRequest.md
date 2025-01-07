# GenericRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** | customer account number | [optional] 
**Classification** | Pointer to **string** | asset classification | [optional] 
**Ibx** | Pointer to **string** | ibx code | [optional] 

## Methods

### NewGenericRequest

`func NewGenericRequest() *GenericRequest`

NewGenericRequest instantiates a new GenericRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericRequestWithDefaults

`func NewGenericRequestWithDefaults() *GenericRequest`

NewGenericRequestWithDefaults instantiates a new GenericRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *GenericRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *GenericRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *GenericRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *GenericRequest) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetClassification

`func (o *GenericRequest) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *GenericRequest) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *GenericRequest) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *GenericRequest) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetIbx

`func (o *GenericRequest) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *GenericRequest) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *GenericRequest) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *GenericRequest) HasIbx() bool`

HasIbx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



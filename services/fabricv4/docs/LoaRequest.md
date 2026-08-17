# LoaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LoaType**](LoaType.md) |  | 
**Name** | **string** | A short, descriptive name for this LOA. | 
**Description** | Pointer to **string** | Additional context about this LOA. | [optional] 
**AuthorizedProductType** | [**LoaProductType**](LoaProductType.md) |  | 

## Methods

### NewLoaRequest

`func NewLoaRequest(type_ LoaType, name string, authorizedProductType LoaProductType, ) *LoaRequest`

NewLoaRequest instantiates a new LoaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaRequestWithDefaults

`func NewLoaRequestWithDefaults() *LoaRequest`

NewLoaRequestWithDefaults instantiates a new LoaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoaRequest) GetType() LoaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoaRequest) GetTypeOk() (*LoaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoaRequest) SetType(v LoaType)`

SetType sets Type field to given value.


### GetName

`func (o *LoaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LoaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthorizedProductType

`func (o *LoaRequest) GetAuthorizedProductType() LoaProductType`

GetAuthorizedProductType returns the AuthorizedProductType field if non-nil, zero value otherwise.

### GetAuthorizedProductTypeOk

`func (o *LoaRequest) GetAuthorizedProductTypeOk() (*LoaProductType, bool)`

GetAuthorizedProductTypeOk returns a tuple with the AuthorizedProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedProductType

`func (o *LoaRequest) SetAuthorizedProductType(v LoaProductType)`

SetAuthorizedProductType sets AuthorizedProductType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



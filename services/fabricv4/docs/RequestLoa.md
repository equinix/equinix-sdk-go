# RequestLoa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LoaType**](LoaType.md) |  | 
**Name** | **string** | A short, descriptive name for this LOA. | 
**Description** | Pointer to **string** | Additional context about this LOA. | [optional] 
**AuthorizedProductType** | [**LoaProductType**](LoaProductType.md) |  | 
**Issuer** | Pointer to [**LoaIssuer**](LoaIssuer.md) |  | [optional] 
**Location** | [**LoaLocation**](LoaLocation.md) |  | 

## Methods

### NewRequestLoa

`func NewRequestLoa(type_ LoaType, name string, authorizedProductType LoaProductType, location LoaLocation, ) *RequestLoa`

NewRequestLoa instantiates a new RequestLoa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestLoaWithDefaults

`func NewRequestLoaWithDefaults() *RequestLoa`

NewRequestLoaWithDefaults instantiates a new RequestLoa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestLoa) GetType() LoaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestLoa) GetTypeOk() (*LoaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestLoa) SetType(v LoaType)`

SetType sets Type field to given value.


### GetName

`func (o *RequestLoa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestLoa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestLoa) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestLoa) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestLoa) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestLoa) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestLoa) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthorizedProductType

`func (o *RequestLoa) GetAuthorizedProductType() LoaProductType`

GetAuthorizedProductType returns the AuthorizedProductType field if non-nil, zero value otherwise.

### GetAuthorizedProductTypeOk

`func (o *RequestLoa) GetAuthorizedProductTypeOk() (*LoaProductType, bool)`

GetAuthorizedProductTypeOk returns a tuple with the AuthorizedProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedProductType

`func (o *RequestLoa) SetAuthorizedProductType(v LoaProductType)`

SetAuthorizedProductType sets AuthorizedProductType field to given value.


### GetIssuer

`func (o *RequestLoa) GetIssuer() LoaIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *RequestLoa) GetIssuerOk() (*LoaIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *RequestLoa) SetIssuer(v LoaIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *RequestLoa) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLocation

`func (o *RequestLoa) GetLocation() LoaLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RequestLoa) GetLocationOk() (*LoaLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RequestLoa) SetLocation(v LoaLocation)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateLoa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LoaType**](LoaType.md) |  | 
**Name** | **string** | A short, descriptive name for this LOA. | 
**Description** | Pointer to **string** | Additional context about this LOA. | [optional] 
**AuthorizedProductType** | [**LoaProductType**](LoaProductType.md) |  | 
**ExpirationDateTime** | Pointer to **time.Time** | Date and time when this LOA expires.&lt;br&gt; Default to 3 months from the creation date  | [optional] 
**Requestor** | Pointer to [**LoaRequestor**](LoaRequestor.md) |  | [optional] 
**DemarcationPoint** | [**LoaDemarcationPoint**](LoaDemarcationPoint.md) |  | 
**Issuer** | Pointer to [**LoaIssuer**](LoaIssuer.md) |  | [optional] 
**Location** | [**LoaLocation**](LoaLocation.md) |  | 

## Methods

### NewCreateLoa

`func NewCreateLoa(type_ LoaType, name string, authorizedProductType LoaProductType, demarcationPoint LoaDemarcationPoint, location LoaLocation, ) *CreateLoa`

NewCreateLoa instantiates a new CreateLoa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoaWithDefaults

`func NewCreateLoaWithDefaults() *CreateLoa`

NewCreateLoaWithDefaults instantiates a new CreateLoa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateLoa) GetType() LoaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLoa) GetTypeOk() (*LoaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLoa) SetType(v LoaType)`

SetType sets Type field to given value.


### GetName

`func (o *CreateLoa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoa) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateLoa) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoa) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoa) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoa) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthorizedProductType

`func (o *CreateLoa) GetAuthorizedProductType() LoaProductType`

GetAuthorizedProductType returns the AuthorizedProductType field if non-nil, zero value otherwise.

### GetAuthorizedProductTypeOk

`func (o *CreateLoa) GetAuthorizedProductTypeOk() (*LoaProductType, bool)`

GetAuthorizedProductTypeOk returns a tuple with the AuthorizedProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedProductType

`func (o *CreateLoa) SetAuthorizedProductType(v LoaProductType)`

SetAuthorizedProductType sets AuthorizedProductType field to given value.


### GetExpirationDateTime

`func (o *CreateLoa) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *CreateLoa) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *CreateLoa) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *CreateLoa) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetRequestor

`func (o *CreateLoa) GetRequestor() LoaRequestor`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *CreateLoa) GetRequestorOk() (*LoaRequestor, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *CreateLoa) SetRequestor(v LoaRequestor)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *CreateLoa) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *CreateLoa) GetDemarcationPoint() LoaDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *CreateLoa) GetDemarcationPointOk() (*LoaDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *CreateLoa) SetDemarcationPoint(v LoaDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.


### GetIssuer

`func (o *CreateLoa) GetIssuer() LoaIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateLoa) GetIssuerOk() (*LoaIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateLoa) SetIssuer(v LoaIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CreateLoa) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLocation

`func (o *CreateLoa) GetLocation() LoaLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateLoa) GetLocationOk() (*LoaLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateLoa) SetLocation(v LoaLocation)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



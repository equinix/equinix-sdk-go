# LoaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI of this LOA resource. | [optional] 
**Uuid** | Pointer to **string** | Unique identifier of this LOA. | [optional] 
**Type** | Pointer to [**LoaType**](LoaType.md) |  | [optional] 
**Name** | Pointer to **string** | A short, descriptive name for this LOA | [optional] 
**Description** | Pointer to **string** | Additional context about this LOA | [optional] 
**AuthorizedProductType** | Pointer to [**LoaProductType**](LoaProductType.md) |  | [optional] 
**State** | Pointer to [**LoaState**](LoaState.md) |  | [optional] 
**Operation** | Pointer to [**LoaResponseOperation**](LoaResponseOperation.md) |  | [optional] 
**Requestor** | Pointer to [**LoaRequestor**](LoaRequestor.md) |  | [optional] 
**Issuer** | Pointer to [**LoaIssuer**](LoaIssuer.md) |  | [optional] 
**DemarcationPoint** | Pointer to [**LoaDemarcationPoint**](LoaDemarcationPoint.md) |  | [optional] 
**Location** | Pointer to [**LoaLocation**](LoaLocation.md) |  | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Date and time when this LOA expires.&lt;br&gt; Default to 3 months from the creation date  | [optional] 
**ChangeLog** | Pointer to [**LoaChangelog**](LoaChangelog.md) |  | [optional] 

## Methods

### NewLoaResponse

`func NewLoaResponse() *LoaResponse`

NewLoaResponse instantiates a new LoaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaResponseWithDefaults

`func NewLoaResponseWithDefaults() *LoaResponse`

NewLoaResponseWithDefaults instantiates a new LoaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LoaResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LoaResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LoaResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LoaResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *LoaResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LoaResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LoaResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LoaResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *LoaResponse) GetType() LoaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoaResponse) GetTypeOk() (*LoaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoaResponse) SetType(v LoaType)`

SetType sets Type field to given value.

### HasType

`func (o *LoaResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *LoaResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoaResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoaResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoaResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LoaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthorizedProductType

`func (o *LoaResponse) GetAuthorizedProductType() LoaProductType`

GetAuthorizedProductType returns the AuthorizedProductType field if non-nil, zero value otherwise.

### GetAuthorizedProductTypeOk

`func (o *LoaResponse) GetAuthorizedProductTypeOk() (*LoaProductType, bool)`

GetAuthorizedProductTypeOk returns a tuple with the AuthorizedProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedProductType

`func (o *LoaResponse) SetAuthorizedProductType(v LoaProductType)`

SetAuthorizedProductType sets AuthorizedProductType field to given value.

### HasAuthorizedProductType

`func (o *LoaResponse) HasAuthorizedProductType() bool`

HasAuthorizedProductType returns a boolean if a field has been set.

### GetState

`func (o *LoaResponse) GetState() LoaState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoaResponse) GetStateOk() (*LoaState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoaResponse) SetState(v LoaState)`

SetState sets State field to given value.

### HasState

`func (o *LoaResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperation

`func (o *LoaResponse) GetOperation() LoaResponseOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LoaResponse) GetOperationOk() (*LoaResponseOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LoaResponse) SetOperation(v LoaResponseOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *LoaResponse) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRequestor

`func (o *LoaResponse) GetRequestor() LoaRequestor`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *LoaResponse) GetRequestorOk() (*LoaRequestor, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *LoaResponse) SetRequestor(v LoaRequestor)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *LoaResponse) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetIssuer

`func (o *LoaResponse) GetIssuer() LoaIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *LoaResponse) GetIssuerOk() (*LoaIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *LoaResponse) SetIssuer(v LoaIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *LoaResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *LoaResponse) GetDemarcationPoint() LoaDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *LoaResponse) GetDemarcationPointOk() (*LoaDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *LoaResponse) SetDemarcationPoint(v LoaDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *LoaResponse) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetLocation

`func (o *LoaResponse) GetLocation() LoaLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoaResponse) GetLocationOk() (*LoaLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoaResponse) SetLocation(v LoaLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LoaResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *LoaResponse) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *LoaResponse) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *LoaResponse) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *LoaResponse) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetChangeLog

`func (o *LoaResponse) GetChangeLog() LoaChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *LoaResponse) GetChangeLogOk() (*LoaChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *LoaResponse) SetChangeLog(v LoaChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *LoaResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IssueLoa

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

## Methods

### NewIssueLoa

`func NewIssueLoa(type_ LoaType, name string, authorizedProductType LoaProductType, demarcationPoint LoaDemarcationPoint, ) *IssueLoa`

NewIssueLoa instantiates a new IssueLoa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueLoaWithDefaults

`func NewIssueLoaWithDefaults() *IssueLoa`

NewIssueLoaWithDefaults instantiates a new IssueLoa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IssueLoa) GetType() LoaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueLoa) GetTypeOk() (*LoaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueLoa) SetType(v LoaType)`

SetType sets Type field to given value.


### GetName

`func (o *IssueLoa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueLoa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueLoa) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IssueLoa) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueLoa) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueLoa) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueLoa) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthorizedProductType

`func (o *IssueLoa) GetAuthorizedProductType() LoaProductType`

GetAuthorizedProductType returns the AuthorizedProductType field if non-nil, zero value otherwise.

### GetAuthorizedProductTypeOk

`func (o *IssueLoa) GetAuthorizedProductTypeOk() (*LoaProductType, bool)`

GetAuthorizedProductTypeOk returns a tuple with the AuthorizedProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedProductType

`func (o *IssueLoa) SetAuthorizedProductType(v LoaProductType)`

SetAuthorizedProductType sets AuthorizedProductType field to given value.


### GetExpirationDateTime

`func (o *IssueLoa) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *IssueLoa) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *IssueLoa) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *IssueLoa) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetRequestor

`func (o *IssueLoa) GetRequestor() LoaRequestor`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *IssueLoa) GetRequestorOk() (*LoaRequestor, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *IssueLoa) SetRequestor(v LoaRequestor)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *IssueLoa) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *IssueLoa) GetDemarcationPoint() LoaDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *IssueLoa) GetDemarcationPointOk() (*LoaDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *IssueLoa) SetDemarcationPoint(v LoaDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



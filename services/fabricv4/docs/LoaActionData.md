# LoaActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DemarcationPoint** | Pointer to [**LoaDemarcationPoint**](LoaDemarcationPoint.md) |  | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Date and time when this LOA expires.&lt;br&gt; Default to 3 months from the creation date  | [optional] 
**PortalUrl** | Pointer to **string** | Portal URL for the LOA to either accept from requestor &lt;br&gt; or authorize from the issuer.  | [optional] [readonly] 

## Methods

### NewLoaActionData

`func NewLoaActionData() *LoaActionData`

NewLoaActionData instantiates a new LoaActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionDataWithDefaults

`func NewLoaActionDataWithDefaults() *LoaActionData`

NewLoaActionDataWithDefaults instantiates a new LoaActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDemarcationPoint

`func (o *LoaActionData) GetDemarcationPoint() LoaDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *LoaActionData) GetDemarcationPointOk() (*LoaDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *LoaActionData) SetDemarcationPoint(v LoaDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *LoaActionData) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *LoaActionData) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *LoaActionData) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *LoaActionData) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *LoaActionData) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetPortalUrl

`func (o *LoaActionData) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *LoaActionData) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *LoaActionData) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *LoaActionData) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



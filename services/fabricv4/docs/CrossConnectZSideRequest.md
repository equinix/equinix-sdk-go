# CrossConnectZSideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DemarcationPoint** | Pointer to [**CrossConnectZSideRequestDemarcationPoint**](CrossConnectZSideRequestDemarcationPoint.md) | Required for a patch panel to patch panel Cross Connect. | [optional] 
**Location** | Pointer to [**CrossConnectLocation**](CrossConnectLocation.md) | Required when connecting to a BMMR. | [optional] 
**Loa** | Pointer to [**CrossConnectLoaRef**](CrossConnectLoaRef.md) | Required for a Cross Connect with LOA and when bmmrType is EQUINIX. | [optional] 

## Methods

### NewCrossConnectZSideRequest

`func NewCrossConnectZSideRequest() *CrossConnectZSideRequest`

NewCrossConnectZSideRequest instantiates a new CrossConnectZSideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectZSideRequestWithDefaults

`func NewCrossConnectZSideRequestWithDefaults() *CrossConnectZSideRequest`

NewCrossConnectZSideRequestWithDefaults instantiates a new CrossConnectZSideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDemarcationPoint

`func (o *CrossConnectZSideRequest) GetDemarcationPoint() CrossConnectZSideRequestDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *CrossConnectZSideRequest) GetDemarcationPointOk() (*CrossConnectZSideRequestDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *CrossConnectZSideRequest) SetDemarcationPoint(v CrossConnectZSideRequestDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *CrossConnectZSideRequest) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetLocation

`func (o *CrossConnectZSideRequest) GetLocation() CrossConnectLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CrossConnectZSideRequest) GetLocationOk() (*CrossConnectLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CrossConnectZSideRequest) SetLocation(v CrossConnectLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CrossConnectZSideRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoa

`func (o *CrossConnectZSideRequest) GetLoa() CrossConnectLoaRef`

GetLoa returns the Loa field if non-nil, zero value otherwise.

### GetLoaOk

`func (o *CrossConnectZSideRequest) GetLoaOk() (*CrossConnectLoaRef, bool)`

GetLoaOk returns a tuple with the Loa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoa

`func (o *CrossConnectZSideRequest) SetLoa(v CrossConnectLoaRef)`

SetLoa sets Loa field to given value.

### HasLoa

`func (o *CrossConnectZSideRequest) HasLoa() bool`

HasLoa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ShipmentUnpackRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboundShipmentOrderNumber** | **string** | Inbound Shipment Order Number | 
**CopyOfPackingSlipNeeded** | **bool** | Copy of Packaging Slip Needed? | 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 
**DiscardShipmentMaterial** | **bool** | Discard Shipment Material? | 

## Methods

### NewShipmentUnpackRequestServiceDetails

`func NewShipmentUnpackRequestServiceDetails(inboundShipmentOrderNumber string, copyOfPackingSlipNeeded bool, scopeOfWork string, discardShipmentMaterial bool, ) *ShipmentUnpackRequestServiceDetails`

NewShipmentUnpackRequestServiceDetails instantiates a new ShipmentUnpackRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentUnpackRequestServiceDetailsWithDefaults

`func NewShipmentUnpackRequestServiceDetailsWithDefaults() *ShipmentUnpackRequestServiceDetails`

NewShipmentUnpackRequestServiceDetailsWithDefaults instantiates a new ShipmentUnpackRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundShipmentOrderNumber

`func (o *ShipmentUnpackRequestServiceDetails) GetInboundShipmentOrderNumber() string`

GetInboundShipmentOrderNumber returns the InboundShipmentOrderNumber field if non-nil, zero value otherwise.

### GetInboundShipmentOrderNumberOk

`func (o *ShipmentUnpackRequestServiceDetails) GetInboundShipmentOrderNumberOk() (*string, bool)`

GetInboundShipmentOrderNumberOk returns a tuple with the InboundShipmentOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundShipmentOrderNumber

`func (o *ShipmentUnpackRequestServiceDetails) SetInboundShipmentOrderNumber(v string)`

SetInboundShipmentOrderNumber sets InboundShipmentOrderNumber field to given value.


### GetCopyOfPackingSlipNeeded

`func (o *ShipmentUnpackRequestServiceDetails) GetCopyOfPackingSlipNeeded() bool`

GetCopyOfPackingSlipNeeded returns the CopyOfPackingSlipNeeded field if non-nil, zero value otherwise.

### GetCopyOfPackingSlipNeededOk

`func (o *ShipmentUnpackRequestServiceDetails) GetCopyOfPackingSlipNeededOk() (*bool, bool)`

GetCopyOfPackingSlipNeededOk returns a tuple with the CopyOfPackingSlipNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOfPackingSlipNeeded

`func (o *ShipmentUnpackRequestServiceDetails) SetCopyOfPackingSlipNeeded(v bool)`

SetCopyOfPackingSlipNeeded sets CopyOfPackingSlipNeeded field to given value.


### GetScopeOfWork

`func (o *ShipmentUnpackRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *ShipmentUnpackRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *ShipmentUnpackRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.


### GetDiscardShipmentMaterial

`func (o *ShipmentUnpackRequestServiceDetails) GetDiscardShipmentMaterial() bool`

GetDiscardShipmentMaterial returns the DiscardShipmentMaterial field if non-nil, zero value otherwise.

### GetDiscardShipmentMaterialOk

`func (o *ShipmentUnpackRequestServiceDetails) GetDiscardShipmentMaterialOk() (*bool, bool)`

GetDiscardShipmentMaterialOk returns a tuple with the DiscardShipmentMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardShipmentMaterial

`func (o *ShipmentUnpackRequestServiceDetails) SetDiscardShipmentMaterial(v bool)`

SetDiscardShipmentMaterial sets DiscardShipmentMaterial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



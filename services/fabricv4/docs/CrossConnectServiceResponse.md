# CrossConnectServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Service URI. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Service Identifier. | [optional] [readonly] 
**Type** | Pointer to [**CrossConnectServiceType**](CrossConnectServiceType.md) |  | [optional] [readonly] 
**State** | Pointer to [**CrossConnectServiceState**](CrossConnectServiceState.md) |  | [optional] [readonly] 
**Device** | Pointer to [**PatchEquipmentResponse**](PatchEquipmentResponse.md) | Patch equipment device details. Returned for the XC_PATCH_EQUIPMENT service. | [optional] [readonly] 
**Description** | Pointer to **string** | Additional notes captured for the service. | [optional] [readonly] 
**ChangeLog** | Pointer to [**CrossConnectChangelog**](CrossConnectChangelog.md) |  | [optional] 

## Methods

### NewCrossConnectServiceResponse

`func NewCrossConnectServiceResponse() *CrossConnectServiceResponse`

NewCrossConnectServiceResponse instantiates a new CrossConnectServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectServiceResponseWithDefaults

`func NewCrossConnectServiceResponseWithDefaults() *CrossConnectServiceResponse`

NewCrossConnectServiceResponseWithDefaults instantiates a new CrossConnectServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectServiceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectServiceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectServiceResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectServiceResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectServiceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectServiceResponse) GetType() CrossConnectServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectServiceResponse) GetTypeOk() (*CrossConnectServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectServiceResponse) SetType(v CrossConnectServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectServiceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *CrossConnectServiceResponse) GetState() CrossConnectServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CrossConnectServiceResponse) GetStateOk() (*CrossConnectServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CrossConnectServiceResponse) SetState(v CrossConnectServiceState)`

SetState sets State field to given value.

### HasState

`func (o *CrossConnectServiceResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDevice

`func (o *CrossConnectServiceResponse) GetDevice() PatchEquipmentResponse`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CrossConnectServiceResponse) GetDeviceOk() (*PatchEquipmentResponse, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CrossConnectServiceResponse) SetDevice(v PatchEquipmentResponse)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CrossConnectServiceResponse) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDescription

`func (o *CrossConnectServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CrossConnectServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CrossConnectServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CrossConnectServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChangeLog

`func (o *CrossConnectServiceResponse) GetChangeLog() CrossConnectChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CrossConnectServiceResponse) GetChangeLogOk() (*CrossConnectChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CrossConnectServiceResponse) SetChangeLog(v CrossConnectChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CrossConnectServiceResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



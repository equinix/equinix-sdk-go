# CheckCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DemarcationPointIbx** | Pointer to **string** | demarcation point IBX | [optional] 
**TetherIbx** | Pointer to **string** | cloud port ibx location | [optional] 
**MetroConnectUsed** | Pointer to **bool** | metroConnect | [optional] 
**CampusCrossConnectUsed** | Pointer to **bool** | campusCrossConnect | [optional] 
**LagEnabled** | Pointer to **bool** | lag or non lag port | [optional] 
**PhysicalPortsSpeed** | Pointer to **int32** | Physical Ports Speed in Mbps | [optional] 
**Redundancy** | Pointer to [**CheckCapacityPortRedundancy**](CheckCapacityPortRedundancy.md) |  | [optional] 
**PhysicalPortsType** | Pointer to [**PortPhysicalPortsType**](PortPhysicalPortsType.md) |  | [optional] 

## Methods

### NewCheckCapacity

`func NewCheckCapacity() *CheckCapacity`

NewCheckCapacity instantiates a new CheckCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCapacityWithDefaults

`func NewCheckCapacityWithDefaults() *CheckCapacity`

NewCheckCapacityWithDefaults instantiates a new CheckCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDemarcationPointIbx

`func (o *CheckCapacity) GetDemarcationPointIbx() string`

GetDemarcationPointIbx returns the DemarcationPointIbx field if non-nil, zero value otherwise.

### GetDemarcationPointIbxOk

`func (o *CheckCapacity) GetDemarcationPointIbxOk() (*string, bool)`

GetDemarcationPointIbxOk returns a tuple with the DemarcationPointIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPointIbx

`func (o *CheckCapacity) SetDemarcationPointIbx(v string)`

SetDemarcationPointIbx sets DemarcationPointIbx field to given value.

### HasDemarcationPointIbx

`func (o *CheckCapacity) HasDemarcationPointIbx() bool`

HasDemarcationPointIbx returns a boolean if a field has been set.

### GetTetherIbx

`func (o *CheckCapacity) GetTetherIbx() string`

GetTetherIbx returns the TetherIbx field if non-nil, zero value otherwise.

### GetTetherIbxOk

`func (o *CheckCapacity) GetTetherIbxOk() (*string, bool)`

GetTetherIbxOk returns a tuple with the TetherIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTetherIbx

`func (o *CheckCapacity) SetTetherIbx(v string)`

SetTetherIbx sets TetherIbx field to given value.

### HasTetherIbx

`func (o *CheckCapacity) HasTetherIbx() bool`

HasTetherIbx returns a boolean if a field has been set.

### GetMetroConnectUsed

`func (o *CheckCapacity) GetMetroConnectUsed() bool`

GetMetroConnectUsed returns the MetroConnectUsed field if non-nil, zero value otherwise.

### GetMetroConnectUsedOk

`func (o *CheckCapacity) GetMetroConnectUsedOk() (*bool, bool)`

GetMetroConnectUsedOk returns a tuple with the MetroConnectUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroConnectUsed

`func (o *CheckCapacity) SetMetroConnectUsed(v bool)`

SetMetroConnectUsed sets MetroConnectUsed field to given value.

### HasMetroConnectUsed

`func (o *CheckCapacity) HasMetroConnectUsed() bool`

HasMetroConnectUsed returns a boolean if a field has been set.

### GetCampusCrossConnectUsed

`func (o *CheckCapacity) GetCampusCrossConnectUsed() bool`

GetCampusCrossConnectUsed returns the CampusCrossConnectUsed field if non-nil, zero value otherwise.

### GetCampusCrossConnectUsedOk

`func (o *CheckCapacity) GetCampusCrossConnectUsedOk() (*bool, bool)`

GetCampusCrossConnectUsedOk returns a tuple with the CampusCrossConnectUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampusCrossConnectUsed

`func (o *CheckCapacity) SetCampusCrossConnectUsed(v bool)`

SetCampusCrossConnectUsed sets CampusCrossConnectUsed field to given value.

### HasCampusCrossConnectUsed

`func (o *CheckCapacity) HasCampusCrossConnectUsed() bool`

HasCampusCrossConnectUsed returns a boolean if a field has been set.

### GetLagEnabled

`func (o *CheckCapacity) GetLagEnabled() bool`

GetLagEnabled returns the LagEnabled field if non-nil, zero value otherwise.

### GetLagEnabledOk

`func (o *CheckCapacity) GetLagEnabledOk() (*bool, bool)`

GetLagEnabledOk returns a tuple with the LagEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagEnabled

`func (o *CheckCapacity) SetLagEnabled(v bool)`

SetLagEnabled sets LagEnabled field to given value.

### HasLagEnabled

`func (o *CheckCapacity) HasLagEnabled() bool`

HasLagEnabled returns a boolean if a field has been set.

### GetPhysicalPortsSpeed

`func (o *CheckCapacity) GetPhysicalPortsSpeed() int32`

GetPhysicalPortsSpeed returns the PhysicalPortsSpeed field if non-nil, zero value otherwise.

### GetPhysicalPortsSpeedOk

`func (o *CheckCapacity) GetPhysicalPortsSpeedOk() (*int32, bool)`

GetPhysicalPortsSpeedOk returns a tuple with the PhysicalPortsSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsSpeed

`func (o *CheckCapacity) SetPhysicalPortsSpeed(v int32)`

SetPhysicalPortsSpeed sets PhysicalPortsSpeed field to given value.

### HasPhysicalPortsSpeed

`func (o *CheckCapacity) HasPhysicalPortsSpeed() bool`

HasPhysicalPortsSpeed returns a boolean if a field has been set.

### GetRedundancy

`func (o *CheckCapacity) GetRedundancy() CheckCapacityPortRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CheckCapacity) GetRedundancyOk() (*CheckCapacityPortRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CheckCapacity) SetRedundancy(v CheckCapacityPortRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CheckCapacity) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetPhysicalPortsType

`func (o *CheckCapacity) GetPhysicalPortsType() PortPhysicalPortsType`

GetPhysicalPortsType returns the PhysicalPortsType field if non-nil, zero value otherwise.

### GetPhysicalPortsTypeOk

`func (o *CheckCapacity) GetPhysicalPortsTypeOk() (*PortPhysicalPortsType, bool)`

GetPhysicalPortsTypeOk returns a tuple with the PhysicalPortsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsType

`func (o *CheckCapacity) SetPhysicalPortsType(v PortPhysicalPortsType)`

SetPhysicalPortsType sets PhysicalPortsType field to given value.

### HasPhysicalPortsType

`func (o *CheckCapacity) HasPhysicalPortsType() bool`

HasPhysicalPortsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



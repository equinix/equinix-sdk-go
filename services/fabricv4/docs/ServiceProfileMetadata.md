# ServiceProfileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Props** | Pointer to **string** |  | [optional] 
**RegEx** | Pointer to **string** |  | [optional] 
**RegExMsg** | Pointer to **string** |  | [optional] 
**VlanRangeMaxValue** | Pointer to **int32** |  | [optional] 
**VlanRangeMinValue** | Pointer to **int32** |  | [optional] 
**MaxQinq** | Pointer to **string** |  | [optional] 
**MaxDot1q** | Pointer to **int32** |  | [optional] 
**VariableBilling** | Pointer to **bool** |  | [optional] 
**GlobalOrganization** | Pointer to **string** |  | [optional] 
**LimitAuthKeyConn** | Pointer to **bool** |  | [optional] 
**AllowSecondaryLocation** | Pointer to **bool** |  | [optional] 
**RedundantProfileId** | Pointer to **string** |  | [optional] 
**AllowVcMigration** | Pointer to **bool** |  | [optional] 
**ConnectionEditable** | Pointer to **bool** |  | [optional] 
**ReleaseVlan** | Pointer to **bool** |  | [optional] 
**MaxConnectionsOnPort** | Pointer to **int32** |  | [optional] 
**PortAssignmentStrategy** | Pointer to **string** |  | [optional] 
**EqxManagedPort** | Pointer to **bool** |  | [optional] 
**ConnectionNameEditable** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceProfileMetadata

`func NewServiceProfileMetadata() *ServiceProfileMetadata`

NewServiceProfileMetadata instantiates a new ServiceProfileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileMetadataWithDefaults

`func NewServiceProfileMetadataWithDefaults() *ServiceProfileMetadata`

NewServiceProfileMetadataWithDefaults instantiates a new ServiceProfileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProps

`func (o *ServiceProfileMetadata) GetProps() string`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *ServiceProfileMetadata) GetPropsOk() (*string, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *ServiceProfileMetadata) SetProps(v string)`

SetProps sets Props field to given value.

### HasProps

`func (o *ServiceProfileMetadata) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetRegEx

`func (o *ServiceProfileMetadata) GetRegEx() string`

GetRegEx returns the RegEx field if non-nil, zero value otherwise.

### GetRegExOk

`func (o *ServiceProfileMetadata) GetRegExOk() (*string, bool)`

GetRegExOk returns a tuple with the RegEx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegEx

`func (o *ServiceProfileMetadata) SetRegEx(v string)`

SetRegEx sets RegEx field to given value.

### HasRegEx

`func (o *ServiceProfileMetadata) HasRegEx() bool`

HasRegEx returns a boolean if a field has been set.

### GetRegExMsg

`func (o *ServiceProfileMetadata) GetRegExMsg() string`

GetRegExMsg returns the RegExMsg field if non-nil, zero value otherwise.

### GetRegExMsgOk

`func (o *ServiceProfileMetadata) GetRegExMsgOk() (*string, bool)`

GetRegExMsgOk returns a tuple with the RegExMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegExMsg

`func (o *ServiceProfileMetadata) SetRegExMsg(v string)`

SetRegExMsg sets RegExMsg field to given value.

### HasRegExMsg

`func (o *ServiceProfileMetadata) HasRegExMsg() bool`

HasRegExMsg returns a boolean if a field has been set.

### GetVlanRangeMaxValue

`func (o *ServiceProfileMetadata) GetVlanRangeMaxValue() int32`

GetVlanRangeMaxValue returns the VlanRangeMaxValue field if non-nil, zero value otherwise.

### GetVlanRangeMaxValueOk

`func (o *ServiceProfileMetadata) GetVlanRangeMaxValueOk() (*int32, bool)`

GetVlanRangeMaxValueOk returns a tuple with the VlanRangeMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRangeMaxValue

`func (o *ServiceProfileMetadata) SetVlanRangeMaxValue(v int32)`

SetVlanRangeMaxValue sets VlanRangeMaxValue field to given value.

### HasVlanRangeMaxValue

`func (o *ServiceProfileMetadata) HasVlanRangeMaxValue() bool`

HasVlanRangeMaxValue returns a boolean if a field has been set.

### GetVlanRangeMinValue

`func (o *ServiceProfileMetadata) GetVlanRangeMinValue() int32`

GetVlanRangeMinValue returns the VlanRangeMinValue field if non-nil, zero value otherwise.

### GetVlanRangeMinValueOk

`func (o *ServiceProfileMetadata) GetVlanRangeMinValueOk() (*int32, bool)`

GetVlanRangeMinValueOk returns a tuple with the VlanRangeMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRangeMinValue

`func (o *ServiceProfileMetadata) SetVlanRangeMinValue(v int32)`

SetVlanRangeMinValue sets VlanRangeMinValue field to given value.

### HasVlanRangeMinValue

`func (o *ServiceProfileMetadata) HasVlanRangeMinValue() bool`

HasVlanRangeMinValue returns a boolean if a field has been set.

### GetMaxQinq

`func (o *ServiceProfileMetadata) GetMaxQinq() string`

GetMaxQinq returns the MaxQinq field if non-nil, zero value otherwise.

### GetMaxQinqOk

`func (o *ServiceProfileMetadata) GetMaxQinqOk() (*string, bool)`

GetMaxQinqOk returns a tuple with the MaxQinq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQinq

`func (o *ServiceProfileMetadata) SetMaxQinq(v string)`

SetMaxQinq sets MaxQinq field to given value.

### HasMaxQinq

`func (o *ServiceProfileMetadata) HasMaxQinq() bool`

HasMaxQinq returns a boolean if a field has been set.

### GetMaxDot1q

`func (o *ServiceProfileMetadata) GetMaxDot1q() int32`

GetMaxDot1q returns the MaxDot1q field if non-nil, zero value otherwise.

### GetMaxDot1qOk

`func (o *ServiceProfileMetadata) GetMaxDot1qOk() (*int32, bool)`

GetMaxDot1qOk returns a tuple with the MaxDot1q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDot1q

`func (o *ServiceProfileMetadata) SetMaxDot1q(v int32)`

SetMaxDot1q sets MaxDot1q field to given value.

### HasMaxDot1q

`func (o *ServiceProfileMetadata) HasMaxDot1q() bool`

HasMaxDot1q returns a boolean if a field has been set.

### GetVariableBilling

`func (o *ServiceProfileMetadata) GetVariableBilling() bool`

GetVariableBilling returns the VariableBilling field if non-nil, zero value otherwise.

### GetVariableBillingOk

`func (o *ServiceProfileMetadata) GetVariableBillingOk() (*bool, bool)`

GetVariableBillingOk returns a tuple with the VariableBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableBilling

`func (o *ServiceProfileMetadata) SetVariableBilling(v bool)`

SetVariableBilling sets VariableBilling field to given value.

### HasVariableBilling

`func (o *ServiceProfileMetadata) HasVariableBilling() bool`

HasVariableBilling returns a boolean if a field has been set.

### GetGlobalOrganization

`func (o *ServiceProfileMetadata) GetGlobalOrganization() string`

GetGlobalOrganization returns the GlobalOrganization field if non-nil, zero value otherwise.

### GetGlobalOrganizationOk

`func (o *ServiceProfileMetadata) GetGlobalOrganizationOk() (*string, bool)`

GetGlobalOrganizationOk returns a tuple with the GlobalOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalOrganization

`func (o *ServiceProfileMetadata) SetGlobalOrganization(v string)`

SetGlobalOrganization sets GlobalOrganization field to given value.

### HasGlobalOrganization

`func (o *ServiceProfileMetadata) HasGlobalOrganization() bool`

HasGlobalOrganization returns a boolean if a field has been set.

### GetLimitAuthKeyConn

`func (o *ServiceProfileMetadata) GetLimitAuthKeyConn() bool`

GetLimitAuthKeyConn returns the LimitAuthKeyConn field if non-nil, zero value otherwise.

### GetLimitAuthKeyConnOk

`func (o *ServiceProfileMetadata) GetLimitAuthKeyConnOk() (*bool, bool)`

GetLimitAuthKeyConnOk returns a tuple with the LimitAuthKeyConn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitAuthKeyConn

`func (o *ServiceProfileMetadata) SetLimitAuthKeyConn(v bool)`

SetLimitAuthKeyConn sets LimitAuthKeyConn field to given value.

### HasLimitAuthKeyConn

`func (o *ServiceProfileMetadata) HasLimitAuthKeyConn() bool`

HasLimitAuthKeyConn returns a boolean if a field has been set.

### GetAllowSecondaryLocation

`func (o *ServiceProfileMetadata) GetAllowSecondaryLocation() bool`

GetAllowSecondaryLocation returns the AllowSecondaryLocation field if non-nil, zero value otherwise.

### GetAllowSecondaryLocationOk

`func (o *ServiceProfileMetadata) GetAllowSecondaryLocationOk() (*bool, bool)`

GetAllowSecondaryLocationOk returns a tuple with the AllowSecondaryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSecondaryLocation

`func (o *ServiceProfileMetadata) SetAllowSecondaryLocation(v bool)`

SetAllowSecondaryLocation sets AllowSecondaryLocation field to given value.

### HasAllowSecondaryLocation

`func (o *ServiceProfileMetadata) HasAllowSecondaryLocation() bool`

HasAllowSecondaryLocation returns a boolean if a field has been set.

### GetRedundantProfileId

`func (o *ServiceProfileMetadata) GetRedundantProfileId() string`

GetRedundantProfileId returns the RedundantProfileId field if non-nil, zero value otherwise.

### GetRedundantProfileIdOk

`func (o *ServiceProfileMetadata) GetRedundantProfileIdOk() (*string, bool)`

GetRedundantProfileIdOk returns a tuple with the RedundantProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantProfileId

`func (o *ServiceProfileMetadata) SetRedundantProfileId(v string)`

SetRedundantProfileId sets RedundantProfileId field to given value.

### HasRedundantProfileId

`func (o *ServiceProfileMetadata) HasRedundantProfileId() bool`

HasRedundantProfileId returns a boolean if a field has been set.

### GetAllowVcMigration

`func (o *ServiceProfileMetadata) GetAllowVcMigration() bool`

GetAllowVcMigration returns the AllowVcMigration field if non-nil, zero value otherwise.

### GetAllowVcMigrationOk

`func (o *ServiceProfileMetadata) GetAllowVcMigrationOk() (*bool, bool)`

GetAllowVcMigrationOk returns a tuple with the AllowVcMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVcMigration

`func (o *ServiceProfileMetadata) SetAllowVcMigration(v bool)`

SetAllowVcMigration sets AllowVcMigration field to given value.

### HasAllowVcMigration

`func (o *ServiceProfileMetadata) HasAllowVcMigration() bool`

HasAllowVcMigration returns a boolean if a field has been set.

### GetConnectionEditable

`func (o *ServiceProfileMetadata) GetConnectionEditable() bool`

GetConnectionEditable returns the ConnectionEditable field if non-nil, zero value otherwise.

### GetConnectionEditableOk

`func (o *ServiceProfileMetadata) GetConnectionEditableOk() (*bool, bool)`

GetConnectionEditableOk returns a tuple with the ConnectionEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionEditable

`func (o *ServiceProfileMetadata) SetConnectionEditable(v bool)`

SetConnectionEditable sets ConnectionEditable field to given value.

### HasConnectionEditable

`func (o *ServiceProfileMetadata) HasConnectionEditable() bool`

HasConnectionEditable returns a boolean if a field has been set.

### GetReleaseVlan

`func (o *ServiceProfileMetadata) GetReleaseVlan() bool`

GetReleaseVlan returns the ReleaseVlan field if non-nil, zero value otherwise.

### GetReleaseVlanOk

`func (o *ServiceProfileMetadata) GetReleaseVlanOk() (*bool, bool)`

GetReleaseVlanOk returns a tuple with the ReleaseVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVlan

`func (o *ServiceProfileMetadata) SetReleaseVlan(v bool)`

SetReleaseVlan sets ReleaseVlan field to given value.

### HasReleaseVlan

`func (o *ServiceProfileMetadata) HasReleaseVlan() bool`

HasReleaseVlan returns a boolean if a field has been set.

### GetMaxConnectionsOnPort

`func (o *ServiceProfileMetadata) GetMaxConnectionsOnPort() int32`

GetMaxConnectionsOnPort returns the MaxConnectionsOnPort field if non-nil, zero value otherwise.

### GetMaxConnectionsOnPortOk

`func (o *ServiceProfileMetadata) GetMaxConnectionsOnPortOk() (*int32, bool)`

GetMaxConnectionsOnPortOk returns a tuple with the MaxConnectionsOnPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionsOnPort

`func (o *ServiceProfileMetadata) SetMaxConnectionsOnPort(v int32)`

SetMaxConnectionsOnPort sets MaxConnectionsOnPort field to given value.

### HasMaxConnectionsOnPort

`func (o *ServiceProfileMetadata) HasMaxConnectionsOnPort() bool`

HasMaxConnectionsOnPort returns a boolean if a field has been set.

### GetPortAssignmentStrategy

`func (o *ServiceProfileMetadata) GetPortAssignmentStrategy() string`

GetPortAssignmentStrategy returns the PortAssignmentStrategy field if non-nil, zero value otherwise.

### GetPortAssignmentStrategyOk

`func (o *ServiceProfileMetadata) GetPortAssignmentStrategyOk() (*string, bool)`

GetPortAssignmentStrategyOk returns a tuple with the PortAssignmentStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAssignmentStrategy

`func (o *ServiceProfileMetadata) SetPortAssignmentStrategy(v string)`

SetPortAssignmentStrategy sets PortAssignmentStrategy field to given value.

### HasPortAssignmentStrategy

`func (o *ServiceProfileMetadata) HasPortAssignmentStrategy() bool`

HasPortAssignmentStrategy returns a boolean if a field has been set.

### GetEqxManagedPort

`func (o *ServiceProfileMetadata) GetEqxManagedPort() bool`

GetEqxManagedPort returns the EqxManagedPort field if non-nil, zero value otherwise.

### GetEqxManagedPortOk

`func (o *ServiceProfileMetadata) GetEqxManagedPortOk() (*bool, bool)`

GetEqxManagedPortOk returns a tuple with the EqxManagedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqxManagedPort

`func (o *ServiceProfileMetadata) SetEqxManagedPort(v bool)`

SetEqxManagedPort sets EqxManagedPort field to given value.

### HasEqxManagedPort

`func (o *ServiceProfileMetadata) HasEqxManagedPort() bool`

HasEqxManagedPort returns a boolean if a field has been set.

### GetConnectionNameEditable

`func (o *ServiceProfileMetadata) GetConnectionNameEditable() bool`

GetConnectionNameEditable returns the ConnectionNameEditable field if non-nil, zero value otherwise.

### GetConnectionNameEditableOk

`func (o *ServiceProfileMetadata) GetConnectionNameEditableOk() (*bool, bool)`

GetConnectionNameEditableOk returns a tuple with the ConnectionNameEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionNameEditable

`func (o *ServiceProfileMetadata) SetConnectionNameEditable(v bool)`

SetConnectionNameEditable sets ConnectionNameEditable field to given value.

### HasConnectionNameEditable

`func (o *ServiceProfileMetadata) HasConnectionNameEditable() bool`

HasConnectionNameEditable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



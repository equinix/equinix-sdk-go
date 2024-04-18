# InterfaceStatsDetailObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | Pointer to **string** | Start time of the duration for which you want stats. | [optional] 
**EndDateTime** | Pointer to **string** | End time of the duration for which you want stats. | [optional] 
**Unit** | Pointer to **string** | Unit. | [optional] 
**Inbound** | Pointer to [**InterfaceStatsofTraffic**](InterfaceStatsofTraffic.md) |  | [optional] 
**Outbound** | Pointer to [**InterfaceStatsofTraffic**](InterfaceStatsofTraffic.md) |  | [optional] 

## Methods

### NewInterfaceStatsDetailObject

`func NewInterfaceStatsDetailObject() *InterfaceStatsDetailObject`

NewInterfaceStatsDetailObject instantiates a new InterfaceStatsDetailObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceStatsDetailObjectWithDefaults

`func NewInterfaceStatsDetailObjectWithDefaults() *InterfaceStatsDetailObject`

NewInterfaceStatsDetailObjectWithDefaults instantiates a new InterfaceStatsDetailObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *InterfaceStatsDetailObject) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *InterfaceStatsDetailObject) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *InterfaceStatsDetailObject) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *InterfaceStatsDetailObject) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *InterfaceStatsDetailObject) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *InterfaceStatsDetailObject) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *InterfaceStatsDetailObject) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *InterfaceStatsDetailObject) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetUnit

`func (o *InterfaceStatsDetailObject) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InterfaceStatsDetailObject) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InterfaceStatsDetailObject) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InterfaceStatsDetailObject) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetInbound

`func (o *InterfaceStatsDetailObject) GetInbound() InterfaceStatsofTraffic`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *InterfaceStatsDetailObject) GetInboundOk() (*InterfaceStatsofTraffic, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *InterfaceStatsDetailObject) SetInbound(v InterfaceStatsofTraffic)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *InterfaceStatsDetailObject) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *InterfaceStatsDetailObject) GetOutbound() InterfaceStatsofTraffic`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *InterfaceStatsDetailObject) GetOutboundOk() (*InterfaceStatsofTraffic, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *InterfaceStatsDetailObject) SetOutbound(v InterfaceStatsofTraffic)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *InterfaceStatsDetailObject) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



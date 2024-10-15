# ConnectionRouteTableEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RouteTableEntryType**](RouteTableEntryType.md) |  | 
**ProtocolType** | Pointer to [**RouteTableEntryProtocolType**](RouteTableEntryProtocolType.md) |  | [optional] 
**State** | [**ConnectionRouteTableEntryState**](ConnectionRouteTableEntryState.md) |  | 
**Age** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**NextHop** | Pointer to **string** |  | [optional] 
**MED** | Pointer to **int32** |  | [optional] 
**LocalPreference** | Pointer to **int32** |  | [optional] 
**AsPath** | Pointer to **[]string** |  | [optional] 
**Connection** | Pointer to [**ConnectionRouteTableEntryConnection**](ConnectionRouteTableEntryConnection.md) |  | [optional] 
**ChangeLog** | [**Changelog**](Changelog.md) |  | 

## Methods

### NewConnectionRouteTableEntry

`func NewConnectionRouteTableEntry(type_ RouteTableEntryType, state ConnectionRouteTableEntryState, changeLog Changelog, ) *ConnectionRouteTableEntry`

NewConnectionRouteTableEntry instantiates a new ConnectionRouteTableEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRouteTableEntryWithDefaults

`func NewConnectionRouteTableEntryWithDefaults() *ConnectionRouteTableEntry`

NewConnectionRouteTableEntryWithDefaults instantiates a new ConnectionRouteTableEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionRouteTableEntry) GetType() RouteTableEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionRouteTableEntry) GetTypeOk() (*RouteTableEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionRouteTableEntry) SetType(v RouteTableEntryType)`

SetType sets Type field to given value.


### GetProtocolType

`func (o *ConnectionRouteTableEntry) GetProtocolType() RouteTableEntryProtocolType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *ConnectionRouteTableEntry) GetProtocolTypeOk() (*RouteTableEntryProtocolType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *ConnectionRouteTableEntry) SetProtocolType(v RouteTableEntryProtocolType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *ConnectionRouteTableEntry) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetState

`func (o *ConnectionRouteTableEntry) GetState() ConnectionRouteTableEntryState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectionRouteTableEntry) GetStateOk() (*ConnectionRouteTableEntryState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectionRouteTableEntry) SetState(v ConnectionRouteTableEntryState)`

SetState sets State field to given value.


### GetAge

`func (o *ConnectionRouteTableEntry) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ConnectionRouteTableEntry) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ConnectionRouteTableEntry) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *ConnectionRouteTableEntry) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetPrefix

`func (o *ConnectionRouteTableEntry) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ConnectionRouteTableEntry) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ConnectionRouteTableEntry) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ConnectionRouteTableEntry) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNextHop

`func (o *ConnectionRouteTableEntry) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *ConnectionRouteTableEntry) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *ConnectionRouteTableEntry) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *ConnectionRouteTableEntry) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetMED

`func (o *ConnectionRouteTableEntry) GetMED() int32`

GetMED returns the MED field if non-nil, zero value otherwise.

### GetMEDOk

`func (o *ConnectionRouteTableEntry) GetMEDOk() (*int32, bool)`

GetMEDOk returns a tuple with the MED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMED

`func (o *ConnectionRouteTableEntry) SetMED(v int32)`

SetMED sets MED field to given value.

### HasMED

`func (o *ConnectionRouteTableEntry) HasMED() bool`

HasMED returns a boolean if a field has been set.

### GetLocalPreference

`func (o *ConnectionRouteTableEntry) GetLocalPreference() int32`

GetLocalPreference returns the LocalPreference field if non-nil, zero value otherwise.

### GetLocalPreferenceOk

`func (o *ConnectionRouteTableEntry) GetLocalPreferenceOk() (*int32, bool)`

GetLocalPreferenceOk returns a tuple with the LocalPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPreference

`func (o *ConnectionRouteTableEntry) SetLocalPreference(v int32)`

SetLocalPreference sets LocalPreference field to given value.

### HasLocalPreference

`func (o *ConnectionRouteTableEntry) HasLocalPreference() bool`

HasLocalPreference returns a boolean if a field has been set.

### GetAsPath

`func (o *ConnectionRouteTableEntry) GetAsPath() []string`

GetAsPath returns the AsPath field if non-nil, zero value otherwise.

### GetAsPathOk

`func (o *ConnectionRouteTableEntry) GetAsPathOk() (*[]string, bool)`

GetAsPathOk returns a tuple with the AsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsPath

`func (o *ConnectionRouteTableEntry) SetAsPath(v []string)`

SetAsPath sets AsPath field to given value.

### HasAsPath

`func (o *ConnectionRouteTableEntry) HasAsPath() bool`

HasAsPath returns a boolean if a field has been set.

### GetConnection

`func (o *ConnectionRouteTableEntry) GetConnection() ConnectionRouteTableEntryConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ConnectionRouteTableEntry) GetConnectionOk() (*ConnectionRouteTableEntryConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ConnectionRouteTableEntry) SetConnection(v ConnectionRouteTableEntryConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ConnectionRouteTableEntry) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetChangeLog

`func (o *ConnectionRouteTableEntry) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ConnectionRouteTableEntry) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ConnectionRouteTableEntry) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



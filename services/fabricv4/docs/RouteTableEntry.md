# RouteTableEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RouteTableEntryType**](RouteTableEntryType.md) |  | 
**ProtocolType** | Pointer to [**RouteTableEntryProtocolType**](RouteTableEntryProtocolType.md) |  | [optional] 
**State** | [**RouteTableEntryState**](RouteTableEntryState.md) |  | 
**Age** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**NextHop** | Pointer to **string** |  | [optional] 
**MED** | Pointer to **int32** |  | [optional] 
**LocalPreference** | Pointer to **int32** |  | [optional] 
**AsPath** | Pointer to **[]string** |  | [optional] 
**Connection** | Pointer to [**ConnectionRouteTableEntryConnection**](ConnectionRouteTableEntryConnection.md) |  | [optional] 
**ChangeLog** | [**Changelog**](Changelog.md) |  | 

## Methods

### NewRouteTableEntry

`func NewRouteTableEntry(type_ RouteTableEntryType, state RouteTableEntryState, changeLog Changelog, ) *RouteTableEntry`

NewRouteTableEntry instantiates a new RouteTableEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableEntryWithDefaults

`func NewRouteTableEntryWithDefaults() *RouteTableEntry`

NewRouteTableEntryWithDefaults instantiates a new RouteTableEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RouteTableEntry) GetType() RouteTableEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteTableEntry) GetTypeOk() (*RouteTableEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteTableEntry) SetType(v RouteTableEntryType)`

SetType sets Type field to given value.


### GetProtocolType

`func (o *RouteTableEntry) GetProtocolType() RouteTableEntryProtocolType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *RouteTableEntry) GetProtocolTypeOk() (*RouteTableEntryProtocolType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *RouteTableEntry) SetProtocolType(v RouteTableEntryProtocolType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *RouteTableEntry) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetState

`func (o *RouteTableEntry) GetState() RouteTableEntryState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteTableEntry) GetStateOk() (*RouteTableEntryState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RouteTableEntry) SetState(v RouteTableEntryState)`

SetState sets State field to given value.


### GetAge

`func (o *RouteTableEntry) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RouteTableEntry) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RouteTableEntry) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *RouteTableEntry) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetPrefix

`func (o *RouteTableEntry) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteTableEntry) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteTableEntry) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RouteTableEntry) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNextHop

`func (o *RouteTableEntry) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *RouteTableEntry) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *RouteTableEntry) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *RouteTableEntry) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetMED

`func (o *RouteTableEntry) GetMED() int32`

GetMED returns the MED field if non-nil, zero value otherwise.

### GetMEDOk

`func (o *RouteTableEntry) GetMEDOk() (*int32, bool)`

GetMEDOk returns a tuple with the MED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMED

`func (o *RouteTableEntry) SetMED(v int32)`

SetMED sets MED field to given value.

### HasMED

`func (o *RouteTableEntry) HasMED() bool`

HasMED returns a boolean if a field has been set.

### GetLocalPreference

`func (o *RouteTableEntry) GetLocalPreference() int32`

GetLocalPreference returns the LocalPreference field if non-nil, zero value otherwise.

### GetLocalPreferenceOk

`func (o *RouteTableEntry) GetLocalPreferenceOk() (*int32, bool)`

GetLocalPreferenceOk returns a tuple with the LocalPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPreference

`func (o *RouteTableEntry) SetLocalPreference(v int32)`

SetLocalPreference sets LocalPreference field to given value.

### HasLocalPreference

`func (o *RouteTableEntry) HasLocalPreference() bool`

HasLocalPreference returns a boolean if a field has been set.

### GetAsPath

`func (o *RouteTableEntry) GetAsPath() []string`

GetAsPath returns the AsPath field if non-nil, zero value otherwise.

### GetAsPathOk

`func (o *RouteTableEntry) GetAsPathOk() (*[]string, bool)`

GetAsPathOk returns a tuple with the AsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsPath

`func (o *RouteTableEntry) SetAsPath(v []string)`

SetAsPath sets AsPath field to given value.

### HasAsPath

`func (o *RouteTableEntry) HasAsPath() bool`

HasAsPath returns a boolean if a field has been set.

### GetConnection

`func (o *RouteTableEntry) GetConnection() ConnectionRouteTableEntryConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *RouteTableEntry) GetConnectionOk() (*ConnectionRouteTableEntryConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *RouteTableEntry) SetConnection(v ConnectionRouteTableEntryConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *RouteTableEntry) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetChangeLog

`func (o *RouteTableEntry) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RouteTableEntry) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RouteTableEntry) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



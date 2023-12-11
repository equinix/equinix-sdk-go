# InterconnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interconnections** | Pointer to [**[]Interconnection**](Interconnection.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewInterconnectionList

`func NewInterconnectionList() *InterconnectionList`

NewInterconnectionList instantiates a new InterconnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionListWithDefaults

`func NewInterconnectionListWithDefaults() *InterconnectionList`

NewInterconnectionListWithDefaults instantiates a new InterconnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterconnections

`func (o *InterconnectionList) GetInterconnections() []Interconnection`

GetInterconnections returns the Interconnections field if non-nil, zero value otherwise.

### GetInterconnectionsOk

`func (o *InterconnectionList) GetInterconnectionsOk() (*[]Interconnection, bool)`

GetInterconnectionsOk returns a tuple with the Interconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnections

`func (o *InterconnectionList) SetInterconnections(v []Interconnection)`

SetInterconnections sets Interconnections field to given value.

### HasInterconnections

`func (o *InterconnectionList) HasInterconnections() bool`

HasInterconnections returns a boolean if a field has been set.

### GetMeta

`func (o *InterconnectionList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InterconnectionList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InterconnectionList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InterconnectionList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



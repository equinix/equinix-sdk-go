/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the ConnectionRouteTableEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionRouteTableEntry{}

// ConnectionRouteTableEntry Advertised and received route table entry object
type ConnectionRouteTableEntry struct {
	Type                 RouteTableEntryType                  `json:"type"`
	ProtocolType         *RouteTableEntryProtocolType         `json:"protocolType,omitempty"`
	State                ConnectionRouteTableEntryState       `json:"state"`
	Age                  *string                              `json:"age,omitempty"`
	Prefix               *string                              `json:"prefix,omitempty"`
	NextHop              *string                              `json:"nextHop,omitempty"`
	MED                  *int32                               `json:"MED,omitempty"`
	LocalPreference      *int32                               `json:"localPreference,omitempty"`
	AsPath               []string                             `json:"asPath,omitempty"`
	Connection           *ConnectionRouteTableEntryConnection `json:"connection,omitempty"`
	ChangeLog            Changelog                            `json:"changeLog"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionRouteTableEntry ConnectionRouteTableEntry

// NewConnectionRouteTableEntry instantiates a new ConnectionRouteTableEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionRouteTableEntry(type_ RouteTableEntryType, state ConnectionRouteTableEntryState, changeLog Changelog) *ConnectionRouteTableEntry {
	this := ConnectionRouteTableEntry{}
	this.Type = type_
	this.State = state
	this.ChangeLog = changeLog
	return &this
}

// NewConnectionRouteTableEntryWithDefaults instantiates a new ConnectionRouteTableEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionRouteTableEntryWithDefaults() *ConnectionRouteTableEntry {
	this := ConnectionRouteTableEntry{}
	return &this
}

// GetType returns the Type field value
func (o *ConnectionRouteTableEntry) GetType() RouteTableEntryType {
	if o == nil {
		var ret RouteTableEntryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetTypeOk() (*RouteTableEntryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectionRouteTableEntry) SetType(v RouteTableEntryType) {
	o.Type = v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetProtocolType() RouteTableEntryProtocolType {
	if o == nil || IsNil(o.ProtocolType) {
		var ret RouteTableEntryProtocolType
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetProtocolTypeOk() (*RouteTableEntryProtocolType, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given RouteTableEntryProtocolType and assigns it to the ProtocolType field.
func (o *ConnectionRouteTableEntry) SetProtocolType(v RouteTableEntryProtocolType) {
	o.ProtocolType = &v
}

// GetState returns the State field value
func (o *ConnectionRouteTableEntry) GetState() ConnectionRouteTableEntryState {
	if o == nil {
		var ret ConnectionRouteTableEntryState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetStateOk() (*ConnectionRouteTableEntryState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ConnectionRouteTableEntry) SetState(v ConnectionRouteTableEntryState) {
	o.State = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetAge() string {
	if o == nil || IsNil(o.Age) {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetAgeOk() (*string, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *ConnectionRouteTableEntry) SetAge(v string) {
	o.Age = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ConnectionRouteTableEntry) SetPrefix(v string) {
	o.Prefix = &v
}

// GetNextHop returns the NextHop field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetNextHop() string {
	if o == nil || IsNil(o.NextHop) {
		var ret string
		return ret
	}
	return *o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetNextHopOk() (*string, bool) {
	if o == nil || IsNil(o.NextHop) {
		return nil, false
	}
	return o.NextHop, true
}

// HasNextHop returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasNextHop() bool {
	if o != nil && !IsNil(o.NextHop) {
		return true
	}

	return false
}

// SetNextHop gets a reference to the given string and assigns it to the NextHop field.
func (o *ConnectionRouteTableEntry) SetNextHop(v string) {
	o.NextHop = &v
}

// GetMED returns the MED field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetMED() int32 {
	if o == nil || IsNil(o.MED) {
		var ret int32
		return ret
	}
	return *o.MED
}

// GetMEDOk returns a tuple with the MED field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetMEDOk() (*int32, bool) {
	if o == nil || IsNil(o.MED) {
		return nil, false
	}
	return o.MED, true
}

// HasMED returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasMED() bool {
	if o != nil && !IsNil(o.MED) {
		return true
	}

	return false
}

// SetMED gets a reference to the given int32 and assigns it to the MED field.
func (o *ConnectionRouteTableEntry) SetMED(v int32) {
	o.MED = &v
}

// GetLocalPreference returns the LocalPreference field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetLocalPreference() int32 {
	if o == nil || IsNil(o.LocalPreference) {
		var ret int32
		return ret
	}
	return *o.LocalPreference
}

// GetLocalPreferenceOk returns a tuple with the LocalPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetLocalPreferenceOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalPreference) {
		return nil, false
	}
	return o.LocalPreference, true
}

// HasLocalPreference returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasLocalPreference() bool {
	if o != nil && !IsNil(o.LocalPreference) {
		return true
	}

	return false
}

// SetLocalPreference gets a reference to the given int32 and assigns it to the LocalPreference field.
func (o *ConnectionRouteTableEntry) SetLocalPreference(v int32) {
	o.LocalPreference = &v
}

// GetAsPath returns the AsPath field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetAsPath() []string {
	if o == nil || IsNil(o.AsPath) {
		var ret []string
		return ret
	}
	return o.AsPath
}

// GetAsPathOk returns a tuple with the AsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetAsPathOk() ([]string, bool) {
	if o == nil || IsNil(o.AsPath) {
		return nil, false
	}
	return o.AsPath, true
}

// HasAsPath returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasAsPath() bool {
	if o != nil && !IsNil(o.AsPath) {
		return true
	}

	return false
}

// SetAsPath gets a reference to the given []string and assigns it to the AsPath field.
func (o *ConnectionRouteTableEntry) SetAsPath(v []string) {
	o.AsPath = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntry) GetConnection() ConnectionRouteTableEntryConnection {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionRouteTableEntryConnection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetConnectionOk() (*ConnectionRouteTableEntryConnection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntry) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionRouteTableEntryConnection and assigns it to the Connection field.
func (o *ConnectionRouteTableEntry) SetConnection(v ConnectionRouteTableEntryConnection) {
	o.Connection = &v
}

// GetChangeLog returns the ChangeLog field value
func (o *ConnectionRouteTableEntry) GetChangeLog() Changelog {
	if o == nil {
		var ret Changelog
		return ret
	}

	return o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntry) GetChangeLogOk() (*Changelog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeLog, true
}

// SetChangeLog sets field value
func (o *ConnectionRouteTableEntry) SetChangeLog(v Changelog) {
	o.ChangeLog = v
}

func (o ConnectionRouteTableEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionRouteTableEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	toSerialize["state"] = o.State
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.NextHop) {
		toSerialize["nextHop"] = o.NextHop
	}
	if !IsNil(o.MED) {
		toSerialize["MED"] = o.MED
	}
	if !IsNil(o.LocalPreference) {
		toSerialize["localPreference"] = o.LocalPreference
	}
	if !IsNil(o.AsPath) {
		toSerialize["asPath"] = o.AsPath
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["changeLog"] = o.ChangeLog

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionRouteTableEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"state",
		"changeLog",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varConnectionRouteTableEntry := _ConnectionRouteTableEntry{}

	err = json.Unmarshal(data, &varConnectionRouteTableEntry)

	if err != nil {
		return err
	}

	*o = ConnectionRouteTableEntry(varConnectionRouteTableEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "protocolType")
		delete(additionalProperties, "state")
		delete(additionalProperties, "age")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "nextHop")
		delete(additionalProperties, "MED")
		delete(additionalProperties, "localPreference")
		delete(additionalProperties, "asPath")
		delete(additionalProperties, "connection")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionRouteTableEntry struct {
	value *ConnectionRouteTableEntry
	isSet bool
}

func (v NullableConnectionRouteTableEntry) Get() *ConnectionRouteTableEntry {
	return v.value
}

func (v *NullableConnectionRouteTableEntry) Set(val *ConnectionRouteTableEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteTableEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteTableEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteTableEntry(val *ConnectionRouteTableEntry) *NullableConnectionRouteTableEntry {
	return &NullableConnectionRouteTableEntry{value: val, isSet: true}
}

func (v NullableConnectionRouteTableEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteTableEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

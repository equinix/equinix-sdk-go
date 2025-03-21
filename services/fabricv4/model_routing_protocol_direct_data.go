/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the RoutingProtocolDirectData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolDirectData{}

// RoutingProtocolDirectData struct for RoutingProtocolDirectData
type RoutingProtocolDirectData struct {
	Type       *RoutingProtocolDirectTypeType `json:"type,omitempty"`
	Name       *string                        `json:"name,omitempty"`
	DirectIpv4 *DirectConnectionIpv4          `json:"directIpv4,omitempty"`
	DirectIpv6 *DirectConnectionIpv6          `json:"directIpv6,omitempty"`
	// Routing Protocol URI
	Href *string `json:"href,omitempty"`
	// Routing protocol identifier
	Uuid                 *string                      `json:"uuid,omitempty"`
	State                *RoutingProtocolBGPDataState `json:"state,omitempty"`
	Operation            *RoutingProtocolOperation    `json:"operation,omitempty"`
	Change               *RoutingProtocolChange       `json:"change,omitempty"`
	Changelog            *Changelog                   `json:"changelog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolDirectData RoutingProtocolDirectData

// NewRoutingProtocolDirectData instantiates a new RoutingProtocolDirectData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolDirectData() *RoutingProtocolDirectData {
	this := RoutingProtocolDirectData{}
	return &this
}

// NewRoutingProtocolDirectDataWithDefaults instantiates a new RoutingProtocolDirectData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolDirectDataWithDefaults() *RoutingProtocolDirectData {
	this := RoutingProtocolDirectData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetType() RoutingProtocolDirectTypeType {
	if o == nil || IsNil(o.Type) {
		var ret RoutingProtocolDirectTypeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetTypeOk() (*RoutingProtocolDirectTypeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RoutingProtocolDirectTypeType and assigns it to the Type field.
func (o *RoutingProtocolDirectData) SetType(v RoutingProtocolDirectTypeType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoutingProtocolDirectData) SetName(v string) {
	o.Name = &v
}

// GetDirectIpv4 returns the DirectIpv4 field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetDirectIpv4() DirectConnectionIpv4 {
	if o == nil || IsNil(o.DirectIpv4) {
		var ret DirectConnectionIpv4
		return ret
	}
	return *o.DirectIpv4
}

// GetDirectIpv4Ok returns a tuple with the DirectIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetDirectIpv4Ok() (*DirectConnectionIpv4, bool) {
	if o == nil || IsNil(o.DirectIpv4) {
		return nil, false
	}
	return o.DirectIpv4, true
}

// HasDirectIpv4 returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasDirectIpv4() bool {
	if o != nil && !IsNil(o.DirectIpv4) {
		return true
	}

	return false
}

// SetDirectIpv4 gets a reference to the given DirectConnectionIpv4 and assigns it to the DirectIpv4 field.
func (o *RoutingProtocolDirectData) SetDirectIpv4(v DirectConnectionIpv4) {
	o.DirectIpv4 = &v
}

// GetDirectIpv6 returns the DirectIpv6 field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetDirectIpv6() DirectConnectionIpv6 {
	if o == nil || IsNil(o.DirectIpv6) {
		var ret DirectConnectionIpv6
		return ret
	}
	return *o.DirectIpv6
}

// GetDirectIpv6Ok returns a tuple with the DirectIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetDirectIpv6Ok() (*DirectConnectionIpv6, bool) {
	if o == nil || IsNil(o.DirectIpv6) {
		return nil, false
	}
	return o.DirectIpv6, true
}

// HasDirectIpv6 returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasDirectIpv6() bool {
	if o != nil && !IsNil(o.DirectIpv6) {
		return true
	}

	return false
}

// SetDirectIpv6 gets a reference to the given DirectConnectionIpv6 and assigns it to the DirectIpv6 field.
func (o *RoutingProtocolDirectData) SetDirectIpv6(v DirectConnectionIpv6) {
	o.DirectIpv6 = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RoutingProtocolDirectData) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RoutingProtocolDirectData) SetUuid(v string) {
	o.Uuid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetState() RoutingProtocolBGPDataState {
	if o == nil || IsNil(o.State) {
		var ret RoutingProtocolBGPDataState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetStateOk() (*RoutingProtocolBGPDataState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given RoutingProtocolBGPDataState and assigns it to the State field.
func (o *RoutingProtocolDirectData) SetState(v RoutingProtocolBGPDataState) {
	o.State = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetOperation() RoutingProtocolOperation {
	if o == nil || IsNil(o.Operation) {
		var ret RoutingProtocolOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetOperationOk() (*RoutingProtocolOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given RoutingProtocolOperation and assigns it to the Operation field.
func (o *RoutingProtocolDirectData) SetOperation(v RoutingProtocolOperation) {
	o.Operation = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetChange() RoutingProtocolChange {
	if o == nil || IsNil(o.Change) {
		var ret RoutingProtocolChange
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetChangeOk() (*RoutingProtocolChange, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given RoutingProtocolChange and assigns it to the Change field.
func (o *RoutingProtocolDirectData) SetChange(v RoutingProtocolChange) {
	o.Change = &v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise.
func (o *RoutingProtocolDirectData) GetChangelog() Changelog {
	if o == nil || IsNil(o.Changelog) {
		var ret Changelog
		return ret
	}
	return *o.Changelog
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolDirectData) GetChangelogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.Changelog) {
		return nil, false
	}
	return o.Changelog, true
}

// HasChangelog returns a boolean if a field has been set.
func (o *RoutingProtocolDirectData) HasChangelog() bool {
	if o != nil && !IsNil(o.Changelog) {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given Changelog and assigns it to the Changelog field.
func (o *RoutingProtocolDirectData) SetChangelog(v Changelog) {
	o.Changelog = &v
}

func (o RoutingProtocolDirectData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolDirectData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DirectIpv4) {
		toSerialize["directIpv4"] = o.DirectIpv4
	}
	if !IsNil(o.DirectIpv6) {
		toSerialize["directIpv6"] = o.DirectIpv6
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Changelog) {
		toSerialize["changelog"] = o.Changelog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolDirectData) UnmarshalJSON(data []byte) (err error) {
	varRoutingProtocolDirectData := _RoutingProtocolDirectData{}

	err = json.Unmarshal(data, &varRoutingProtocolDirectData)

	if err != nil {
		return err
	}

	*o = RoutingProtocolDirectData(varRoutingProtocolDirectData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "directIpv4")
		delete(additionalProperties, "directIpv6")
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "state")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "change")
		delete(additionalProperties, "changelog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolDirectData struct {
	value *RoutingProtocolDirectData
	isSet bool
}

func (v NullableRoutingProtocolDirectData) Get() *RoutingProtocolDirectData {
	return v.value
}

func (v *NullableRoutingProtocolDirectData) Set(val *RoutingProtocolDirectData) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolDirectData) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolDirectData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolDirectData(val *RoutingProtocolDirectData) *NullableRoutingProtocolDirectData {
	return &NullableRoutingProtocolDirectData{value: val, isSet: true}
}

func (v NullableRoutingProtocolDirectData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolDirectData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

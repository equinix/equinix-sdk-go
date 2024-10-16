/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the SimplifiedNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedNetwork{}

// SimplifiedNetwork Network specification
type SimplifiedNetwork struct {
	// Network URI
	Href *string `json:"href,omitempty"`
	// Equinix-assigned network identifier
	Uuid string `json:"uuid"`
	// Customer-assigned network name
	Name      *string                  `json:"name,omitempty"`
	State     *NetworkState            `json:"state,omitempty"`
	Account   *SimplifiedAccount       `json:"account,omitempty"`
	Change    *SimplifiedNetworkChange `json:"change,omitempty"`
	Operation *NetworkOperation        `json:"operation,omitempty"`
	ChangeLog *Changelog               `json:"changeLog,omitempty"`
	// Network sub-resources links
	Links                []Link              `json:"links,omitempty"`
	Type                 *NetworkType        `json:"type,omitempty"`
	Scope                *NetworkScope       `json:"scope,omitempty"`
	Location             *SimplifiedLocation `json:"location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedNetwork SimplifiedNetwork

// NewSimplifiedNetwork instantiates a new SimplifiedNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedNetwork(uuid string) *SimplifiedNetwork {
	this := SimplifiedNetwork{}
	this.Uuid = uuid
	return &this
}

// NewSimplifiedNetworkWithDefaults instantiates a new SimplifiedNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedNetworkWithDefaults() *SimplifiedNetwork {
	this := SimplifiedNetwork{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedNetwork) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value
func (o *SimplifiedNetwork) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SimplifiedNetwork) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimplifiedNetwork) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetState() NetworkState {
	if o == nil || IsNil(o.State) {
		var ret NetworkState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetStateOk() (*NetworkState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given NetworkState and assigns it to the State field.
func (o *SimplifiedNetwork) SetState(v NetworkState) {
	o.State = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *SimplifiedNetwork) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetChange() SimplifiedNetworkChange {
	if o == nil || IsNil(o.Change) {
		var ret SimplifiedNetworkChange
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetChangeOk() (*SimplifiedNetworkChange, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given SimplifiedNetworkChange and assigns it to the Change field.
func (o *SimplifiedNetwork) SetChange(v SimplifiedNetworkChange) {
	o.Change = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetOperation() NetworkOperation {
	if o == nil || IsNil(o.Operation) {
		var ret NetworkOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetOperationOk() (*NetworkOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NetworkOperation and assigns it to the Operation field.
func (o *SimplifiedNetwork) SetOperation(v NetworkOperation) {
	o.Operation = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetChangeLog() Changelog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret Changelog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetChangeLogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given Changelog and assigns it to the ChangeLog field.
func (o *SimplifiedNetwork) SetChangeLog(v Changelog) {
	o.ChangeLog = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SimplifiedNetwork) SetLinks(v []Link) {
	o.Links = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetType() NetworkType {
	if o == nil || IsNil(o.Type) {
		var ret NetworkType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetTypeOk() (*NetworkType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NetworkType and assigns it to the Type field.
func (o *SimplifiedNetwork) SetType(v NetworkType) {
	o.Type = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetScope() NetworkScope {
	if o == nil || IsNil(o.Scope) {
		var ret NetworkScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetScopeOk() (*NetworkScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given NetworkScope and assigns it to the Scope field.
func (o *SimplifiedNetwork) SetScope(v NetworkScope) {
	o.Scope = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SimplifiedNetwork) GetLocation() SimplifiedLocation {
	if o == nil || IsNil(o.Location) {
		var ret SimplifiedLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNetwork) GetLocationOk() (*SimplifiedLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SimplifiedNetwork) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SimplifiedLocation and assigns it to the Location field.
func (o *SimplifiedNetwork) SetLocation(v SimplifiedLocation) {
	o.Location = &v
}

func (o SimplifiedNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
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

	varSimplifiedNetwork := _SimplifiedNetwork{}

	err = json.Unmarshal(data, &varSimplifiedNetwork)

	if err != nil {
		return err
	}

	*o = SimplifiedNetwork(varSimplifiedNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "account")
		delete(additionalProperties, "change")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "changeLog")
		delete(additionalProperties, "links")
		delete(additionalProperties, "type")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedNetwork struct {
	value *SimplifiedNetwork
	isSet bool
}

func (v NullableSimplifiedNetwork) Get() *SimplifiedNetwork {
	return v.value
}

func (v *NullableSimplifiedNetwork) Set(val *SimplifiedNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedNetwork(val *SimplifiedNetwork) *NullableSimplifiedNetwork {
	return &NullableSimplifiedNetwork{value: val, isSet: true}
}

func (v NullableSimplifiedNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

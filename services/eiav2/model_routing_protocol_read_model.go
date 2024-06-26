/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutingProtocolReadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolReadModel{}

// RoutingProtocolReadModel struct for RoutingProtocolReadModel
type RoutingProtocolReadModel struct {
	// Routing protocol URI
	Href *string `json:"href,omitempty"`
	// Routing protocol identifier
	Uuid *string             `json:"uuid,omitempty"`
	Type RoutingProtocolType `json:"type"`
	// Name of the routing protocol instance.
	Name string `json:"name"`
	// Description of the routing protocol instance
	Description          *string              `json:"description,omitempty"`
	Ipv4                 *RoutingProtocolIpv4 `json:"ipv4,omitempty"`
	Ipv6                 *RoutingProtocolIpv6 `json:"ipv6,omitempty"`
	ChangeLog            ChangeLog            `json:"changeLog"`
	Tags                 []string             `json:"tags,omitempty"`
	Links                []Link               `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolReadModel RoutingProtocolReadModel

// NewRoutingProtocolReadModel instantiates a new RoutingProtocolReadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolReadModel(type_ RoutingProtocolType, name string, changeLog ChangeLog, links []Link) *RoutingProtocolReadModel {
	this := RoutingProtocolReadModel{}
	this.Type = type_
	this.Name = name
	this.ChangeLog = changeLog
	this.Links = links
	return &this
}

// NewRoutingProtocolReadModelWithDefaults instantiates a new RoutingProtocolReadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolReadModelWithDefaults() *RoutingProtocolReadModel {
	this := RoutingProtocolReadModel{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RoutingProtocolReadModel) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RoutingProtocolReadModel) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RoutingProtocolReadModel) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RoutingProtocolReadModel) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RoutingProtocolReadModel) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RoutingProtocolReadModel) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value
func (o *RoutingProtocolReadModel) GetType() RoutingProtocolType {
	if o == nil {
		var ret RoutingProtocolType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetTypeOk() (*RoutingProtocolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingProtocolReadModel) SetType(v RoutingProtocolType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *RoutingProtocolReadModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoutingProtocolReadModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoutingProtocolReadModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoutingProtocolReadModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoutingProtocolReadModel) SetDescription(v string) {
	o.Description = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *RoutingProtocolReadModel) GetIpv4() RoutingProtocolIpv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret RoutingProtocolIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetIpv4Ok() (*RoutingProtocolIpv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *RoutingProtocolReadModel) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given RoutingProtocolIpv4 and assigns it to the Ipv4 field.
func (o *RoutingProtocolReadModel) SetIpv4(v RoutingProtocolIpv4) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *RoutingProtocolReadModel) GetIpv6() RoutingProtocolIpv6 {
	if o == nil || IsNil(o.Ipv6) {
		var ret RoutingProtocolIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetIpv6Ok() (*RoutingProtocolIpv6, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *RoutingProtocolReadModel) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given RoutingProtocolIpv6 and assigns it to the Ipv6 field.
func (o *RoutingProtocolReadModel) SetIpv6(v RoutingProtocolIpv6) {
	o.Ipv6 = &v
}

// GetChangeLog returns the ChangeLog field value
func (o *RoutingProtocolReadModel) GetChangeLog() ChangeLog {
	if o == nil {
		var ret ChangeLog
		return ret
	}

	return o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetChangeLogOk() (*ChangeLog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeLog, true
}

// SetChangeLog sets field value
func (o *RoutingProtocolReadModel) SetChangeLog(v ChangeLog) {
	o.ChangeLog = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RoutingProtocolReadModel) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RoutingProtocolReadModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RoutingProtocolReadModel) SetTags(v []string) {
	o.Tags = v
}

// GetLinks returns the Links field value
func (o *RoutingProtocolReadModel) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolReadModel) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *RoutingProtocolReadModel) SetLinks(v []Link) {
	o.Links = v
}

func (o RoutingProtocolReadModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolReadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	toSerialize["changeLog"] = o.ChangeLog
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolReadModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"changeLog",
		"links",
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

	varRoutingProtocolReadModel := _RoutingProtocolReadModel{}

	err = json.Unmarshal(data, &varRoutingProtocolReadModel)

	if err != nil {
		return err
	}

	*o = RoutingProtocolReadModel(varRoutingProtocolReadModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "ipv6")
		delete(additionalProperties, "changeLog")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolReadModel struct {
	value *RoutingProtocolReadModel
	isSet bool
}

func (v NullableRoutingProtocolReadModel) Get() *RoutingProtocolReadModel {
	return v.value
}

func (v *NullableRoutingProtocolReadModel) Set(val *RoutingProtocolReadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolReadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolReadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolReadModel(val *RoutingProtocolReadModel) *NullableRoutingProtocolReadModel {
	return &NullableRoutingProtocolReadModel{value: val, isSet: true}
}

func (v NullableRoutingProtocolReadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolReadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

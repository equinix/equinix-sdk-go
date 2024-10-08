/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"time"
)

// checks if the Vrf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vrf{}

// Vrf struct for Vrf
type Vrf struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Optional field that can be set to describe the VRF
	Description *string `json:"description,omitempty"`
	// True if the VRF is being billed. Usage will start when the first VRF Virtual Circuit is active, and will only stop when the VRF has been deleted.
	Bill *bool `json:"bill,omitempty"`
	// Toggle to enable the dynamic bgp neighbors feature on the VRF
	BgpDynamicNeighborsEnabled *bool `json:"bgp_dynamic_neighbors_enabled,omitempty"`
	// Toggle to export the VRF route-map to the dynamic bgp neighbors
	BgpDynamicNeighborsExportRouteMap *bool `json:"bgp_dynamic_neighbors_export_route_map,omitempty"`
	// Toggle BFD on dynamic bgp neighbors sessions
	BgpDynamicNeighborsBfdEnabled *bool `json:"bgp_dynamic_neighbors_bfd_enabled,omitempty"`
	// A 4-byte ASN associated with the VRF.
	LocalAsn *int64 `json:"local_asn,omitempty"`
	// Virtual circuits that are in the VRF
	VirtualCircuits []VrfVirtualCircuit `json:"virtual_circuits,omitempty"`
	// A list of CIDR network addresses. Like [\"10.0.0.0/16\", \"2001:d78::/59\"].
	IpRanges             []string   `json:"ip_ranges,omitempty"`
	Project              *Project   `json:"project,omitempty"`
	Metro                *Metro     `json:"metro,omitempty"`
	CreatedBy            *User      `json:"created_by,omitempty"`
	Href                 *string    `json:"href,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	Tags                 []string   `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Vrf Vrf

// NewVrf instantiates a new Vrf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrf() *Vrf {
	this := Vrf{}
	var bill bool = false
	this.Bill = &bill
	return &this
}

// NewVrfWithDefaults instantiates a new Vrf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfWithDefaults() *Vrf {
	this := Vrf{}
	var bill bool = false
	this.Bill = &bill
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Vrf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Vrf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Vrf) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vrf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vrf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vrf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Vrf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Vrf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Vrf) SetDescription(v string) {
	o.Description = &v
}

// GetBill returns the Bill field value if set, zero value otherwise.
func (o *Vrf) GetBill() bool {
	if o == nil || IsNil(o.Bill) {
		var ret bool
		return ret
	}
	return *o.Bill
}

// GetBillOk returns a tuple with the Bill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetBillOk() (*bool, bool) {
	if o == nil || IsNil(o.Bill) {
		return nil, false
	}
	return o.Bill, true
}

// HasBill returns a boolean if a field has been set.
func (o *Vrf) HasBill() bool {
	if o != nil && !IsNil(o.Bill) {
		return true
	}

	return false
}

// SetBill gets a reference to the given bool and assigns it to the Bill field.
func (o *Vrf) SetBill(v bool) {
	o.Bill = &v
}

// GetBgpDynamicNeighborsEnabled returns the BgpDynamicNeighborsEnabled field value if set, zero value otherwise.
func (o *Vrf) GetBgpDynamicNeighborsEnabled() bool {
	if o == nil || IsNil(o.BgpDynamicNeighborsEnabled) {
		var ret bool
		return ret
	}
	return *o.BgpDynamicNeighborsEnabled
}

// GetBgpDynamicNeighborsEnabledOk returns a tuple with the BgpDynamicNeighborsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetBgpDynamicNeighborsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BgpDynamicNeighborsEnabled) {
		return nil, false
	}
	return o.BgpDynamicNeighborsEnabled, true
}

// HasBgpDynamicNeighborsEnabled returns a boolean if a field has been set.
func (o *Vrf) HasBgpDynamicNeighborsEnabled() bool {
	if o != nil && !IsNil(o.BgpDynamicNeighborsEnabled) {
		return true
	}

	return false
}

// SetBgpDynamicNeighborsEnabled gets a reference to the given bool and assigns it to the BgpDynamicNeighborsEnabled field.
func (o *Vrf) SetBgpDynamicNeighborsEnabled(v bool) {
	o.BgpDynamicNeighborsEnabled = &v
}

// GetBgpDynamicNeighborsExportRouteMap returns the BgpDynamicNeighborsExportRouteMap field value if set, zero value otherwise.
func (o *Vrf) GetBgpDynamicNeighborsExportRouteMap() bool {
	if o == nil || IsNil(o.BgpDynamicNeighborsExportRouteMap) {
		var ret bool
		return ret
	}
	return *o.BgpDynamicNeighborsExportRouteMap
}

// GetBgpDynamicNeighborsExportRouteMapOk returns a tuple with the BgpDynamicNeighborsExportRouteMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetBgpDynamicNeighborsExportRouteMapOk() (*bool, bool) {
	if o == nil || IsNil(o.BgpDynamicNeighborsExportRouteMap) {
		return nil, false
	}
	return o.BgpDynamicNeighborsExportRouteMap, true
}

// HasBgpDynamicNeighborsExportRouteMap returns a boolean if a field has been set.
func (o *Vrf) HasBgpDynamicNeighborsExportRouteMap() bool {
	if o != nil && !IsNil(o.BgpDynamicNeighborsExportRouteMap) {
		return true
	}

	return false
}

// SetBgpDynamicNeighborsExportRouteMap gets a reference to the given bool and assigns it to the BgpDynamicNeighborsExportRouteMap field.
func (o *Vrf) SetBgpDynamicNeighborsExportRouteMap(v bool) {
	o.BgpDynamicNeighborsExportRouteMap = &v
}

// GetBgpDynamicNeighborsBfdEnabled returns the BgpDynamicNeighborsBfdEnabled field value if set, zero value otherwise.
func (o *Vrf) GetBgpDynamicNeighborsBfdEnabled() bool {
	if o == nil || IsNil(o.BgpDynamicNeighborsBfdEnabled) {
		var ret bool
		return ret
	}
	return *o.BgpDynamicNeighborsBfdEnabled
}

// GetBgpDynamicNeighborsBfdEnabledOk returns a tuple with the BgpDynamicNeighborsBfdEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetBgpDynamicNeighborsBfdEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BgpDynamicNeighborsBfdEnabled) {
		return nil, false
	}
	return o.BgpDynamicNeighborsBfdEnabled, true
}

// HasBgpDynamicNeighborsBfdEnabled returns a boolean if a field has been set.
func (o *Vrf) HasBgpDynamicNeighborsBfdEnabled() bool {
	if o != nil && !IsNil(o.BgpDynamicNeighborsBfdEnabled) {
		return true
	}

	return false
}

// SetBgpDynamicNeighborsBfdEnabled gets a reference to the given bool and assigns it to the BgpDynamicNeighborsBfdEnabled field.
func (o *Vrf) SetBgpDynamicNeighborsBfdEnabled(v bool) {
	o.BgpDynamicNeighborsBfdEnabled = &v
}

// GetLocalAsn returns the LocalAsn field value if set, zero value otherwise.
func (o *Vrf) GetLocalAsn() int64 {
	if o == nil || IsNil(o.LocalAsn) {
		var ret int64
		return ret
	}
	return *o.LocalAsn
}

// GetLocalAsnOk returns a tuple with the LocalAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetLocalAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalAsn) {
		return nil, false
	}
	return o.LocalAsn, true
}

// HasLocalAsn returns a boolean if a field has been set.
func (o *Vrf) HasLocalAsn() bool {
	if o != nil && !IsNil(o.LocalAsn) {
		return true
	}

	return false
}

// SetLocalAsn gets a reference to the given int64 and assigns it to the LocalAsn field.
func (o *Vrf) SetLocalAsn(v int64) {
	o.LocalAsn = &v
}

// GetVirtualCircuits returns the VirtualCircuits field value if set, zero value otherwise.
func (o *Vrf) GetVirtualCircuits() []VrfVirtualCircuit {
	if o == nil || IsNil(o.VirtualCircuits) {
		var ret []VrfVirtualCircuit
		return ret
	}
	return o.VirtualCircuits
}

// GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetVirtualCircuitsOk() ([]VrfVirtualCircuit, bool) {
	if o == nil || IsNil(o.VirtualCircuits) {
		return nil, false
	}
	return o.VirtualCircuits, true
}

// HasVirtualCircuits returns a boolean if a field has been set.
func (o *Vrf) HasVirtualCircuits() bool {
	if o != nil && !IsNil(o.VirtualCircuits) {
		return true
	}

	return false
}

// SetVirtualCircuits gets a reference to the given []VrfVirtualCircuit and assigns it to the VirtualCircuits field.
func (o *Vrf) SetVirtualCircuits(v []VrfVirtualCircuit) {
	o.VirtualCircuits = v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *Vrf) GetIpRanges() []string {
	if o == nil || IsNil(o.IpRanges) {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetIpRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *Vrf) HasIpRanges() bool {
	if o != nil && !IsNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *Vrf) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Vrf) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Vrf) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *Vrf) SetProject(v Project) {
	o.Project = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *Vrf) GetMetro() Metro {
	if o == nil || IsNil(o.Metro) {
		var ret Metro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetMetroOk() (*Metro, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *Vrf) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given Metro and assigns it to the Metro field.
func (o *Vrf) SetMetro(v Metro) {
	o.Metro = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Vrf) GetCreatedBy() User {
	if o == nil || IsNil(o.CreatedBy) {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetCreatedByOk() (*User, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Vrf) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *Vrf) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Vrf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Vrf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Vrf) SetHref(v string) {
	o.Href = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Vrf) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Vrf) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Vrf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Vrf) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Vrf) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Vrf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Vrf) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vrf) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Vrf) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Vrf) SetTags(v []string) {
	o.Tags = v
}

func (o Vrf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vrf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Bill) {
		toSerialize["bill"] = o.Bill
	}
	if !IsNil(o.BgpDynamicNeighborsEnabled) {
		toSerialize["bgp_dynamic_neighbors_enabled"] = o.BgpDynamicNeighborsEnabled
	}
	if !IsNil(o.BgpDynamicNeighborsExportRouteMap) {
		toSerialize["bgp_dynamic_neighbors_export_route_map"] = o.BgpDynamicNeighborsExportRouteMap
	}
	if !IsNil(o.BgpDynamicNeighborsBfdEnabled) {
		toSerialize["bgp_dynamic_neighbors_bfd_enabled"] = o.BgpDynamicNeighborsBfdEnabled
	}
	if !IsNil(o.LocalAsn) {
		toSerialize["local_asn"] = o.LocalAsn
	}
	if !IsNil(o.VirtualCircuits) {
		toSerialize["virtual_circuits"] = o.VirtualCircuits
	}
	if !IsNil(o.IpRanges) {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Vrf) UnmarshalJSON(data []byte) (err error) {
	varVrf := _Vrf{}

	err = json.Unmarshal(data, &varVrf)

	if err != nil {
		return err
	}

	*o = Vrf(varVrf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "bill")
		delete(additionalProperties, "bgp_dynamic_neighbors_enabled")
		delete(additionalProperties, "bgp_dynamic_neighbors_export_route_map")
		delete(additionalProperties, "bgp_dynamic_neighbors_bfd_enabled")
		delete(additionalProperties, "local_asn")
		delete(additionalProperties, "virtual_circuits")
		delete(additionalProperties, "ip_ranges")
		delete(additionalProperties, "project")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "href")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrf struct {
	value *Vrf
	isSet bool
}

func (v NullableVrf) Get() *Vrf {
	return v.value
}

func (v *NullableVrf) Set(val *Vrf) {
	v.value = val
	v.isSet = true
}

func (v NullableVrf) IsSet() bool {
	return v.isSet
}

func (v *NullableVrf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrf(val *Vrf) *NullableVrf {
	return &NullableVrf{value: val, isSet: true}
}

func (v NullableVrf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

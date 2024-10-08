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

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project struct for Project
type Project struct {
	BackendTransferEnabled *bool                  `json:"backend_transfer_enabled,omitempty"`
	BgpConfig              *Href                  `json:"bgp_config,omitempty"`
	CreatedAt              *time.Time             `json:"created_at,omitempty"`
	Customdata             map[string]interface{} `json:"customdata,omitempty"`
	Devices                []Href                 `json:"devices,omitempty"`
	Href                   *string                `json:"href,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	Invitations            []Href                 `json:"invitations,omitempty"`
	MaxDevices             map[string]interface{} `json:"max_devices,omitempty"`
	Members                []Href                 `json:"members,omitempty"`
	Memberships            []Href                 `json:"memberships,omitempty"`
	// The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis.
	Name                 *string                `json:"name,omitempty"`
	NetworkStatus        map[string]interface{} `json:"network_status,omitempty"`
	Organization         *Organization          `json:"organization,omitempty"`
	PaymentMethod        *Href                  `json:"payment_method,omitempty"`
	SshKeys              []Href                 `json:"ssh_keys,omitempty"`
	UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
	Url                  *string                `json:"url,omitempty"`
	Volumes              []Href                 `json:"volumes,omitempty"`
	Type                 *ProjectType           `json:"type,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Project Project

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject() *Project {
	this := Project{}
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetBackendTransferEnabled returns the BackendTransferEnabled field value if set, zero value otherwise.
func (o *Project) GetBackendTransferEnabled() bool {
	if o == nil || IsNil(o.BackendTransferEnabled) {
		var ret bool
		return ret
	}
	return *o.BackendTransferEnabled
}

// GetBackendTransferEnabledOk returns a tuple with the BackendTransferEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetBackendTransferEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackendTransferEnabled) {
		return nil, false
	}
	return o.BackendTransferEnabled, true
}

// HasBackendTransferEnabled returns a boolean if a field has been set.
func (o *Project) HasBackendTransferEnabled() bool {
	if o != nil && !IsNil(o.BackendTransferEnabled) {
		return true
	}

	return false
}

// SetBackendTransferEnabled gets a reference to the given bool and assigns it to the BackendTransferEnabled field.
func (o *Project) SetBackendTransferEnabled(v bool) {
	o.BackendTransferEnabled = &v
}

// GetBgpConfig returns the BgpConfig field value if set, zero value otherwise.
func (o *Project) GetBgpConfig() Href {
	if o == nil || IsNil(o.BgpConfig) {
		var ret Href
		return ret
	}
	return *o.BgpConfig
}

// GetBgpConfigOk returns a tuple with the BgpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetBgpConfigOk() (*Href, bool) {
	if o == nil || IsNil(o.BgpConfig) {
		return nil, false
	}
	return o.BgpConfig, true
}

// HasBgpConfig returns a boolean if a field has been set.
func (o *Project) HasBgpConfig() bool {
	if o != nil && !IsNil(o.BgpConfig) {
		return true
	}

	return false
}

// SetBgpConfig gets a reference to the given Href and assigns it to the BgpConfig field.
func (o *Project) SetBgpConfig(v Href) {
	o.BgpConfig = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Project) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Project) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *Project) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *Project) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *Project) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *Project) GetDevices() []Href {
	if o == nil || IsNil(o.Devices) {
		var ret []Href
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDevicesOk() ([]Href, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *Project) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []Href and assigns it to the Devices field.
func (o *Project) SetDevices(v []Href) {
	o.Devices = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Project) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Project) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Project) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Project) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Project) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Project) SetId(v string) {
	o.Id = &v
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *Project) GetInvitations() []Href {
	if o == nil || IsNil(o.Invitations) {
		var ret []Href
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetInvitationsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Invitations) {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *Project) HasInvitations() bool {
	if o != nil && !IsNil(o.Invitations) {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []Href and assigns it to the Invitations field.
func (o *Project) SetInvitations(v []Href) {
	o.Invitations = v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *Project) GetMaxDevices() map[string]interface{} {
	if o == nil || IsNil(o.MaxDevices) {
		var ret map[string]interface{}
		return ret
	}
	return o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMaxDevicesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MaxDevices) {
		return map[string]interface{}{}, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *Project) HasMaxDevices() bool {
	if o != nil && !IsNil(o.MaxDevices) {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given map[string]interface{} and assigns it to the MaxDevices field.
func (o *Project) SetMaxDevices(v map[string]interface{}) {
	o.MaxDevices = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Project) GetMembers() []Href {
	if o == nil || IsNil(o.Members) {
		var ret []Href
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMembersOk() ([]Href, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Project) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Href and assigns it to the Members field.
func (o *Project) SetMembers(v []Href) {
	o.Members = v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *Project) GetMemberships() []Href {
	if o == nil || IsNil(o.Memberships) {
		var ret []Href
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetMembershipsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *Project) HasMemberships() bool {
	if o != nil && !IsNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []Href and assigns it to the Memberships field.
func (o *Project) SetMemberships(v []Href) {
	o.Memberships = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Project) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Project) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Project) SetName(v string) {
	o.Name = &v
}

// GetNetworkStatus returns the NetworkStatus field value if set, zero value otherwise.
func (o *Project) GetNetworkStatus() map[string]interface{} {
	if o == nil || IsNil(o.NetworkStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetNetworkStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NetworkStatus) {
		return map[string]interface{}{}, false
	}
	return o.NetworkStatus, true
}

// HasNetworkStatus returns a boolean if a field has been set.
func (o *Project) HasNetworkStatus() bool {
	if o != nil && !IsNil(o.NetworkStatus) {
		return true
	}

	return false
}

// SetNetworkStatus gets a reference to the given map[string]interface{} and assigns it to the NetworkStatus field.
func (o *Project) SetNetworkStatus(v map[string]interface{}) {
	o.NetworkStatus = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Project) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Project) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *Project) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *Project) GetPaymentMethod() Href {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret Href
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetPaymentMethodOk() (*Href, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *Project) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given Href and assigns it to the PaymentMethod field.
func (o *Project) SetPaymentMethod(v Href) {
	o.PaymentMethod = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Project) GetSshKeys() []Href {
	if o == nil || IsNil(o.SshKeys) {
		var ret []Href
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSshKeysOk() ([]Href, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Project) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []Href and assigns it to the SshKeys field.
func (o *Project) SetSshKeys(v []Href) {
	o.SshKeys = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Project) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Project) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Project) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Project) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Project) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Project) SetUrl(v string) {
	o.Url = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *Project) GetVolumes() []Href {
	if o == nil || IsNil(o.Volumes) {
		var ret []Href
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetVolumesOk() ([]Href, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *Project) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Href and assigns it to the Volumes field.
func (o *Project) SetVolumes(v []Href) {
	o.Volumes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Project) GetType() ProjectType {
	if o == nil || IsNil(o.Type) {
		var ret ProjectType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetTypeOk() (*ProjectType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Project) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProjectType and assigns it to the Type field.
func (o *Project) SetType(v ProjectType) {
	o.Type = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Project) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Project) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Project) SetTags(v []string) {
	o.Tags = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackendTransferEnabled) {
		toSerialize["backend_transfer_enabled"] = o.BackendTransferEnabled
	}
	if !IsNil(o.BgpConfig) {
		toSerialize["bgp_config"] = o.BgpConfig
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Invitations) {
		toSerialize["invitations"] = o.Invitations
	}
	if !IsNil(o.MaxDevices) {
		toSerialize["max_devices"] = o.MaxDevices
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkStatus) {
		toSerialize["network_status"] = o.NetworkStatus
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Project) UnmarshalJSON(data []byte) (err error) {
	varProject := _Project{}

	err = json.Unmarshal(data, &varProject)

	if err != nil {
		return err
	}

	*o = Project(varProject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "backend_transfer_enabled")
		delete(additionalProperties, "bgp_config")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "devices")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invitations")
		delete(additionalProperties, "max_devices")
		delete(additionalProperties, "members")
		delete(additionalProperties, "memberships")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network_status")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "ssh_keys")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "url")
		delete(additionalProperties, "volumes")
		delete(additionalProperties, "type")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

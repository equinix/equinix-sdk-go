/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"os"
	"time"
)

// OrganizationInput struct for OrganizationInput
type OrganizationInput struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Website *string `json:"website,omitempty"`
	Twitter *string `json:"twitter,omitempty"`
	Logo **os.File `json:"logo,omitempty"`
	Address *Address `json:"address,omitempty"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
	// Force to all members to have enabled the two factor authentication after that date, unless the value is null
	Enforce2faAt *time.Time `json:"enforce_2fa_at,omitempty"`
}

// NewOrganizationInput instantiates a new OrganizationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInput() *OrganizationInput {
	this := OrganizationInput{}
	return &this
}

// NewOrganizationInputWithDefaults instantiates a new OrganizationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInputWithDefaults() *OrganizationInput {
	this := OrganizationInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationInput) SetDescription(v string) {
	o.Description = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *OrganizationInput) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *OrganizationInput) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *OrganizationInput) SetWebsite(v string) {
	o.Website = &v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *OrganizationInput) GetTwitter() string {
	if o == nil || o.Twitter == nil {
		var ret string
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetTwitterOk() (*string, bool) {
	if o == nil || o.Twitter == nil {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *OrganizationInput) HasTwitter() bool {
	if o != nil && o.Twitter != nil {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given string and assigns it to the Twitter field.
func (o *OrganizationInput) SetTwitter(v string) {
	o.Twitter = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *OrganizationInput) GetLogo() *os.File {
	if o == nil || o.Logo == nil {
		var ret *os.File
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetLogoOk() (**os.File, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *OrganizationInput) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given *os.File and assigns it to the Logo field.
func (o *OrganizationInput) SetLogo(v *os.File) {
	o.Logo = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationInput) GetAddress() Address {
	if o == nil || o.Address == nil {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetAddressOk() (*Address, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationInput) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *OrganizationInput) SetAddress(v Address) {
	o.Address = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *OrganizationInput) GetBillingAddress() Address {
	if o == nil || o.BillingAddress == nil {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetBillingAddressOk() (*Address, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *OrganizationInput) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *OrganizationInput) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *OrganizationInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *OrganizationInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *OrganizationInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

// GetEnforce2faAt returns the Enforce2faAt field value if set, zero value otherwise.
func (o *OrganizationInput) GetEnforce2faAt() time.Time {
	if o == nil || o.Enforce2faAt == nil {
		var ret time.Time
		return ret
	}
	return *o.Enforce2faAt
}

// GetEnforce2faAtOk returns a tuple with the Enforce2faAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetEnforce2faAtOk() (*time.Time, bool) {
	if o == nil || o.Enforce2faAt == nil {
		return nil, false
	}
	return o.Enforce2faAt, true
}

// HasEnforce2faAt returns a boolean if a field has been set.
func (o *OrganizationInput) HasEnforce2faAt() bool {
	if o != nil && o.Enforce2faAt != nil {
		return true
	}

	return false
}

// SetEnforce2faAt gets a reference to the given time.Time and assigns it to the Enforce2faAt field.
func (o *OrganizationInput) SetEnforce2faAt(v time.Time) {
	o.Enforce2faAt = &v
}

func (o OrganizationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Twitter != nil {
		toSerialize["twitter"] = o.Twitter
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.Enforce2faAt != nil {
		toSerialize["enforce_2fa_at"] = o.Enforce2faAt
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationInput struct {
	value *OrganizationInput
	isSet bool
}

func (v NullableOrganizationInput) Get() *OrganizationInput {
	return v.value
}

func (v *NullableOrganizationInput) Set(val *OrganizationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInput(val *OrganizationInput) *NullableOrganizationInput {
	return &NullableOrganizationInput{value: val, isSet: true}
}

func (v NullableOrganizationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


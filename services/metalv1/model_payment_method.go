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

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethod{}

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	BillingAddress       *PaymentMethodBillingAddress `json:"billing_address,omitempty"`
	CardType             *string                      `json:"card_type,omitempty"`
	CardholderName       *string                      `json:"cardholder_name,omitempty"`
	CreatedAt            *time.Time                   `json:"created_at,omitempty"`
	CreatedByUser        *Href                        `json:"created_by_user,omitempty"`
	Default              *bool                        `json:"default,omitempty"`
	Email                *string                      `json:"email,omitempty"`
	ExpirationMonth      *string                      `json:"expiration_month,omitempty"`
	ExpirationYear       *string                      `json:"expiration_year,omitempty"`
	Id                   *string                      `json:"id,omitempty"`
	Name                 *string                      `json:"name,omitempty"`
	Organization         *Href                        `json:"organization,omitempty"`
	Projects             []Href                       `json:"projects,omitempty"`
	Type                 *string                      `json:"type,omitempty"`
	UpdatedAt            *time.Time                   `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethod PaymentMethod

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *PaymentMethod) GetBillingAddress() PaymentMethodBillingAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret PaymentMethodBillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBillingAddressOk() (*PaymentMethodBillingAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PaymentMethod) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given PaymentMethodBillingAddress and assigns it to the BillingAddress field.
func (o *PaymentMethod) SetBillingAddress(v PaymentMethodBillingAddress) {
	o.BillingAddress = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *PaymentMethod) SetCardType(v string) {
	o.CardType = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise.
func (o *PaymentMethod) GetCardholderName() string {
	if o == nil || IsNil(o.CardholderName) {
		var ret string
		return ret
	}
	return *o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardholderNameOk() (*string, bool) {
	if o == nil || IsNil(o.CardholderName) {
		return nil, false
	}
	return o.CardholderName, true
}

// HasCardholderName returns a boolean if a field has been set.
func (o *PaymentMethod) HasCardholderName() bool {
	if o != nil && !IsNil(o.CardholderName) {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given string and assigns it to the CardholderName field.
func (o *PaymentMethod) SetCardholderName(v string) {
	o.CardholderName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentMethod) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreatedByUser() Href {
	if o == nil || IsNil(o.CreatedByUser) {
		var ret Href
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedByUserOk() (*Href, bool) {
	if o == nil || IsNil(o.CreatedByUser) {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreatedByUser() bool {
	if o != nil && !IsNil(o.CreatedByUser) {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given Href and assigns it to the CreatedByUser field.
func (o *PaymentMethod) SetCreatedByUser(v Href) {
	o.CreatedByUser = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethod) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethod) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethod) SetDefault(v bool) {
	o.Default = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PaymentMethod) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PaymentMethod) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PaymentMethod) SetEmail(v string) {
	o.Email = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PaymentMethod) GetExpirationMonth() string {
	if o == nil || IsNil(o.ExpirationMonth) {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetExpirationMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationMonth) {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PaymentMethod) HasExpirationMonth() bool {
	if o != nil && !IsNil(o.ExpirationMonth) {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PaymentMethod) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PaymentMethod) GetExpirationYear() string {
	if o == nil || IsNil(o.ExpirationYear) {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetExpirationYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationYear) {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PaymentMethod) HasExpirationYear() bool {
	if o != nil && !IsNil(o.ExpirationYear) {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *PaymentMethod) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethod) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethod) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethod) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethod) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PaymentMethod) GetOrganization() Href {
	if o == nil || IsNil(o.Organization) {
		var ret Href
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetOrganizationOk() (*Href, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PaymentMethod) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Href and assigns it to the Organization field.
func (o *PaymentMethod) SetOrganization(v Href) {
	o.Organization = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PaymentMethod) GetProjects() []Href {
	if o == nil || IsNil(o.Projects) {
		var ret []Href
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetProjectsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PaymentMethod) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Href and assigns it to the Projects field.
func (o *PaymentMethod) SetProjects(v []Href) {
	o.Projects = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PaymentMethod) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.CardholderName) {
		toSerialize["cardholder_name"] = o.CardholderName
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedByUser) {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExpirationMonth) {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if !IsNil(o.ExpirationYear) {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethod) UnmarshalJSON(data []byte) (err error) {
	varPaymentMethod := _PaymentMethod{}

	err = json.Unmarshal(data, &varPaymentMethod)

	if err != nil {
		return err
	}

	*o = PaymentMethod(varPaymentMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_address")
		delete(additionalProperties, "card_type")
		delete(additionalProperties, "cardholder_name")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by_user")
		delete(additionalProperties, "default")
		delete(additionalProperties, "email")
		delete(additionalProperties, "expiration_month")
		delete(additionalProperties, "expiration_year")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "projects")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

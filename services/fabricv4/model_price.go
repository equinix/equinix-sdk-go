/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price struct for Price
type Price struct {
	// An absolute URL that returns specified pricing data
	Href *string      `json:"href,omitempty"`
	Type *ProductType `json:"type,omitempty"`
	// Equinix-assigned product code
	Code *string `json:"code,omitempty"`
	// Full product name
	Name *string `json:"name,omitempty"`
	// Product description
	Description *string            `json:"description,omitempty"`
	Account     *SimplifiedAccount `json:"account,omitempty"`
	Charges     []PriceCharge      `json:"charges,omitempty"`
	// Product offering price currency
	Currency             *string                 `json:"currency,omitempty"`
	TermLength           *PriceTermLength        `json:"termLength,omitempty"`
	Catgory              *PriceCategory          `json:"catgory,omitempty"`
	Connection           *VirtualConnectionPrice `json:"connection,omitempty"`
	IpBlock              *IpBlockPrice           `json:"ipBlock,omitempty"`
	Router               *FabricCloudRouterPrice `json:"router,omitempty"`
	Port                 *VirtualPortPrice       `json:"port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Price Price

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice() *Price {
	this := Price{}
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Price) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Price) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Price) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Price) GetType() ProductType {
	if o == nil || IsNil(o.Type) {
		var ret ProductType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetTypeOk() (*ProductType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Price) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProductType and assigns it to the Type field.
func (o *Price) SetType(v ProductType) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Price) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Price) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Price) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Price) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Price) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Price) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Price) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Price) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Price) SetDescription(v string) {
	o.Description = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Price) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Price) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *Price) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *Price) GetCharges() []PriceCharge {
	if o == nil || IsNil(o.Charges) {
		var ret []PriceCharge
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetChargesOk() ([]PriceCharge, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *Price) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []PriceCharge and assigns it to the Charges field.
func (o *Price) SetCharges(v []PriceCharge) {
	o.Charges = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Price) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Price) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Price) SetCurrency(v string) {
	o.Currency = &v
}

// GetTermLength returns the TermLength field value if set, zero value otherwise.
func (o *Price) GetTermLength() PriceTermLength {
	if o == nil || IsNil(o.TermLength) {
		var ret PriceTermLength
		return ret
	}
	return *o.TermLength
}

// GetTermLengthOk returns a tuple with the TermLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetTermLengthOk() (*PriceTermLength, bool) {
	if o == nil || IsNil(o.TermLength) {
		return nil, false
	}
	return o.TermLength, true
}

// HasTermLength returns a boolean if a field has been set.
func (o *Price) HasTermLength() bool {
	if o != nil && !IsNil(o.TermLength) {
		return true
	}

	return false
}

// SetTermLength gets a reference to the given PriceTermLength and assigns it to the TermLength field.
func (o *Price) SetTermLength(v PriceTermLength) {
	o.TermLength = &v
}

// GetCatgory returns the Catgory field value if set, zero value otherwise.
func (o *Price) GetCatgory() PriceCategory {
	if o == nil || IsNil(o.Catgory) {
		var ret PriceCategory
		return ret
	}
	return *o.Catgory
}

// GetCatgoryOk returns a tuple with the Catgory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCatgoryOk() (*PriceCategory, bool) {
	if o == nil || IsNil(o.Catgory) {
		return nil, false
	}
	return o.Catgory, true
}

// HasCatgory returns a boolean if a field has been set.
func (o *Price) HasCatgory() bool {
	if o != nil && !IsNil(o.Catgory) {
		return true
	}

	return false
}

// SetCatgory gets a reference to the given PriceCategory and assigns it to the Catgory field.
func (o *Price) SetCatgory(v PriceCategory) {
	o.Catgory = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *Price) GetConnection() VirtualConnectionPrice {
	if o == nil || IsNil(o.Connection) {
		var ret VirtualConnectionPrice
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetConnectionOk() (*VirtualConnectionPrice, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *Price) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given VirtualConnectionPrice and assigns it to the Connection field.
func (o *Price) SetConnection(v VirtualConnectionPrice) {
	o.Connection = &v
}

// GetIpBlock returns the IpBlock field value if set, zero value otherwise.
func (o *Price) GetIpBlock() IpBlockPrice {
	if o == nil || IsNil(o.IpBlock) {
		var ret IpBlockPrice
		return ret
	}
	return *o.IpBlock
}

// GetIpBlockOk returns a tuple with the IpBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetIpBlockOk() (*IpBlockPrice, bool) {
	if o == nil || IsNil(o.IpBlock) {
		return nil, false
	}
	return o.IpBlock, true
}

// HasIpBlock returns a boolean if a field has been set.
func (o *Price) HasIpBlock() bool {
	if o != nil && !IsNil(o.IpBlock) {
		return true
	}

	return false
}

// SetIpBlock gets a reference to the given IpBlockPrice and assigns it to the IpBlock field.
func (o *Price) SetIpBlock(v IpBlockPrice) {
	o.IpBlock = &v
}

// GetRouter returns the Router field value if set, zero value otherwise.
func (o *Price) GetRouter() FabricCloudRouterPrice {
	if o == nil || IsNil(o.Router) {
		var ret FabricCloudRouterPrice
		return ret
	}
	return *o.Router
}

// GetRouterOk returns a tuple with the Router field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetRouterOk() (*FabricCloudRouterPrice, bool) {
	if o == nil || IsNil(o.Router) {
		return nil, false
	}
	return o.Router, true
}

// HasRouter returns a boolean if a field has been set.
func (o *Price) HasRouter() bool {
	if o != nil && !IsNil(o.Router) {
		return true
	}

	return false
}

// SetRouter gets a reference to the given FabricCloudRouterPrice and assigns it to the Router field.
func (o *Price) SetRouter(v FabricCloudRouterPrice) {
	o.Router = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Price) GetPort() VirtualPortPrice {
	if o == nil || IsNil(o.Port) {
		var ret VirtualPortPrice
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetPortOk() (*VirtualPortPrice, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Price) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given VirtualPortPrice and assigns it to the Port field.
func (o *Price) SetPort(v VirtualPortPrice) {
	o.Port = &v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.TermLength) {
		toSerialize["termLength"] = o.TermLength
	}
	if !IsNil(o.Catgory) {
		toSerialize["catgory"] = o.Catgory
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.IpBlock) {
		toSerialize["ipBlock"] = o.IpBlock
	}
	if !IsNil(o.Router) {
		toSerialize["router"] = o.Router
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Price) UnmarshalJSON(data []byte) (err error) {
	varPrice := _Price{}

	err = json.Unmarshal(data, &varPrice)

	if err != nil {
		return err
	}

	*o = Price(varPrice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "code")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "account")
		delete(additionalProperties, "charges")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "termLength")
		delete(additionalProperties, "catgory")
		delete(additionalProperties, "connection")
		delete(additionalProperties, "ipBlock")
		delete(additionalProperties, "router")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
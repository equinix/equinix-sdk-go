/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterActionsSearchExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterActionsSearchExpression{}

// CloudRouterActionsSearchExpression struct for CloudRouterActionsSearchExpression
type CloudRouterActionsSearchExpression struct {
	// Possible field names to use on filters:  * `/type` - type of update  * `/state` - action state  * `/connection/uuid` - connection uuid associated  * `/_*` - all-category search
	Property *string `json:"property,omitempty"`
	// Possible operators to use on filters:  * `=` - equal  * `!=` - not equal  * `>` - greater than  * `>=` - greater than or equal to  * `<` - less than  * `<=` - less than or equal to  * `[NOT] BETWEEN` - (not) between  * `[NOT] LIKE` - (not) like  * `[NOT] IN` - (not) in  * `~*` - case-insensitive like
	Operator             *string  `json:"operator,omitempty"`
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterActionsSearchExpression CloudRouterActionsSearchExpression

// NewCloudRouterActionsSearchExpression instantiates a new CloudRouterActionsSearchExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterActionsSearchExpression() *CloudRouterActionsSearchExpression {
	this := CloudRouterActionsSearchExpression{}
	return &this
}

// NewCloudRouterActionsSearchExpressionWithDefaults instantiates a new CloudRouterActionsSearchExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterActionsSearchExpressionWithDefaults() *CloudRouterActionsSearchExpression {
	this := CloudRouterActionsSearchExpression{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *CloudRouterActionsSearchExpression) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterActionsSearchExpression) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *CloudRouterActionsSearchExpression) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *CloudRouterActionsSearchExpression) SetProperty(v string) {
	o.Property = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *CloudRouterActionsSearchExpression) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterActionsSearchExpression) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *CloudRouterActionsSearchExpression) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *CloudRouterActionsSearchExpression) SetOperator(v string) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *CloudRouterActionsSearchExpression) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterActionsSearchExpression) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *CloudRouterActionsSearchExpression) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *CloudRouterActionsSearchExpression) SetValues(v []string) {
	o.Values = v
}

func (o CloudRouterActionsSearchExpression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterActionsSearchExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterActionsSearchExpression) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterActionsSearchExpression := _CloudRouterActionsSearchExpression{}

	err = json.Unmarshal(data, &varCloudRouterActionsSearchExpression)

	if err != nil {
		return err
	}

	*o = CloudRouterActionsSearchExpression(varCloudRouterActionsSearchExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "property")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterActionsSearchExpression struct {
	value *CloudRouterActionsSearchExpression
	isSet bool
}

func (v NullableCloudRouterActionsSearchExpression) Get() *CloudRouterActionsSearchExpression {
	return v.value
}

func (v *NullableCloudRouterActionsSearchExpression) Set(val *CloudRouterActionsSearchExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterActionsSearchExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterActionsSearchExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterActionsSearchExpression(val *CloudRouterActionsSearchExpression) *NullableCloudRouterActionsSearchExpression {
	return &NullableCloudRouterActionsSearchExpression{value: val, isSet: true}
}

func (v NullableCloudRouterActionsSearchExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterActionsSearchExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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
	"fmt"
)

// RouteFilterRulesDataAction the model 'RouteFilterRulesDataAction'
type RouteFilterRulesDataAction string

// List of RouteFilterRulesData_action
const (
	ROUTEFILTERRULESDATAACTION_PERMIT RouteFilterRulesDataAction = "PERMIT"
	ROUTEFILTERRULESDATAACTION_DENY   RouteFilterRulesDataAction = "DENY"
)

// All allowed values of RouteFilterRulesDataAction enum
var AllowedRouteFilterRulesDataActionEnumValues = []RouteFilterRulesDataAction{
	"PERMIT",
	"DENY",
}

func (v *RouteFilterRulesDataAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteFilterRulesDataAction(value)
	for _, existing := range AllowedRouteFilterRulesDataActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteFilterRulesDataAction", value)
}

// NewRouteFilterRulesDataActionFromValue returns a pointer to a valid RouteFilterRulesDataAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteFilterRulesDataActionFromValue(v string) (*RouteFilterRulesDataAction, error) {
	ev := RouteFilterRulesDataAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteFilterRulesDataAction: valid values are %v", v, AllowedRouteFilterRulesDataActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteFilterRulesDataAction) IsValid() bool {
	for _, existing := range AllowedRouteFilterRulesDataActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteFilterRulesData_action value
func (v RouteFilterRulesDataAction) Ptr() *RouteFilterRulesDataAction {
	return &v
}

type NullableRouteFilterRulesDataAction struct {
	value *RouteFilterRulesDataAction
	isSet bool
}

func (v NullableRouteFilterRulesDataAction) Get() *RouteFilterRulesDataAction {
	return v.value
}

func (v *NullableRouteFilterRulesDataAction) Set(val *RouteFilterRulesDataAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesDataAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesDataAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesDataAction(val *RouteFilterRulesDataAction) *NullableRouteFilterRulesDataAction {
	return &NullableRouteFilterRulesDataAction{value: val, isSet: true}
}

func (v NullableRouteFilterRulesDataAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesDataAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
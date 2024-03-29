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

// RouteFilterRulesChangeOperationType type of filter rule
type RouteFilterRulesChangeOperationType string

// List of RouteFilterRulesChangeOperation_type
const (
	ROUTEFILTERRULESCHANGEOPERATIONTYPE_IPV4_PREFIX_FILTER_RULE_UPDATE   RouteFilterRulesChangeOperationType = "BGP_IPv4_PREFIX_FILTER_RULE_UPDATE"
	ROUTEFILTERRULESCHANGEOPERATIONTYPE_IPV4_PREFIX_FILTER_RULE_CREATION RouteFilterRulesChangeOperationType = "BGP_IPv4_PREFIX_FILTER_RULE_CREATION"
	ROUTEFILTERRULESCHANGEOPERATIONTYPE_IPV4_PREFIX_FILTER_RULE_DELETION RouteFilterRulesChangeOperationType = "BGP_IPv4_PREFIX_FILTER_RULE_DELETION"
	ROUTEFILTERRULESCHANGEOPERATIONTYPE_IPV6_PREFIX_FILTER_RULE_UPDATE   RouteFilterRulesChangeOperationType = "BGP_IPv6_PREFIX_FILTER_RULE_UPDATE"
	ROUTEFILTERRULESCHANGEOPERATIONTYPE_IPV6_PREFIX_FILTER_RULE_CREATION RouteFilterRulesChangeOperationType = "BGP_IPv6_PREFIX_FILTER_RULE_CREATION"
	ROUTEFILTERRULESCHANGEOPERATIONTYPE_IPV6_PREFIX_FILTER_RULE_DELETION RouteFilterRulesChangeOperationType = "BGP_IPv6_PREFIX_FILTER_RULE_DELETION"
)

// All allowed values of RouteFilterRulesChangeOperationType enum
var AllowedRouteFilterRulesChangeOperationTypeEnumValues = []RouteFilterRulesChangeOperationType{
	"BGP_IPv4_PREFIX_FILTER_RULE_UPDATE",
	"BGP_IPv4_PREFIX_FILTER_RULE_CREATION",
	"BGP_IPv4_PREFIX_FILTER_RULE_DELETION",
	"BGP_IPv6_PREFIX_FILTER_RULE_UPDATE",
	"BGP_IPv6_PREFIX_FILTER_RULE_CREATION",
	"BGP_IPv6_PREFIX_FILTER_RULE_DELETION",
}

func (v *RouteFilterRulesChangeOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteFilterRulesChangeOperationType(value)
	for _, existing := range AllowedRouteFilterRulesChangeOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteFilterRulesChangeOperationType", value)
}

// NewRouteFilterRulesChangeOperationTypeFromValue returns a pointer to a valid RouteFilterRulesChangeOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteFilterRulesChangeOperationTypeFromValue(v string) (*RouteFilterRulesChangeOperationType, error) {
	ev := RouteFilterRulesChangeOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteFilterRulesChangeOperationType: valid values are %v", v, AllowedRouteFilterRulesChangeOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteFilterRulesChangeOperationType) IsValid() bool {
	for _, existing := range AllowedRouteFilterRulesChangeOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteFilterRulesChangeOperation_type value
func (v RouteFilterRulesChangeOperationType) Ptr() *RouteFilterRulesChangeOperationType {
	return &v
}

type NullableRouteFilterRulesChangeOperationType struct {
	value *RouteFilterRulesChangeOperationType
	isSet bool
}

func (v NullableRouteFilterRulesChangeOperationType) Get() *RouteFilterRulesChangeOperationType {
	return v.value
}

func (v *NullableRouteFilterRulesChangeOperationType) Set(val *RouteFilterRulesChangeOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesChangeOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesChangeOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesChangeOperationType(val *RouteFilterRulesChangeOperationType) *NullableRouteFilterRulesChangeOperationType {
	return &NullableRouteFilterRulesChangeOperationType{value: val, isSet: true}
}

func (v NullableRouteFilterRulesChangeOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesChangeOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

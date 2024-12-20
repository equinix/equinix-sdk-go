/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// RouteFiltersSearchFilterItemProperty the model 'RouteFiltersSearchFilterItemProperty'
type RouteFiltersSearchFilterItemProperty string

// List of RouteFiltersSearchFilterItem_property
const (
	ROUTEFILTERSSEARCHFILTERITEMPROPERTY_TYPE               RouteFiltersSearchFilterItemProperty = "/type"
	ROUTEFILTERSSEARCHFILTERITEMPROPERTY_NAME               RouteFiltersSearchFilterItemProperty = "/name"
	ROUTEFILTERSSEARCHFILTERITEMPROPERTY_PROJECT_PROJECT_ID RouteFiltersSearchFilterItemProperty = "/project/projectId"
	ROUTEFILTERSSEARCHFILTERITEMPROPERTY_UUID               RouteFiltersSearchFilterItemProperty = "/uuid"
	ROUTEFILTERSSEARCHFILTERITEMPROPERTY_STATE              RouteFiltersSearchFilterItemProperty = "/state"
)

// All allowed values of RouteFiltersSearchFilterItemProperty enum
var AllowedRouteFiltersSearchFilterItemPropertyEnumValues = []RouteFiltersSearchFilterItemProperty{
	"/type",
	"/name",
	"/project/projectId",
	"/uuid",
	"/state",
}

func (v *RouteFiltersSearchFilterItemProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteFiltersSearchFilterItemProperty(value)
	for _, existing := range AllowedRouteFiltersSearchFilterItemPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteFiltersSearchFilterItemProperty", value)
}

// NewRouteFiltersSearchFilterItemPropertyFromValue returns a pointer to a valid RouteFiltersSearchFilterItemProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteFiltersSearchFilterItemPropertyFromValue(v string) (*RouteFiltersSearchFilterItemProperty, error) {
	ev := RouteFiltersSearchFilterItemProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteFiltersSearchFilterItemProperty: valid values are %v", v, AllowedRouteFiltersSearchFilterItemPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteFiltersSearchFilterItemProperty) IsValid() bool {
	for _, existing := range AllowedRouteFiltersSearchFilterItemPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteFiltersSearchFilterItem_property value
func (v RouteFiltersSearchFilterItemProperty) Ptr() *RouteFiltersSearchFilterItemProperty {
	return &v
}

type NullableRouteFiltersSearchFilterItemProperty struct {
	value *RouteFiltersSearchFilterItemProperty
	isSet bool
}

func (v NullableRouteFiltersSearchFilterItemProperty) Get() *RouteFiltersSearchFilterItemProperty {
	return v.value
}

func (v *NullableRouteFiltersSearchFilterItemProperty) Set(val *RouteFiltersSearchFilterItemProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFiltersSearchFilterItemProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFiltersSearchFilterItemProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFiltersSearchFilterItemProperty(val *RouteFiltersSearchFilterItemProperty) *NullableRouteFiltersSearchFilterItemProperty {
	return &NullableRouteFiltersSearchFilterItemProperty{value: val, isSet: true}
}

func (v NullableRouteFiltersSearchFilterItemProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFiltersSearchFilterItemProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

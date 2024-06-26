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

// PurchaseOrderType Purchase order type
type PurchaseOrderType string

// List of PurchaseOrderType
const (
	PURCHASEORDERTYPE_STANDARD_PURCHASE_ORDER PurchaseOrderType = "STANDARD_PURCHASE_ORDER"
	PURCHASEORDERTYPE_BLANKET_PURCHASE_ORDER  PurchaseOrderType = "BLANKET_PURCHASE_ORDER"
)

// All allowed values of PurchaseOrderType enum
var AllowedPurchaseOrderTypeEnumValues = []PurchaseOrderType{
	"STANDARD_PURCHASE_ORDER",
	"BLANKET_PURCHASE_ORDER",
}

func (v *PurchaseOrderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PurchaseOrderType(value)
	for _, existing := range AllowedPurchaseOrderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PurchaseOrderType", value)
}

// NewPurchaseOrderTypeFromValue returns a pointer to a valid PurchaseOrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPurchaseOrderTypeFromValue(v string) (*PurchaseOrderType, error) {
	ev := PurchaseOrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PurchaseOrderType: valid values are %v", v, AllowedPurchaseOrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PurchaseOrderType) IsValid() bool {
	for _, existing := range AllowedPurchaseOrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PurchaseOrderType value
func (v PurchaseOrderType) Ptr() *PurchaseOrderType {
	return &v
}

type NullablePurchaseOrderType struct {
	value *PurchaseOrderType
	isSet bool
}

func (v NullablePurchaseOrderType) Get() *PurchaseOrderType {
	return v.value
}

func (v *NullablePurchaseOrderType) Set(val *PurchaseOrderType) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderType) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderType(val *PurchaseOrderType) *NullablePurchaseOrderType {
	return &NullablePurchaseOrderType{value: val, isSet: true}
}

func (v NullablePurchaseOrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

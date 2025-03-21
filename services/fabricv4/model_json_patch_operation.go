/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// JsonPatchOperation - Service Profile Access Points
type JsonPatchOperation struct {
	AddOperation     *AddOperation
	RemoveOperation  *RemoveOperation
	ReplaceOperation *ReplaceOperation
}

// AddOperationAsJsonPatchOperation is a convenience function that returns AddOperation wrapped in JsonPatchOperation
func AddOperationAsJsonPatchOperation(v *AddOperation) JsonPatchOperation {
	return JsonPatchOperation{
		AddOperation: v,
	}
}

// RemoveOperationAsJsonPatchOperation is a convenience function that returns RemoveOperation wrapped in JsonPatchOperation
func RemoveOperationAsJsonPatchOperation(v *RemoveOperation) JsonPatchOperation {
	return JsonPatchOperation{
		RemoveOperation: v,
	}
}

// ReplaceOperationAsJsonPatchOperation is a convenience function that returns ReplaceOperation wrapped in JsonPatchOperation
func ReplaceOperationAsJsonPatchOperation(v *ReplaceOperation) JsonPatchOperation {
	return JsonPatchOperation{
		ReplaceOperation: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *JsonPatchOperation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddOperation
	err = newStrictDecoder(data).Decode(&dst.AddOperation)
	if err == nil {
		jsonAddOperation, _ := json.Marshal(dst.AddOperation)
		if string(jsonAddOperation) == "{}" { // empty struct
			dst.AddOperation = nil
		} else {
			if err = validator.Validate(dst.AddOperation); err != nil {
				dst.AddOperation = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddOperation = nil
	}

	// try to unmarshal data into RemoveOperation
	err = newStrictDecoder(data).Decode(&dst.RemoveOperation)
	if err == nil {
		jsonRemoveOperation, _ := json.Marshal(dst.RemoveOperation)
		if string(jsonRemoveOperation) == "{}" { // empty struct
			dst.RemoveOperation = nil
		} else {
			if err = validator.Validate(dst.RemoveOperation); err != nil {
				dst.RemoveOperation = nil
			} else {
				match++
			}
		}
	} else {
		dst.RemoveOperation = nil
	}

	// try to unmarshal data into ReplaceOperation
	err = newStrictDecoder(data).Decode(&dst.ReplaceOperation)
	if err == nil {
		jsonReplaceOperation, _ := json.Marshal(dst.ReplaceOperation)
		if string(jsonReplaceOperation) == "{}" { // empty struct
			dst.ReplaceOperation = nil
		} else {
			if err = validator.Validate(dst.ReplaceOperation); err != nil {
				dst.ReplaceOperation = nil
			} else {
				match++
			}
		}
	} else {
		dst.ReplaceOperation = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddOperation = nil
		dst.RemoveOperation = nil
		dst.ReplaceOperation = nil

		return fmt.Errorf("data matches more than one schema in oneOf(JsonPatchOperation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(JsonPatchOperation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src JsonPatchOperation) MarshalJSON() ([]byte, error) {
	if src.AddOperation != nil {
		return json.Marshal(&src.AddOperation)
	}

	if src.RemoveOperation != nil {
		return json.Marshal(&src.RemoveOperation)
	}

	if src.ReplaceOperation != nil {
		return json.Marshal(&src.ReplaceOperation)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *JsonPatchOperation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddOperation != nil {
		return obj.AddOperation
	}

	if obj.RemoveOperation != nil {
		return obj.RemoveOperation
	}

	if obj.ReplaceOperation != nil {
		return obj.ReplaceOperation
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj JsonPatchOperation) GetActualInstanceValue() interface{} {
	if obj.AddOperation != nil {
		return *obj.AddOperation
	}

	if obj.RemoveOperation != nil {
		return *obj.RemoveOperation
	}

	if obj.ReplaceOperation != nil {
		return *obj.ReplaceOperation
	}

	// all schemas are nil
	return nil
}

type NullableJsonPatchOperation struct {
	value *JsonPatchOperation
	isSet bool
}

func (v NullableJsonPatchOperation) Get() *JsonPatchOperation {
	return v.value
}

func (v *NullableJsonPatchOperation) Set(val *JsonPatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchOperation(val *JsonPatchOperation) *NullableJsonPatchOperation {
	return &NullableJsonPatchOperation{value: val, isSet: true}
}

func (v NullableJsonPatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

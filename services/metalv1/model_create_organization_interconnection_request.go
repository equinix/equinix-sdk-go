/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// CreateOrganizationInterconnectionRequest - struct for CreateOrganizationInterconnectionRequest
type CreateOrganizationInterconnectionRequest struct {
	DedicatedPortCreateInput     *DedicatedPortCreateInput
	SharedPortVCVlanCreateInput  *SharedPortVCVlanCreateInput
	VlanCSPConnectionCreateInput *VlanCSPConnectionCreateInput
	VlanFabricVcCreateInput      *VlanFabricVcCreateInput
	VrfFabricVcCreateInput       *VrfFabricVcCreateInput
}

// DedicatedPortCreateInputAsCreateOrganizationInterconnectionRequest is a convenience function that returns DedicatedPortCreateInput wrapped in CreateOrganizationInterconnectionRequest
func DedicatedPortCreateInputAsCreateOrganizationInterconnectionRequest(v *DedicatedPortCreateInput) CreateOrganizationInterconnectionRequest {
	return CreateOrganizationInterconnectionRequest{
		DedicatedPortCreateInput: v,
	}
}

// SharedPortVCVlanCreateInputAsCreateOrganizationInterconnectionRequest is a convenience function that returns SharedPortVCVlanCreateInput wrapped in CreateOrganizationInterconnectionRequest
func SharedPortVCVlanCreateInputAsCreateOrganizationInterconnectionRequest(v *SharedPortVCVlanCreateInput) CreateOrganizationInterconnectionRequest {
	return CreateOrganizationInterconnectionRequest{
		SharedPortVCVlanCreateInput: v,
	}
}

// VlanCSPConnectionCreateInputAsCreateOrganizationInterconnectionRequest is a convenience function that returns VlanCSPConnectionCreateInput wrapped in CreateOrganizationInterconnectionRequest
func VlanCSPConnectionCreateInputAsCreateOrganizationInterconnectionRequest(v *VlanCSPConnectionCreateInput) CreateOrganizationInterconnectionRequest {
	return CreateOrganizationInterconnectionRequest{
		VlanCSPConnectionCreateInput: v,
	}
}

// VlanFabricVcCreateInputAsCreateOrganizationInterconnectionRequest is a convenience function that returns VlanFabricVcCreateInput wrapped in CreateOrganizationInterconnectionRequest
func VlanFabricVcCreateInputAsCreateOrganizationInterconnectionRequest(v *VlanFabricVcCreateInput) CreateOrganizationInterconnectionRequest {
	return CreateOrganizationInterconnectionRequest{
		VlanFabricVcCreateInput: v,
	}
}

// VrfFabricVcCreateInputAsCreateOrganizationInterconnectionRequest is a convenience function that returns VrfFabricVcCreateInput wrapped in CreateOrganizationInterconnectionRequest
func VrfFabricVcCreateInputAsCreateOrganizationInterconnectionRequest(v *VrfFabricVcCreateInput) CreateOrganizationInterconnectionRequest {
	return CreateOrganizationInterconnectionRequest{
		VrfFabricVcCreateInput: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateOrganizationInterconnectionRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DedicatedPortCreateInput
	err = newStrictDecoder(data).Decode(&dst.DedicatedPortCreateInput)
	if err == nil {
		jsonDedicatedPortCreateInput, _ := json.Marshal(dst.DedicatedPortCreateInput)
		if string(jsonDedicatedPortCreateInput) == "{}" { // empty struct
			dst.DedicatedPortCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.DedicatedPortCreateInput = nil
	}

	// try to unmarshal data into SharedPortVCVlanCreateInput
	err = newStrictDecoder(data).Decode(&dst.SharedPortVCVlanCreateInput)
	if err == nil {
		jsonSharedPortVCVlanCreateInput, _ := json.Marshal(dst.SharedPortVCVlanCreateInput)
		if string(jsonSharedPortVCVlanCreateInput) == "{}" { // empty struct
			dst.SharedPortVCVlanCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.SharedPortVCVlanCreateInput = nil
	}

	// try to unmarshal data into VlanCSPConnectionCreateInput
	err = newStrictDecoder(data).Decode(&dst.VlanCSPConnectionCreateInput)
	if err == nil {
		jsonVlanCSPConnectionCreateInput, _ := json.Marshal(dst.VlanCSPConnectionCreateInput)
		if string(jsonVlanCSPConnectionCreateInput) == "{}" { // empty struct
			dst.VlanCSPConnectionCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.VlanCSPConnectionCreateInput = nil
	}

	// try to unmarshal data into VlanFabricVcCreateInput
	err = newStrictDecoder(data).Decode(&dst.VlanFabricVcCreateInput)
	if err == nil {
		jsonVlanFabricVcCreateInput, _ := json.Marshal(dst.VlanFabricVcCreateInput)
		if string(jsonVlanFabricVcCreateInput) == "{}" { // empty struct
			dst.VlanFabricVcCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.VlanFabricVcCreateInput = nil
	}

	// try to unmarshal data into VrfFabricVcCreateInput
	err = newStrictDecoder(data).Decode(&dst.VrfFabricVcCreateInput)
	if err == nil {
		jsonVrfFabricVcCreateInput, _ := json.Marshal(dst.VrfFabricVcCreateInput)
		if string(jsonVrfFabricVcCreateInput) == "{}" { // empty struct
			dst.VrfFabricVcCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.VrfFabricVcCreateInput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DedicatedPortCreateInput = nil
		dst.SharedPortVCVlanCreateInput = nil
		dst.VlanCSPConnectionCreateInput = nil
		dst.VlanFabricVcCreateInput = nil
		dst.VrfFabricVcCreateInput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateOrganizationInterconnectionRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateOrganizationInterconnectionRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateOrganizationInterconnectionRequest) MarshalJSON() ([]byte, error) {
	if src.DedicatedPortCreateInput != nil {
		return json.Marshal(&src.DedicatedPortCreateInput)
	}

	if src.SharedPortVCVlanCreateInput != nil {
		return json.Marshal(&src.SharedPortVCVlanCreateInput)
	}

	if src.VlanCSPConnectionCreateInput != nil {
		return json.Marshal(&src.VlanCSPConnectionCreateInput)
	}

	if src.VlanFabricVcCreateInput != nil {
		return json.Marshal(&src.VlanFabricVcCreateInput)
	}

	if src.VrfFabricVcCreateInput != nil {
		return json.Marshal(&src.VrfFabricVcCreateInput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateOrganizationInterconnectionRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DedicatedPortCreateInput != nil {
		return obj.DedicatedPortCreateInput
	}

	if obj.SharedPortVCVlanCreateInput != nil {
		return obj.SharedPortVCVlanCreateInput
	}

	if obj.VlanCSPConnectionCreateInput != nil {
		return obj.VlanCSPConnectionCreateInput
	}

	if obj.VlanFabricVcCreateInput != nil {
		return obj.VlanFabricVcCreateInput
	}

	if obj.VrfFabricVcCreateInput != nil {
		return obj.VrfFabricVcCreateInput
	}

	// all schemas are nil
	return nil
}

type NullableCreateOrganizationInterconnectionRequest struct {
	value *CreateOrganizationInterconnectionRequest
	isSet bool
}

func (v NullableCreateOrganizationInterconnectionRequest) Get() *CreateOrganizationInterconnectionRequest {
	return v.value
}

func (v *NullableCreateOrganizationInterconnectionRequest) Set(val *CreateOrganizationInterconnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInterconnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInterconnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInterconnectionRequest(val *CreateOrganizationInterconnectionRequest) *NullableCreateOrganizationInterconnectionRequest {
	return &NullableCreateOrganizationInterconnectionRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInterconnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInterconnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

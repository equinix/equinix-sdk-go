/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterCommand{}

// CloudRouterCommand Get Fabric Cloud Router Command response object
type CloudRouterCommand struct {
	Href *string                 `json:"href,omitempty"`
	Type *CloudRouterCommandType `json:"type,omitempty"`
	Uuid *string                 `json:"uuid,omitempty"`
	// Customer-provided Cloud Router name
	Name                 *string                     `json:"name,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	State                *CloudRouterCommandState    `json:"state,omitempty"`
	Project              *Project                    `json:"project,omitempty"`
	Request              *CloudRouterCommandRequest  `json:"request,omitempty"`
	Response             *CloudRouterCommandResponse `json:"response,omitempty"`
	ChangeLog            *Changelog                  `json:"changeLog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterCommand CloudRouterCommand

// NewCloudRouterCommand instantiates a new CloudRouterCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterCommand() *CloudRouterCommand {
	this := CloudRouterCommand{}
	return &this
}

// NewCloudRouterCommandWithDefaults instantiates a new CloudRouterCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterCommandWithDefaults() *CloudRouterCommand {
	this := CloudRouterCommand{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CloudRouterCommand) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetType() CloudRouterCommandType {
	if o == nil || IsNil(o.Type) {
		var ret CloudRouterCommandType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetTypeOk() (*CloudRouterCommandType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CloudRouterCommandType and assigns it to the Type field.
func (o *CloudRouterCommand) SetType(v CloudRouterCommandType) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CloudRouterCommand) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudRouterCommand) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudRouterCommand) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetState() CloudRouterCommandState {
	if o == nil || IsNil(o.State) {
		var ret CloudRouterCommandState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetStateOk() (*CloudRouterCommandState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CloudRouterCommandState and assigns it to the State field.
func (o *CloudRouterCommand) SetState(v CloudRouterCommandState) {
	o.State = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *CloudRouterCommand) SetProject(v Project) {
	o.Project = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetRequest() CloudRouterCommandRequest {
	if o == nil || IsNil(o.Request) {
		var ret CloudRouterCommandRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetRequestOk() (*CloudRouterCommandRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given CloudRouterCommandRequest and assigns it to the Request field.
func (o *CloudRouterCommand) SetRequest(v CloudRouterCommandRequest) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetResponse() CloudRouterCommandResponse {
	if o == nil || IsNil(o.Response) {
		var ret CloudRouterCommandResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetResponseOk() (*CloudRouterCommandResponse, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given CloudRouterCommandResponse and assigns it to the Response field.
func (o *CloudRouterCommand) SetResponse(v CloudRouterCommandResponse) {
	o.Response = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *CloudRouterCommand) GetChangeLog() Changelog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret Changelog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommand) GetChangeLogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *CloudRouterCommand) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given Changelog and assigns it to the ChangeLog field.
func (o *CloudRouterCommand) SetChangeLog(v Changelog) {
	o.ChangeLog = &v
}

func (o CloudRouterCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterCommand) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterCommand := _CloudRouterCommand{}

	err = json.Unmarshal(data, &varCloudRouterCommand)

	if err != nil {
		return err
	}

	*o = CloudRouterCommand(varCloudRouterCommand)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "project")
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterCommand struct {
	value *CloudRouterCommand
	isSet bool
}

func (v NullableCloudRouterCommand) Get() *CloudRouterCommand {
	return v.value
}

func (v *NullableCloudRouterCommand) Set(val *CloudRouterCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterCommand(val *CloudRouterCommand) *NullableCloudRouterCommand {
	return &NullableCloudRouterCommand{value: val, isSet: true}
}

func (v NullableCloudRouterCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

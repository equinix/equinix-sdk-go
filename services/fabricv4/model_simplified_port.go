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

// checks if the SimplifiedPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedPort{}

// SimplifiedPort Port specification
type SimplifiedPort struct {
	Type *PortType `json:"type,omitempty"`
	// Equinix assigned response attribute for Port Id
	Id *int32 `json:"id,omitempty"`
	// Equinix assigned response attribute for an absolute URL that is the subject of the link's context.
	Href *string `json:"href,omitempty"`
	// Equinix assigned response attribute for  port identifier
	Uuid *string `json:"uuid,omitempty"`
	// Equinix assigned response attribute for Port name
	Name *string `json:"name,omitempty"`
	// Equinix assigned response attribute for Port description
	Description *string `json:"description,omitempty"`
	// Physical Ports Speed in Mbps
	PhysicalPortsSpeed *int32 `json:"physicalPortsSpeed,omitempty"`
	// Equinix assigned response attribute for Connection count
	ConnectionsCount *int32     `json:"connectionsCount,omitempty"`
	Project          *Project   `json:"project,omitempty"`
	State            *PortState `json:"state,omitempty"`
	// Equinix assigned response attribute for Unique ID for a virtual port.
	CvpId       *string            `json:"cvpId,omitempty"`
	Operation   *PortOperation     `json:"operation,omitempty"`
	Account     *SimplifiedAccount `json:"account,omitempty"`
	ServiceType *PortServiceType   `json:"serviceType,omitempty"`
	// Equinix assigned response attribute for Port bandwidth in Mbps
	Bandwidth *int32 `json:"bandwidth,omitempty"`
	// Equinix assigned response attribute for Port available bandwidth in Mbps
	AvailableBandwidth *int32 `json:"availableBandwidth,omitempty"`
	// Equinix assigned response attribute for Port used bandwidth in Mbps
	UsedBandwidth    *int32                `json:"usedBandwidth,omitempty"`
	Location         *SimplifiedLocation   `json:"location,omitempty"`
	Device           *PortDevice           `json:"device,omitempty"`
	Interface        *PortInterface        `json:"interface,omitempty"`
	Tether           *PortTether           `json:"tether,omitempty"`
	DemarcationPoint *PortDemarcationPoint `json:"demarcationPoint,omitempty"`
	Redundancy       *PortRedundancy       `json:"redundancy,omitempty"`
	Encapsulation    *PortEncapsulation    `json:"encapsulation,omitempty"`
	// If LAG enabled
	LagEnabled *bool         `json:"lagEnabled,omitempty"`
	Settings   *PortSettings `json:"settings,omitempty"`
	// Number of physical ports
	PhysicalPortQuantity *int32 `json:"physicalPortQuantity,omitempty"`
	// Port additional information
	AdditionalInfo []PortAdditionalInfo `json:"additionalInfo,omitempty"`
	// Physical ports that implement this port
	PhysicalPorts        []PhysicalPort `json:"physicalPorts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedPort SimplifiedPort

// NewSimplifiedPort instantiates a new SimplifiedPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedPort() *SimplifiedPort {
	this := SimplifiedPort{}
	return &this
}

// NewSimplifiedPortWithDefaults instantiates a new SimplifiedPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedPortWithDefaults() *SimplifiedPort {
	this := SimplifiedPort{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimplifiedPort) GetType() PortType {
	if o == nil || IsNil(o.Type) {
		var ret PortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetTypeOk() (*PortType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimplifiedPort) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PortType and assigns it to the Type field.
func (o *SimplifiedPort) SetType(v PortType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimplifiedPort) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimplifiedPort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SimplifiedPort) SetId(v int32) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedPort) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedPort) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedPort) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SimplifiedPort) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SimplifiedPort) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SimplifiedPort) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimplifiedPort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimplifiedPort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimplifiedPort) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SimplifiedPort) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SimplifiedPort) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SimplifiedPort) SetDescription(v string) {
	o.Description = &v
}

// GetPhysicalPortsSpeed returns the PhysicalPortsSpeed field value if set, zero value otherwise.
func (o *SimplifiedPort) GetPhysicalPortsSpeed() int32 {
	if o == nil || IsNil(o.PhysicalPortsSpeed) {
		var ret int32
		return ret
	}
	return *o.PhysicalPortsSpeed
}

// GetPhysicalPortsSpeedOk returns a tuple with the PhysicalPortsSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetPhysicalPortsSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.PhysicalPortsSpeed) {
		return nil, false
	}
	return o.PhysicalPortsSpeed, true
}

// HasPhysicalPortsSpeed returns a boolean if a field has been set.
func (o *SimplifiedPort) HasPhysicalPortsSpeed() bool {
	if o != nil && !IsNil(o.PhysicalPortsSpeed) {
		return true
	}

	return false
}

// SetPhysicalPortsSpeed gets a reference to the given int32 and assigns it to the PhysicalPortsSpeed field.
func (o *SimplifiedPort) SetPhysicalPortsSpeed(v int32) {
	o.PhysicalPortsSpeed = &v
}

// GetConnectionsCount returns the ConnectionsCount field value if set, zero value otherwise.
func (o *SimplifiedPort) GetConnectionsCount() int32 {
	if o == nil || IsNil(o.ConnectionsCount) {
		var ret int32
		return ret
	}
	return *o.ConnectionsCount
}

// GetConnectionsCountOk returns a tuple with the ConnectionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetConnectionsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectionsCount) {
		return nil, false
	}
	return o.ConnectionsCount, true
}

// HasConnectionsCount returns a boolean if a field has been set.
func (o *SimplifiedPort) HasConnectionsCount() bool {
	if o != nil && !IsNil(o.ConnectionsCount) {
		return true
	}

	return false
}

// SetConnectionsCount gets a reference to the given int32 and assigns it to the ConnectionsCount field.
func (o *SimplifiedPort) SetConnectionsCount(v int32) {
	o.ConnectionsCount = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *SimplifiedPort) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *SimplifiedPort) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *SimplifiedPort) SetProject(v Project) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SimplifiedPort) GetState() PortState {
	if o == nil || IsNil(o.State) {
		var ret PortState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetStateOk() (*PortState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SimplifiedPort) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given PortState and assigns it to the State field.
func (o *SimplifiedPort) SetState(v PortState) {
	o.State = &v
}

// GetCvpId returns the CvpId field value if set, zero value otherwise.
func (o *SimplifiedPort) GetCvpId() string {
	if o == nil || IsNil(o.CvpId) {
		var ret string
		return ret
	}
	return *o.CvpId
}

// GetCvpIdOk returns a tuple with the CvpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetCvpIdOk() (*string, bool) {
	if o == nil || IsNil(o.CvpId) {
		return nil, false
	}
	return o.CvpId, true
}

// HasCvpId returns a boolean if a field has been set.
func (o *SimplifiedPort) HasCvpId() bool {
	if o != nil && !IsNil(o.CvpId) {
		return true
	}

	return false
}

// SetCvpId gets a reference to the given string and assigns it to the CvpId field.
func (o *SimplifiedPort) SetCvpId(v string) {
	o.CvpId = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SimplifiedPort) GetOperation() PortOperation {
	if o == nil || IsNil(o.Operation) {
		var ret PortOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetOperationOk() (*PortOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SimplifiedPort) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given PortOperation and assigns it to the Operation field.
func (o *SimplifiedPort) SetOperation(v PortOperation) {
	o.Operation = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SimplifiedPort) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SimplifiedPort) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *SimplifiedPort) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *SimplifiedPort) GetServiceType() PortServiceType {
	if o == nil || IsNil(o.ServiceType) {
		var ret PortServiceType
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetServiceTypeOk() (*PortServiceType, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *SimplifiedPort) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given PortServiceType and assigns it to the ServiceType field.
func (o *SimplifiedPort) SetServiceType(v PortServiceType) {
	o.ServiceType = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *SimplifiedPort) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *SimplifiedPort) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *SimplifiedPort) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

// GetAvailableBandwidth returns the AvailableBandwidth field value if set, zero value otherwise.
func (o *SimplifiedPort) GetAvailableBandwidth() int32 {
	if o == nil || IsNil(o.AvailableBandwidth) {
		var ret int32
		return ret
	}
	return *o.AvailableBandwidth
}

// GetAvailableBandwidthOk returns a tuple with the AvailableBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetAvailableBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableBandwidth) {
		return nil, false
	}
	return o.AvailableBandwidth, true
}

// HasAvailableBandwidth returns a boolean if a field has been set.
func (o *SimplifiedPort) HasAvailableBandwidth() bool {
	if o != nil && !IsNil(o.AvailableBandwidth) {
		return true
	}

	return false
}

// SetAvailableBandwidth gets a reference to the given int32 and assigns it to the AvailableBandwidth field.
func (o *SimplifiedPort) SetAvailableBandwidth(v int32) {
	o.AvailableBandwidth = &v
}

// GetUsedBandwidth returns the UsedBandwidth field value if set, zero value otherwise.
func (o *SimplifiedPort) GetUsedBandwidth() int32 {
	if o == nil || IsNil(o.UsedBandwidth) {
		var ret int32
		return ret
	}
	return *o.UsedBandwidth
}

// GetUsedBandwidthOk returns a tuple with the UsedBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetUsedBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.UsedBandwidth) {
		return nil, false
	}
	return o.UsedBandwidth, true
}

// HasUsedBandwidth returns a boolean if a field has been set.
func (o *SimplifiedPort) HasUsedBandwidth() bool {
	if o != nil && !IsNil(o.UsedBandwidth) {
		return true
	}

	return false
}

// SetUsedBandwidth gets a reference to the given int32 and assigns it to the UsedBandwidth field.
func (o *SimplifiedPort) SetUsedBandwidth(v int32) {
	o.UsedBandwidth = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SimplifiedPort) GetLocation() SimplifiedLocation {
	if o == nil || IsNil(o.Location) {
		var ret SimplifiedLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetLocationOk() (*SimplifiedLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SimplifiedPort) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SimplifiedLocation and assigns it to the Location field.
func (o *SimplifiedPort) SetLocation(v SimplifiedLocation) {
	o.Location = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SimplifiedPort) GetDevice() PortDevice {
	if o == nil || IsNil(o.Device) {
		var ret PortDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetDeviceOk() (*PortDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SimplifiedPort) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given PortDevice and assigns it to the Device field.
func (o *SimplifiedPort) SetDevice(v PortDevice) {
	o.Device = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *SimplifiedPort) GetInterface() PortInterface {
	if o == nil || IsNil(o.Interface) {
		var ret PortInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetInterfaceOk() (*PortInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *SimplifiedPort) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given PortInterface and assigns it to the Interface field.
func (o *SimplifiedPort) SetInterface(v PortInterface) {
	o.Interface = &v
}

// GetTether returns the Tether field value if set, zero value otherwise.
func (o *SimplifiedPort) GetTether() PortTether {
	if o == nil || IsNil(o.Tether) {
		var ret PortTether
		return ret
	}
	return *o.Tether
}

// GetTetherOk returns a tuple with the Tether field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetTetherOk() (*PortTether, bool) {
	if o == nil || IsNil(o.Tether) {
		return nil, false
	}
	return o.Tether, true
}

// HasTether returns a boolean if a field has been set.
func (o *SimplifiedPort) HasTether() bool {
	if o != nil && !IsNil(o.Tether) {
		return true
	}

	return false
}

// SetTether gets a reference to the given PortTether and assigns it to the Tether field.
func (o *SimplifiedPort) SetTether(v PortTether) {
	o.Tether = &v
}

// GetDemarcationPoint returns the DemarcationPoint field value if set, zero value otherwise.
func (o *SimplifiedPort) GetDemarcationPoint() PortDemarcationPoint {
	if o == nil || IsNil(o.DemarcationPoint) {
		var ret PortDemarcationPoint
		return ret
	}
	return *o.DemarcationPoint
}

// GetDemarcationPointOk returns a tuple with the DemarcationPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetDemarcationPointOk() (*PortDemarcationPoint, bool) {
	if o == nil || IsNil(o.DemarcationPoint) {
		return nil, false
	}
	return o.DemarcationPoint, true
}

// HasDemarcationPoint returns a boolean if a field has been set.
func (o *SimplifiedPort) HasDemarcationPoint() bool {
	if o != nil && !IsNil(o.DemarcationPoint) {
		return true
	}

	return false
}

// SetDemarcationPoint gets a reference to the given PortDemarcationPoint and assigns it to the DemarcationPoint field.
func (o *SimplifiedPort) SetDemarcationPoint(v PortDemarcationPoint) {
	o.DemarcationPoint = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *SimplifiedPort) GetRedundancy() PortRedundancy {
	if o == nil || IsNil(o.Redundancy) {
		var ret PortRedundancy
		return ret
	}
	return *o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetRedundancyOk() (*PortRedundancy, bool) {
	if o == nil || IsNil(o.Redundancy) {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *SimplifiedPort) HasRedundancy() bool {
	if o != nil && !IsNil(o.Redundancy) {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given PortRedundancy and assigns it to the Redundancy field.
func (o *SimplifiedPort) SetRedundancy(v PortRedundancy) {
	o.Redundancy = &v
}

// GetEncapsulation returns the Encapsulation field value if set, zero value otherwise.
func (o *SimplifiedPort) GetEncapsulation() PortEncapsulation {
	if o == nil || IsNil(o.Encapsulation) {
		var ret PortEncapsulation
		return ret
	}
	return *o.Encapsulation
}

// GetEncapsulationOk returns a tuple with the Encapsulation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetEncapsulationOk() (*PortEncapsulation, bool) {
	if o == nil || IsNil(o.Encapsulation) {
		return nil, false
	}
	return o.Encapsulation, true
}

// HasEncapsulation returns a boolean if a field has been set.
func (o *SimplifiedPort) HasEncapsulation() bool {
	if o != nil && !IsNil(o.Encapsulation) {
		return true
	}

	return false
}

// SetEncapsulation gets a reference to the given PortEncapsulation and assigns it to the Encapsulation field.
func (o *SimplifiedPort) SetEncapsulation(v PortEncapsulation) {
	o.Encapsulation = &v
}

// GetLagEnabled returns the LagEnabled field value if set, zero value otherwise.
func (o *SimplifiedPort) GetLagEnabled() bool {
	if o == nil || IsNil(o.LagEnabled) {
		var ret bool
		return ret
	}
	return *o.LagEnabled
}

// GetLagEnabledOk returns a tuple with the LagEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetLagEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LagEnabled) {
		return nil, false
	}
	return o.LagEnabled, true
}

// HasLagEnabled returns a boolean if a field has been set.
func (o *SimplifiedPort) HasLagEnabled() bool {
	if o != nil && !IsNil(o.LagEnabled) {
		return true
	}

	return false
}

// SetLagEnabled gets a reference to the given bool and assigns it to the LagEnabled field.
func (o *SimplifiedPort) SetLagEnabled(v bool) {
	o.LagEnabled = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SimplifiedPort) GetSettings() PortSettings {
	if o == nil || IsNil(o.Settings) {
		var ret PortSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetSettingsOk() (*PortSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SimplifiedPort) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given PortSettings and assigns it to the Settings field.
func (o *SimplifiedPort) SetSettings(v PortSettings) {
	o.Settings = &v
}

// GetPhysicalPortQuantity returns the PhysicalPortQuantity field value if set, zero value otherwise.
func (o *SimplifiedPort) GetPhysicalPortQuantity() int32 {
	if o == nil || IsNil(o.PhysicalPortQuantity) {
		var ret int32
		return ret
	}
	return *o.PhysicalPortQuantity
}

// GetPhysicalPortQuantityOk returns a tuple with the PhysicalPortQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetPhysicalPortQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.PhysicalPortQuantity) {
		return nil, false
	}
	return o.PhysicalPortQuantity, true
}

// HasPhysicalPortQuantity returns a boolean if a field has been set.
func (o *SimplifiedPort) HasPhysicalPortQuantity() bool {
	if o != nil && !IsNil(o.PhysicalPortQuantity) {
		return true
	}

	return false
}

// SetPhysicalPortQuantity gets a reference to the given int32 and assigns it to the PhysicalPortQuantity field.
func (o *SimplifiedPort) SetPhysicalPortQuantity(v int32) {
	o.PhysicalPortQuantity = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *SimplifiedPort) GetAdditionalInfo() []PortAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []PortAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetAdditionalInfoOk() ([]PortAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *SimplifiedPort) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []PortAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *SimplifiedPort) SetAdditionalInfo(v []PortAdditionalInfo) {
	o.AdditionalInfo = v
}

// GetPhysicalPorts returns the PhysicalPorts field value if set, zero value otherwise.
func (o *SimplifiedPort) GetPhysicalPorts() []PhysicalPort {
	if o == nil || IsNil(o.PhysicalPorts) {
		var ret []PhysicalPort
		return ret
	}
	return o.PhysicalPorts
}

// GetPhysicalPortsOk returns a tuple with the PhysicalPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedPort) GetPhysicalPortsOk() ([]PhysicalPort, bool) {
	if o == nil || IsNil(o.PhysicalPorts) {
		return nil, false
	}
	return o.PhysicalPorts, true
}

// HasPhysicalPorts returns a boolean if a field has been set.
func (o *SimplifiedPort) HasPhysicalPorts() bool {
	if o != nil && !IsNil(o.PhysicalPorts) {
		return true
	}

	return false
}

// SetPhysicalPorts gets a reference to the given []PhysicalPort and assigns it to the PhysicalPorts field.
func (o *SimplifiedPort) SetPhysicalPorts(v []PhysicalPort) {
	o.PhysicalPorts = v
}

func (o SimplifiedPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
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
	if !IsNil(o.PhysicalPortsSpeed) {
		toSerialize["physicalPortsSpeed"] = o.PhysicalPortsSpeed
	}
	if !IsNil(o.ConnectionsCount) {
		toSerialize["connectionsCount"] = o.ConnectionsCount
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CvpId) {
		toSerialize["cvpId"] = o.CvpId
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.ServiceType) {
		toSerialize["serviceType"] = o.ServiceType
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.AvailableBandwidth) {
		toSerialize["availableBandwidth"] = o.AvailableBandwidth
	}
	if !IsNil(o.UsedBandwidth) {
		toSerialize["usedBandwidth"] = o.UsedBandwidth
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Tether) {
		toSerialize["tether"] = o.Tether
	}
	if !IsNil(o.DemarcationPoint) {
		toSerialize["demarcationPoint"] = o.DemarcationPoint
	}
	if !IsNil(o.Redundancy) {
		toSerialize["redundancy"] = o.Redundancy
	}
	if !IsNil(o.Encapsulation) {
		toSerialize["encapsulation"] = o.Encapsulation
	}
	if !IsNil(o.LagEnabled) {
		toSerialize["lagEnabled"] = o.LagEnabled
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.PhysicalPortQuantity) {
		toSerialize["physicalPortQuantity"] = o.PhysicalPortQuantity
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.PhysicalPorts) {
		toSerialize["physicalPorts"] = o.PhysicalPorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedPort) UnmarshalJSON(data []byte) (err error) {
	varSimplifiedPort := _SimplifiedPort{}

	err = json.Unmarshal(data, &varSimplifiedPort)

	if err != nil {
		return err
	}

	*o = SimplifiedPort(varSimplifiedPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "physicalPortsSpeed")
		delete(additionalProperties, "connectionsCount")
		delete(additionalProperties, "project")
		delete(additionalProperties, "state")
		delete(additionalProperties, "cvpId")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "account")
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "availableBandwidth")
		delete(additionalProperties, "usedBandwidth")
		delete(additionalProperties, "location")
		delete(additionalProperties, "device")
		delete(additionalProperties, "interface")
		delete(additionalProperties, "tether")
		delete(additionalProperties, "demarcationPoint")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "encapsulation")
		delete(additionalProperties, "lagEnabled")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "physicalPortQuantity")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "physicalPorts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedPort struct {
	value *SimplifiedPort
	isSet bool
}

func (v NullableSimplifiedPort) Get() *SimplifiedPort {
	return v.value
}

func (v *NullableSimplifiedPort) Set(val *SimplifiedPort) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedPort) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedPort(val *SimplifiedPort) *NullableSimplifiedPort {
	return &NullableSimplifiedPort{value: val, isSet: true}
}

func (v NullableSimplifiedPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
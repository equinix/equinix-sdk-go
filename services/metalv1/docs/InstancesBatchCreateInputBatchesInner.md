# InstancesBatchCreateInputBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Quantity** | Pointer to **int32** | The number of devices to create in this batch. The hostname may contain an &#x60;{{index}}&#x60; placeholder, which will be replaced with the index of the device in the batch. For example, if the hostname is &#x60;device-{{index}}&#x60;, the first device in the batch will have the hostname &#x60;device-01&#x60;, the second device will have the hostname &#x60;device-02&#x60;, and so on. | [optional] 
**Metro** | **string** | Metro code or ID of where the device should be provisioned in, or it can be instructed to create the device in the best available metro with &#x60;{ \&quot;metro\&quot;: \&quot;any\&quot; }&#x60;. The special metro value of any means anywhere, any metro. When any is chosen in the request, the metro location is picked per our scheduling algorithms that favor the following factors: hardware reservation location (if requesting reserved hardware), ip reservations, spot instances, etc. The any keyword *does not* optimize for cost, this means that usage costs (instance, transfer, other features dependent on location) will vary. Please check metro value in response to see where the device was created. Either metro or facility must be provided. | 
**AlwaysPxe** | Pointer to **bool** | When true, devices with a &#x60;custom_ipxe&#x60; OS will always boot to iPXE. The default setting of false ensures that iPXE will be used on only the first boot. | [optional] [default to false]
**BillingCycle** | Pointer to [**DeviceCreateInputBillingCycle**](DeviceCreateInputBillingCycle.md) |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** | Customdata is an arbitrary JSON value that can be accessed via the metadata service. | [optional] [default to {}]
**Description** | Pointer to **string** | Any description of the device or how it will be used. This may be used to inform other API consumers with project access. | [optional] 
**Features** | Pointer to **[]string** | The features attribute allows you to optionally specify what features your server should have.  In the API shorthand syntax, all features listed are &#x60;required&#x60;:  &#x60;&#x60;&#x60; { \&quot;features\&quot;: [\&quot;tpm\&quot;] } &#x60;&#x60;&#x60;  Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a &#x60;preferred&#x60; value. The request will not fail if we have no servers with that feature in our inventory. The API offers an alternative syntax for mixing preferred and required features:  &#x60;&#x60;&#x60; { \&quot;features\&quot;: { \&quot;tpm\&quot;: \&quot;required\&quot;, \&quot;raid\&quot;: \&quot;preferred\&quot; } } &#x60;&#x60;&#x60;  The request will only fail if there are no available servers matching the required &#x60;tpm&#x60; criteria. | [optional] 
**HardwareReservationId** | Pointer to **string** | The Hardware Reservation UUID to provision. Alternatively, &#x60;next-available&#x60; can be specified to select from any of the available hardware reservations. An error will be returned if the requested reservation option is not available.  See [Reserved Hardware](https://metal.equinix.com/developers/docs/deploy/reserved/) for more details. | [optional] [default to ""]
**Hostname** | Pointer to **string** | The hostname to use within the operating system. The same hostname may be used on multiple devices within a project. | [optional] 
**IpAddresses** | Pointer to [**[]IPAddress**](IPAddress.md) | The &#x60;ip_addresses attribute will allow you to specify the addresses you want created with your device.  The default value configures public IPv4, public IPv6, and private IPv4.  Private IPv4 address is required. When specifying &#x60;ip_addresses&#x60;, one of the array items must enable private IPv4.  Some operating systems require public IPv4 address. In those cases you will receive an error message if public IPv4 is not enabled.  For example, to only configure your server with a private IPv4 address, you can send &#x60;{ \&quot;ip_addresses\&quot;: [{ \&quot;address_family\&quot;: 4, \&quot;public\&quot;: false }] }&#x60;.  It is possible to request a subnet size larger than a &#x60;/30&#x60; by assigning addresses using the UUID(s) of ip_reservations in your project.  For example, &#x60;{ \&quot;ip_addresses\&quot;: [..., {\&quot;address_family\&quot;: 4, \&quot;public\&quot;: true, \&quot;ip_reservations\&quot;: [\&quot;uuid1\&quot;, \&quot;uuid2\&quot;]}] }&#x60;  To access a server without public IPs, you can use our Out-of-Band console access (SOS) or proxy through another server in the project with public IPs enabled. | [optional] [default to [{address_family=4, public=true}, {address_family=4, public=false}, {address_family=6, public=true}]]
**IpxeScriptUrl** | Pointer to **string** | When set, the device will chainload an iPXE Script at boot fetched from the supplied URL.  See [Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/) for more details. | [optional] 
**Locked** | Pointer to **bool** | Whether the device should be locked, preventing accidental deletion. | [optional] [default to false]
**NetworkFrozen** | Pointer to **bool** | If true, this instance can not be converted to a different network type. | [optional] 
**NoSshKeys** | Pointer to **bool** | Overrides default behaviour of attaching all of the organization members ssh keys and project ssh keys to device if no specific keys specified | [optional] [default to false]
**OperatingSystem** | **string** | The slug of the operating system to provision. Check the Equinix Metal operating system documentation for rules that may be imposed per operating system, including restrictions on IP address options and device plans. | 
**Plan** | **string** | The slug of the device plan to provision. | 
**PrivateIpv4SubnetSize** | Pointer to **int32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. | [optional] [default to 28]
**ProjectSshKeys** | Pointer to **[]string** | A list of UUIDs identifying the device parent project that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  If no SSH keys are specified (&#x60;user_ssh_keys&#x60;, &#x60;project_ssh_keys&#x60;, and &#x60;ssh_keys&#x60; are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with &#39;no_ssh_keys&#39; option to omit any SSH key being added.  | [optional] 
**PublicIpv4SubnetSize** | Pointer to **int32** | Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request. | [optional] [default to 31]
**SpotInstance** | Pointer to **bool** | Create a spot instance. Spot instances are created with a maximum bid price. If the bid price is not met, the spot instance will be terminated as indicated by the &#x60;termination_time&#x60; field. | [optional] 
**SpotPriceMax** | Pointer to **float32** | The maximum amount to bid for a spot instance. | [optional] 
**SshKeys** | Pointer to [**[]SSHKeyInput**](SSHKeyInput.md) | A list of new or existing project ssh_keys that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  These keys are added in addition to any keys defined by   &#x60;project_ssh_keys&#x60; and &#x60;user_ssh_keys&#x60;.  | [optional] 
**Storage** | Pointer to [**Storage**](Storage.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** | When the device will be terminated. If you don&#39;t supply timezone info, the timestamp is assumed to be in UTC.  This is commonly set in advance for ephemeral spot market instances but this field may also be set with on-demand and reservation instances to automatically delete the resource at a given time. The termination time can also be used to release a hardware reservation instance at a given time, keeping the reservation open for other uses.  On a spot market device, the termination time will be set automatically when outbid.  | [optional] 
**UserSshKeys** | Pointer to **[]string** | A list of UUIDs identifying the users that should be authorized to access this device (typically via /root/.ssh/authorized_keys).  These keys will also appear in the device metadata.  The users must be members of the project or organization.  If no SSH keys are specified (&#x60;user_ssh_keys&#x60;, &#x60;project_ssh_keys&#x60;, and &#x60;ssh_keys&#x60; are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with &#39;no_ssh_keys&#39; option to omit any SSH key being added.  | [optional] 
**Userdata** | Pointer to **string** | The userdata presented in the metadata service for this device.  Userdata is fetched and interpreted by the operating system installed on the device. Acceptable formats are determined by the operating system, with the exception of a special iPXE enabling syntax which is handled before the operating system starts.  See [Server User Data](https://metal.equinix.com/developers/docs/servers/user-data/) and [Provisioning with Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/#provisioning-with-custom-ipxe) for more details. | [optional] 
**Facility** | **[]string** | The datacenter where the device should be created.  Either metro or facility must be provided.  The API will accept either a single facility &#x60;{ \&quot;facility\&quot;: \&quot;f1\&quot; }&#x60;, or it can be instructed to create the device in the best available datacenter &#x60;{ \&quot;facility\&quot;: \&quot;any\&quot; }&#x60;.  Additionally it is possible to set a prioritized location selection. For example &#x60;{ \&quot;facility\&quot;: [\&quot;f3\&quot;, \&quot;f2\&quot;, \&quot;any\&quot;] }&#x60; can be used to prioritize &#x60;f3&#x60; and then &#x60;f2&#x60; before accepting &#x60;any&#x60; facility. If none of the facilities provided have availability for the requested device the request will fail. | 

## Methods

### NewInstancesBatchCreateInputBatchesInner

`func NewInstancesBatchCreateInputBatchesInner(metro string, operatingSystem string, plan string, facility []string, ) *InstancesBatchCreateInputBatchesInner`

NewInstancesBatchCreateInputBatchesInner instantiates a new InstancesBatchCreateInputBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesBatchCreateInputBatchesInnerWithDefaults

`func NewInstancesBatchCreateInputBatchesInnerWithDefaults() *InstancesBatchCreateInputBatchesInner`

NewInstancesBatchCreateInputBatchesInnerWithDefaults instantiates a new InstancesBatchCreateInputBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnames

`func (o *InstancesBatchCreateInputBatchesInner) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *InstancesBatchCreateInputBatchesInner) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *InstancesBatchCreateInputBatchesInner) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *InstancesBatchCreateInputBatchesInner) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetQuantity

`func (o *InstancesBatchCreateInputBatchesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InstancesBatchCreateInputBatchesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InstancesBatchCreateInputBatchesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InstancesBatchCreateInputBatchesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetMetro

`func (o *InstancesBatchCreateInputBatchesInner) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *InstancesBatchCreateInputBatchesInner) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *InstancesBatchCreateInputBatchesInner) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetAlwaysPxe

`func (o *InstancesBatchCreateInputBatchesInner) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *InstancesBatchCreateInputBatchesInner) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *InstancesBatchCreateInputBatchesInner) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *InstancesBatchCreateInputBatchesInner) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *InstancesBatchCreateInputBatchesInner) GetBillingCycle() DeviceCreateInputBillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *InstancesBatchCreateInputBatchesInner) GetBillingCycleOk() (*DeviceCreateInputBillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *InstancesBatchCreateInputBatchesInner) SetBillingCycle(v DeviceCreateInputBillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *InstancesBatchCreateInputBatchesInner) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *InstancesBatchCreateInputBatchesInner) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *InstancesBatchCreateInputBatchesInner) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *InstancesBatchCreateInputBatchesInner) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *InstancesBatchCreateInputBatchesInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *InstancesBatchCreateInputBatchesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstancesBatchCreateInputBatchesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstancesBatchCreateInputBatchesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstancesBatchCreateInputBatchesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *InstancesBatchCreateInputBatchesInner) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InstancesBatchCreateInputBatchesInner) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InstancesBatchCreateInputBatchesInner) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InstancesBatchCreateInputBatchesInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHardwareReservationId

`func (o *InstancesBatchCreateInputBatchesInner) GetHardwareReservationId() string`

GetHardwareReservationId returns the HardwareReservationId field if non-nil, zero value otherwise.

### GetHardwareReservationIdOk

`func (o *InstancesBatchCreateInputBatchesInner) GetHardwareReservationIdOk() (*string, bool)`

GetHardwareReservationIdOk returns a tuple with the HardwareReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReservationId

`func (o *InstancesBatchCreateInputBatchesInner) SetHardwareReservationId(v string)`

SetHardwareReservationId sets HardwareReservationId field to given value.

### HasHardwareReservationId

`func (o *InstancesBatchCreateInputBatchesInner) HasHardwareReservationId() bool`

HasHardwareReservationId returns a boolean if a field has been set.

### GetHostname

`func (o *InstancesBatchCreateInputBatchesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstancesBatchCreateInputBatchesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstancesBatchCreateInputBatchesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstancesBatchCreateInputBatchesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddresses

`func (o *InstancesBatchCreateInputBatchesInner) GetIpAddresses() []IPAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *InstancesBatchCreateInputBatchesInner) GetIpAddressesOk() (*[]IPAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *InstancesBatchCreateInputBatchesInner) SetIpAddresses(v []IPAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *InstancesBatchCreateInputBatchesInner) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *InstancesBatchCreateInputBatchesInner) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *InstancesBatchCreateInputBatchesInner) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *InstancesBatchCreateInputBatchesInner) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *InstancesBatchCreateInputBatchesInner) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

### GetLocked

`func (o *InstancesBatchCreateInputBatchesInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstancesBatchCreateInputBatchesInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstancesBatchCreateInputBatchesInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstancesBatchCreateInputBatchesInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNetworkFrozen

`func (o *InstancesBatchCreateInputBatchesInner) GetNetworkFrozen() bool`

GetNetworkFrozen returns the NetworkFrozen field if non-nil, zero value otherwise.

### GetNetworkFrozenOk

`func (o *InstancesBatchCreateInputBatchesInner) GetNetworkFrozenOk() (*bool, bool)`

GetNetworkFrozenOk returns a tuple with the NetworkFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFrozen

`func (o *InstancesBatchCreateInputBatchesInner) SetNetworkFrozen(v bool)`

SetNetworkFrozen sets NetworkFrozen field to given value.

### HasNetworkFrozen

`func (o *InstancesBatchCreateInputBatchesInner) HasNetworkFrozen() bool`

HasNetworkFrozen returns a boolean if a field has been set.

### GetNoSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *InstancesBatchCreateInputBatchesInner) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *InstancesBatchCreateInputBatchesInner) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *InstancesBatchCreateInputBatchesInner) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetPlan

`func (o *InstancesBatchCreateInputBatchesInner) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstancesBatchCreateInputBatchesInner) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstancesBatchCreateInputBatchesInner) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetPrivateIpv4SubnetSize

`func (o *InstancesBatchCreateInputBatchesInner) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *InstancesBatchCreateInputBatchesInner) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *InstancesBatchCreateInputBatchesInner) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *InstancesBatchCreateInputBatchesInner) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetPublicIpv4SubnetSize

`func (o *InstancesBatchCreateInputBatchesInner) GetPublicIpv4SubnetSize() int32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *InstancesBatchCreateInputBatchesInner) GetPublicIpv4SubnetSizeOk() (*int32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *InstancesBatchCreateInputBatchesInner) SetPublicIpv4SubnetSize(v int32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *InstancesBatchCreateInputBatchesInner) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

### GetSpotInstance

`func (o *InstancesBatchCreateInputBatchesInner) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *InstancesBatchCreateInputBatchesInner) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *InstancesBatchCreateInputBatchesInner) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *InstancesBatchCreateInputBatchesInner) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetSpotPriceMax

`func (o *InstancesBatchCreateInputBatchesInner) GetSpotPriceMax() float32`

GetSpotPriceMax returns the SpotPriceMax field if non-nil, zero value otherwise.

### GetSpotPriceMaxOk

`func (o *InstancesBatchCreateInputBatchesInner) GetSpotPriceMaxOk() (*float32, bool)`

GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPriceMax

`func (o *InstancesBatchCreateInputBatchesInner) SetSpotPriceMax(v float32)`

SetSpotPriceMax sets SpotPriceMax field to given value.

### HasSpotPriceMax

`func (o *InstancesBatchCreateInputBatchesInner) HasSpotPriceMax() bool`

HasSpotPriceMax returns a boolean if a field has been set.

### GetSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetSshKeys() []SSHKeyInput`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetSshKeysOk() (*[]SSHKeyInput, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetSshKeys(v []SSHKeyInput)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetStorage

`func (o *InstancesBatchCreateInputBatchesInner) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstancesBatchCreateInputBatchesInner) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstancesBatchCreateInputBatchesInner) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstancesBatchCreateInputBatchesInner) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTags

`func (o *InstancesBatchCreateInputBatchesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstancesBatchCreateInputBatchesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstancesBatchCreateInputBatchesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstancesBatchCreateInputBatchesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationTime

`func (o *InstancesBatchCreateInputBatchesInner) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *InstancesBatchCreateInputBatchesInner) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *InstancesBatchCreateInputBatchesInner) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *InstancesBatchCreateInputBatchesInner) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *InstancesBatchCreateInputBatchesInner) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *InstancesBatchCreateInputBatchesInner) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetUserdata

`func (o *InstancesBatchCreateInputBatchesInner) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *InstancesBatchCreateInputBatchesInner) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *InstancesBatchCreateInputBatchesInner) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *InstancesBatchCreateInputBatchesInner) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetFacility

`func (o *InstancesBatchCreateInputBatchesInner) GetFacility() []string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *InstancesBatchCreateInputBatchesInner) GetFacilityOk() (*[]string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *InstancesBatchCreateInputBatchesInner) SetFacility(v []string)`

SetFacility sets Facility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Interconnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**Facility**](Facility.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**Mode** | Pointer to [**InterconnectionMode**](InterconnectionMode.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Ports** | Pointer to [**[]InterconnectionPort**](InterconnectionPort.md) | For Fabric VCs, these represent Virtual Port(s) created for the interconnection. For dedicated interconnections, these represent the Dedicated Port(s). | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Redundancy** | Pointer to [**InterconnectionRedundancy**](InterconnectionRedundancy.md) |  | [optional] 
**ServiceTokens** | Pointer to [**[]FabricServiceToken**](FabricServiceToken.md) | For Fabric VCs (Metal Billed), this will show details of the A-Side service tokens issued for the interconnection. For Fabric VCs (Fabric Billed), this will show the details of the Z-Side service tokens issued for the interconnection. Dedicated interconnections will not have any service tokens issued. There will be one per interconnection, so for redundant interconnections, there should be two service tokens issued. | [optional] 
**AuthorizationCode** | Pointer to **string** | For Fabric VCs (Metal Billed), this allows Fabric to connect the Metal network to any connection Fabric facilitates. Fabric uses this token to be able to give more detailed information about the Metal end of the network, when viewing resources from within Fabric. | [optional] 
**Speed** | Pointer to **int64** | For interconnections on Dedicated Ports and shared connections, this represents the interconnection&#39;s speed in bps. For Fabric VCs, this field refers to the maximum speed of the interconnection in bps. This value will default to 10Gbps for Fabric VCs (Fabric Billed). | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** | This token is used for shared interconnections to be used as the Fabric Token. This field is entirely deprecated. | [optional] 
**Type** | Pointer to [**InterconnectionType**](InterconnectionType.md) |  | [optional] 
**FabricProvider** | Pointer to [**AWSFabricProvider**](AWSFabricProvider.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RequestedBy** | Pointer to [**Href**](Href.md) |  | [optional] 

## Methods

### NewInterconnection

`func NewInterconnection() *Interconnection`

NewInterconnection instantiates a new Interconnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionWithDefaults

`func NewInterconnectionWithDefaults() *Interconnection`

NewInterconnectionWithDefaults instantiates a new Interconnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *Interconnection) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Interconnection) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Interconnection) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Interconnection) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *Interconnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Interconnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Interconnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Interconnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *Interconnection) GetFacility() Facility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *Interconnection) GetFacilityOk() (*Facility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *Interconnection) SetFacility(v Facility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *Interconnection) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetId

`func (o *Interconnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interconnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interconnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Interconnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetro

`func (o *Interconnection) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Interconnection) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Interconnection) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Interconnection) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMode

`func (o *Interconnection) GetMode() InterconnectionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Interconnection) GetModeOk() (*InterconnectionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Interconnection) SetMode(v InterconnectionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Interconnection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *Interconnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interconnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interconnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Interconnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *Interconnection) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Interconnection) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Interconnection) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Interconnection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPorts

`func (o *Interconnection) GetPorts() []InterconnectionPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Interconnection) GetPortsOk() (*[]InterconnectionPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Interconnection) SetPorts(v []InterconnectionPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Interconnection) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProject

`func (o *Interconnection) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Interconnection) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Interconnection) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Interconnection) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRedundancy

`func (o *Interconnection) GetRedundancy() InterconnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *Interconnection) GetRedundancyOk() (*InterconnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *Interconnection) SetRedundancy(v InterconnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *Interconnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetServiceTokens

`func (o *Interconnection) GetServiceTokens() []FabricServiceToken`

GetServiceTokens returns the ServiceTokens field if non-nil, zero value otherwise.

### GetServiceTokensOk

`func (o *Interconnection) GetServiceTokensOk() (*[]FabricServiceToken, bool)`

GetServiceTokensOk returns a tuple with the ServiceTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokens

`func (o *Interconnection) SetServiceTokens(v []FabricServiceToken)`

SetServiceTokens sets ServiceTokens field to given value.

### HasServiceTokens

`func (o *Interconnection) HasServiceTokens() bool`

HasServiceTokens returns a boolean if a field has been set.

### GetAuthorizationCode

`func (o *Interconnection) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *Interconnection) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *Interconnection) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *Interconnection) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetSpeed

`func (o *Interconnection) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Interconnection) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Interconnection) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Interconnection) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *Interconnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Interconnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Interconnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Interconnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Interconnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Interconnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Interconnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Interconnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *Interconnection) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Interconnection) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Interconnection) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Interconnection) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *Interconnection) GetType() InterconnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interconnection) GetTypeOk() (*InterconnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interconnection) SetType(v InterconnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *Interconnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFabricProvider

`func (o *Interconnection) GetFabricProvider() AWSFabricProvider`

GetFabricProvider returns the FabricProvider field if non-nil, zero value otherwise.

### GetFabricProviderOk

`func (o *Interconnection) GetFabricProviderOk() (*AWSFabricProvider, bool)`

GetFabricProviderOk returns a tuple with the FabricProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricProvider

`func (o *Interconnection) SetFabricProvider(v AWSFabricProvider)`

SetFabricProvider sets FabricProvider field to given value.

### HasFabricProvider

`func (o *Interconnection) HasFabricProvider() bool`

HasFabricProvider returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Interconnection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Interconnection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Interconnection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Interconnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Interconnection) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Interconnection) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Interconnection) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Interconnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *Interconnection) GetRequestedBy() Href`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *Interconnection) GetRequestedByOk() (*Href, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *Interconnection) SetRequestedBy(v Href)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *Interconnection) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



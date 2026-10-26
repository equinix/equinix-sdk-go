# CreateCrossConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CrossConnectType**](CrossConnectType.md) |  | 
**Name** | **string** | A short, descriptive name for this Cross Connect. | 
**BmmrType** | Pointer to [**BmmrType**](BmmrType.md) | Required when connecting to a BMMR. &lt;br&gt; Use EQUINIX when BMMR Extension is required and Equinix installs it, or SELF when you install it yourself.  | [optional] 
**CircuitId** | Pointer to **string** | Circuit identifier of the Cross Connect. | [optional] 
**Redundancy** | Pointer to [**CrossConnectRedundancy**](CrossConnectRedundancy.md) | Provide it to place a diverse Cross Connect against an existing one - set group to the redundancy group id or the serial number of the existing Cross Connect, and priority to SECONDARY.  | [optional] 
**ASide** | [**CrossConnectASideRequest**](CrossConnectASideRequest.md) | A-Side connection details. | 
**ZSide** | [**CrossConnectZSideRequest**](CrossConnectZSideRequest.md) | Z-Side connection details. &lt;br&gt; Provide the DemarcationPoint, the LOA reference or the BMMR location depending on the Cross Connect being requested.  | 
**Schedule** | Pointer to [**ScheduleRequest**](ScheduleRequest.md) | Preferred installation schedule. When omitted, the Cross Connect is scheduled with the STANDARD schedule code.  | [optional] 
**Services** | Pointer to [**[]CrossConnectServiceRequest**](CrossConnectServiceRequest.md) | Cross Connect Services. | [optional] 
**Notifications** | Pointer to [**[]CrossConnectNotification**](CrossConnectNotification.md) | Preferences for notifications on Cross Connect configuration or status changes.  | [optional] 
**Order** | Pointer to [**OrderRequest**](OrderRequest.md) |  | [optional] 
**Notes** | Pointer to [**[]CrossConnectNoteRequest**](CrossConnectNoteRequest.md) | Additional Note. | [optional] 
**Attachments** | Pointer to [**[]CrossConnectAttachment**](CrossConnectAttachment.md) | Additional Attachments. | [optional] 

## Methods

### NewCreateCrossConnect

`func NewCreateCrossConnect(type_ CrossConnectType, name string, aSide CrossConnectASideRequest, zSide CrossConnectZSideRequest, ) *CreateCrossConnect`

NewCreateCrossConnect instantiates a new CreateCrossConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCrossConnectWithDefaults

`func NewCreateCrossConnectWithDefaults() *CreateCrossConnect`

NewCreateCrossConnectWithDefaults instantiates a new CreateCrossConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateCrossConnect) GetType() CrossConnectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCrossConnect) GetTypeOk() (*CrossConnectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCrossConnect) SetType(v CrossConnectType)`

SetType sets Type field to given value.


### GetName

`func (o *CreateCrossConnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCrossConnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCrossConnect) SetName(v string)`

SetName sets Name field to given value.


### GetBmmrType

`func (o *CreateCrossConnect) GetBmmrType() BmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *CreateCrossConnect) GetBmmrTypeOk() (*BmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *CreateCrossConnect) SetBmmrType(v BmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *CreateCrossConnect) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.

### GetCircuitId

`func (o *CreateCrossConnect) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *CreateCrossConnect) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *CreateCrossConnect) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *CreateCrossConnect) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetRedundancy

`func (o *CreateCrossConnect) GetRedundancy() CrossConnectRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateCrossConnect) GetRedundancyOk() (*CrossConnectRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateCrossConnect) SetRedundancy(v CrossConnectRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CreateCrossConnect) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *CreateCrossConnect) GetASide() CrossConnectASideRequest`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *CreateCrossConnect) GetASideOk() (*CrossConnectASideRequest, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *CreateCrossConnect) SetASide(v CrossConnectASideRequest)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *CreateCrossConnect) GetZSide() CrossConnectZSideRequest`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *CreateCrossConnect) GetZSideOk() (*CrossConnectZSideRequest, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *CreateCrossConnect) SetZSide(v CrossConnectZSideRequest)`

SetZSide sets ZSide field to given value.


### GetSchedule

`func (o *CreateCrossConnect) GetSchedule() ScheduleRequest`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateCrossConnect) GetScheduleOk() (*ScheduleRequest, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateCrossConnect) SetSchedule(v ScheduleRequest)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateCrossConnect) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServices

`func (o *CreateCrossConnect) GetServices() []CrossConnectServiceRequest`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CreateCrossConnect) GetServicesOk() (*[]CrossConnectServiceRequest, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CreateCrossConnect) SetServices(v []CrossConnectServiceRequest)`

SetServices sets Services field to given value.

### HasServices

`func (o *CreateCrossConnect) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetNotifications

`func (o *CreateCrossConnect) GetNotifications() []CrossConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CreateCrossConnect) GetNotificationsOk() (*[]CrossConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CreateCrossConnect) SetNotifications(v []CrossConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CreateCrossConnect) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOrder

`func (o *CreateCrossConnect) GetOrder() OrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateCrossConnect) GetOrderOk() (*OrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateCrossConnect) SetOrder(v OrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateCrossConnect) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNotes

`func (o *CreateCrossConnect) GetNotes() []CrossConnectNoteRequest`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateCrossConnect) GetNotesOk() (*[]CrossConnectNoteRequest, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateCrossConnect) SetNotes(v []CrossConnectNoteRequest)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateCrossConnect) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAttachments

`func (o *CreateCrossConnect) GetAttachments() []CrossConnectAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateCrossConnect) GetAttachmentsOk() (*[]CrossConnectAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateCrossConnect) SetAttachments(v []CrossConnectAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateCrossConnect) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



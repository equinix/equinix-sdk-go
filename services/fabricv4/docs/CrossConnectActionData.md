# CrossConnectActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to [**[]CrossConnectActionNote**](CrossConnectActionNote.md) | Applicable for XC_INFORMATION_UPDATE and XC_DEINSTALL. | [optional] 
**Attachments** | Pointer to [**[]CrossConnectAttachment**](CrossConnectAttachment.md) | Applicable for XC_INFORMATION_UPDATE and XC_DEINSTALL. | [optional] 
**Schedule** | Pointer to [**DeinstallScheduleRequest**](DeinstallScheduleRequest.md) | Removal schedule. Required for XC_DEINSTALL - the requestedDateTime is mandatory and only the STANDARD schedule type is supported.  | [optional] 
**LiveTrafficConsentProvided** | Pointer to **bool** | Applicable for XC_DEINSTALL. Confirms the customer consents to the removal even though the Cross Connect may still be carrying live traffic.  | [optional] 
**Services** | Pointer to [**[]CrossConnectServiceRequest**](CrossConnectServiceRequest.md) | Applicable for XC_DEINSTALL. Only XC_PATCH_EQUIPMENT is supported for deinstall.  | [optional] 
**Notifications** | Pointer to [**[]CrossConnectNotification**](CrossConnectNotification.md) | Applicable for XC_DEINSTALL. The contacts are appended to the ones already held on the Cross Connect.  | [optional] 
**Order** | Pointer to [**OrderRequest**](OrderRequest.md) | Applicable for XC_DEINSTALL. The purchase order and the customer reference identifier replace the ones already held on the Cross Connect.  | [optional] 

## Methods

### NewCrossConnectActionData

`func NewCrossConnectActionData() *CrossConnectActionData`

NewCrossConnectActionData instantiates a new CrossConnectActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionDataWithDefaults

`func NewCrossConnectActionDataWithDefaults() *CrossConnectActionData`

NewCrossConnectActionDataWithDefaults instantiates a new CrossConnectActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *CrossConnectActionData) GetNotes() []CrossConnectActionNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CrossConnectActionData) GetNotesOk() (*[]CrossConnectActionNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CrossConnectActionData) SetNotes(v []CrossConnectActionNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CrossConnectActionData) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAttachments

`func (o *CrossConnectActionData) GetAttachments() []CrossConnectAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CrossConnectActionData) GetAttachmentsOk() (*[]CrossConnectAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CrossConnectActionData) SetAttachments(v []CrossConnectAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CrossConnectActionData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSchedule

`func (o *CrossConnectActionData) GetSchedule() DeinstallScheduleRequest`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CrossConnectActionData) GetScheduleOk() (*DeinstallScheduleRequest, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CrossConnectActionData) SetSchedule(v DeinstallScheduleRequest)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CrossConnectActionData) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetLiveTrafficConsentProvided

`func (o *CrossConnectActionData) GetLiveTrafficConsentProvided() bool`

GetLiveTrafficConsentProvided returns the LiveTrafficConsentProvided field if non-nil, zero value otherwise.

### GetLiveTrafficConsentProvidedOk

`func (o *CrossConnectActionData) GetLiveTrafficConsentProvidedOk() (*bool, bool)`

GetLiveTrafficConsentProvidedOk returns a tuple with the LiveTrafficConsentProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveTrafficConsentProvided

`func (o *CrossConnectActionData) SetLiveTrafficConsentProvided(v bool)`

SetLiveTrafficConsentProvided sets LiveTrafficConsentProvided field to given value.

### HasLiveTrafficConsentProvided

`func (o *CrossConnectActionData) HasLiveTrafficConsentProvided() bool`

HasLiveTrafficConsentProvided returns a boolean if a field has been set.

### GetServices

`func (o *CrossConnectActionData) GetServices() []CrossConnectServiceRequest`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CrossConnectActionData) GetServicesOk() (*[]CrossConnectServiceRequest, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CrossConnectActionData) SetServices(v []CrossConnectServiceRequest)`

SetServices sets Services field to given value.

### HasServices

`func (o *CrossConnectActionData) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetNotifications

`func (o *CrossConnectActionData) GetNotifications() []CrossConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CrossConnectActionData) GetNotificationsOk() (*[]CrossConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CrossConnectActionData) SetNotifications(v []CrossConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CrossConnectActionData) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOrder

`func (o *CrossConnectActionData) GetOrder() OrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CrossConnectActionData) GetOrderOk() (*OrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CrossConnectActionData) SetOrder(v OrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CrossConnectActionData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



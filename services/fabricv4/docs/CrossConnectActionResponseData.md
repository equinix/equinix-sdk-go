# CrossConnectActionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to [**[]CrossConnectActionNoteResponse**](CrossConnectActionNoteResponse.md) | Applicable for XC_INFORMATION_UPDATE and XC_DEINSTALL. | [optional] [readonly] 
**Attachments** | Pointer to [**[]CrossConnectActionAttachmentResponse**](CrossConnectActionAttachmentResponse.md) | Applicable for XC_INFORMATION_UPDATE and XC_DEINSTALL. | [optional] [readonly] 
**Schedule** | Pointer to [**CrossConnectScheduleResponse**](CrossConnectScheduleResponse.md) | Applicable for XC_DEINSTALL. | [optional] [readonly] 
**LiveTrafficConsentProvided** | Pointer to **bool** | Applicable for XC_DEINSTALL. Confirms the customer consents to the removal even though the Cross Connect may still be carrying live traffic.  | [optional] [readonly] 
**Services** | Pointer to [**[]CrossConnectServiceResponse**](CrossConnectServiceResponse.md) | Applicable for XC_DEINSTALL. Only XC_PATCH_EQUIPMENT is supported for deinstall.  | [optional] [readonly] 
**Notifications** | Pointer to [**[]CrossConnectNotification**](CrossConnectNotification.md) | Applicable for XC_DEINSTALL. The contacts are appended to the ones already held on the Cross Connect.  | [optional] [readonly] 
**Order** | Pointer to [**OrderResponse**](OrderResponse.md) | Applicable for XC_DEINSTALL. The purchase order and the customer reference identifier replace the ones already held on the Cross Connect.  | [optional] [readonly] 

## Methods

### NewCrossConnectActionResponseData

`func NewCrossConnectActionResponseData() *CrossConnectActionResponseData`

NewCrossConnectActionResponseData instantiates a new CrossConnectActionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionResponseDataWithDefaults

`func NewCrossConnectActionResponseDataWithDefaults() *CrossConnectActionResponseData`

NewCrossConnectActionResponseDataWithDefaults instantiates a new CrossConnectActionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *CrossConnectActionResponseData) GetNotes() []CrossConnectActionNoteResponse`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CrossConnectActionResponseData) GetNotesOk() (*[]CrossConnectActionNoteResponse, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CrossConnectActionResponseData) SetNotes(v []CrossConnectActionNoteResponse)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CrossConnectActionResponseData) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAttachments

`func (o *CrossConnectActionResponseData) GetAttachments() []CrossConnectActionAttachmentResponse`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CrossConnectActionResponseData) GetAttachmentsOk() (*[]CrossConnectActionAttachmentResponse, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CrossConnectActionResponseData) SetAttachments(v []CrossConnectActionAttachmentResponse)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CrossConnectActionResponseData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSchedule

`func (o *CrossConnectActionResponseData) GetSchedule() CrossConnectScheduleResponse`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CrossConnectActionResponseData) GetScheduleOk() (*CrossConnectScheduleResponse, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CrossConnectActionResponseData) SetSchedule(v CrossConnectScheduleResponse)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CrossConnectActionResponseData) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetLiveTrafficConsentProvided

`func (o *CrossConnectActionResponseData) GetLiveTrafficConsentProvided() bool`

GetLiveTrafficConsentProvided returns the LiveTrafficConsentProvided field if non-nil, zero value otherwise.

### GetLiveTrafficConsentProvidedOk

`func (o *CrossConnectActionResponseData) GetLiveTrafficConsentProvidedOk() (*bool, bool)`

GetLiveTrafficConsentProvidedOk returns a tuple with the LiveTrafficConsentProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveTrafficConsentProvided

`func (o *CrossConnectActionResponseData) SetLiveTrafficConsentProvided(v bool)`

SetLiveTrafficConsentProvided sets LiveTrafficConsentProvided field to given value.

### HasLiveTrafficConsentProvided

`func (o *CrossConnectActionResponseData) HasLiveTrafficConsentProvided() bool`

HasLiveTrafficConsentProvided returns a boolean if a field has been set.

### GetServices

`func (o *CrossConnectActionResponseData) GetServices() []CrossConnectServiceResponse`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CrossConnectActionResponseData) GetServicesOk() (*[]CrossConnectServiceResponse, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CrossConnectActionResponseData) SetServices(v []CrossConnectServiceResponse)`

SetServices sets Services field to given value.

### HasServices

`func (o *CrossConnectActionResponseData) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetNotifications

`func (o *CrossConnectActionResponseData) GetNotifications() []CrossConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CrossConnectActionResponseData) GetNotificationsOk() (*[]CrossConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CrossConnectActionResponseData) SetNotifications(v []CrossConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CrossConnectActionResponseData) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOrder

`func (o *CrossConnectActionResponseData) GetOrder() OrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CrossConnectActionResponseData) GetOrderOk() (*OrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CrossConnectActionResponseData) SetOrder(v OrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CrossConnectActionResponseData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



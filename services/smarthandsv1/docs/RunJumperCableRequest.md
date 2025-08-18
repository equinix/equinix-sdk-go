# RunJumperCableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerReferenceNumber** | Pointer to **string** | You may use numbers and text in this field to enter reference information for your records. This will also appear in your reports and details. You may use this information to search for this content on the submitted requests page. | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrder**](PurchaseOrder.md) |  | [optional] 
**IbxLocation** | [**IbxLocation**](IbxLocation.md) |  | 
**Contacts** | [**[]ContactInfo**](ContactInfo.md) | Use this array to pass ordering contact, notification contacts and technical contact. Only one ordering contact, technical contact is allowed. One or more notification contacts are allowed. Ordering and notification contacts are always registered customers with the customer portal. | 
**Schedule** | [**ScheduleInfo**](ScheduleInfo.md) |  | 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Use this to pass uploaded attachments. Attachments need to be uploaded using the attachments API. Maximum size of an attachment is 2MB with the following formats - bmp, jpg, jpeg, gif, png, tif, tiff, txt, doc, docx, xls, xlsx, ppt, pps, ppsx, pdf, and vsd. | [optional] 
**ServiceDetails** | [**RunJumperCableRequestServiceDetails**](RunJumperCableRequestServiceDetails.md) |  | 

## Methods

### NewRunJumperCableRequest

`func NewRunJumperCableRequest(ibxLocation IbxLocation, contacts []ContactInfo, schedule ScheduleInfo, serviceDetails RunJumperCableRequestServiceDetails, ) *RunJumperCableRequest`

NewRunJumperCableRequest instantiates a new RunJumperCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunJumperCableRequestWithDefaults

`func NewRunJumperCableRequestWithDefaults() *RunJumperCableRequest`

NewRunJumperCableRequestWithDefaults instantiates a new RunJumperCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerReferenceNumber

`func (o *RunJumperCableRequest) GetCustomerReferenceNumber() string`

GetCustomerReferenceNumber returns the CustomerReferenceNumber field if non-nil, zero value otherwise.

### GetCustomerReferenceNumberOk

`func (o *RunJumperCableRequest) GetCustomerReferenceNumberOk() (*string, bool)`

GetCustomerReferenceNumberOk returns a tuple with the CustomerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceNumber

`func (o *RunJumperCableRequest) SetCustomerReferenceNumber(v string)`

SetCustomerReferenceNumber sets CustomerReferenceNumber field to given value.

### HasCustomerReferenceNumber

`func (o *RunJumperCableRequest) HasCustomerReferenceNumber() bool`

HasCustomerReferenceNumber returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *RunJumperCableRequest) GetPurchaseOrder() PurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *RunJumperCableRequest) GetPurchaseOrderOk() (*PurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *RunJumperCableRequest) SetPurchaseOrder(v PurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *RunJumperCableRequest) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetIbxLocation

`func (o *RunJumperCableRequest) GetIbxLocation() IbxLocation`

GetIbxLocation returns the IbxLocation field if non-nil, zero value otherwise.

### GetIbxLocationOk

`func (o *RunJumperCableRequest) GetIbxLocationOk() (*IbxLocation, bool)`

GetIbxLocationOk returns a tuple with the IbxLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxLocation

`func (o *RunJumperCableRequest) SetIbxLocation(v IbxLocation)`

SetIbxLocation sets IbxLocation field to given value.


### GetContacts

`func (o *RunJumperCableRequest) GetContacts() []ContactInfo`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *RunJumperCableRequest) GetContactsOk() (*[]ContactInfo, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *RunJumperCableRequest) SetContacts(v []ContactInfo)`

SetContacts sets Contacts field to given value.


### GetSchedule

`func (o *RunJumperCableRequest) GetSchedule() ScheduleInfo`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RunJumperCableRequest) GetScheduleOk() (*ScheduleInfo, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RunJumperCableRequest) SetSchedule(v ScheduleInfo)`

SetSchedule sets Schedule field to given value.


### GetAttachments

`func (o *RunJumperCableRequest) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *RunJumperCableRequest) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *RunJumperCableRequest) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *RunJumperCableRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetServiceDetails

`func (o *RunJumperCableRequest) GetServiceDetails() RunJumperCableRequestServiceDetails`

GetServiceDetails returns the ServiceDetails field if non-nil, zero value otherwise.

### GetServiceDetailsOk

`func (o *RunJumperCableRequest) GetServiceDetailsOk() (*RunJumperCableRequestServiceDetails, bool)`

GetServiceDetailsOk returns a tuple with the ServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDetails

`func (o *RunJumperCableRequest) SetServiceDetails(v RunJumperCableRequestServiceDetails)`

SetServiceDetails sets ServiceDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



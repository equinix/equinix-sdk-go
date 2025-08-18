# ShipmentUnpackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerReferenceNumber** | Pointer to **string** | You may use numbers and text in this field to enter reference information for your records. This will also appear in your reports and details. You may use this information to search for this content on the submitted requests page. | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrder**](PurchaseOrder.md) |  | [optional] 
**IbxLocation** | [**IbxLocation**](IbxLocation.md) |  | 
**Contacts** | [**[]ContactInfo**](ContactInfo.md) | Use this array to pass ordering contact, notification contacts and technical contact. Only one ordering contact, technical contact is allowed. One or more notification contacts are allowed. Ordering and notification contacts are always registered customers with the customer portal. | 
**Schedule** | [**ScheduleInfo**](ScheduleInfo.md) |  | 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Use this to pass uploaded attachments. Attachments need to be uploaded using the attachments API. Maximum size of an attachment is 2MB with the following formats - bmp, jpg, jpeg, gif, png, tif, tiff, txt, doc, docx, xls, xlsx, ppt, pps, ppsx, pdf, and vsd. | [optional] 
**ServiceDetails** | [**ShipmentUnpackRequestServiceDetails**](ShipmentUnpackRequestServiceDetails.md) |  | 

## Methods

### NewShipmentUnpackRequest

`func NewShipmentUnpackRequest(ibxLocation IbxLocation, contacts []ContactInfo, schedule ScheduleInfo, serviceDetails ShipmentUnpackRequestServiceDetails, ) *ShipmentUnpackRequest`

NewShipmentUnpackRequest instantiates a new ShipmentUnpackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentUnpackRequestWithDefaults

`func NewShipmentUnpackRequestWithDefaults() *ShipmentUnpackRequest`

NewShipmentUnpackRequestWithDefaults instantiates a new ShipmentUnpackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerReferenceNumber

`func (o *ShipmentUnpackRequest) GetCustomerReferenceNumber() string`

GetCustomerReferenceNumber returns the CustomerReferenceNumber field if non-nil, zero value otherwise.

### GetCustomerReferenceNumberOk

`func (o *ShipmentUnpackRequest) GetCustomerReferenceNumberOk() (*string, bool)`

GetCustomerReferenceNumberOk returns a tuple with the CustomerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceNumber

`func (o *ShipmentUnpackRequest) SetCustomerReferenceNumber(v string)`

SetCustomerReferenceNumber sets CustomerReferenceNumber field to given value.

### HasCustomerReferenceNumber

`func (o *ShipmentUnpackRequest) HasCustomerReferenceNumber() bool`

HasCustomerReferenceNumber returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *ShipmentUnpackRequest) GetPurchaseOrder() PurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *ShipmentUnpackRequest) GetPurchaseOrderOk() (*PurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *ShipmentUnpackRequest) SetPurchaseOrder(v PurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *ShipmentUnpackRequest) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetIbxLocation

`func (o *ShipmentUnpackRequest) GetIbxLocation() IbxLocation`

GetIbxLocation returns the IbxLocation field if non-nil, zero value otherwise.

### GetIbxLocationOk

`func (o *ShipmentUnpackRequest) GetIbxLocationOk() (*IbxLocation, bool)`

GetIbxLocationOk returns a tuple with the IbxLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxLocation

`func (o *ShipmentUnpackRequest) SetIbxLocation(v IbxLocation)`

SetIbxLocation sets IbxLocation field to given value.


### GetContacts

`func (o *ShipmentUnpackRequest) GetContacts() []ContactInfo`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ShipmentUnpackRequest) GetContactsOk() (*[]ContactInfo, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ShipmentUnpackRequest) SetContacts(v []ContactInfo)`

SetContacts sets Contacts field to given value.


### GetSchedule

`func (o *ShipmentUnpackRequest) GetSchedule() ScheduleInfo`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ShipmentUnpackRequest) GetScheduleOk() (*ScheduleInfo, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ShipmentUnpackRequest) SetSchedule(v ScheduleInfo)`

SetSchedule sets Schedule field to given value.


### GetAttachments

`func (o *ShipmentUnpackRequest) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ShipmentUnpackRequest) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ShipmentUnpackRequest) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ShipmentUnpackRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetServiceDetails

`func (o *ShipmentUnpackRequest) GetServiceDetails() ShipmentUnpackRequestServiceDetails`

GetServiceDetails returns the ServiceDetails field if non-nil, zero value otherwise.

### GetServiceDetailsOk

`func (o *ShipmentUnpackRequest) GetServiceDetailsOk() (*ShipmentUnpackRequestServiceDetails, bool)`

GetServiceDetailsOk returns a tuple with the ServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDetails

`func (o *ShipmentUnpackRequest) SetServiceDetails(v ShipmentUnpackRequestServiceDetails)`

SetServiceDetails sets ServiceDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



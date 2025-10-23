# GatewayAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GatewayAttachmentResponseType**](GatewayAttachmentResponseType.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AttachmentStatus** | Pointer to [**GatewayAttachmentResponseAttachmentStatus**](GatewayAttachmentResponseAttachmentStatus.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewGatewayAttachmentResponse

`func NewGatewayAttachmentResponse() *GatewayAttachmentResponse`

NewGatewayAttachmentResponse instantiates a new GatewayAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAttachmentResponseWithDefaults

`func NewGatewayAttachmentResponseWithDefaults() *GatewayAttachmentResponse`

NewGatewayAttachmentResponseWithDefaults instantiates a new GatewayAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *GatewayAttachmentResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GatewayAttachmentResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GatewayAttachmentResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GatewayAttachmentResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *GatewayAttachmentResponse) GetType() GatewayAttachmentResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayAttachmentResponse) GetTypeOk() (*GatewayAttachmentResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayAttachmentResponse) SetType(v GatewayAttachmentResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayAttachmentResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *GatewayAttachmentResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GatewayAttachmentResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GatewayAttachmentResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GatewayAttachmentResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *GatewayAttachmentResponse) GetAttachmentStatus() GatewayAttachmentResponseAttachmentStatus`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *GatewayAttachmentResponse) GetAttachmentStatusOk() (*GatewayAttachmentResponseAttachmentStatus, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *GatewayAttachmentResponse) SetAttachmentStatus(v GatewayAttachmentResponseAttachmentStatus)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *GatewayAttachmentResponse) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.

### GetErrors

`func (o *GatewayAttachmentResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GatewayAttachmentResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GatewayAttachmentResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GatewayAttachmentResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetChangeLog

`func (o *GatewayAttachmentResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *GatewayAttachmentResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *GatewayAttachmentResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *GatewayAttachmentResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



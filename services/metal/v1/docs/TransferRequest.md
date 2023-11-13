# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**TargetOrganization** | Pointer to [**Href**](Href.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTransferRequest

`func NewTransferRequest() *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TransferRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransferRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransferRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransferRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *TransferRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TransferRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TransferRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TransferRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *TransferRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *TransferRequest) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TransferRequest) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TransferRequest) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *TransferRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTargetOrganization

`func (o *TransferRequest) GetTargetOrganization() Href`

GetTargetOrganization returns the TargetOrganization field if non-nil, zero value otherwise.

### GetTargetOrganizationOk

`func (o *TransferRequest) GetTargetOrganizationOk() (*Href, bool)`

GetTargetOrganizationOk returns a tuple with the TargetOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrganization

`func (o *TransferRequest) SetTargetOrganization(v Href)`

SetTargetOrganization sets TargetOrganization field to given value.

### HasTargetOrganization

`func (o *TransferRequest) HasTargetOrganization() bool`

HasTargetOrganization returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TransferRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransferRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransferRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransferRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



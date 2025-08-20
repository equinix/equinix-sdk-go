# PicturesDocumentRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CameraProvidedBy** | Pointer to [**PicturesDocumentRequestServiceDetailsCameraProvidedBy**](PicturesDocumentRequestServiceDetailsCameraProvidedBy.md) |  | [optional] 
**SpecificDateAndTime** | Pointer to **bool** | Would you like to request a specific date and time? If true, Scheduling should not be STANDARD. If documentOnly is false, this field is mandatory. | [optional] 
**DocumentOnly** | **bool** | Do you need documents or pictures? Set true if you need documents. | 
**Description** | Pointer to **string** | Descrption of the Photo/Document. If documentOnly is false, this field is mandatory. | [optional] 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 

## Methods

### NewPicturesDocumentRequestServiceDetails

`func NewPicturesDocumentRequestServiceDetails(documentOnly bool, scopeOfWork string, ) *PicturesDocumentRequestServiceDetails`

NewPicturesDocumentRequestServiceDetails instantiates a new PicturesDocumentRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPicturesDocumentRequestServiceDetailsWithDefaults

`func NewPicturesDocumentRequestServiceDetailsWithDefaults() *PicturesDocumentRequestServiceDetails`

NewPicturesDocumentRequestServiceDetailsWithDefaults instantiates a new PicturesDocumentRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCameraProvidedBy

`func (o *PicturesDocumentRequestServiceDetails) GetCameraProvidedBy() PicturesDocumentRequestServiceDetailsCameraProvidedBy`

GetCameraProvidedBy returns the CameraProvidedBy field if non-nil, zero value otherwise.

### GetCameraProvidedByOk

`func (o *PicturesDocumentRequestServiceDetails) GetCameraProvidedByOk() (*PicturesDocumentRequestServiceDetailsCameraProvidedBy, bool)`

GetCameraProvidedByOk returns a tuple with the CameraProvidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraProvidedBy

`func (o *PicturesDocumentRequestServiceDetails) SetCameraProvidedBy(v PicturesDocumentRequestServiceDetailsCameraProvidedBy)`

SetCameraProvidedBy sets CameraProvidedBy field to given value.

### HasCameraProvidedBy

`func (o *PicturesDocumentRequestServiceDetails) HasCameraProvidedBy() bool`

HasCameraProvidedBy returns a boolean if a field has been set.

### GetSpecificDateAndTime

`func (o *PicturesDocumentRequestServiceDetails) GetSpecificDateAndTime() bool`

GetSpecificDateAndTime returns the SpecificDateAndTime field if non-nil, zero value otherwise.

### GetSpecificDateAndTimeOk

`func (o *PicturesDocumentRequestServiceDetails) GetSpecificDateAndTimeOk() (*bool, bool)`

GetSpecificDateAndTimeOk returns a tuple with the SpecificDateAndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificDateAndTime

`func (o *PicturesDocumentRequestServiceDetails) SetSpecificDateAndTime(v bool)`

SetSpecificDateAndTime sets SpecificDateAndTime field to given value.

### HasSpecificDateAndTime

`func (o *PicturesDocumentRequestServiceDetails) HasSpecificDateAndTime() bool`

HasSpecificDateAndTime returns a boolean if a field has been set.

### GetDocumentOnly

`func (o *PicturesDocumentRequestServiceDetails) GetDocumentOnly() bool`

GetDocumentOnly returns the DocumentOnly field if non-nil, zero value otherwise.

### GetDocumentOnlyOk

`func (o *PicturesDocumentRequestServiceDetails) GetDocumentOnlyOk() (*bool, bool)`

GetDocumentOnlyOk returns a tuple with the DocumentOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentOnly

`func (o *PicturesDocumentRequestServiceDetails) SetDocumentOnly(v bool)`

SetDocumentOnly sets DocumentOnly field to given value.


### GetDescription

`func (o *PicturesDocumentRequestServiceDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PicturesDocumentRequestServiceDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PicturesDocumentRequestServiceDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PicturesDocumentRequestServiceDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopeOfWork

`func (o *PicturesDocumentRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *PicturesDocumentRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *PicturesDocumentRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



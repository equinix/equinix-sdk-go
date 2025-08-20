# CageEscortRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationVisit** | [**CageEscortRequestServiceDetailsDurationVisit**](CageEscortRequestServiceDetailsDurationVisit.md) |  | 
**OpenCabinetForVisitor** | **bool** | Open the Cabinet for the Visitor? | 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 
**SupervisionReqForVisitor** | **bool** | Supervision Required For the Visitor? | 
**WorkVisitOrderNumber** | **string** | Work Visit Order Number | 

## Methods

### NewCageEscortRequestServiceDetails

`func NewCageEscortRequestServiceDetails(durationVisit CageEscortRequestServiceDetailsDurationVisit, openCabinetForVisitor bool, scopeOfWork string, supervisionReqForVisitor bool, workVisitOrderNumber string, ) *CageEscortRequestServiceDetails`

NewCageEscortRequestServiceDetails instantiates a new CageEscortRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCageEscortRequestServiceDetailsWithDefaults

`func NewCageEscortRequestServiceDetailsWithDefaults() *CageEscortRequestServiceDetails`

NewCageEscortRequestServiceDetailsWithDefaults instantiates a new CageEscortRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationVisit

`func (o *CageEscortRequestServiceDetails) GetDurationVisit() CageEscortRequestServiceDetailsDurationVisit`

GetDurationVisit returns the DurationVisit field if non-nil, zero value otherwise.

### GetDurationVisitOk

`func (o *CageEscortRequestServiceDetails) GetDurationVisitOk() (*CageEscortRequestServiceDetailsDurationVisit, bool)`

GetDurationVisitOk returns a tuple with the DurationVisit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationVisit

`func (o *CageEscortRequestServiceDetails) SetDurationVisit(v CageEscortRequestServiceDetailsDurationVisit)`

SetDurationVisit sets DurationVisit field to given value.


### GetOpenCabinetForVisitor

`func (o *CageEscortRequestServiceDetails) GetOpenCabinetForVisitor() bool`

GetOpenCabinetForVisitor returns the OpenCabinetForVisitor field if non-nil, zero value otherwise.

### GetOpenCabinetForVisitorOk

`func (o *CageEscortRequestServiceDetails) GetOpenCabinetForVisitorOk() (*bool, bool)`

GetOpenCabinetForVisitorOk returns a tuple with the OpenCabinetForVisitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenCabinetForVisitor

`func (o *CageEscortRequestServiceDetails) SetOpenCabinetForVisitor(v bool)`

SetOpenCabinetForVisitor sets OpenCabinetForVisitor field to given value.


### GetScopeOfWork

`func (o *CageEscortRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *CageEscortRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *CageEscortRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.


### GetSupervisionReqForVisitor

`func (o *CageEscortRequestServiceDetails) GetSupervisionReqForVisitor() bool`

GetSupervisionReqForVisitor returns the SupervisionReqForVisitor field if non-nil, zero value otherwise.

### GetSupervisionReqForVisitorOk

`func (o *CageEscortRequestServiceDetails) GetSupervisionReqForVisitorOk() (*bool, bool)`

GetSupervisionReqForVisitorOk returns a tuple with the SupervisionReqForVisitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisionReqForVisitor

`func (o *CageEscortRequestServiceDetails) SetSupervisionReqForVisitor(v bool)`

SetSupervisionReqForVisitor sets SupervisionReqForVisitor field to given value.


### GetWorkVisitOrderNumber

`func (o *CageEscortRequestServiceDetails) GetWorkVisitOrderNumber() string`

GetWorkVisitOrderNumber returns the WorkVisitOrderNumber field if non-nil, zero value otherwise.

### GetWorkVisitOrderNumberOk

`func (o *CageEscortRequestServiceDetails) GetWorkVisitOrderNumberOk() (*string, bool)`

GetWorkVisitOrderNumberOk returns a tuple with the WorkVisitOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkVisitOrderNumber

`func (o *CageEscortRequestServiceDetails) SetWorkVisitOrderNumber(v string)`

SetWorkVisitOrderNumber sets WorkVisitOrderNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



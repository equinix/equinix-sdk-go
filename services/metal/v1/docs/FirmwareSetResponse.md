# FirmwareSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**FirmwareSet**](FirmwareSet.md) |  | [optional] 

## Methods

### NewFirmwareSetResponse

`func NewFirmwareSetResponse() *FirmwareSetResponse`

NewFirmwareSetResponse instantiates a new FirmwareSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareSetResponseWithDefaults

`func NewFirmwareSetResponseWithDefaults() *FirmwareSetResponse`

NewFirmwareSetResponseWithDefaults instantiates a new FirmwareSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *FirmwareSetResponse) GetRecord() FirmwareSet`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *FirmwareSetResponse) GetRecordOk() (*FirmwareSet, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *FirmwareSetResponse) SetRecord(v FirmwareSet)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *FirmwareSetResponse) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



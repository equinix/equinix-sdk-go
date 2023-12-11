# FirmwareSetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **int32** | Max number of items returned in a page | [optional] 
**Page** | Pointer to **int32** | Page returned | [optional] 
**PageCount** | Pointer to **int32** | Items returned in current page | [optional] 
**TotalPages** | Pointer to **int32** | Total count of pages | [optional] 
**TotalRecordCount** | Pointer to **int32** | Total count of items | [optional] 
**Records** | Pointer to [**[]FirmwareSet**](FirmwareSet.md) | Represents a list of FirmwareSets | [optional] 

## Methods

### NewFirmwareSetListResponse

`func NewFirmwareSetListResponse() *FirmwareSetListResponse`

NewFirmwareSetListResponse instantiates a new FirmwareSetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareSetListResponseWithDefaults

`func NewFirmwareSetListResponseWithDefaults() *FirmwareSetListResponse`

NewFirmwareSetListResponseWithDefaults instantiates a new FirmwareSetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *FirmwareSetListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FirmwareSetListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FirmwareSetListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FirmwareSetListResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPage

`func (o *FirmwareSetListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FirmwareSetListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FirmwareSetListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *FirmwareSetListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageCount

`func (o *FirmwareSetListResponse) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *FirmwareSetListResponse) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *FirmwareSetListResponse) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *FirmwareSetListResponse) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *FirmwareSetListResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *FirmwareSetListResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *FirmwareSetListResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *FirmwareSetListResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *FirmwareSetListResponse) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *FirmwareSetListResponse) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *FirmwareSetListResponse) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *FirmwareSetListResponse) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetRecords

`func (o *FirmwareSetListResponse) GetRecords() []FirmwareSet`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *FirmwareSetListResponse) GetRecordsOk() (*[]FirmwareSet, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *FirmwareSetListResponse) SetRecords(v []FirmwareSet)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *FirmwareSetListResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



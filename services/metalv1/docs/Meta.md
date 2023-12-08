# Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**Href**](Href.md) |  | [optional] 
**Last** | Pointer to [**Href**](Href.md) |  | [optional] 
**Next** | Pointer to [**Href**](Href.md) |  | [optional] 
**Previous** | Pointer to [**Href**](Href.md) |  | [optional] 
**Self** | Pointer to [**Href**](Href.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**LastPage** | Pointer to **int32** |  | [optional] 

## Methods

### NewMeta

`func NewMeta() *Meta`

NewMeta instantiates a new Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaWithDefaults

`func NewMetaWithDefaults() *Meta`

NewMetaWithDefaults instantiates a new Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *Meta) GetFirst() Href`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *Meta) GetFirstOk() (*Href, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *Meta) SetFirst(v Href)`

SetFirst sets First field to given value.

### HasFirst

`func (o *Meta) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *Meta) GetLast() Href`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Meta) GetLastOk() (*Href, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Meta) SetLast(v Href)`

SetLast sets Last field to given value.

### HasLast

`func (o *Meta) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *Meta) GetNext() Href`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Meta) GetNextOk() (*Href, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Meta) SetNext(v Href)`

SetNext sets Next field to given value.

### HasNext

`func (o *Meta) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *Meta) GetPrevious() Href`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Meta) GetPreviousOk() (*Href, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Meta) SetPrevious(v Href)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *Meta) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetSelf

`func (o *Meta) GetSelf() Href`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Meta) GetSelfOk() (*Href, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Meta) SetSelf(v Href)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Meta) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTotal

`func (o *Meta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Meta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Meta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Meta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCurrentPage

`func (o *Meta) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *Meta) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *Meta) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *Meta) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetLastPage

`func (o *Meta) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *Meta) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *Meta) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *Meta) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



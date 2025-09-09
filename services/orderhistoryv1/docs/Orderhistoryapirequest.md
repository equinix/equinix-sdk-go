# Orderhistoryapirequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**Filters**](Filters.md) |  | [optional] 
**Source** | Pointer to [**[]OrderhistoryapirequestSourceInner**](OrderhistoryapirequestSourceInner.md) | &lt;b&gt;ORDER_NUMBER:&lt;/b&gt; Search by order number(1-123456789).&lt;br&gt;&lt;b&gt;CUSTOMER_REFERENCE_NUMBER:&lt;/b&gt; Search by customer reference number which was entered as part place order.&lt;br&gt;&lt;b&gt;TROUBLE_TICKET_NUMBER:&lt;/b&gt; Search by trouble ticket numnber(5-123456).&lt;br&gt;&lt;b&gt;WORK_ACTIVITY_NUMBER:&lt;/b&gt; Search by work order activity number(3-123456).  | [optional] 
**Q** | Pointer to **string** | Query value to be queried against source values(Or operation against all sources).&lt;br/&gt;Supports partial text search | [optional] 
**Sorts** | Pointer to [**[]Sort**](Sort.md) |  | [optional] 
**Page** | Pointer to [**PageRequestModel**](PageRequestModel.md) |  | [optional] 

## Methods

### NewOrderhistoryapirequest

`func NewOrderhistoryapirequest() *Orderhistoryapirequest`

NewOrderhistoryapirequest instantiates a new Orderhistoryapirequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderhistoryapirequestWithDefaults

`func NewOrderhistoryapirequestWithDefaults() *Orderhistoryapirequest`

NewOrderhistoryapirequestWithDefaults instantiates a new Orderhistoryapirequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *Orderhistoryapirequest) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Orderhistoryapirequest) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Orderhistoryapirequest) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Orderhistoryapirequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSource

`func (o *Orderhistoryapirequest) GetSource() []OrderhistoryapirequestSourceInner`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Orderhistoryapirequest) GetSourceOk() (*[]OrderhistoryapirequestSourceInner, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Orderhistoryapirequest) SetSource(v []OrderhistoryapirequestSourceInner)`

SetSource sets Source field to given value.

### HasSource

`func (o *Orderhistoryapirequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetQ

`func (o *Orderhistoryapirequest) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *Orderhistoryapirequest) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *Orderhistoryapirequest) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *Orderhistoryapirequest) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetSorts

`func (o *Orderhistoryapirequest) GetSorts() []Sort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *Orderhistoryapirequest) GetSortsOk() (*[]Sort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *Orderhistoryapirequest) SetSorts(v []Sort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *Orderhistoryapirequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetPage

`func (o *Orderhistoryapirequest) GetPage() PageRequestModel`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Orderhistoryapirequest) GetPageOk() (*PageRequestModel, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Orderhistoryapirequest) SetPage(v PageRequestModel)`

SetPage sets Page field to given value.

### HasPage

`func (o *Orderhistoryapirequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



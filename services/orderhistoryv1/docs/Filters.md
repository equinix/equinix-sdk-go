# Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibxs** | Pointer to **[]string** | ibxs filter | [optional] 
**ProductTypes** | Pointer to [**[]ProductTypesEnum**](ProductTypesEnum.md) | Product(order type) filter | [optional] 
**OrderStatus** | Pointer to [**[]OrderHeaderStatusEnum**](OrderHeaderStatusEnum.md) | order status filter | [optional] 
**DateRange** | Pointer to [**FiltersDateRange**](FiltersDateRange.md) |  | [optional] 
**FromDate** | Pointer to **string** | Date Format Should be mm/dd/yyyy.&lt;br/&gt; Not applicable when dateRange is provided. | [optional] 
**ToDate** | Pointer to **string** | Date Format Should be mm/dd/yyyy.&lt;br/&gt; Not applicable when dateRange is provided | [optional] 

## Methods

### NewFilters

`func NewFilters() *Filters`

NewFilters instantiates a new Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersWithDefaults

`func NewFiltersWithDefaults() *Filters`

NewFiltersWithDefaults instantiates a new Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbxs

`func (o *Filters) GetIbxs() []string`

GetIbxs returns the Ibxs field if non-nil, zero value otherwise.

### GetIbxsOk

`func (o *Filters) GetIbxsOk() (*[]string, bool)`

GetIbxsOk returns a tuple with the Ibxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxs

`func (o *Filters) SetIbxs(v []string)`

SetIbxs sets Ibxs field to given value.

### HasIbxs

`func (o *Filters) HasIbxs() bool`

HasIbxs returns a boolean if a field has been set.

### GetProductTypes

`func (o *Filters) GetProductTypes() []ProductTypesEnum`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *Filters) GetProductTypesOk() (*[]ProductTypesEnum, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *Filters) SetProductTypes(v []ProductTypesEnum)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *Filters) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetOrderStatus

`func (o *Filters) GetOrderStatus() []OrderHeaderStatusEnum`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *Filters) GetOrderStatusOk() (*[]OrderHeaderStatusEnum, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *Filters) SetOrderStatus(v []OrderHeaderStatusEnum)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *Filters) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetDateRange

`func (o *Filters) GetDateRange() FiltersDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Filters) GetDateRangeOk() (*FiltersDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Filters) SetDateRange(v FiltersDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *Filters) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFromDate

`func (o *Filters) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *Filters) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *Filters) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *Filters) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *Filters) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *Filters) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *Filters) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *Filters) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



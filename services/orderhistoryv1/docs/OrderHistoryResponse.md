# OrderHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]OrderHeader**](OrderHeader.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**Page** | Pointer to [**PageDetailsForResponse**](PageDetailsForResponse.md) |  | [optional] 

## Methods

### NewOrderHistoryResponse

`func NewOrderHistoryResponse() *OrderHistoryResponse`

NewOrderHistoryResponse instantiates a new OrderHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderHistoryResponseWithDefaults

`func NewOrderHistoryResponseWithDefaults() *OrderHistoryResponse`

NewOrderHistoryResponseWithDefaults instantiates a new OrderHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *OrderHistoryResponse) GetContent() []OrderHeader`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OrderHistoryResponse) GetContentOk() (*[]OrderHeader, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OrderHistoryResponse) SetContent(v []OrderHeader)`

SetContent sets Content field to given value.

### HasContent

`func (o *OrderHistoryResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLinks

`func (o *OrderHistoryResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderHistoryResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderHistoryResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderHistoryResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPage

`func (o *OrderHistoryResponse) GetPage() PageDetailsForResponse`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrderHistoryResponse) GetPageOk() (*PageDetailsForResponse, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrderHistoryResponse) SetPage(v PageDetailsForResponse)`

SetPage sets Page field to given value.

### HasPage

`func (o *OrderHistoryResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



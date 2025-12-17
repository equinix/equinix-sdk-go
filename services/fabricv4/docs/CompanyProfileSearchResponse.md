# CompanyProfileSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]CompanyProfileResponse**](CompanyProfileResponse.md) |  | 

## Methods

### NewCompanyProfileSearchResponse

`func NewCompanyProfileSearchResponse(pagination Pagination, data []CompanyProfileResponse, ) *CompanyProfileSearchResponse`

NewCompanyProfileSearchResponse instantiates a new CompanyProfileSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileSearchResponseWithDefaults

`func NewCompanyProfileSearchResponseWithDefaults() *CompanyProfileSearchResponse`

NewCompanyProfileSearchResponseWithDefaults instantiates a new CompanyProfileSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *CompanyProfileSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CompanyProfileSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CompanyProfileSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *CompanyProfileSearchResponse) GetData() []CompanyProfileResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompanyProfileSearchResponse) GetDataOk() (*[]CompanyProfileResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompanyProfileSearchResponse) SetData(v []CompanyProfileResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ProviderSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**Data** | Pointer to [**[]ProvidersSearchResponse**](ProvidersSearchResponse.md) |  | [optional] 

## Methods

### NewProviderSearchResponse

`func NewProviderSearchResponse() *ProviderSearchResponse`

NewProviderSearchResponse instantiates a new ProviderSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSearchResponseWithDefaults

`func NewProviderSearchResponseWithDefaults() *ProviderSearchResponse`

NewProviderSearchResponseWithDefaults instantiates a new ProviderSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProviderSearchResponse) GetType() ProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderSearchResponse) GetTypeOk() (*ProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderSearchResponse) SetType(v ProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *ProviderSearchResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *ProviderSearchResponse) GetData() []ProvidersSearchResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProviderSearchResponse) GetDataOk() (*[]ProvidersSearchResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProviderSearchResponse) SetData(v []ProvidersSearchResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ProviderSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



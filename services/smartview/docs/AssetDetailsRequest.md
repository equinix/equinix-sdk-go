# AssetDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** | customer account number | [optional] 
**Classification** | Pointer to **string** | asset classification | [optional] 
**Ibx** | Pointer to **string** | ibx code | [optional] 
**AssetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAssetDetailsRequest

`func NewAssetDetailsRequest() *AssetDetailsRequest`

NewAssetDetailsRequest instantiates a new AssetDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDetailsRequestWithDefaults

`func NewAssetDetailsRequestWithDefaults() *AssetDetailsRequest`

NewAssetDetailsRequestWithDefaults instantiates a new AssetDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *AssetDetailsRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *AssetDetailsRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *AssetDetailsRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *AssetDetailsRequest) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetClassification

`func (o *AssetDetailsRequest) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *AssetDetailsRequest) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *AssetDetailsRequest) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *AssetDetailsRequest) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetIbx

`func (o *AssetDetailsRequest) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *AssetDetailsRequest) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *AssetDetailsRequest) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *AssetDetailsRequest) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetAssetIds

`func (o *AssetDetailsRequest) GetAssetIds() []string`

GetAssetIds returns the AssetIds field if non-nil, zero value otherwise.

### GetAssetIdsOk

`func (o *AssetDetailsRequest) GetAssetIdsOk() (*[]string, bool)`

GetAssetIdsOk returns a tuple with the AssetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIds

`func (o *AssetDetailsRequest) SetAssetIds(v []string)`

SetAssetIds sets AssetIds field to given value.

### HasAssetIds

`func (o *AssetDetailsRequest) HasAssetIds() bool`

HasAssetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



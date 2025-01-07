# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**AssetClassification** | Pointer to [**[]AssetAssetClassificationInner**](AssetAssetClassificationInner.md) | Asset Classification value if specified, otherwise all allowable classifications | [optional] 
**AssetId** | Pointer to **[]string** |  | [optional] 
**Ibx** | **[]string** |  | 

## Methods

### NewAsset

`func NewAsset(accountNumber string, ibx []string, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Asset) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Asset) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Asset) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAssetClassification

`func (o *Asset) GetAssetClassification() []AssetAssetClassificationInner`

GetAssetClassification returns the AssetClassification field if non-nil, zero value otherwise.

### GetAssetClassificationOk

`func (o *Asset) GetAssetClassificationOk() (*[]AssetAssetClassificationInner, bool)`

GetAssetClassificationOk returns a tuple with the AssetClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClassification

`func (o *Asset) SetAssetClassification(v []AssetAssetClassificationInner)`

SetAssetClassification sets AssetClassification field to given value.

### HasAssetClassification

`func (o *Asset) HasAssetClassification() bool`

HasAssetClassification returns a boolean if a field has been set.

### GetAssetId

`func (o *Asset) GetAssetId() []string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Asset) GetAssetIdOk() (*[]string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Asset) SetAssetId(v []string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Asset) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetIbx

`func (o *Asset) GetIbx() []string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Asset) GetIbxOk() (*[]string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Asset) SetIbx(v []string)`

SetIbx sets Ibx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



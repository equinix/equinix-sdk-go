# AssetsArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetClassification** | Pointer to **string** | Asset   classification is electrical or mechanical  | [optional] 
**AssetId** | Pointer to **string** | The assetid is the circuit number, sensor id, asset id, asset id  for type circuit, sensor, electrical and mechanical resp.  | [optional] 
**AssetLabel** | Pointer to **string** | Asset Label is the Circuit display label, Sensor ID, and Asset ID for types circuit, sensor, electrical and mechanical resp.   | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetsArray

`func NewAssetsArray() *AssetsArray`

NewAssetsArray instantiates a new AssetsArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsArrayWithDefaults

`func NewAssetsArrayWithDefaults() *AssetsArray`

NewAssetsArrayWithDefaults instantiates a new AssetsArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetClassification

`func (o *AssetsArray) GetAssetClassification() string`

GetAssetClassification returns the AssetClassification field if non-nil, zero value otherwise.

### GetAssetClassificationOk

`func (o *AssetsArray) GetAssetClassificationOk() (*string, bool)`

GetAssetClassificationOk returns a tuple with the AssetClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClassification

`func (o *AssetsArray) SetAssetClassification(v string)`

SetAssetClassification sets AssetClassification field to given value.

### HasAssetClassification

`func (o *AssetsArray) HasAssetClassification() bool`

HasAssetClassification returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetsArray) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetsArray) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetsArray) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetsArray) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetLabel

`func (o *AssetsArray) GetAssetLabel() string`

GetAssetLabel returns the AssetLabel field if non-nil, zero value otherwise.

### GetAssetLabelOk

`func (o *AssetsArray) GetAssetLabelOk() (*string, bool)`

GetAssetLabelOk returns a tuple with the AssetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLabel

`func (o *AssetsArray) SetAssetLabel(v string)`

SetAssetLabel sets AssetLabel field to given value.

### HasAssetLabel

`func (o *AssetsArray) HasAssetLabel() bool`

HasAssetLabel returns a boolean if a field has been set.

### GetType

`func (o *AssetsArray) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetsArray) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetsArray) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssetsArray) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



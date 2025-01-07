# Assets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**AssetsPayLoad**](AssetsPayLoad.md) |  | [optional] 
**Status** | Pointer to [**AssetsStatus**](AssetsStatus.md) |  | [optional] 

## Methods

### NewAssets

`func NewAssets() *Assets`

NewAssets instantiates a new Assets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsWithDefaults

`func NewAssetsWithDefaults() *Assets`

NewAssetsWithDefaults instantiates a new Assets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *Assets) GetPayLoad() AssetsPayLoad`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *Assets) GetPayLoadOk() (*AssetsPayLoad, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *Assets) SetPayLoad(v AssetsPayLoad)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *Assets) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *Assets) GetStatus() AssetsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Assets) GetStatusOk() (*AssetsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Assets) SetStatus(v AssetsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Assets) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



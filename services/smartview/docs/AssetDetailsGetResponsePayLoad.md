# AssetDetailsGetResponsePayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmLastProcessedTime** | Pointer to **string** | Datetime when the latest alarm was processed on the asset | [optional] 
**AlarmLastTriggeredTime** | Pointer to **string** | Datetime when the latest alarm was triggered on the asset | [optional] 
**AssetId** | Pointer to **string** | asset id for the Asset | [optional] 
**AssetType** | Pointer to **string** | Template Name for the asset | [optional] 
**EquipmentModelNumber** | Pointer to **string** | Equipment Model Number | [optional] 
**EquipmentSerialNumber** | Pointer to **string** | Equipment Serial Number | [optional] 
**LastMaintenanceDate** | Pointer to **string** | Datetime when the machine had its last maintenance | [optional] 
**ManufacturerName** | Pointer to **string** | Manufacturer name for the Asset | [optional] 
**Tags** | Pointer to [**[]TagPointDataArray**](TagPointDataArray.md) | List of tag points for the Asset | [optional] 
**UserPrefTimeZone** | Pointer to **string** | Time zone for the user | [optional] 

## Methods

### NewAssetDetailsGetResponsePayLoad

`func NewAssetDetailsGetResponsePayLoad() *AssetDetailsGetResponsePayLoad`

NewAssetDetailsGetResponsePayLoad instantiates a new AssetDetailsGetResponsePayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDetailsGetResponsePayLoadWithDefaults

`func NewAssetDetailsGetResponsePayLoadWithDefaults() *AssetDetailsGetResponsePayLoad`

NewAssetDetailsGetResponsePayLoadWithDefaults instantiates a new AssetDetailsGetResponsePayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmLastProcessedTime

`func (o *AssetDetailsGetResponsePayLoad) GetAlarmLastProcessedTime() string`

GetAlarmLastProcessedTime returns the AlarmLastProcessedTime field if non-nil, zero value otherwise.

### GetAlarmLastProcessedTimeOk

`func (o *AssetDetailsGetResponsePayLoad) GetAlarmLastProcessedTimeOk() (*string, bool)`

GetAlarmLastProcessedTimeOk returns a tuple with the AlarmLastProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmLastProcessedTime

`func (o *AssetDetailsGetResponsePayLoad) SetAlarmLastProcessedTime(v string)`

SetAlarmLastProcessedTime sets AlarmLastProcessedTime field to given value.

### HasAlarmLastProcessedTime

`func (o *AssetDetailsGetResponsePayLoad) HasAlarmLastProcessedTime() bool`

HasAlarmLastProcessedTime returns a boolean if a field has been set.

### GetAlarmLastTriggeredTime

`func (o *AssetDetailsGetResponsePayLoad) GetAlarmLastTriggeredTime() string`

GetAlarmLastTriggeredTime returns the AlarmLastTriggeredTime field if non-nil, zero value otherwise.

### GetAlarmLastTriggeredTimeOk

`func (o *AssetDetailsGetResponsePayLoad) GetAlarmLastTriggeredTimeOk() (*string, bool)`

GetAlarmLastTriggeredTimeOk returns a tuple with the AlarmLastTriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmLastTriggeredTime

`func (o *AssetDetailsGetResponsePayLoad) SetAlarmLastTriggeredTime(v string)`

SetAlarmLastTriggeredTime sets AlarmLastTriggeredTime field to given value.

### HasAlarmLastTriggeredTime

`func (o *AssetDetailsGetResponsePayLoad) HasAlarmLastTriggeredTime() bool`

HasAlarmLastTriggeredTime returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetDetailsGetResponsePayLoad) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetDetailsGetResponsePayLoad) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetDetailsGetResponsePayLoad) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetDetailsGetResponsePayLoad) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetType

`func (o *AssetDetailsGetResponsePayLoad) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetDetailsGetResponsePayLoad) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetDetailsGetResponsePayLoad) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AssetDetailsGetResponsePayLoad) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetEquipmentModelNumber

`func (o *AssetDetailsGetResponsePayLoad) GetEquipmentModelNumber() string`

GetEquipmentModelNumber returns the EquipmentModelNumber field if non-nil, zero value otherwise.

### GetEquipmentModelNumberOk

`func (o *AssetDetailsGetResponsePayLoad) GetEquipmentModelNumberOk() (*string, bool)`

GetEquipmentModelNumberOk returns a tuple with the EquipmentModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentModelNumber

`func (o *AssetDetailsGetResponsePayLoad) SetEquipmentModelNumber(v string)`

SetEquipmentModelNumber sets EquipmentModelNumber field to given value.

### HasEquipmentModelNumber

`func (o *AssetDetailsGetResponsePayLoad) HasEquipmentModelNumber() bool`

HasEquipmentModelNumber returns a boolean if a field has been set.

### GetEquipmentSerialNumber

`func (o *AssetDetailsGetResponsePayLoad) GetEquipmentSerialNumber() string`

GetEquipmentSerialNumber returns the EquipmentSerialNumber field if non-nil, zero value otherwise.

### GetEquipmentSerialNumberOk

`func (o *AssetDetailsGetResponsePayLoad) GetEquipmentSerialNumberOk() (*string, bool)`

GetEquipmentSerialNumberOk returns a tuple with the EquipmentSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSerialNumber

`func (o *AssetDetailsGetResponsePayLoad) SetEquipmentSerialNumber(v string)`

SetEquipmentSerialNumber sets EquipmentSerialNumber field to given value.

### HasEquipmentSerialNumber

`func (o *AssetDetailsGetResponsePayLoad) HasEquipmentSerialNumber() bool`

HasEquipmentSerialNumber returns a boolean if a field has been set.

### GetLastMaintenanceDate

`func (o *AssetDetailsGetResponsePayLoad) GetLastMaintenanceDate() string`

GetLastMaintenanceDate returns the LastMaintenanceDate field if non-nil, zero value otherwise.

### GetLastMaintenanceDateOk

`func (o *AssetDetailsGetResponsePayLoad) GetLastMaintenanceDateOk() (*string, bool)`

GetLastMaintenanceDateOk returns a tuple with the LastMaintenanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMaintenanceDate

`func (o *AssetDetailsGetResponsePayLoad) SetLastMaintenanceDate(v string)`

SetLastMaintenanceDate sets LastMaintenanceDate field to given value.

### HasLastMaintenanceDate

`func (o *AssetDetailsGetResponsePayLoad) HasLastMaintenanceDate() bool`

HasLastMaintenanceDate returns a boolean if a field has been set.

### GetManufacturerName

`func (o *AssetDetailsGetResponsePayLoad) GetManufacturerName() string`

GetManufacturerName returns the ManufacturerName field if non-nil, zero value otherwise.

### GetManufacturerNameOk

`func (o *AssetDetailsGetResponsePayLoad) GetManufacturerNameOk() (*string, bool)`

GetManufacturerNameOk returns a tuple with the ManufacturerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerName

`func (o *AssetDetailsGetResponsePayLoad) SetManufacturerName(v string)`

SetManufacturerName sets ManufacturerName field to given value.

### HasManufacturerName

`func (o *AssetDetailsGetResponsePayLoad) HasManufacturerName() bool`

HasManufacturerName returns a boolean if a field has been set.

### GetTags

`func (o *AssetDetailsGetResponsePayLoad) GetTags() []TagPointDataArray`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetDetailsGetResponsePayLoad) GetTagsOk() (*[]TagPointDataArray, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetDetailsGetResponsePayLoad) SetTags(v []TagPointDataArray)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetDetailsGetResponsePayLoad) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserPrefTimeZone

`func (o *AssetDetailsGetResponsePayLoad) GetUserPrefTimeZone() string`

GetUserPrefTimeZone returns the UserPrefTimeZone field if non-nil, zero value otherwise.

### GetUserPrefTimeZoneOk

`func (o *AssetDetailsGetResponsePayLoad) GetUserPrefTimeZoneOk() (*string, bool)`

GetUserPrefTimeZoneOk returns a tuple with the UserPrefTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefTimeZone

`func (o *AssetDetailsGetResponsePayLoad) SetUserPrefTimeZone(v string)`

SetUserPrefTimeZone sets UserPrefTimeZone field to given value.

### HasUserPrefTimeZone

`func (o *AssetDetailsGetResponsePayLoad) HasUserPrefTimeZone() bool`

HasUserPrefTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



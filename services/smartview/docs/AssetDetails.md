# AssetDetails

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

### NewAssetDetails

`func NewAssetDetails() *AssetDetails`

NewAssetDetails instantiates a new AssetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDetailsWithDefaults

`func NewAssetDetailsWithDefaults() *AssetDetails`

NewAssetDetailsWithDefaults instantiates a new AssetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmLastProcessedTime

`func (o *AssetDetails) GetAlarmLastProcessedTime() string`

GetAlarmLastProcessedTime returns the AlarmLastProcessedTime field if non-nil, zero value otherwise.

### GetAlarmLastProcessedTimeOk

`func (o *AssetDetails) GetAlarmLastProcessedTimeOk() (*string, bool)`

GetAlarmLastProcessedTimeOk returns a tuple with the AlarmLastProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmLastProcessedTime

`func (o *AssetDetails) SetAlarmLastProcessedTime(v string)`

SetAlarmLastProcessedTime sets AlarmLastProcessedTime field to given value.

### HasAlarmLastProcessedTime

`func (o *AssetDetails) HasAlarmLastProcessedTime() bool`

HasAlarmLastProcessedTime returns a boolean if a field has been set.

### GetAlarmLastTriggeredTime

`func (o *AssetDetails) GetAlarmLastTriggeredTime() string`

GetAlarmLastTriggeredTime returns the AlarmLastTriggeredTime field if non-nil, zero value otherwise.

### GetAlarmLastTriggeredTimeOk

`func (o *AssetDetails) GetAlarmLastTriggeredTimeOk() (*string, bool)`

GetAlarmLastTriggeredTimeOk returns a tuple with the AlarmLastTriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmLastTriggeredTime

`func (o *AssetDetails) SetAlarmLastTriggeredTime(v string)`

SetAlarmLastTriggeredTime sets AlarmLastTriggeredTime field to given value.

### HasAlarmLastTriggeredTime

`func (o *AssetDetails) HasAlarmLastTriggeredTime() bool`

HasAlarmLastTriggeredTime returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetDetails) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetDetails) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetDetails) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetDetails) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetType

`func (o *AssetDetails) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetDetails) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetDetails) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AssetDetails) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetEquipmentModelNumber

`func (o *AssetDetails) GetEquipmentModelNumber() string`

GetEquipmentModelNumber returns the EquipmentModelNumber field if non-nil, zero value otherwise.

### GetEquipmentModelNumberOk

`func (o *AssetDetails) GetEquipmentModelNumberOk() (*string, bool)`

GetEquipmentModelNumberOk returns a tuple with the EquipmentModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentModelNumber

`func (o *AssetDetails) SetEquipmentModelNumber(v string)`

SetEquipmentModelNumber sets EquipmentModelNumber field to given value.

### HasEquipmentModelNumber

`func (o *AssetDetails) HasEquipmentModelNumber() bool`

HasEquipmentModelNumber returns a boolean if a field has been set.

### GetEquipmentSerialNumber

`func (o *AssetDetails) GetEquipmentSerialNumber() string`

GetEquipmentSerialNumber returns the EquipmentSerialNumber field if non-nil, zero value otherwise.

### GetEquipmentSerialNumberOk

`func (o *AssetDetails) GetEquipmentSerialNumberOk() (*string, bool)`

GetEquipmentSerialNumberOk returns a tuple with the EquipmentSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSerialNumber

`func (o *AssetDetails) SetEquipmentSerialNumber(v string)`

SetEquipmentSerialNumber sets EquipmentSerialNumber field to given value.

### HasEquipmentSerialNumber

`func (o *AssetDetails) HasEquipmentSerialNumber() bool`

HasEquipmentSerialNumber returns a boolean if a field has been set.

### GetLastMaintenanceDate

`func (o *AssetDetails) GetLastMaintenanceDate() string`

GetLastMaintenanceDate returns the LastMaintenanceDate field if non-nil, zero value otherwise.

### GetLastMaintenanceDateOk

`func (o *AssetDetails) GetLastMaintenanceDateOk() (*string, bool)`

GetLastMaintenanceDateOk returns a tuple with the LastMaintenanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMaintenanceDate

`func (o *AssetDetails) SetLastMaintenanceDate(v string)`

SetLastMaintenanceDate sets LastMaintenanceDate field to given value.

### HasLastMaintenanceDate

`func (o *AssetDetails) HasLastMaintenanceDate() bool`

HasLastMaintenanceDate returns a boolean if a field has been set.

### GetManufacturerName

`func (o *AssetDetails) GetManufacturerName() string`

GetManufacturerName returns the ManufacturerName field if non-nil, zero value otherwise.

### GetManufacturerNameOk

`func (o *AssetDetails) GetManufacturerNameOk() (*string, bool)`

GetManufacturerNameOk returns a tuple with the ManufacturerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerName

`func (o *AssetDetails) SetManufacturerName(v string)`

SetManufacturerName sets ManufacturerName field to given value.

### HasManufacturerName

`func (o *AssetDetails) HasManufacturerName() bool`

HasManufacturerName returns a boolean if a field has been set.

### GetTags

`func (o *AssetDetails) GetTags() []TagPointDataArray`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetDetails) GetTagsOk() (*[]TagPointDataArray, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetDetails) SetTags(v []TagPointDataArray)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserPrefTimeZone

`func (o *AssetDetails) GetUserPrefTimeZone() string`

GetUserPrefTimeZone returns the UserPrefTimeZone field if non-nil, zero value otherwise.

### GetUserPrefTimeZoneOk

`func (o *AssetDetails) GetUserPrefTimeZoneOk() (*string, bool)`

GetUserPrefTimeZoneOk returns a tuple with the UserPrefTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefTimeZone

`func (o *AssetDetails) SetUserPrefTimeZone(v string)`

SetUserPrefTimeZone sets UserPrefTimeZone field to given value.

### HasUserPrefTimeZone

`func (o *AssetDetails) HasUserPrefTimeZone() bool`

HasUserPrefTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



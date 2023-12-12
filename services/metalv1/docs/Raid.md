# Raid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to **[]string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewRaid

`func NewRaid() *Raid`

NewRaid instantiates a new Raid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaidWithDefaults

`func NewRaidWithDefaults() *Raid`

NewRaidWithDefaults instantiates a new Raid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *Raid) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Raid) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Raid) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Raid) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetLevel

`func (o *Raid) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Raid) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Raid) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Raid) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *Raid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Raid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Raid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Raid) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



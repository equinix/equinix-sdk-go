# Mount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Point** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMount

`func NewMount() *Mount`

NewMount instantiates a new Mount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountWithDefaults

`func NewMountWithDefaults() *Mount`

NewMountWithDefaults instantiates a new Mount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *Mount) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Mount) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Mount) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Mount) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFormat

`func (o *Mount) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Mount) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Mount) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Mount) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPoint

`func (o *Mount) GetPoint() string`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *Mount) GetPointOk() (*string, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *Mount) SetPoint(v string)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *Mount) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetOptions

`func (o *Mount) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Mount) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Mount) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Mount) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



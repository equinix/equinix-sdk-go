# LogoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | ***os.File** | Logo image file | 
**Name** | **string** | Name of the Logo | 
**Description** | **string** | Description of the logo | 
**Type** | **string** | Type of logo | 

## Methods

### NewLogoRequest

`func NewLogoRequest(logo *os.File, name string, description string, type_ string, ) *LogoRequest`

NewLogoRequest instantiates a new LogoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoRequestWithDefaults

`func NewLogoRequestWithDefaults() *LogoRequest`

NewLogoRequestWithDefaults instantiates a new LogoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *LogoRequest) GetLogo() *os.File`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *LogoRequest) GetLogoOk() (**os.File, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *LogoRequest) SetLogo(v *os.File)`

SetLogo sets Logo field to given value.


### GetName

`func (o *LogoRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogoRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogoRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LogoRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogoRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogoRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *LogoRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogoRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogoRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



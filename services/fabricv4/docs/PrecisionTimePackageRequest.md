# PrecisionTimePackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Code** | [**GetTimeServicesPackageByCodePackageCodeParameter**](GetTimeServicesPackageByCodePackageCodeParameter.md) |  | 

## Methods

### NewPrecisionTimePackageRequest

`func NewPrecisionTimePackageRequest(code GetTimeServicesPackageByCodePackageCodeParameter, ) *PrecisionTimePackageRequest`

NewPrecisionTimePackageRequest instantiates a new PrecisionTimePackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimePackageRequestWithDefaults

`func NewPrecisionTimePackageRequestWithDefaults() *PrecisionTimePackageRequest`

NewPrecisionTimePackageRequestWithDefaults instantiates a new PrecisionTimePackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PrecisionTimePackageRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrecisionTimePackageRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrecisionTimePackageRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrecisionTimePackageRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCode

`func (o *PrecisionTimePackageRequest) GetCode() GetTimeServicesPackageByCodePackageCodeParameter`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrecisionTimePackageRequest) GetCodeOk() (*GetTimeServicesPackageByCodePackageCodeParameter, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrecisionTimePackageRequest) SetCode(v GetTimeServicesPackageByCodePackageCodeParameter)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



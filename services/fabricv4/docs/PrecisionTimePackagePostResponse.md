# PrecisionTimePackagePostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Code** | [**GetTimeServicesPackageByCodePackageCodeParameter**](GetTimeServicesPackageByCodePackageCodeParameter.md) |  | 

## Methods

### NewPrecisionTimePackagePostResponse

`func NewPrecisionTimePackagePostResponse(code GetTimeServicesPackageByCodePackageCodeParameter, ) *PrecisionTimePackagePostResponse`

NewPrecisionTimePackagePostResponse instantiates a new PrecisionTimePackagePostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimePackagePostResponseWithDefaults

`func NewPrecisionTimePackagePostResponseWithDefaults() *PrecisionTimePackagePostResponse`

NewPrecisionTimePackagePostResponseWithDefaults instantiates a new PrecisionTimePackagePostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PrecisionTimePackagePostResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrecisionTimePackagePostResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrecisionTimePackagePostResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrecisionTimePackagePostResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCode

`func (o *PrecisionTimePackagePostResponse) GetCode() GetTimeServicesPackageByCodePackageCodeParameter`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrecisionTimePackagePostResponse) GetCodeOk() (*GetTimeServicesPackageByCodePackageCodeParameter, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrecisionTimePackagePostResponse) SetCode(v GetTimeServicesPackageByCodePackageCodeParameter)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



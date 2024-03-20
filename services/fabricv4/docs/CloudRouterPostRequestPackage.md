# CloudRouterPostRequestPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Fabric Cloud Router URI | [optional] 
**Type** | Pointer to [**CloudRouterPostRequestPackageType**](CloudRouterPostRequestPackageType.md) |  | [optional] 
**Code** | [**CloudRouterPostRequestPackageCode**](CloudRouterPostRequestPackageCode.md) |  | 

## Methods

### NewCloudRouterPostRequestPackage

`func NewCloudRouterPostRequestPackage(code CloudRouterPostRequestPackageCode, ) *CloudRouterPostRequestPackage`

NewCloudRouterPostRequestPackage instantiates a new CloudRouterPostRequestPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterPostRequestPackageWithDefaults

`func NewCloudRouterPostRequestPackageWithDefaults() *CloudRouterPostRequestPackage`

NewCloudRouterPostRequestPackageWithDefaults instantiates a new CloudRouterPostRequestPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CloudRouterPostRequestPackage) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudRouterPostRequestPackage) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudRouterPostRequestPackage) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudRouterPostRequestPackage) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *CloudRouterPostRequestPackage) GetType() CloudRouterPostRequestPackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterPostRequestPackage) GetTypeOk() (*CloudRouterPostRequestPackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterPostRequestPackage) SetType(v CloudRouterPostRequestPackageType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouterPostRequestPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *CloudRouterPostRequestPackage) GetCode() CloudRouterPostRequestPackageCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudRouterPostRequestPackage) GetCodeOk() (*CloudRouterPostRequestPackageCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudRouterPostRequestPackage) SetCode(v CloudRouterPostRequestPackageCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



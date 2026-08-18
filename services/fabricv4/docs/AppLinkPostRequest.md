# AppLinkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AppLinkType**](AppLinkType.md) |  | [default to APPLINKTYPE_APP_LINK]
**Name** | **string** | Customer-provided App Link name | 
**Description** | Pointer to **string** | Customer-provided App Link description | [optional] 
**Router** | [**AppLinkPostRequestRouter**](AppLinkPostRequestRouter.md) |  | 
**Ipv4Address** | Pointer to **string** | AppLink IP address | [optional] 
**Bandwidth** | **int32** | App Link aggregated data transfer capacity in Mbps | 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewAppLinkPostRequest

`func NewAppLinkPostRequest(type_ AppLinkType, name string, router AppLinkPostRequestRouter, bandwidth int32, project Project, ) *AppLinkPostRequest`

NewAppLinkPostRequest instantiates a new AppLinkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkPostRequestWithDefaults

`func NewAppLinkPostRequestWithDefaults() *AppLinkPostRequest`

NewAppLinkPostRequestWithDefaults instantiates a new AppLinkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppLinkPostRequest) GetType() AppLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkPostRequest) GetTypeOk() (*AppLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkPostRequest) SetType(v AppLinkType)`

SetType sets Type field to given value.


### GetName

`func (o *AppLinkPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLinkPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLinkPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppLinkPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppLinkPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppLinkPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppLinkPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRouter

`func (o *AppLinkPostRequest) GetRouter() AppLinkPostRequestRouter`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *AppLinkPostRequest) GetRouterOk() (*AppLinkPostRequestRouter, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *AppLinkPostRequest) SetRouter(v AppLinkPostRequestRouter)`

SetRouter sets Router field to given value.


### GetIpv4Address

`func (o *AppLinkPostRequest) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *AppLinkPostRequest) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *AppLinkPostRequest) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *AppLinkPostRequest) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetBandwidth

`func (o *AppLinkPostRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *AppLinkPostRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *AppLinkPostRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetProject

`func (o *AppLinkPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppLinkPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppLinkPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



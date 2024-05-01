# ServiceOrderReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**Uuid** | **string** |  | 
**Type** | [**ServiceOrderType**](ServiceOrderType.md) |  | 

## Methods

### NewServiceOrderReference

`func NewServiceOrderReference(href string, uuid string, type_ ServiceOrderType, ) *ServiceOrderReference`

NewServiceOrderReference instantiates a new ServiceOrderReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOrderReferenceWithDefaults

`func NewServiceOrderReferenceWithDefaults() *ServiceOrderReference`

NewServiceOrderReferenceWithDefaults instantiates a new ServiceOrderReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ServiceOrderReference) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceOrderReference) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceOrderReference) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *ServiceOrderReference) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceOrderReference) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceOrderReference) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *ServiceOrderReference) GetType() ServiceOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceOrderReference) GetTypeOk() (*ServiceOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceOrderReference) SetType(v ServiceOrderType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



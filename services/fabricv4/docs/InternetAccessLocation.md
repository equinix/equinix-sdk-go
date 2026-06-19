# InternetAccessLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetroHref** | **string** | Metro URL path for the linked resource | 
**MetroCode** | **string** | Code representing the metro | 
**Region** | [**InternetAccessLocationRegion**](InternetAccessLocationRegion.md) |  | 
**Ibx** | **string** | IBX data center code | 

## Methods

### NewInternetAccessLocation

`func NewInternetAccessLocation(metroHref string, metroCode string, region InternetAccessLocationRegion, ibx string, ) *InternetAccessLocation`

NewInternetAccessLocation instantiates a new InternetAccessLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessLocationWithDefaults

`func NewInternetAccessLocationWithDefaults() *InternetAccessLocation`

NewInternetAccessLocationWithDefaults instantiates a new InternetAccessLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetroHref

`func (o *InternetAccessLocation) GetMetroHref() string`

GetMetroHref returns the MetroHref field if non-nil, zero value otherwise.

### GetMetroHrefOk

`func (o *InternetAccessLocation) GetMetroHrefOk() (*string, bool)`

GetMetroHrefOk returns a tuple with the MetroHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroHref

`func (o *InternetAccessLocation) SetMetroHref(v string)`

SetMetroHref sets MetroHref field to given value.


### GetMetroCode

`func (o *InternetAccessLocation) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *InternetAccessLocation) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *InternetAccessLocation) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.


### GetRegion

`func (o *InternetAccessLocation) GetRegion() InternetAccessLocationRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InternetAccessLocation) GetRegionOk() (*InternetAccessLocationRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InternetAccessLocation) SetRegion(v InternetAccessLocationRegion)`

SetRegion sets Region field to given value.


### GetIbx

`func (o *InternetAccessLocation) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *InternetAccessLocation) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *InternetAccessLocation) SetIbx(v string)`

SetIbx sets Ibx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



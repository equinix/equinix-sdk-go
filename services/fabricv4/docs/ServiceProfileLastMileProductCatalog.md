# ServiceProfileLastMileProductCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Last-mile provider or catalog name. | [optional] 
**Id** | Pointer to **string** | Last-mile provider or catalog identifier. | [optional] 
**WebUrl** | Pointer to **string** | Last-mile catalog or provider website URL. | [optional] 
**DeliveryDate** | Pointer to [**ServiceProfileLastMileDeliveryDateRange**](ServiceProfileLastMileDeliveryDateRange.md) |  | [optional] 
**BillingTerms** | Pointer to [**[]ServiceProfileLastMileProductCatalogBillingTermsInner**](ServiceProfileLastMileProductCatalogBillingTermsInner.md) | Available billing terms in months. | [optional] 
**Offerings** | Pointer to [**[]ServiceProfileLastMileOffering**](ServiceProfileLastMileOffering.md) | Available last-mile offerings. | [optional] 

## Methods

### NewServiceProfileLastMileProductCatalog

`func NewServiceProfileLastMileProductCatalog() *ServiceProfileLastMileProductCatalog`

NewServiceProfileLastMileProductCatalog instantiates a new ServiceProfileLastMileProductCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileLastMileProductCatalogWithDefaults

`func NewServiceProfileLastMileProductCatalogWithDefaults() *ServiceProfileLastMileProductCatalog`

NewServiceProfileLastMileProductCatalogWithDefaults instantiates a new ServiceProfileLastMileProductCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceProfileLastMileProductCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceProfileLastMileProductCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceProfileLastMileProductCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceProfileLastMileProductCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ServiceProfileLastMileProductCatalog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceProfileLastMileProductCatalog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceProfileLastMileProductCatalog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceProfileLastMileProductCatalog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWebUrl

`func (o *ServiceProfileLastMileProductCatalog) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ServiceProfileLastMileProductCatalog) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ServiceProfileLastMileProductCatalog) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ServiceProfileLastMileProductCatalog) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *ServiceProfileLastMileProductCatalog) GetDeliveryDate() ServiceProfileLastMileDeliveryDateRange`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ServiceProfileLastMileProductCatalog) GetDeliveryDateOk() (*ServiceProfileLastMileDeliveryDateRange, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ServiceProfileLastMileProductCatalog) SetDeliveryDate(v ServiceProfileLastMileDeliveryDateRange)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ServiceProfileLastMileProductCatalog) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetBillingTerms

`func (o *ServiceProfileLastMileProductCatalog) GetBillingTerms() []ServiceProfileLastMileProductCatalogBillingTermsInner`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *ServiceProfileLastMileProductCatalog) GetBillingTermsOk() (*[]ServiceProfileLastMileProductCatalogBillingTermsInner, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *ServiceProfileLastMileProductCatalog) SetBillingTerms(v []ServiceProfileLastMileProductCatalogBillingTermsInner)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *ServiceProfileLastMileProductCatalog) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### SetBillingTermsNil

`func (o *ServiceProfileLastMileProductCatalog) SetBillingTermsNil(b bool)`

 SetBillingTermsNil sets the value for BillingTerms to be an explicit nil

### UnsetBillingTerms
`func (o *ServiceProfileLastMileProductCatalog) UnsetBillingTerms()`

UnsetBillingTerms ensures that no value is present for BillingTerms, not even an explicit nil
### GetOfferings

`func (o *ServiceProfileLastMileProductCatalog) GetOfferings() []ServiceProfileLastMileOffering`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *ServiceProfileLastMileProductCatalog) GetOfferingsOk() (*[]ServiceProfileLastMileOffering, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *ServiceProfileLastMileProductCatalog) SetOfferings(v []ServiceProfileLastMileOffering)`

SetOfferings sets Offerings field to given value.

### HasOfferings

`func (o *ServiceProfileLastMileProductCatalog) HasOfferings() bool`

HasOfferings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



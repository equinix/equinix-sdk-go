# ServiceProfileLastMileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiIntegration** | Pointer to [**ServiceProfileLastMileApiIntegration**](ServiceProfileLastMileApiIntegration.md) |  | [optional] 
**ProductCatalogs** | Pointer to [**[]ServiceProfileLastMileProductCatalog**](ServiceProfileLastMileProductCatalog.md) | Last-mile provider catalogs. | [optional] 

## Methods

### NewServiceProfileLastMileConfig

`func NewServiceProfileLastMileConfig() *ServiceProfileLastMileConfig`

NewServiceProfileLastMileConfig instantiates a new ServiceProfileLastMileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileLastMileConfigWithDefaults

`func NewServiceProfileLastMileConfigWithDefaults() *ServiceProfileLastMileConfig`

NewServiceProfileLastMileConfigWithDefaults instantiates a new ServiceProfileLastMileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiIntegration

`func (o *ServiceProfileLastMileConfig) GetApiIntegration() ServiceProfileLastMileApiIntegration`

GetApiIntegration returns the ApiIntegration field if non-nil, zero value otherwise.

### GetApiIntegrationOk

`func (o *ServiceProfileLastMileConfig) GetApiIntegrationOk() (*ServiceProfileLastMileApiIntegration, bool)`

GetApiIntegrationOk returns a tuple with the ApiIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIntegration

`func (o *ServiceProfileLastMileConfig) SetApiIntegration(v ServiceProfileLastMileApiIntegration)`

SetApiIntegration sets ApiIntegration field to given value.

### HasApiIntegration

`func (o *ServiceProfileLastMileConfig) HasApiIntegration() bool`

HasApiIntegration returns a boolean if a field has been set.

### GetProductCatalogs

`func (o *ServiceProfileLastMileConfig) GetProductCatalogs() []ServiceProfileLastMileProductCatalog`

GetProductCatalogs returns the ProductCatalogs field if non-nil, zero value otherwise.

### GetProductCatalogsOk

`func (o *ServiceProfileLastMileConfig) GetProductCatalogsOk() (*[]ServiceProfileLastMileProductCatalog, bool)`

GetProductCatalogsOk returns a tuple with the ProductCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCatalogs

`func (o *ServiceProfileLastMileConfig) SetProductCatalogs(v []ServiceProfileLastMileProductCatalog)`

SetProductCatalogs sets ProductCatalogs field to given value.

### HasProductCatalogs

`func (o *ServiceProfileLastMileConfig) HasProductCatalogs() bool`

HasProductCatalogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



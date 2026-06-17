# ServiceProfileLastMileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**ServiceProfileLastMileAddress**](ServiceProfileLastMileAddress.md) |  | [optional] 
**ProductCatalogs** | Pointer to [**[]ServiceProfileLastMileProductCatalog**](ServiceProfileLastMileProductCatalog.md) | Last-mile provider catalogs. | [optional] 
**Notifications** | Pointer to [**[]ServiceProfileLastMileNotification**](ServiceProfileLastMileNotification.md) | Contact details for notifications related to last-mile provisioning and ordering. | [optional] 

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

### GetAddress

`func (o *ServiceProfileLastMileConfig) GetAddress() ServiceProfileLastMileAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServiceProfileLastMileConfig) GetAddressOk() (*ServiceProfileLastMileAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServiceProfileLastMileConfig) SetAddress(v ServiceProfileLastMileAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ServiceProfileLastMileConfig) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

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

### GetNotifications

`func (o *ServiceProfileLastMileConfig) GetNotifications() []ServiceProfileLastMileNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ServiceProfileLastMileConfig) GetNotificationsOk() (*[]ServiceProfileLastMileNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ServiceProfileLastMileConfig) SetNotifications(v []ServiceProfileLastMileNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ServiceProfileLastMileConfig) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



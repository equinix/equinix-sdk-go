# AppLinkAttachProcessorSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**AppLinkAttachProcessorSortDirection**](AppLinkAttachProcessorSortDirection.md) |  | [optional] [default to APPLINKATTACHPROCESSORSORTDIRECTION_DESC]
**Property** | Pointer to **string** | Possible field names to use on &#39;400_InvalidSorting&#39;:   * &#x60;/uuid&#x60; - App Processor attach to App Link uuid   * &#x60;/attachmentStatus&#x60; - App Processor attach to App Link status   * &#x60;/changeLog/createdDateTime&#x60; - Date and time when change flow starts   * &#x60;/changeLog/updatedDateTime&#x60; - Date and time when change object is updated  | [optional] [default to "/changeLog/updatedDateTime"]

## Methods

### NewAppLinkAttachProcessorSortCriteria

`func NewAppLinkAttachProcessorSortCriteria() *AppLinkAttachProcessorSortCriteria`

NewAppLinkAttachProcessorSortCriteria instantiates a new AppLinkAttachProcessorSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachProcessorSortCriteriaWithDefaults

`func NewAppLinkAttachProcessorSortCriteriaWithDefaults() *AppLinkAttachProcessorSortCriteria`

NewAppLinkAttachProcessorSortCriteriaWithDefaults instantiates a new AppLinkAttachProcessorSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *AppLinkAttachProcessorSortCriteria) GetDirection() AppLinkAttachProcessorSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AppLinkAttachProcessorSortCriteria) GetDirectionOk() (*AppLinkAttachProcessorSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AppLinkAttachProcessorSortCriteria) SetDirection(v AppLinkAttachProcessorSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AppLinkAttachProcessorSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *AppLinkAttachProcessorSortCriteria) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppLinkAttachProcessorSortCriteria) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppLinkAttachProcessorSortCriteria) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppLinkAttachProcessorSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



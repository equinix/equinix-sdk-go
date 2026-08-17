# LoaReplaceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**LoaOpEnum**](LoaOpEnum.md) |  | 
**Path** | **string** | ### A JSON Pointer path.   * &#x60;/name&#x60; - Name of the Loa   * &#x60;/description&#x60; - Description of the Loa   * &#x60;/requestor/email&#x60; - Email of the requestor contact   * &#x60;/requestor/orgName&#x60; - Organization name of the requestor   * &#x60;/issuer/email&#x60; - Email of the issuer contact   * &#x60;/issuer/orgName&#x60; - Organization name of the issuer   * &#x60;/demarcationPoint/patchPanelId&#x60; - Patch Panel Identifier   * &#x60;/demarcationPoint/patchPanelPortA&#x60; - Patch Panel Port A   * &#x60;/demarcationPoint/patchPanelPortB&#x60; - Patch Panel Port B   * &#x60;/demarcationPoint/connectorType&#x60; - Connector Type   * &#x60;/demarcationPoint/cageUniqueSpaceId&#x60; - Cage Unique Space Identifier   * &#x60;/expirationDateTime&#x60; - Expiration date and time of the Loa   * &#x60;/location/ibxCode&#x60; - IBX code of the Eligible location  | 
**Value** | **map[string]interface{}** | New value for updated parameter | 

## Methods

### NewLoaReplaceOperation

`func NewLoaReplaceOperation(op LoaOpEnum, path string, value map[string]interface{}, ) *LoaReplaceOperation`

NewLoaReplaceOperation instantiates a new LoaReplaceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaReplaceOperationWithDefaults

`func NewLoaReplaceOperationWithDefaults() *LoaReplaceOperation`

NewLoaReplaceOperationWithDefaults instantiates a new LoaReplaceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *LoaReplaceOperation) GetOp() LoaOpEnum`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *LoaReplaceOperation) GetOpOk() (*LoaOpEnum, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *LoaReplaceOperation) SetOp(v LoaOpEnum)`

SetOp sets Op field to given value.


### GetPath

`func (o *LoaReplaceOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoaReplaceOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoaReplaceOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *LoaReplaceOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoaReplaceOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoaReplaceOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


